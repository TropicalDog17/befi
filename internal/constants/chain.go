package constants

import (
	"math/big"
	"time"
)

const (
	TransactionTimeout = 30              // seconds
	GasPriceMultiplier = 2.0             // Multiplier for gas price to ensure faster transactions
	RetryDelay         = 3 * time.Second // seconds
)

var (
	ChainID = big.NewInt(80094) // Berachain mainnet chain ID
)
