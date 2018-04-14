package peer

import (
	"bufio"
	"context"
	"fmt"

	"github.com/ellcrys/druid/wire"

	"github.com/ellcrys/druid/util"
	net "github.com/libp2p/go-libp2p-net"
	pc "github.com/multiformats/go-multicodec/protobuf"
)

func (protoc *Inception) sendPing(remotePeer *Peer) error {

	remotePeerID := remotePeer.IDPretty()
	s, err := protoc.LocalPeer().addToPeerStore(remotePeer).newStream(context.Background(), remotePeer.ID(), util.PingVersion)
	if err != nil {
		protocLog.Debugw("Ping failed. failed to connect to peer", "Err", err, "PeerID", remotePeerID)
		return fmt.Errorf("ping failed. failed to connect to peer. %s", err)
	}
	defer s.Close()

	w := bufio.NewWriter(s)
	msg := &wire.Ping{}
	msg.Sig = protoc.sign(msg)
	if err := pc.Multicodec(nil).Encoder(w).Encode(msg); err != nil {
		protocLog.Debugw("ping failed. failed to write to stream", "Err", err, "PeerID", remotePeerID)
		return fmt.Errorf("ping failed. failed to write to stream")
	}
	w.Flush()

	protoc.log.Infow("Sent ping to peer", "PeerID", remotePeerID)

	// receive pong response
	pongMsg := &wire.Pong{}
	decoder := pc.Multicodec(nil).Decoder(bufio.NewReader(s))
	if err := decoder.Decode(pongMsg); err != nil {
		protocLog.Debugw("Failed to read pong response", "Err", err, "PeerID", remotePeerID)
		return fmt.Errorf("failed to read pong response")
	}

	sig := pongMsg.Sig
	pongMsg.Sig = nil
	if err := protoc.verify(pongMsg, sig, s.Conn().RemotePublicKey()); err != nil {
		protoc.log.Debugw("failed to verify message signature", "Err", err, "PeerID", remotePeerID)
		return fmt.Errorf("failed to verify message signature")
	}

	protoc.PM().AddOrUpdatePeer(NewRemotePeer(util.FullRemoteAddressFromStream(s), protoc.LocalPeer()))

	protoc.log.Infow("Received pong response from peer", "PeerID", s.Conn().RemotePeer().Pretty())

	return nil
}

// SendPing sends a ping message
func (protoc *Inception) SendPing(remotePeers []*Peer) {
	protoc.log.Infow("Sending ping to peer(s)", "NumPeers", len(remotePeers))
	for _, remotePeer := range remotePeers {
		go protoc.sendPing(remotePeer)
	}
}

// OnPing handles incoming ping message
func (protoc *Inception) OnPing(s net.Stream) {

	remotePeerID := s.Conn().RemotePeer().Pretty()
	defer s.Close()

	protoc.log.Infow("Received ping message", "PeerID", remotePeerID)

	msg := &wire.Ping{}
	if err := pc.Multicodec(nil).Decoder(bufio.NewReader(s)).Decode(msg); err != nil {
		protoc.log.Errorw("failed to read ping message", "Err", err, "PeerID", remotePeerID)
		return
	}

	// verify signature
	sig := msg.Sig
	msg.Sig = nil
	if err := protoc.verify(msg, sig, s.Conn().RemotePublicKey()); err != nil {
		protoc.log.Debugw("failed to verify message signature", "Err", err, "PeerID", remotePeerID)
		return
	}

	protoc.PM().AddOrUpdatePeer(NewRemotePeer(util.FullRemoteAddressFromStream(s), protoc.LocalPeer()))

	// send pong message
	pongMsg := &wire.Pong{}
	pongMsg.Sig = protoc.sign(pongMsg)
	w := bufio.NewWriter(s)
	enc := pc.Multicodec(nil).Encoder(w)
	if err := enc.Encode(pongMsg); err != nil {
		protoc.log.Errorw("failed to send pong response", "Err", err)
		return
	}

	protoc.log.Infow("Sent pong response to peer", "PeerID", s.Conn().RemotePeer().Pretty())

	w.Flush()
}
