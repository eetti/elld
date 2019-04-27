package gossip

import (
	"fmt"

	"github.com/ellcrys/elld/node/common"
	"github.com/ellcrys/elld/util/cache"
	"github.com/jinzhu/copier"

	"github.com/ellcrys/elld/config"
	"github.com/ellcrys/elld/params"
	"github.com/ellcrys/elld/types"
	"github.com/ellcrys/elld/types/core"
	"github.com/ellcrys/elld/util"
	net "github.com/libp2p/go-libp2p-net"
)

// BroadcastBlock sends a given block to remote peers.
// The block is encapsulated in a BlockBody message.
func (g *Manager) BroadcastBlock(block types.Block, remotePeers []core.Engine) []error {

	var sent int
	var errs []error
	broadcastPeers := g.PickBroadcastersFromPeers(g.broadcasters, remotePeers, 3)
	for _, peer := range broadcastPeers.Peers() {

		// We need to remove the broadcast peer
		// if it is no longer connected
		if !peer.Connected() {
			broadcastPeers.Remove(peer)
			continue
		}

		// Check if we have an history of receiving
		// this block from this peer recently. If yes,
		// we will not proceed further
		hk := common.KeyBlock(block.GetHashAsHex(), peer)
		if g.engine.GetHistory().HasMulti(hk...) {
			continue
		}

		g.log.Debug("Broadcasting a block",
			"PeerID", peer.ShortID(),
			"BlockNo", block.GetNumber(),
			"BlockHash", block.GetHash().SS(),
			"NumPeers", len(remotePeers))

		s, c, err := g.NewStream(peer, config.GetVersions().BlockInfo)
		if err != nil {
			errs = append(errs, err)
			g.logConnectErr(err, peer, "[BroadcastBlock] Failed to connect")
			continue
		}
		defer c()

		// Send a message describing the block.
		// If the peer accepts the block, we can send the full block.
		blockInfo := core.BlockInfo{Hash: block.GetHash()}
		if err := WriteStream(s, blockInfo); err != nil {
			errs = append(errs, err)
			s.Reset()
			g.logErr(err, peer, "[BroadcastBlock] Failed to write to stream")
			continue
		}

		// Read BlockOk message to know whether to send the block
		blockOk := &core.BlockOk{}
		if err := ReadStream(s, blockOk); err != nil {
			errs = append(errs, err)
			s.Reset()
			g.logErr(err, peer, "[BroadcastBlock] Failed to read BlockOk message")
			continue
		}

		if !blockOk.Ok {
			s.Close()
			errs = append(errs, fmt.Errorf("block rejected by remote peer"))
			g.log.Debug("Peer rejected our intent to broadcast a block",
				"PeerID", peer.ShortID(),
				"BlockNo", block.GetNumber(),
				"BlockHash", block.GetHash().SS())
			continue
		}

		s.Close()

		// At this point, we can send the block to the peer.
		// First we need to create a new stream targeting the
		// BlockBody handler
		s2, c2, err := g.NewStream(peer, config.GetVersions().BlockBody)
		if err != nil {
			errs = append(errs, err)
			g.logConnectErr(err, peer, "[BroadcastBlock] Failed to connect to peer")
			continue
		}
		defer c2()

		var blockBody core.BlockBody
		copier.Copy(&blockBody, block)
		if err := WriteStream(s2, blockBody); err != nil {
			errs = append(errs, err)
			s2.Reset()
			g.logErr(err, peer, "[BroadcastBlock] Failed to write BlockBody")
			continue
		}

		s2.Close()

		sent++
	}

	g.log.Debug("Block broadcast completed",
		"BlockNo", block.GetNumber(),
		"BlockHash", block.GetHash().SS(),
		"NumPeersSentTo", sent)

	return errs
}

// OnBlockInfo handles incoming BlockInfo messages.
// BlockInfo messages describe a block that a peer
// intends to send. The local peer responds with a
// BlockOk message if it accepts the block.
func (g *Manager) OnBlockInfo(s net.Stream, rp core.Engine) error {

	msg := &core.BlockInfo{}
	if err := ReadStream(s, msg); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[OnBlockInfo] Failed to read BlockInfo message")
	}

	// If synchronization is disabled, do not accept the block
	if g.engine.GetSyncMode().IsDisabled() {
		goto blk_not_ok
	}

	// We can't accept a block we already know
	if existingBlock, _ := g.engine.GetBlockchain().HaveBlock(msg.Hash); existingBlock {
		goto blk_not_ok
	}

	// Send back BlockOk message indicating readiness
	// to receive the block
	if err := WriteStream(s, &core.BlockOk{Ok: true}); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[OnBlockInfo] Failed to write BlockOk message")
	}

