package profit

import (
	"math/big"

	"github.com/TropicalDog17/befi/internal/types"
)

func calculateProfit(initialPosition types.AmmPosition, finalPosition types.AmmPosition, accumulatedReward types.TokenWithAmount) (*big.Float, error) {
	// Calculate profit based on the initial and final positions
	// This is a placeholder implementation and should be replaced with actual logic
	impermanentLossValue := new(big.Float).Sub(finalPosition.GetValueInUSD(), initialPosition.GetValueInUSD())

	accumulatedRewardValue := new(big.Float).SetInt(accumulatedReward.Amount).Mul(accumulatedReward.PriceInUSD, big.NewFloat(1e-18))

	entryFeeValue := new(big.Float).Mul(initialPosition.GetValueInUSD(), new(big.Float).SetInt(initialPosition.FeeRatio))

	exitFeeValue := new(big.Float).Mul(finalPosition.GetValueInUSD(), new(big.Float).SetInt(finalPosition.FeeRatio))

	// Calculate profit
	profit := new(big.Float).Sub(impermanentLossValue, entryFeeValue)
	profit = new(big.Float).Add(profit, exitFeeValue)
	profit = new(big.Float).Add(profit, accumulatedRewardValue)

	return profit, nil
}

func CalculateApr(initialPosition types.AmmPosition, finalPosition types.AmmPosition, accumulatedReward types.TokenWithAmount) (*big.Float, error) {
	// Calculate APR based on the initial and final positions
	// This is a placeholder implementation and should be replaced with actual logic
	profit, err := calculateProfit(initialPosition, finalPosition, accumulatedReward)
	if err != nil {
		return nil, err
	}

	// Calculate APR
	apr := new(big.Float).Quo(profit, initialPosition.GetValueInUSD())
	apr = new(big.Float).Mul(apr, big.NewFloat(100))

	return apr, nil
}
