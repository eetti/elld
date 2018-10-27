package gossip_test

import (
	"context"
	"testing"

	"github.com/ellcrys/elld/crypto"
	"github.com/ellcrys/elld/node"
	"github.com/ellcrys/elld/node/gossip"
	"github.com/ellcrys/elld/params"
	"github.com/ellcrys/elld/types/core"
	"github.com/ellcrys/elld/util"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	. "github.com/ncodes/goblin"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
)

func TestIntro(t *testing.T) {
	params.FeePerByte = decimal.NewFromFloat(0.01)
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Intro", func() {

		var lp, rp *node.Node
		var sender, _ = crypto.NewKey(nil)
		var lpPort, rpPort int

		g.BeforeEach(func() {
			lpPort = getPort()
			rpPort = getPort()

			lp = makeTestNode(lpPort)
			Expect(lp.GetBlockchain().Up()).To(BeNil())

			rp = makeTestNode(rpPort)
			Expect(rp.GetBlockchain().Up()).To(BeNil())

			// Create sender account on the remote peer
			Expect(rp.GetBlockchain().CreateAccount(1, rp.GetBlockchain().GetBestChain(), &core.Account{
				Type:    core.AccountTypeBalance,
				Address: util.String(sender.Addr()),
				Balance: "100",
			})).To(BeNil())

			// Create sender account on the local peer
			Expect(lp.GetBlockchain().CreateAccount(1, lp.GetBlockchain().GetBestChain(), &core.Account{
				Type:    core.AccountTypeBalance,
				Address: util.String(sender.Addr()),
				Balance: "100",
			})).To(BeNil())
		})

		g.AfterEach(func() {
			closeNode(lp)
			closeNode(rp)
		})

		g.Describe(".SendIntro", func() {

			// Connect lp and rp as peers
			g.BeforeEach(func() {
				lp.GetHost().Peerstore().AddAddr(rp.ID(), rp.GetAddress().DecapIPFS(), pstore.PermanentAddrTTL)
				err := lp.GetHost().Connect(context.TODO(), lp.GetHost().Peerstore().PeerInfo(rp.ID()))
				Expect(err).To(BeNil())
				lp.PM().AddOrUpdateNode(rp)
				rp.SetLocalNode(lp)
			})

			g.When("intro is successfully sent, remote peer should receive intro", func() {
				g.Specify("remote peer's intro count must be 1", func(done Done) {
					go func() {
						<-rp.GetEventEmitter().On(gossip.EventIntroReceived)
						Expect(rp.CountIntros()).To(Equal(1))
						done()
					}()
					lp.Gossip().SendIntro(nil)
				})
			})

			g.When("the intro message is explicitly passed", func() {
				g.When("intro is successfully sent, remote peer should receive intro", func() {
					g.Specify("remote peer's intro count must be 1", func(done Done) {
						go func() {
							<-rp.GetEventEmitter().On(gossip.EventIntroReceived)
							Expect(rp.CountIntros()).To(Equal(1))
							done()
						}()
						lp.Gossip().SendIntro(&core.Intro{PeerID: "12D3KooWPR29KSgCH9QkgUEkxyEkKo6Ehg6ubZxD3T74No97RW6L"})
					})
				})
			})
		})
	})
}