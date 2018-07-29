package wire

import (
	"github.com/ellcrys/elld/crypto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction", func() {

	Describe(".Bytes", func() {

		It("should return expected bytes", func() {
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: "some_pub_key"}
			bs := tx.Bytes()
			Expect(bs).ToNot(BeEmpty())
			Expect(bs).To(Equal([]byte{48, 45, 2, 1, 1, 2, 1, 1, 12, 12, 115, 111, 109, 101, 95, 97, 100, 100, 114, 101, 115, 115, 12, 12, 115, 111, 109, 101, 95, 112, 117, 98, 95, 107, 101, 121, 19, 0, 19, 0, 19, 0, 2, 1, 0, 4, 0}))
		})
	})

	Describe(".TxSign", func() {
		It("should return error = 'nil tx' when tx is nil", func() {
			_, err := TxSign(nil, "private key")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("nil tx"))
		})

		It("should expected signature", func() {
			seed := int64(1)
			a, _ := crypto.NewKey(&seed)
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: "some_pub_key"}
			sig, err := TxSign(tx, a.PrivKey().Base58())
			Expect(err).To(BeNil())
			Expect(sig).ToNot(BeEmpty())
			Expect(sig).To(Equal([]byte{142, 45, 87, 222, 14, 201, 253, 209, 97, 166, 83, 182, 211, 219, 218, 254, 223, 33, 43, 230, 121, 55, 213, 66, 120, 48, 243, 90, 132, 228, 241, 187, 248, 124, 70, 192, 96, 31, 0, 31, 252, 165, 196, 65, 37, 8, 244, 131, 239, 87, 227, 97, 67, 226, 235, 157, 153, 40, 39, 127, 97, 66, 197, 12}))
		})
	})

	Describe(".TxVerify", func() {
		It("should return error = 'nil tx' when nil is passed", func() {
			err := TxVerify(nil)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("nil tx"))
		})

		It("should return err = 'sender public not set' when sender public key is not set on the transaction", func() {
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address"}
			err := TxVerify(tx)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("field:senderPubKey, error:sender public not set"))
		})

		It("should return err = 'signature not set' when signature is unset", func() {
			tx := &Transaction{Type: 1, Nonce: 1, SenderPubKey: "pub key", To: "some_address"}
			err := TxVerify(tx)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("field:sig, error:signature not set"))
		})

		It("should return error if senderPubKey is invalid", func() {
			tx := &Transaction{Type: 1, Nonce: 1, SenderPubKey: "pub key", To: "some_address", Sig: "0xsig"}
			err := TxVerify(tx)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal("field:senderPubKey, error:invalid format: version and/or checksum bytes missing"))
		})

		It("should verify successfully", func() {
			seed := int64(1)
			a, _ := crypto.NewKey(&seed)
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: a.PubKey().Base58()}
			sig, err := TxSign(tx, a.PrivKey().Base58())
			Expect(err).To(BeNil())
			Expect(sig).ToNot(BeEmpty())

			tx.Sig = ToHex(sig)
			err = TxVerify(tx)
			Expect(err).To(BeNil())
		})

		It("should return err = 'verify failed' when verification failed", func() {
			seed := int64(1)
			a, _ := crypto.NewKey(&seed)
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: a.PubKey().Base58()}
			sig, err := TxSign(tx, a.PrivKey().Base58())
			Expect(err).To(BeNil())
			Expect(sig).ToNot(BeEmpty())

			tx.Sig = ToHex(sig)
			tx.To = "altered_address"
			err = TxVerify(tx)
			Expect(err).ToNot(BeNil())
			Expect(err).To(Equal(crypto.ErrTxVerificationFailed))
		})
	})

	Describe(".ComputeHash", func() {
		It("should successfully return hash", func() {
			seed := int64(1)
			a, _ := crypto.NewKey(&seed)
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: a.PubKey().Base58()}
			Expect(tx.ComputeHash()).To(Equal([]byte{216, 173, 132, 254, 182, 58, 50, 25, 139, 64, 171, 196, 62, 53, 230, 26, 75, 197, 156, 17, 182, 48, 42, 243, 187, 248, 173, 34, 164, 114, 76, 42}))
		})
	})

	Describe(".ID", func() {
		It("should return '0xd8ad84feb63a32198b40abc43e35e61a4bc59c11b6302af3bbf8ad22a4724c2a'", func() {
			seed := int64(1)
			a, _ := crypto.NewKey(&seed)
			tx := &Transaction{Type: 1, Nonce: 1, To: "some_address", SenderPubKey: a.PubKey().Base58()}
			Expect(tx.ID()).To(Equal("0xd8ad84feb63a32198b40abc43e35e61a4bc59c11b6302af3bbf8ad22a4724c2a"))
		})
	})

})
