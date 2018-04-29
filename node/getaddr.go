package node

import (
	"bufio"
	"context"
	"fmt"

	"github.com/ellcrys/druid/util"
	"github.com/ellcrys/druid/wire"
	net "github.com/libp2p/go-libp2p-net"
	pc "github.com/multiformats/go-multicodec/protobuf"
)

// sendGetAddr sends a wire.GetAddr message to a remote peer.
// The remote peer will respond with a wire.Addr message which the function
// must process using the OnAddr handler and return the response.
func (pt *Inception) sendGetAddr(remotePeer *Node) ([]*wire.Address, error) {

	remotePeerIDShort := remotePeer.ShortID()
	s, err := pt.LocalPeer().addToPeerStore(remotePeer).newStream(context.Background(), remotePeer.ID(), util.GetAddrVersion)
	if err != nil {
		pt.log.Debug("GetAddr message failed. failed to connect to peer", "Err", err, "PeerID", remotePeerIDShort)
		return nil, fmt.Errorf("getaddr failed. failed to connect to peer. %s", err)
	}
	defer s.Close()

	w := bufio.NewWriter(s)
	msg := &wire.GetAddr{}
	if err := pc.Multicodec(nil).Encoder(w).Encode(msg); err != nil {
		pt.log.Debug("GetAddr failed. failed to write to stream", "Err", err, "PeerID", remotePeerIDShort)
		return nil, fmt.Errorf("getaddr failed. failed to write to stream")
	}
	w.Flush()

	pt.log.Debug("GetAddr message sent to peer", "PeerID", remotePeerIDShort)

	return pt.onAddr(s)
}

// SendGetAddr sends GetAddr message to peers in separate goroutines.
// GetAddr returns with a list of addr that should be relayed to other peers.
func (pt *Inception) SendGetAddr(remotePeers []*Node) error {

	if !pt.PM().NeedMorePeers() {
		return nil
	}

	for _, remotePeer := range remotePeers {
		rp := remotePeer
		go func() {
			addressToRelay, err := pt.sendGetAddr(rp)
			if err != nil {
				return
			}
			if len(addressToRelay) > 0 {
				pt.RelayAddr(addressToRelay)
			}
		}()
	}

	return nil
}

// OnGetAddr processes a wire.GetAddr request.
// Sends a list of active addresses to the sender
func (pt *Inception) OnGetAddr(s net.Stream) {
	defer s.Close()

	remotePeerIDShort := util.ShortID(s.Conn().RemotePeer())
	remoteAddr := util.FullRemoteAddressFromStream(s)
	remotePeer := NewRemoteNode(remoteAddr, pt.LocalPeer())

	if pt.LocalPeer().isDevMode() && !util.IsDevAddr(remotePeer.IP) {
		s.Reset()
		pt.log.Debug("Can't accept message from non local or private IP in development mode", "Addr", remotePeer.GetMultiAddr(), "Msg", "GetAddr")
		return
	}

	if !remotePeer.IsKnown() && !pt.LocalPeer().isDevMode() {
		s.Conn().Close()
		return
	}

	pt.log.Debug("Received GetAddr message", "PeerID", remotePeerIDShort)

	msg := &wire.GetAddr{}
	if err := pc.Multicodec(nil).Decoder(bufio.NewReader(s)).Decode(msg); err != nil {
		s.Reset()
		pt.log.Error("failed to read getaddr message", "Err", err, "PeerID", remotePeerIDShort)
		return
	}

	activePeers := pt.PM().GetActivePeers(0)
	if len(activePeers) > 2500 {
		activePeers = pt.PM().GetRandomActivePeers(2500)
	}

	addr := &wire.Addr{}
	for _, peer := range activePeers {
		if !pt.PM().IsLocalNode(peer) && !peer.IsSame(remotePeer) && !peer.isHardcodedSeed {
			addr.Addresses = append(addr.Addresses, &wire.Address{
				Address:   peer.GetMultiAddr(),
				Timestamp: peer.Timestamp.Unix(),
			})
		}
	}

	w := bufio.NewWriter(s)
	enc := pc.Multicodec(nil).Encoder(w)
	if err := enc.Encode(addr); err != nil {
		s.Reset()
		pt.log.Error("failed to send GetAddr response", "Err", err)
		return
	}

	pt.log.Debug("Sent GetAddr response to peer", "PeerID", remotePeerIDShort)
	w.Flush()
}
