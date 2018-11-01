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
	DurationLimit = big.NewInt(60)

	// DifficultyEpoch is the number of block between difficulty calculation
	DifficultyEpoch = uint64(1000)

	// GenesisDifficulty is the difficulty of the Genesis block.
	GenesisDifficulty = big.NewInt(10000000)

	// MinimumDifficulty is the minimum that the difficulty may ever be.
	MinimumDifficulty = big.NewInt(100000)

	// MinimumDurationIncrease is the minimum percent increase
	// a block's time can be when compared to its parent's
	MinimumDurationIncrease = big.NewFloat(2)
)

var (
	// MaxGetBlockHashes is the max number of block headers to request
	// from a remote peer per request.
	MaxGetBlockHashes = int64(500)

	// MaxGetBlockBodiesHashes is the max number of block bodies
	// to requests
	MaxGetBlockBodiesHashes = int64(10)

	// NumBlockBodiesRequesters is the number of workers that will
	// fetch block bodies
	NumBlockBodiesRequesters = 10
)

var (
	// Decimals is the number of coin decimal places
	Decimals = int32(18)
)

var (
	// AllowedFutureBlockTime is the number of seconds
	// a block's timestamp can have beyond the current timestamp
	AllowedFutureBlockTime = 15 * time.Second

	// MaxBlockNonTxsSize is the maximum size
	// of the non-transactional data a block
	// can have (headers, signature, hash).
	MaxBlockNonTxsSize = int64(1024)

	// MaxBlockTxsSize is the maximum size of
	// transactions that can fit in a block
	MaxBlockTxsSize = int64(10000000)

	// PoolCapacity is the max. number of transaction
	// that can be added to the transaction pool at
	// any given time.
	PoolCapacity = int64(10000)
)
