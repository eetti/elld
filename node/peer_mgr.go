package node

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ellcrys/elld/elldb"
	"github.com/ellcrys/elld/types"
	"github.com/ellcrys/elld/util/logger"

	"github.com/ellcrys/elld/config"

	"github.com/ellcrys/elld/util"
)

// Manager manages known peers connected to the local peer.
// It is responsible for initiating and managing peers
// according to the current protocol and engine rules.
type Manager struct {
	mtx         *sync.RWMutex           // general mutex
	localNode   *Node                   // local node
	peers       map[string]types.Engine // peers known to the peer manager
	log         logger.Logger           // manager's logger
	config      *config.EngineConfig    // manager's configuration
	connMgr     *ConnectionManager      // connection manager
	stop        bool                    // signifies the start of the manager
	tickersDone chan bool
}

// NewManager creates an instance of the peer manager
func NewManager(cfg *config.EngineConfig, localPeer *Node, log logger.Logger) *Manager {

	if cfg == nil {
		cfg = &config.EngineConfig{}
		cfg.Node = &config.PeerConfig{}
	}

	m := &Manager{
		mtx:         new(sync.RWMutex),
		localNode:   localPeer,
		log:         log,
		peers:       make(map[string]types.Engine),
		config:      cfg,
		tickersDone: make(chan bool),
	}

	m.connMgr = NewConnMrg(m, log)
	m.localNode.host.Network().Notify(m.connMgr)
	return m
}

// PeerExist checks whether a peer is a known peer
func (m *Manager) PeerExist(peerID string) bool {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	_, exist := m.peers[peerID]
	return exist
}

// GetPeer returns a peer
func (m *Manager) GetPeer(peerID string) types.Engine {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if !m.PeerExist(peerID) {
		return nil
	}

	peer, _ := m.peers[peerID]
	return peer
}

// OnPeerDisconnect is called when peer disconnects.
// Active peer count is decreased by one.
// The disconnected peer is not removed from the
// known peer list because it might come back in a
// short time but subtract 2 hours from its current
// timestamp.
// Eventually, it will be removed if it does not reconnect.
func (m *Manager) OnPeerDisconnect(peerAddr util.NodeAddr) {
	peer := m.GetPeer(peerAddr.StringID())
	if peer == nil {
		return
	}
	m.HasDisconnected(peer)
	m.log.Info("Peer has disconnected", "PeerID", peer.ShortID())
	m.CleanPeers()
}

// AddPeer adds a peer
func (m *Manager) AddPeer(peer types.Engine) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.peers[peer.StringID()] = peer
}

// ConnectToPeer attempts to establish
// a connection to a peer with the given id
func (m *Manager) ConnectToPeer(peerID string) error {
	peer := m.GetPeer(peerID)
	if peer == nil {
		return fmt.Errorf("peer not found")
	}
	return m.localNode.connectToNode(peer)
}

// GetUnconnectedPeers returns the peers that
// are currently not connected to the local peer.
func (m *Manager) GetUnconnectedPeers() (peers []types.Engine) {
	for _, p := range m.GetActivePeers(0) {
		if !p.Connected() {
			peers = append(peers, p)
		}
	}
	return
}

// Manage starts managing peer connections.
// Load peers that were serialized and stored in database.
// Start connection manager
// Start periodic self advertisement to other peers
// Start periodic clean up of known peer list
// Start periodic ping messages to peers
func (m *Manager) Manage() {

	if err := m.LoadPeers(); err != nil {
		m.log.Error("failed to load peer addresses from database", "Err", err.Error())
	}

	go m.connMgr.Manage()
	go m.periodicSelfAdvertisement(m.tickersDone)
	go m.periodicCleanUp(m.tickersDone)
	go m.periodicPingMsgs(m.tickersDone)
	go m.sendPeriodicGetAddrMsg(m.tickersDone)
}

// sendPeriodicGetAddrMsg sends "getaddr"
// message to all known active peers
func (m *Manager) sendPeriodicGetAddrMsg(done chan bool) {
	ticker := time.NewTicker(time.Duration(m.config.Node.GetAddrInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			m.localNode.gProtoc.SendGetAddr(m.GetActivePeers(0))
		case <-done:
			ticker.Stop()
			return
		}
	}
}

