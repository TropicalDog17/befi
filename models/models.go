package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Weight represents a reward allocation weight
type Weight struct {
	Receiver            common.Address
	PercentageNumerator *big.Int
}

// RewardAllocation represents a reward allocation
type RewardAllocation struct {
	StartBlock *big.Int
	Weights    []Weight
}

// QueuedCommissionRateChange represents a queued commission rate change
type QueuedCommissionRateChange struct {
	BlockNumberLast *big.Int
	CommissionRate  *big.Int
}

// WhitelistedVault represents vault information
type WhitelistedVault struct {
	Address  common.Address
	Metadata string
}
