package blockchain

import (
	"fmt"
	"time"

	"github.com/ellcrys/elld/blockchain/testdata"
	"github.com/ellcrys/elld/wire"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var BlockTest = func() bool {
	return Describe("Block", func() {

		Describe(".HaveBlock", func() {

			var block *wire.Block

			BeforeEach(func() {
				block, err = wire.BlockFromString(testdata.BlockchainDotGoJSON[1])
				Expect(err).To(BeNil())
			})

			It("should return false when block does not exist in any known chain", func() {
				has, err := bc.HaveBlock(block.GetHash())
				Expect(err).To(BeNil())
				Expect(has).To(BeFalse())
			})

			It("should return true of block exists in a chain", func() {
				chain2 := NewChain("chain2", store, cfg, log)
				Expect(err).To(BeNil())
				err = chain2.append(block)
				Expect(err).To(BeNil())

				err = bc.addChain(chain2)
				Expect(err).To(BeNil())
				err = chain2.store.PutBlock(chain2.id, block)
				Expect(err).To(BeNil())

				has, err := bc.HaveBlock(block.GetHash())
				Expect(err).To(BeNil())
				Expect(has).To(BeTrue())
			})
		})

		Describe(".IsKnownBlock", func() {
			var block *wire.Block

			BeforeEach(func() {
				block, err = wire.BlockFromString(testdata.BlockchainDotGoJSON[1])
				Expect(err).To(BeNil())
			})

			It("should return false when block does not exist in any known chain or caches", func() {
				exist, reason, err := bc.IsKnownBlock(block.GetHash())
				Expect(err).To(BeNil())
				Expect(exist).To(BeFalse())
				Expect(reason).To(BeEmpty())
			})

			It("should return true when block exists in a chain", func() {
				chain2 := NewChain("chain2", store, cfg, log)
				Expect(err).To(BeNil())
				err = chain2.append(block)
				Expect(err).To(BeNil())

				err = bc.addChain(chain2)
				Expect(err).To(BeNil())
				err = chain2.store.PutBlock(chain2.id, block)
				Expect(err).To(BeNil())

				has, err := bc.HaveBlock(block.GetHash())
				Expect(err).To(BeNil())
				Expect(has).To(BeTrue())
			})

			It("should return true when block exist as an orphan", func() {
				bc.addOrphanBlock(block)
				known, reason, err := bc.IsKnownBlock(block.GetHash())
				Expect(err).To(BeNil())
				Expect(known).To(BeTrue())
				Expect(reason).To(Equal("orphan cache"))
			})
		})

		Describe(".GenerateBlock", func() {

			var txs []*wire.Transaction

			BeforeEach(func() {
				bc.bestChain = chain
				txs = []*wire.Transaction{wire.NewTx(wire.TxTypeBalance, 123, receiver.Addr(), sender, "0.1", "0.1", time.Now().Unix())}
			})

			It("should validate params", func() {
				var cases = map[*GenerateBlockParams]interface{}{
					&GenerateBlockParams{}:                                                       fmt.Errorf("at least one transaction is required"),
					&GenerateBlockParams{Transactions: txs}:                                      fmt.Errorf("creator's key is required"),
					&GenerateBlockParams{Transactions: txs, Creator: sender}:                     fmt.Errorf("difficulty is required"),
					&GenerateBlockParams{Transactions: txs, Creator: sender, Difficulty: "1234"}: fmt.Errorf("mix hash is required"),
				}

				for m, r := range cases {
					_, err = bc.GenerateBlock(m)
					Expect(err).To(Equal(r))
				}
			})

			It("should successfully create a new and valid block", func() {
				blk, err := bc.GenerateBlock(&GenerateBlockParams{
					Transactions: txs,
					Creator:      sender,
					Nonce:        384772,
					MixHash:      "0x1cdf0e214bcdb7af36885316506f7388f262f7b710a28a00d21706550cdd72c2",
					Difficulty:   "102993",
				})
				Expect(err).To(BeNil())
				Expect(blk).ToNot(BeNil())
				Expect(blk.Header.StateRoot).ToNot(BeEmpty())
				Expect(blk.Header.Number).To(Equal(uint64(2)))
				Expect(blk.Header.ParentHash).To(Equal(block.Hash))
			})

			It("should return error if final block failed general block validation rules", func() {
				blk, err := bc.GenerateBlock(&GenerateBlockParams{
					Transactions: txs,
					Creator:      sender,
					Nonce:        384772,
					MixHash:      "invalid_mix_hash",
					Difficulty:   "102993",
				})
				Expect(err).ToNot(BeNil())
				Expect(blk).To(BeNil())
				Expect(err.Error()).To(Equal("failed final validation: field:header.mixHash, error:state root has invalid length. 66 chars. required"))
			})

			When("chain is directly passed", func() {
				It("should successfully create a new and valid block", func() {
					blk, err := bc.GenerateBlock(&GenerateBlockParams{
						Transactions: txs,
						Creator:      sender,
						Nonce:        384772,
						MixHash:      "0x1cdf0e214bcdb7af36885316506f7388f262f7b710a28a00d21706550cdd72c2",
						Difficulty:   "102993",
					}, ChainOp{Chain: chain})
					Expect(err).To(BeNil())
					Expect(blk).ToNot(BeNil())
					Expect(blk.Header.StateRoot).ToNot(BeEmpty())
					Expect(blk.Header.Number).To(Equal(uint64(2)))
					Expect(blk.Header.ParentHash).To(Equal(block.Hash))
				})
			})

			When("best chain is nil and no chain is passed directly", func() {

				BeforeEach(func() {
					bc.bestChain = nil
				})

				It("should return error if not target chain", func() {
					blk, err := bc.GenerateBlock(&GenerateBlockParams{
						Transactions: txs,
						Creator:      sender,
						Nonce:        384772,
						MixHash:      "abc",
						Difficulty:   "102993",
					})
					Expect(err).ToNot(BeNil())
					Expect(blk).To(BeNil())
					Expect(err.Error()).To(Equal("target chain not set"))
				})
			})

			When("target chain state does not include the sender account", func() {

				var targetChain *Chain

				BeforeEach(func() {
					targetChain = NewChain("abc", store, cfg, log)
					targetChain.parentBlock = block
				})

				It("should return error if not target chain", func() {
					blk, err := bc.GenerateBlock(&GenerateBlockParams{
						Transactions: txs,
						Creator:      sender,
						Nonce:        384772,
						MixHash:      "abc",
						Difficulty:   "102993",
					}, ChainOp{Chain: targetChain})
					Expect(err).ToNot(BeNil())
					Expect(blk).To(BeNil())
					Expect(err.Error()).To(Equal("exec: transaction error: index{0}: failed to get sender's account: account not found"))
				})
			})

			When("target chain has no block but has a parent block attached", func() {

				var targetChain *Chain

				BeforeEach(func() {
					targetChain = NewChain("abc", store, cfg, log)
					// block.Header.ParentHash = "0x1cdf0e214bcdb7af36885316506f7388f262f7b710a28a00d21706550cdd72c2"
					targetChain.parentBlock = block
				})

				BeforeEach(func() {
					err = bc.putAccount(1, targetChain, &wire.Account{
						Type:    wire.AccountTypeBalance,
						Address: sender.Addr(),
						Balance: "100",
					})
					Expect(err).To(BeNil())
				})

				It("should successfully create a new and valid block", func() {
					blk, err := bc.GenerateBlock(&GenerateBlockParams{
						Transactions: txs,
						Creator:      sender,
						Nonce:        384772,
						MixHash:      "0x1cdf0e214bcdb7af36885316506f7388f262f7b710a28a00d21706550cdd72c2",
						Difficulty:   "102993",
					}, ChainOp{Chain: targetChain})
					Expect(err).To(BeNil())
					Expect(blk).ToNot(BeNil())
					Expect(blk.Header.StateRoot).ToNot(BeEmpty())
					Expect(blk.Header.Number).To(Equal(uint64(2)))
					Expect(blk.Header.ParentHash).To(Equal(block.Hash))
				})
			})

			When("target chain has no tip block and no parent block", func() {

				var targetChain *Chain

				BeforeEach(func() {
					targetChain = NewChain("abc", store, cfg, log)
				})

				BeforeEach(func() {
					err = bc.putAccount(1, targetChain, &wire.Account{
						Type:    wire.AccountTypeBalance,
						Address: sender.Addr(),
						Balance: "100",
					})
					Expect(err).To(BeNil())
				})

				It("should create a 'genesis' block", func() {
					blk, err := bc.GenerateBlock(&GenerateBlockParams{
						Transactions: txs,
						Creator:      sender,
						Nonce:        384772,
						MixHash:      "0x1cdf0e214bcdb7af36885316506f7388f262f7b710a28a00d21706550cdd72c2",
						Difficulty:   "102993",
					}, ChainOp{Chain: targetChain})
					Expect(err).To(BeNil())
					Expect(blk).ToNot(BeNil())
					Expect(blk.Header.StateRoot).ToNot(BeEmpty())
					Expect(blk.Header.Number).To(Equal(uint64(1)))
					Expect(blk.Header.ParentHash).To(BeEmpty())
				})
			})
		})
	})
}