// periodicPingMsgs sends "ping" messages to all peers
// as a basic health check routine.
func (m *Manager) periodicPingMsgs(done chan bool) {
	ticker := time.NewTicker(time.Duration(m.config.Node.PingInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			m.localNode.gProtoc.SendPing(m.GetPeers())
		case <-done:
			ticker.Stop()
			return
		}
	}
}

// periodicSelfAdvertisement send an Addr message containing only the
// local peer address to all connected peers
func (m *Manager) periodicSelfAdvertisement(done chan bool) {
	ticker := time.NewTicker(time.Duration(m.config.Node.SelfAdvInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			connectedPeers := []types.Engine{}
			for _, p := range m.GetPeers() {
				if p.Connected() {
					connectedPeers = append(connectedPeers, p)
				}
			}
			m.localNode.gProtoc.SelfAdvertise(connectedPeers)
			m.CleanPeers()
		case <-done:
			ticker.Stop()
			return
		}
	}
}

// periodicCleanUp performs peer clean up such as
// removing old know peers.
func (m *Manager) periodicCleanUp(done chan bool) {
	ticker := time.NewTicker(time.Duration(m.config.Node.CleanUpInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			nCleaned := m.CleanPeers()
			m.log.Debug("Cleaned up old peers", "NumKnownPeers", len(m.peers), "NumPeersCleaned", nCleaned)
		case <-done:
			ticker.Stop()
			return
		}
	}
}

// UpdateLastSeen updates a peer's
// last seen time to the current time
func (m *Manager) UpdateLastSeen(p types.Engine) error {

	defer func() {
		m.CleanPeers()
		m.SavePeers()
	}()

	// Get a peer matching the ID from the
	// list of peers. if it does not
	// exist, we add it immediately
	existingPeer := m.GetPeer(p.StringID())
	if existingPeer == nil {

		// Update the timestamp only if
		// the address is not set
		if p.GetLastSeen().IsZero() {
			p.SetLastSeen(time.Now())
		}

		m.AddPeer(p)
		return nil
	}

	existingPeer.SetLastSeen(time.Now())

	return nil
}

// Peers returns the map of known peers
func (m *Manager) Peers() map[string]types.Engine {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return m.peers
}

// SetPeers sets the known peers
func (m *Manager) SetPeers(d map[string]types.Engine) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.peers = d
}

// RequirePeers checks whether
// we need more peers
func (m *Manager) RequirePeers() bool {
	return len(m.GetActivePeers(0)) < 1000 && m.connMgr.needMoreConnections()
}

// IsLocalNode checks if a peer
// is the local peer
func (m *Manager) IsLocalNode(p types.Engine) bool {
	return p != nil && m.localNode != nil && m.localNode.IsSame(p)
}

// SetLocalNode sets the local node
func (m *Manager) SetLocalNode(n *Node) {
	m.localNode = n
}

// SetNumActiveConnections sets the number of active
// connections.
func (m *Manager) SetNumActiveConnections(n int64) {
	m.connMgr.activeConn = n
}

// IsActive returns true of a peer is
// considered active. First rule,
// its timestamp must be within
// the last 3 hours
func (m *Manager) IsActive(p types.Engine) bool {
	return time.Now().UTC().Add(-3 * (60 * 60) * time.Second).
		Before(p.GetLastSeen())
}

// HasDisconnected reduces the timestamp of
// a disconnected peer such that its time
// of removal is expedited. It also cleans
// up the known peers list removing peers
// that are unconnected and old.
func (m *Manager) HasDisconnected(remotePeer types.Engine) error {
	if remotePeer == nil {
		return fmt.Errorf("nil passed")
	}
	remotePeer.SetLastSeen(remotePeer.GetLastSeen().Add(-1 * time.Hour))
	m.CleanPeers()
	return nil
}

