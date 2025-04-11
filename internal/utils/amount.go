package utils

import (
	"fmt"
	"math"
	"math/big"
)

func GetReadableAmount(amount *big.Int, decimals int) string {
	// Convert the amount to a float64
	amountFloat := new(big.Float).SetInt(amount)

	// Divide by 10^decimals
	decimalDivisor := new(big.Float).SetFloat64(math.Pow(10, float64(decimals)))
	amountFloat.Quo(amountFloat, decimalDivisor)

	// Format the amount with 18 decimal places
	return fmt.Sprintf("%.18f", amountFloat)
}