blk_not_ok:
	if err := WriteStream(s, &core.BlockOk{Ok: false}); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[OnBlockInfo] Failed to write BlockInfo message")
	}

	return s.Close()
}

// OnBlockBody handles incoming BlockBody messages.
// BlockBody messages contain information about a
// block. It will attempt to process the received
// block.
func (g *Manager) OnBlockBody(s net.Stream, rp core.Engine) error {

	defer s.Close()

	blockBody := &core.BlockBody{}
	if err := ReadStream(s, &blockBody); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[OnBlockBody] Failed to read")
	}

	var block core.Block
	copier.Copy(&block, blockBody)
	block.SetBroadcaster(rp)

	g.log.Info("Received a block",
		"BlockNo", block.GetNumber(),
		"BlockHash", block.GetHash().SS(),
		"Difficulty", block.GetHeader().GetDifficulty())

	// Keep a record of the receipt of this block.
	// This will help us avoiding broadcasting the same
	// block to the sender.
	hk := common.KeyBlock(block.GetHashAsHex(), rp)
	g.engine.GetHistory().AddMulti(cache.Sec(600), hk...)

	// Emit core.EventRelayedBlock to have the block
	// processed by the block manager.
	go g.engine.GetEventEmitter().Emit(core.EventProcessBlock, &block)

	return nil
}

// RequestBlock sends a RequestBlock message to remote peer.
// A RequestBlock message includes information about a
// specific block. It will attempt to process the requested
// block after receiving it from the remote peer.
// The block's validation context is set to ContextBlockSync
// which cause the transactions to not be required to exist
// in the transaction pool.
func (g *Manager) RequestBlock(rp core.Engine, blockHash util.Hash) error {

	hk := common.KeyOrphanBlock(blockHash, rp)
	if g.engine.GetHistory().HasMulti(hk...) {
		return nil
	}

	s, c, err := g.NewStream(rp, config.GetVersions().RequestBlock)
	if err != nil {
		return g.logConnectErr(err, rp, "[RequestBlock] Failed to connect to peer")
	}
	defer c()
	defer s.Reset()

	msg := &core.RequestBlock{Hash: blockHash.HexStr()}
	if err := WriteStream(s, msg); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[RequestBlock] Failed to write to peer")
	}

	var blockBody core.BlockBody
	if err := ReadStream(s, &blockBody); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[RequestBlock] Failed to read")
	}

	// Emit core.EventProcessBlock to have
	// the block processed by the block manager.
	var block core.Block
	copier.Copy(&block, blockBody)
	block.SetBroadcaster(rp)
	go g.engine.GetEventEmitter().Emit(core.EventProcessBlock, &block)

	g.engine.GetHistory().AddMulti(cache.Sec(600), hk...)

	return nil
}

// OnRequestBlock handles RequestBlock message.
// A RequestBlock message includes information
// a bout a block that a remote node needs.
func (g *Manager) OnRequestBlock(s net.Stream, rp core.Engine) error {

	defer s.Close()

	msg := &core.RequestBlock{}
	if err := ReadStream(s, msg); err != nil {
		s.Reset()
		return g.logErr(err, rp, "[OnRequestBlock] Failed to read")
	}

	g.log.Debug("Received request for block",
		"RequestedBlockHash", util.StrToHash(msg.Hash).SS())

	if msg.Hash == "" {
		s.Reset()
		err := fmt.Errorf("Invalid RequestBlock message: empty 'Hash' field")
		g.log.Debug(err.Error(), "PeerID", rp.ShortID())
		return err
	}

	var block types.Block

	// decode the hex into a util.Hash
	blockHash, err := util.HexToHash(msg.Hash)
	if err != nil {
		s.Reset()
		g.log.Debug("Invalid hash supplied in requestblock message",
			"PeerID", rp.ShortID(), "Hash", msg.Hash)
		return err
	}

	// find the block
	block, err = g.GetBlockchain().GetBlockByHash(blockHash)
	if err != nil {
		if err != core.ErrBlockNotFound {
			s.Reset()
			g.log.Error(err.Error())
			return err
		}
		s.Reset()
		g.log.Debug("Requested block is not found", "PeerID", rp.ShortID(),
			"Hash", util.StrToHash(msg.Hash).SS())
		return err
	}

	var blockBody core.BlockBody
	copier.Copy(&blockBody, block)
	if err := WriteStream(s, blockBody); err != nil {
		s.Reset()
		g.logErr(err, rp, "[OnRequestBlock] Failed to write")
	}

	return nil
}

