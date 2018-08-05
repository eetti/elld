package wire

import (
	"math/big"

	"github.com/ellcrys/elld/crypto"
	"github.com/ellcrys/elld/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Block & Header", func() {

	var key = crypto.NewKeyFromIntSeed(1)

	Describe("Header.Bytes", func() {
		It("should successfully return bytes", func() {
			h := Header{
				ParentHash:       util.BytesToHash([]byte("parent_hash")),
				Number:           1,
				TransactionsRoot: util.BytesToHash([]byte("state_root_hash")),
				StateRoot:        util.BytesToHash([]byte("tx_hash")),
				Nonce:            EncodeNonce(1),
				MixHash:          util.BytesToHash([]byte("mixHash")),
				Difficulty:       new(big.Int).SetUint64(100),
				Timestamp:        1500000,
			}
			expected := []byte{152, 196, 32, 112, 97, 114, 101, 110, 116, 95, 104, 97, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 196, 32, 115, 116, 97, 116, 101, 95, 114, 111, 111, 116, 95, 104, 97, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 196, 32, 116, 120, 95, 104, 97, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 196, 8, 0, 0, 0, 0, 0, 0, 0, 1, 196, 32, 109, 105, 120, 72, 97, 115, 104, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 206, 0, 22, 227, 96}
			Expect(h.Bytes()).To(Equal(expected))
		})
	})

	Describe("Header.ComputeHash", func() {
		It("should successfully return 32 bytes digest", func() {
			h := Header{
				ParentHash:       util.BytesToHash([]byte("parent_hash")),
				Number:           1,
				TransactionsRoot: util.BytesToHash([]byte("state_root_hash")),
				StateRoot:        util.BytesToHash([]byte("tx_hash")),
				Nonce:            EncodeNonce(1),
				MixHash:          util.BytesToHash([]byte("mixHash")),
				Difficulty:       new(big.Int).SetUint64(100),
				Timestamp:        1500000,
			}
			actual := h.ComputeHash()
			expected := util.Hash([32]byte{146, 253, 217, 97, 144, 129, 246, 248, 132, 81, 54, 23, 53, 212, 175, 117, 135, 145, 156, 159, 235, 114, 67, 181, 24, 51, 51, 92, 36, 189, 193, 101})
			Expect(actual).To(HaveLen(32))
			Expect(actual).To(Equal(expected))
		})
	})

	Describe(".BlockSign", func() {

		var header1 = &Header{
			ParentHash:       util.BytesToHash([]byte("parent_hash")),
			Number:           1,
			TransactionsRoot: util.BytesToHash([]byte("state_root_hash")),
			StateRoot:        util.BytesToHash([]byte("tx_hash")),
			Nonce:            EncodeNonce(1),
			MixHash:          util.BytesToHash([]byte("mixHash")),
			Difficulty:       new(big.Int).SetUint64(100),
			Timestamp:        1500000,
		}

		It("should successfully sign a block", func() {

			b := Block{
				Transactions: []*Transaction{
					&Transaction{Type: TxTypeBalance, SenderPubKey: key.PubKey().Base58(), To: "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad", Value: "100.30", Timestamp: 1529670647, Fee: "0.10"},
				},
				Header: header1,
			}

			t1 := b.Transactions[0]
			t1.Hash = t1.ComputeHash()

			t1Sig, err := TxSign(t1, key.PrivKey().Base58())
			Expect(err).To(BeNil())
			t1.Sig = t1Sig

			b.Hash = b.ComputeHash()
			bs, err := BlockSign(&b, key.PrivKey().Base58())
			Expect(err).To(BeNil())
			Expect(bs).ToNot(BeEmpty())
		})
	})

	Describe(".BlockVerify", func() {

		var header1 = &Header{
			ParentHash:       util.BytesToHash([]byte("parent_hash")),
			Number:           1,
			TransactionsRoot: util.BytesToHash([]byte("state_root_hash")),
			StateRoot:        util.BytesToHash([]byte("tx_hash")),
			Nonce:            EncodeNonce(1),
			MixHash:          util.BytesToHash([]byte("mixHash")),
			Difficulty:       new(big.Int).SetUint64(100),
			Timestamp:        1500000,
		}

		It("should return err if creator pub key is not set in header", func() {

			b := Block{
				Transactions: []*Transaction{
					&Transaction{Type: TxTypeBalance, SenderPubKey: key.PubKey().Base58(), To: "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad", Value: "100.30", Timestamp: 1529670647, Fee: "0.10"},
				},
				Header: header1,
			}

			err := BlockVerify(&b)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("field:header.creatorPubKey, error:creator public not set"))
		})

		It("should return err if creator signature is not set in header", func() {

			b := Block{
				Transactions: []*Transaction{
					&Transaction{Type: TxTypeBalance, SenderPubKey: key.PubKey().Base58(), To: "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad", Value: "100.30", Timestamp: 1529670647, Fee: "0.10"},
				},
				Header: header1,
			}
			b.Header.CreatorPubKey = key.PubKey().Base58()

			err := BlockVerify(&b)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("field:sig, error:signature not set"))
		})

		It("should return error if signature is invalid", func() {

			b := Block{
				Transactions: []*Transaction{
					&Transaction{Type: TxTypeBalance, SenderPubKey: key.PubKey().Base58(), To: "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad", Value: "100.30", Timestamp: 1529670647, Fee: "0.10"},
				},
				Header: header1,
			}
			b.Header.CreatorPubKey = key.PubKey().Base58()

			t1 := b.Transactions[0]
			t1.Hash = t1.ComputeHash()

			t1Sig, err := TxSign(t1, key.PrivKey().Base58())
			Expect(err).To(BeNil())
			t1.Sig = t1Sig

			b.Hash = b.ComputeHash()

			b.Sig = []byte("invalid")
			err = BlockVerify(&b)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("block verification failed"))
		})

		It("should successfully verify a block", func() {

			b := Block{
				Transactions: []*Transaction{
					&Transaction{Type: TxTypeBalance, SenderPubKey: key.PubKey().Base58(), To: "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad", Value: "100.30", Timestamp: 1529670647, Fee: "0.10"},
				},
				Header: header1,
			}
			b.Header.CreatorPubKey = key.PubKey().Base58()

			t1 := b.Transactions[0]
			t1.Hash = t1.ComputeHash()

			t1Sig, err := TxSign(t1, key.PrivKey().Base58())
			Expect(err).To(BeNil())
			t1.Sig = t1Sig

			b.Hash = b.ComputeHash()
			bs, err := BlockSign(&b, key.PrivKey().Base58())
			Expect(err).To(BeNil())
			Expect(bs).ToNot(BeEmpty())

			b.Sig = bs
			err = BlockVerify(&b)
			Expect(err).To(BeNil())
		})
	})

	Describe("Block.ComputeHash", func() {
		It("should successfully return 32 bytes digest", func() {
			b := Block{
				Transactions: []*Transaction{
					&Transaction{
						Type:         TxTypeBalance,
						SenderPubKey: "48qgD4WR71u2fMJJNdsXmfDKNqNmiFdVo3YfnGjTA915cArTUTw",
						To:           "eGzzf1HtQL7M9Eh792iGHTvb6fsnnPipad",
						From:         "e9L9UNbcxCrnAEc8ARUnkiVrDJjW57MKdq",
						Value:        "100.30",
						Timestamp:    156707944,
						Fee:          "0.10"},
				},
				Header: &Header{
					ParentHash:       util.BytesToHash([]byte("parent_hash")),
					Number:           1,
					TransactionsRoot: util.BytesToHash([]byte("state_root_hash")),
					StateRoot:        util.BytesToHash([]byte("tx_hash")),
					Nonce:            EncodeNonce(1),
					MixHash:          util.BytesToHash([]byte("mixHash")),
					Difficulty:       new(big.Int).SetUint64(100),
					Timestamp:        1500000,
				},
			}

			actual := b.ComputeHash()
			expected := util.BytesToHash([]byte{249, 126, 96, 65, 239, 247, 137, 10, 213, 92, 48, 202, 215, 66, 213, 118, 114, 158, 241, 199, 148, 33, 193, 228, 42, 104, 53, 172, 204, 173, 44, 50})
			Expect(actual).To(Equal(expected))
		})
	})
})
