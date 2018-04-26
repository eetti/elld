package database

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

// PrefixAddress is the prefixed used to scope address data in the db
var PrefixAddress = "address-"

func makeKey(k string) []byte {
	bs := hex.EncodeToString(sha3.New256().Sum([]byte(k)))
	return append([]byte(PrefixAddress), []byte(bs)...)[:40]
}

// AddressStore provides query and storage capabilities for addresses
type AddressStore struct {
	db DB
}

// NewAddressStore creates an instance of AddressStore
func NewAddressStore(db DB) *AddressStore {
	a := new(AddressStore)
	a.db = db
	return a
}

// SaveAll accepts addresses to save
func (as *AddressStore) SaveAll(addresses []string) error {
	key := make([][]byte, len(addresses))
	value := make([][]byte, len(addresses))

	for _, addr := range addresses {
		key = append(key, makeKey(addr))
		value = append(value, []byte(addr))
	}

	return as.db.WriteBatch(key, value)
}

// GetAll returns all saved addresses
func (as *AddressStore) GetAll() (addresses []string, err error) {
	_, values := as.db.GetByPrefix([]byte(PrefixAddress))
	for _, addr := range values {
		addresses = append(addresses, string(addr))
	}
	return
}

// ClearAll deletes all addresses
func (as *AddressStore) ClearAll() error {
	return as.db.DeleteByPrefix([]byte(PrefixAddress))
}