// SendGetBlockHashes sends a GetBlockHashes message to
// the remotePeer asking for block hashes beginning from
// a block they share in common. The local peer sends the
// remote peer a list of hashes (locators) while the
// remote peer use the locators to find the highest
// block height they share in common, then it gathers
// and sends block hashes after the chosen shared block.
//
// If the locators is not provided via the locator argument,
// they will be collected from the main chain.
func (g *Manager) SendGetBlockHashes(rp core.Engine,
	locators []util.Hash, seek util.Hash) (*core.BlockHashes, error) {
	rpID := rp.ShortID()
	g.log.Debug("Requesting block headers", "PeerID", rpID)

	s, c, err := g.NewStream(rp, config.GetVersions().GetBlockHashes)
	if err != nil {
		return nil, g.logConnectErr(err, rp, "[SendGetBlockHashes] Failed to connect")
	}
	defer c()
	defer s.Close()

	if len(locators) == 0 {
		locators, err = g.GetBlockchain().GetLocators()
		if err != nil {
			g.log.Error("failed to get locators", "Err", err)
			return nil, err
		}
	}

	msg := core.GetBlockHashes{
		Locators:  locators,
		Seek:      seek,
		MaxBlocks: params.MaxGetBlockHashes,
	}

	if err := WriteStream(s, msg); err != nil {
		return nil, g.logErr(err, rp, "[SendGetBlockHashes] Failed to write")
	}

	go g.engine.GetEventEmitter().Emit(EventRequestedBlockHashes,
		msg.Locators, msg.MaxBlocks)

	// Read the return block hashes
	var blockHashes core.BlockHashes
	if err := ReadStream(s, &blockHashes); err != nil {
		return nil, g.logErr(err, rp, "[SendGetBlockHashes] Failed to read")
	}

	go g.engine.GetEventEmitter().Emit(EventReceivedBlockHashes)
	g.log.Info("Successfully requested block headers", "PeerID", rpID, "NumLocators",
		len(msg.Locators))

	return &blockHashes, nil
}

