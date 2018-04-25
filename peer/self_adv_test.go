package peer

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ellcrys/druid/configdir"
	"github.com/ellcrys/druid/util"
)

var _ = Describe("SelfAdv", func() {

	var config = &configdir.Config{
		Peer: &configdir.PeerConfig{
			Dev:              true,
			MaxAddrsExpected: 5,
		},
	}

	Describe(".SelfAdvertise", func() {

		var err error
		var lp *Peer
		var lpProtoc *Inception
		var mgr *Manager

		BeforeEach(func() {
			lp, err = NewPeer(config, "127.0.0.1:30010", 0, log)
			Expect(err).To(BeNil())
			lpProtoc = NewInception(lp, log)
			lp.SetProtocol(lpProtoc)
			lp.SetProtocolHandler(util.AddrVersion, lpProtoc.OnAddr)
			mgr = lp.PM()
		})

		It("should successfully self advertise peer; remote peer must add the advertised peer", func() {
			p2, err := NewPeer(config, "127.0.0.1:30011", 1, log)
			Expect(err).To(BeNil())
			p2.Timestamp = time.Now()
			pt := NewInception(p2, log)
			p2.SetProtocol(pt)
			p2.SetProtocolHandler(util.AddrVersion, pt.OnAddr)

			Expect(p2.PM().knownPeers).To(HaveLen(0))
			n := lpProtoc.SelfAdvertise([]*Peer{p2})
			Expect(n).To(Equal(1))
			time.Sleep(5 * time.Millisecond)
			Expect(p2.PM().knownPeers).To(HaveLen(1))
			Expect(p2.PM().knownPeers).To(HaveKey(lp.StringID()))
		})

		AfterEach(func() {
			lp.host.Close()
		})
	})
})