// Package blakimoto provides functionalities for performing
// and verifying proof-of-work computation. It uses
// Blake2b-256 as the hash function.
package blakimoto

import (
	"encoding/binary"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/ellcrys/elld/util"
	"github.com/ellcrys/elld/util/logger"
)

var (
	// maxUint256 is a big integer representing 2^256-1
	maxUint256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
)

// Mode defines the type and amount of PoW verification an blakimoto engine makes.
type Mode uint

const (
	// ModeNormal refers to normal mode
	ModeNormal Mode = iota
	// ModeTest refers to test mode
	ModeTest
)

// Config are the configuration parameters of the blakimoto.
type Config struct {
	PowMode Mode
}

// Blakimoto is a consensus engine based on proof-of-work implementing the blakimoto
// algorithm.
type Blakimoto struct {
	config Config

	log logger.Logger

	// Mining related fields
	rand   *rand.Rand    // Properly seeded random source for nonces
	update chan struct{} // Notification channel to update mining parameters

	// The fields below are hooks for testing
	fakeDelay time.Duration // Time delay to sleep for before returning from verify

	lock sync.Mutex // Ensures thread safety for the in-memory caches and mining fields
}

// New creates a full sized blakimoto PoW scheme.
func New(config Config, log logger.Logger) *Blakimoto {
	return &Blakimoto{
		config: config,
		update: make(chan struct{}),
		log:    log,
	}
}

// ConfiguredBlakimoto creates an Blakimoto instance pre-configured
// using the engine configuration.
func ConfiguredBlakimoto(mode Mode, log logger.Logger) *Blakimoto {
	return New(Config{
		PowMode: mode,
	}, log)
}

// SetFakeDelay sets the delay duration for ModeFake
func (blakimoto *Blakimoto) SetFakeDelay(d time.Duration) {
	blakimoto.fakeDelay = d
}

// BlakeHash combines the header's hash and nonce
// and hashes the value using blake2b-256 to provide
// an output that is checked against a difficulty target
func BlakeHash(headerHash []byte, nonce uint64) []byte {

	// Combine header+nonce into a 64 byte seed
	seed := make([]byte, 40)
	copy(seed, headerHash)
	binary.LittleEndian.PutUint64(seed[32:], nonce)

	return util.Blake2b256(seed)
}
