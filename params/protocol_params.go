package params

import (
	"math/big"
	"time"
)

var (
	// MaximumExtraDataSize is the size of extra data a block can contain.
	MaximumExtraDataSize uint64 = 32

	// DifficultyBoundDivisor is the bound divisor of the difficulty,
	// used in the update calculations.
	DifficultyBoundDivisor = big.NewInt(2048)

	// DurationLimit is the decision boundary on the blocktime duration used to
	// determine whether difficulty should go up or not.
	DurationLimit = big.NewInt(13)

	// GenesisDifficulty is the difficulty of the Genesis block.
	GenesisDifficulty = big.NewInt(131072)

	// MinimumDifficulty is the minimum that the difficulty may ever be.
	MinimumDifficulty = big.NewInt(100000)
)

var (
	// MaxGetBlockHeader is the max number of block headers to request
	// from a remote peer per request.
	MaxGetBlockHeader = int64(500)

	// MaxGetBlockBodiesHashes is the max number of block bodies
	// to requests
	MaxGetBlockBodiesHashes = int64(2)

	// NumBlockBodiesRequesters is the number of workers that will
	// fetch block bodies
	NumBlockBodiesRequesters = 1
)

var (
	// Decimals is the number of coin decimal places
	Decimals = int32(18)
)

var (
	// AllowedFutureBlockTime is the number of seconds
	// a block's timestamp can have beyond the current timestamp
	AllowedFutureBlockTime = 15 * time.Second

	// MaxBlockTransactionsSize is the total size of
	// transactions that can fit in a block
	MaxBlockTransactionsSize = int64(10000000)
)