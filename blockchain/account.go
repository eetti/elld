package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/ellcrys/druid/util"
	common "github.com/ellcrys/elld/blockchain/common"
	"github.com/ellcrys/elld/database"
	"github.com/ellcrys/elld/wire"
)

// ErrAccountNotFound refers to a missing account
var ErrAccountNotFound = fmt.Errorf("account not found")

func (b *Blockchain) putAccount(blockNo uint64, chain *Chain, account *wire.Account) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	key := common.MakeAccountKey(blockNo, chain.id, account.Address)
	return chain.store.Put(key, util.StructToBytes(account))
}

// GetAccount fetches an account on the provided chain with the matching address.
// The most recent account key representing the most recent version of the account is returned.
func (b *Blockchain) GetAccount(chain *Chain, address string) (*wire.Account, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()

	// make an account key, then query the database for an account
	// matching the key
	var result database.KVObject
	chain.store.GetFirstOrLast(false, common.QueryAccountKey(chain.id, address), &result)
	if len(result.Key) == 0 {
		return nil, ErrAccountNotFound
	}

	var account wire.Account

	return &account, json.Unmarshal(result.Value, &account)
}