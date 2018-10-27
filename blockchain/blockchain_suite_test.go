package blockchain

import (
	"math/big"

	"github.com/ellcrys/elld/params"
	"github.com/ellcrys/elld/util/logger"
	"github.com/shopspring/decimal"
)

var log logger.Logger

func init() {
	log = logger.NewLogrusNoOp()
	params.FeePerByte = decimal.NewFromFloat(0.01)
	params.MinimumDifficulty = new(big.Int).SetInt64(100000)
}