// CleanPeers removes old peers from the list
// of peers known by the local peer. Typically,
// we remove peers based on the last time
// they were seen. At least 3 connections must
// be active before we can proceed.
// It returns the number of peers removed
func (m *Manager) CleanPeers() int {
	if m.connMgr.connectionCount() < 3 {
		return 0
	}

	peers := m.GetPeers()
	before := len(peers)
	newKnownPeers := make(map[string]types.Engine)
	for k, p := range m.peers {
		if m.IsActive(p) {
			newKnownPeers[k] = p
		}
	}

	m.mtx.Lock()
	after := len(newKnownPeers)
	m.peers = newKnownPeers
	m.mtx.Unlock()

	return before - after
}

// GetPeers gets all the known
// peers (active or inactive)
func (m *Manager) GetPeers() (peers []types.Engine) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.peers {
		peers = append(peers, p)
	}

	return
}

// GetActivePeers returns active peers.
// Passing a zero or negative value
// as limit means no limit is applied.
func (m *Manager) GetActivePeers(limit int) (peers []types.Engine) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	for _, p := range m.peers {
		if limit > 0 && len(peers) >= limit {
			return
		}
		if m.IsActive(p) {
			peers = append(peers, p)
		}
	}
	return
}

// CopyActivePeers is like GetActivePeers
// but a different slice is returned
func (m *Manager) CopyActivePeers(limit int) (peers []types.Engine) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	activePeers := m.GetActivePeers(limit)
	copiedActivePeers := make([]types.Engine, len(activePeers))
	copy(copiedActivePeers, activePeers)
	return copiedActivePeers
}

// GetRandomActivePeers returns a slice
// of randomly selected peers whose
// timestamp is within 3 hours ago.
func (m *Manager) GetRandomActivePeers(limit int) []types.Engine {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	peers := m.CopyActivePeers(0)

	// shuffle known peer slice
	for i := range peers {
		j := rand.Intn(i + 1)
		peers[i], peers[j] = peers[j], peers[i]
	}

	if len(peers) <= limit {
		return peers
	}

	return peers[:limit]
}

// SavePeers stores active peer addresses
func (m *Manager) SavePeers() error {

	var numAddrs = 0
	var kvObjs []*elldb.KVObject

	// Hardcoded seed peers and peers that are
	// not up to 20 minutes old are also not saved.
	peers := m.CopyActivePeers(0)
	for _, p := range peers {
		isOldEnough := time.Now().UTC().Sub(p.CreatedAt().UTC()).Minutes() >= 20
		if !p.IsHardcodedSeed() && isOldEnough {
			key := []byte(p.StringID())
			value := util.ObjectToBytes(map[string]interface{}{
				"address":   p.GetAddress(),
				"createdAt": p.CreatedAt().Unix(),
				"lastSeen":  p.GetLastSeen().Unix(),
			})
			obj := elldb.NewKVObject(key, value, []byte("address"))
			kvObjs = append(kvObjs, obj)
			numAddrs++
		}
	}

	if err := m.localNode.db.Put(kvObjs); err != nil {
		return err
	}

	m.log.Debug("Saved peer addresses", "NumAddrs", numAddrs)

	return nil
}

// LoadPeers loads peers stored in
// the local database
func (m *Manager) LoadPeers() error {

	kvObjs := m.localNode.db.GetByPrefix([]byte("address"))

	// create remote node to represent
	// the addresses and add them to the
	// managers active peer list
	for _, o := range kvObjs {

		var addrData map[string]interface{}
		if err := o.Scan(&addrData); err != nil {
			return err
		}

		addr := util.NodeAddr(addrData["address"].(string))
		peer := NewRemoteNode(addr, m.localNode)
		peer.createdAt = time.Unix(int64(addrData["createdAt"].(uint32)), 0)
		peer.lastSeen = time.Unix(int64(addrData["lastSeen"].(uint32)), 0)
		m.AddPeer(peer)
	}

	return nil
}

// Stop gracefully stops running
// routines managed by the manager
func (m *Manager) Stop() {
	m.SavePeers()

	m.mtx.Lock()
	defer m.mtx.Unlock()
	if m.stop {
		return
	}

	if m.tickersDone != nil {
		close(m.tickersDone)
	}

	if m.connMgr.tickerDone != nil {
		close(m.connMgr.tickerDone)
	}

	m.stop = true
	m.log.Info("Peer manager has stopped")
}
