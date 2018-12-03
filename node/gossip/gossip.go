package gossip

import (
	"bufio"
	"context"
	"sync"
	"time"

	"github.com/ellcrys/elld/config"
	"github.com/ellcrys/elld/util"

	"github.com/vmihailenco/msgpack"

	"github.com/ellcrys/elld/node/peermanager"
	"github.com/ellcrys/elld/types"
	"github.com/ellcrys/elld/types/core"
	"github.com/ellcrys/elld/util/logger"
	net "github.com/libp2p/go-libp2p-net"
	protocol "github.com/libp2p/go-libp2p-protocol"
)

const (
	// EventReceivedBlockHashes describes an event about
	// a receiving block hashes
	EventReceivedBlockHashes = "event.receivedBlockHashes"
	// EventRequestedBlockHashes describes an event about
	// sending a request for block hashes
	EventRequestedBlockHashes = "event.requestedBlockHashes"
	// EventTransactionProcessed describes an event about
	// a processed transaction
	EventTransactionProcessed = "event.transactionProcessed"
	// EventBlockProcessed describes an event about
	// a processed block
	EventBlockProcessed = "event.blockProcessed"
	// EventBlockBodiesProcessed describes an event about
	// processed block bodies
	EventBlockBodiesProcessed = "event.blockBodiesProcessed"
	//EventAddrProcessed describes an event about
	// a processed address
	EventAddrProcessed = "event.addrProcessed"
	// EventAddressesRelayed describes an event about
	// relayed addresses
	EventAddressesRelayed = "event.addressesRelayed"
	// EventReceivedAddr describes an event about a
	// received addresses
	EventReceivedAddr = "event.receivedAddr"
	// EventIntroReceived describes an event about
	// a received intro
	EventIntroReceived = "event.receivedIntro"
)

// GossipManager represents the peer protocol
type GossipManager struct {

	// mtx is the general mutex
	mtx sync.RWMutex

	// engine represents the local node
	engine core.Engine

	// log is used for logging events
	log logger.Logger

	// broadcastersUpdatedAt is the time the
	// last relay peers where selected
	broadcastersUpdatedAt time.Time

	// broadcasters contains randomly selected
	// peers to broadcast messages to.
	broadcasters *core.BroadcastPeers

	// pm is the peer manager
	pm *peermanager.Manager
}

// NewGossip creates a new instance of the Gossip protocol
func NewGossip(p core.Engine, log logger.Logger) *GossipManager {
	return &GossipManager{
		engine:       p,
		log:          log,
		mtx:          sync.RWMutex{},
		broadcasters: core.NewBroadcastPeers(),
	}
}

// SetPeerManager sets the peer manager
func (g *GossipManager) SetPeerManager(pm *peermanager.Manager) {
	g.pm = pm
}

// PM returns the local peer's peer manager
func (g *GossipManager) PM() *peermanager.Manager {
	return g.pm
}

// GetBlockchain returns the blockchain manager
func (g *GossipManager) GetBlockchain() types.Blockchain {
	return g.engine.GetBlockchain()
}

// NewStream creates a stream for a given protocol
// ID and between the local peer and the given remote peer.
func (g *GossipManager) NewStream(remotePeer core.Engine, msgVersion string) (net.Stream,
	context.CancelFunc, error) {
	ctxDur := time.Second * time.Duration(g.engine.GetCfg().Node.MessageTimeout)
	ctx, cf := context.WithTimeout(context.TODO(), ctxDur)
	g.engine.AddToPeerStore(remotePeer)
	s, err := g.engine.GetHost().NewStream(ctx, remotePeer.ID(), protocol.ID(msgVersion))
	if err != nil {
		cf()
	}
	return s, cf, err
}

// CheckRemotePeer performs validation against the remote peer.
func (g *GossipManager) CheckRemotePeer(ws *core.WrappedStream, rp core.Engine) error {

	s := ws.Stream
	skipAcquaintanceCheck := false

	// Perform no checks for handshake messages
	if s.Protocol() == protocol.ID(config.Versions.Handshake) {
		return nil
	}

	// If we receive an Addr message from an unknown peer,
	// temporarily skip acquaintance check and allow
	// message to be processed.
	// We need to accept this unsolicited message so
	// that peer discovery will be more effective.
	if s.Protocol() == protocol.ID(config.Versions.Addr) &&
		!g.PM().PeerExist(rp.StringID()) {
		skipAcquaintanceCheck = true
	}

	// Check whether the local peer is allowed to receive
	// incoming messages from this remote peer
	if ok, err := g.PM().CanAcceptNode(rp, skipAcquaintanceCheck); !ok {
		return err
	}

	return nil
}

// Handle wrappers a protocol handler providing an
// interface to perform pre and post handling operations.
func (g *GossipManager) Handle(handler func(s net.Stream, remotePeer core.Engine) error) func(net.Stream) {
	return func(s net.Stream) {

		remoteAddr := util.RemoteAddrFromStream(s)
		rp := g.engine.NewRemoteNode(remoteAddr)

		// Check whether we are allowed to receive from this peer
		ws := &core.WrappedStream{Stream: s, Extra: make(map[string]interface{})}
		if err := g.CheckRemotePeer(ws, rp); err != nil {
			g.logErr(err, rp, "message ("+string(s.Protocol())+") unaccepted")
			s.Reset()
			return
		}

		// Update the last seen time of this peer
		g.PM().AddOrUpdateNode(rp)

		// Handle the message
		handler(s, rp)
	}
}

// ReadStream reads the content of a steam into dest
func ReadStream(s net.Stream, dest interface{}) error {
	return msgpack.NewDecoder(bufio.NewReader(s)).Decode(dest)
}

// WriteStream writes msg to the given stream
func WriteStream(s net.Stream, msg interface{}) error {
	w := bufio.NewWriter(s)
	if err := msgpack.NewEncoder(w).Encode(msg); err != nil {
		return err
	}
	w.Flush()
	return nil
}

func (g *GossipManager) logErr(err error, rp core.Engine, msg string) error {
	g.log.Debug(msg, "Err", err, "PeerID", rp.ShortID())
	return err
}

// logConnectErr updates the failure count record of a node
// that failed to connect. It will also add a 1 hour ban time
// if the node failed to connect after n tries.
func (g *GossipManager) logConnectErr(err error, rp core.Engine, msg string) error {

	// Increase connection fail count
	g.PM().IncrConnFailCount(rp.GetAddress())

	// When the peer reaches the max allowed
	// failure count, add a ban time fo 3 hours
	if !rp.IsHardcodedSeed() && g.PM().GetConnFailCount(rp.GetAddress()) >= 3 {
		g.PM().AddTimeBan(rp, 15*time.Minute)
	}

	g.log.Debug(msg, "Err", err, "PeerID", rp.ShortID())

	return types.ConnectError(err.Error())
}

// GetBroadcasters returns the broadcasters
func (g *GossipManager) GetBroadcasters() *core.BroadcastPeers {
	return g.broadcasters
}
