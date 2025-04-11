package utils

import (
	"fmt"
	"math/big"
)

// CalculateBufferedGasPrice calculates the gas price with a multiplier applied.
func CalculateBufferedGasPrice(gasPrice *big.Int, multiplier float64) (*big.Int, error) {
	if gasPrice == nil {
		return nil, fmt.Errorf("gas price cannot be nil")
	}

	if multiplier < 1.0 {
		multiplier = 1.0 // Ensure a valid multiplier
	}

	// Apply the multiplier
	bufferedGasPrice := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), big.NewFloat(multiplier))
	finalGasPrice := new(big.Int)
	bufferedGasPrice.Int(finalGasPrice)

	return finalGasPrice, nil
}
