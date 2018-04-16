package peer

import (
	"bufio"
	"context"
	"fmt"

	"github.com/ellcrys/druid/util"
	"github.com/ellcrys/druid/wire"
	net "github.com/libp2p/go-libp2p-net"
	pc "github.com/multiformats/go-multicodec/protobuf"
)

// SendHandshake sends an introduction message to a peer
func (protoc *Inception) SendHandshake(remotePeer *Peer) error {

	remotePeerIDShort := remotePeer.ShortID()
	protoc.log.Infow("Sending handshake to peer", "PeerID", remotePeerIDShort)

	s, err := protoc.LocalPeer().addToPeerStore(remotePeer).newStream(context.Background(), remotePeer.ID(), util.HandshakeVersion)
	if err != nil {
		protocLog.Debugw("Handshake failed. failed to connect to peer", "Err", err, "PeerID", remotePeerIDShort)
		return fmt.Errorf("handshake failed. failed to connect to peer. %s", err)
	}
	defer s.Close()

	// send handshake
	w := bufio.NewWriter(s)
	msg := &wire.Handshake{Address: protoc.LocalPeer().GetMultiAddr(), SubVersion: util.ClientVersion}
	msg.Sig = protoc.sign(msg)
	if err := pc.Multicodec(nil).Encoder(w).Encode(msg); err != nil {
		protocLog.Debugw("Handshake failed. failed to write to stream", "Err", err, "PeerID", remotePeerIDShort)
		return fmt.Errorf("handshake failed. failed to write to stream")
	}
	w.Flush()

	protoc.log.Debugw("Sent handshake to peer", "PeerID", remotePeerIDShort)

	// receive response
	resp := &wire.HandshakeResponse{}
	decoder := pc.Multicodec(nil).Decoder(bufio.NewReader(s))
	if err := decoder.Decode(resp); err != nil {
		protocLog.Debugw("Failed to read handshake response", "Err", err, "PeerID", remotePeerIDShort)
		return fmt.Errorf("failed to read handshake response")
	}

	// verify message signature
	sig := resp.Sig
	resp.Sig = nil
	if err := protoc.verify(resp, sig, s.Conn().RemotePublicKey()); err != nil {
		protoc.log.Debugw("failed to verify message signature", "Err", err, "PeerID", remotePeerIDShort)
		return fmt.Errorf("failed to verify message signature")
	}

	protoc.PM().AddOrUpdatePeer(NewRemotePeer(util.FullRemoteAddressFromStream(s), protoc.LocalPeer()))

	// validate address before adding
	invalidAddrs := 0
	for _, addr := range resp.Addresses {
		if !util.IsValidAddress(addr) {
			invalidAddrs++
			protoc.log.Debugw("Found invalid address in handshake", "Addr", addr)
			continue
		}
		p, _ := protoc.LocalPeer().PeerFromAddr(addr, true)
		protoc.PM().AddOrUpdatePeer(p)
	}

	protoc.log.Infow("Received handshake response from peer", "PeerID", remotePeerIDShort, "NumAddrs", len(resp.Addresses), "InvalidAddrs", invalidAddrs)
	return nil
}

// OnHandshake handles incoming handshake request
func (protoc *Inception) OnHandshake(s net.Stream) {

	remotePeerIDShort := util.ShortID(s.Conn().RemotePeer())
	remotePeerID := s.Conn().RemotePeer().Pretty()
	defer s.Close()

	protoc.log.Infow("Received handshake message", "PeerID", remotePeerIDShort)

	msg := &wire.Handshake{}
	if err := pc.Multicodec(nil).Decoder(bufio.NewReader(s)).Decode(msg); err != nil {
		protoc.log.Errorw("failed to read handshake message", "Err", err, "PeerID", remotePeerIDShort)
		return
	}

	// verify signature
	sig := msg.Sig
	msg.Sig = nil
	if err := protoc.verify(msg, sig, s.Conn().RemotePublicKey()); err != nil {
		protoc.log.Debugw("failed to verify handshake message signature", "Err", err, "PeerID", remotePeerIDShort)
		return
	}

	protoc.PM().AddOrUpdatePeer(NewRemotePeer(util.FullRemoteAddressFromStream(s), protoc.LocalPeer()))

	// get active peers
	var addresses []string
	peers := protoc.PM().CopyActivePeers(1000)
	for _, p := range peers {
		if p.StringID() != remotePeerID {
			addresses = append(addresses, p.GetMultiAddr())
		}
	}

	// create response message, sign it and add the signature to the message
	addrMsg := &wire.HandshakeResponse{Addresses: addresses}
	addrMsg.Sig = protoc.sign(addrMsg)
	w := bufio.NewWriter(s)
	enc := pc.Multicodec(nil).Encoder(w)
	if err := enc.Encode(addrMsg); err != nil {
		protoc.log.Errorw("failed to send handshake response", "Err", err)
		return
	}

	protoc.log.Infow("Sent handshake response to peer", "PeerID", remotePeerIDShort, "NumAddr", len(addresses))

	w.Flush()
}
