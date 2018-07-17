package blockchain

import (
	"encoding/json"

	"github.com/ellcrys/elld/blockchain/common"
	"github.com/ellcrys/elld/database"
	"github.com/ellcrys/elld/util"
	"github.com/imdario/mergo"
)

// GetMeta returns the metadata of the blockchain
func (b *Blockchain) GetMeta() *common.BlockchainMeta {
	var result database.KVObject
	b.store.GetFirstOrLast(false, common.MakeBlockchainMetadataKey(), &result)
	var meta common.BlockchainMeta
	json.Unmarshal(result.Value, &meta)
	return &meta
}

// updateMeta updates the metadata of this chain
func (b *Blockchain) updateMeta(upd *common.BlockchainMeta) error {
	existingMeta := b.GetMeta()
	mergo.Merge(existingMeta, upd)
	return b.store.Put(common.MakeBlockchainMetadataKey(), util.ObjectToBytes(existingMeta))
}

// updateMetaWithTx is like updateMeta except it accepts a transaction
func (b *Blockchain) updateMetaWithTx(tx database.Tx, upd *common.BlockchainMeta) error {
	existingMeta := b.GetMeta()
	mergo.Merge(existingMeta, upd)
	return b.store.PutWithTx(tx, common.MakeBlockchainMetadataKey(), util.ObjectToBytes(existingMeta))
}