// OnGetBlockHashes processes a core.GetBlockHashes request.
// It will attempt to find a chain it shares in common using
// the locator block hashes provided in the message.
//
// If it does not find a chain that is shared with the remote
// chain, it will assume the chains are not off same network
// and as such send an empty block hash response.
//
// If it finds that the remote peer has a chain that is
// not the same as its main chain (a branch), it will
// send block hashes starting from the root parent block (oldest
// ancestor) which exists on the main chain.
func (g *Manager) OnGetBlockHashes(s net.Stream, rp core.Engine) error {

	defer s.Close()

	// Read the message
	msg := &core.GetBlockHashes{}
	if err := ReadStream(s, msg); err != nil {
		return g.logErr(err, rp, "[OnGetBlockHashes] Failed to read")
	}

	var blockHashes = core.BlockHashes{}
	var startBlock types.Block
	var blockCursor uint64
	var locatorChain types.ChainReaderFactory
	var locatorHash util.Hash
	var mainChain = g.GetBlockchain().GetBestChain()

	// If there is a seek hash,
	if !msg.Seek.IsEmpty() {
		// Find the chain where a block matches the seek hash.
		// If no such chain exist or the chain is not the main chain,
		// We must fall back to locators, otherwise,
		locatorChain = g.GetBlockchain().GetChainReaderByHash(msg.Seek)
		if locatorChain != nil && locatorChain.GetID().Equal(mainChain.GetID()) {
			// Discard all locators and use the seek hash as the sole locator
			msg.Locators = []util.Hash{msg.Seek}
		}
	}

	// Using the provided locator hashes, find a chain
	// where one of the locator block exists. Expects the
	// order of the locator to begin with the highest
	// tip block hash of the remote node
	for _, hash := range msg.Locators {
		locatorChain = g.GetBlockchain().GetChainReaderByHash(hash)
		if locatorChain != nil {
			locatorHash = hash
			break
		}
	}

	// Since we didn't find any common chain,
	// we will assume the node does not share
	// any similarity with the local peer's network
	// as such return nothing
	if locatorChain == nil {
		blockHashes = core.BlockHashes{}
		goto send
	}

	// Check whether the locator's chain is the main
	// chain. If it is not, we need to get the root
	// parent block from which the chain (and its parent)
	// sprouted from. Otherwise, get the locator block
	// and use as the start block.
	if mainChain.GetID() != locatorChain.GetID() {
		startBlock = locatorChain.GetRoot()
	} else {
		startBlock, _ = locatorChain.GetBlockByHash(locatorHash)
	}

	// This should only be true when chain tree
	// structure has been corrupted on disk.
	if startBlock == nil {
		g.log.Warn("Could not get the sync start block. " +
			"Possible chain tree corruption.")
		return nil
	}

	// Fetch block hashes starting from the block
	// after the start block
	blockCursor = startBlock.GetNumber() + 1
	for int64(len(blockHashes.Hashes)) <= msg.MaxBlocks {
		block, err := g.GetBlockchain().ChainReader().GetBlock(blockCursor)
		if err != nil {
			if err != core.ErrBlockNotFound {
				g.log.Error("Failed to fetch block header", "Err", err)
			}
			break
		}
		blockHashes.Hashes = append(blockHashes.Hashes, block.GetHash())
		blockCursor++
	}

send:
	if err := WriteStream(s, blockHashes); err != nil {
		g.logErr(err, rp, "[OnGetBlockHashes] Failed to write")
		return err
	}

	return nil
}

// SendGetBlockBodies sends a GetBlockBodies message
// requesting for whole bodies of a collection blocks.
func (g *Manager) SendGetBlockBodies(rp core.Engine, hashes []util.Hash) (*core.BlockBodies, error) {

	rpID := rp.ShortID()
	g.log.Debug("Requesting block bodies", "PeerID", rpID, "NumHashes", len(hashes))

	s, c, err := g.NewStream(rp, config.GetVersions().GetBlockBodies)
	if err != nil {
		return nil, g.logConnectErr(err, rp, "[SendGetBlockBodies] Failed to connect")
	}
	defer c()
	defer s.Close()

	// do nothing if no hash is given
	if len(hashes) == 0 {
		return &core.BlockBodies{}, nil
	}

	msg := core.GetBlockBodies{
		Hashes: hashes,
	}

	// write to the stream
	if err := WriteStream(s, msg); err != nil {
		return nil, g.logErr(err, rp, "[SendGetBlockBodies] Failed to write")
	}

	// Read the return block bodies
	var blockBodies core.BlockBodies
	if err := ReadStream(s, &blockBodies); err != nil {
		return nil, g.logErr(err, rp, "[SendGetBlockBodies] Failed to read")
	}

	return &blockBodies, nil
}

// OnGetBlockBodies handles GetBlockBodies requests
func (g *Manager) OnGetBlockBodies(s net.Stream, rp core.Engine) error {
	defer s.Close()

	// Read the message
	msg := &core.GetBlockBodies{}
	if err := ReadStream(s, msg); err != nil {
		return g.logErr(err, rp, "[OnGetBlockBodies] Failed to read")
	}

	var bestChain = g.GetBlockchain().ChainReader()
	var blockBodies = new(core.BlockBodies)
	for _, hash := range msg.Hashes {
		block, err := bestChain.GetBlockByHash(hash)
		if err != nil {
			if err != core.ErrBlockNotFound {
				g.log.Error("Failed fetch block body of a given hash", "Err", err,
					"Hash", hash)
				return err
			}
			continue
		}
		var blockBody core.BlockBody
		copier.Copy(&blockBody, block)
		blockBodies.Blocks = append(blockBodies.Blocks, &blockBody)
	}

	// send the block bodies
	if err := WriteStream(s, blockBodies); err != nil {
		g.logErr(err, rp, "[OnGetBlockBodies] Failed to write")
		return err
	}

	return nil
}
