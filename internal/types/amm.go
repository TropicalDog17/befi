package types

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// AmmPosition represents an AMM position
type AmmPosition struct {
	TokenA   TokenWithAmount
	TokenB   TokenWithAmount
	Pool     common.Address
	FeeRatio *big.Int
}

// TokenWithAmount represents a token with its amount and price information
type TokenWithAmount struct {
	Token      common.Address
	Amount     *big.Int
	PriceInUSD *big.Float
	Decimals   uint8
}

// NewAmmPosition creates a new AMM position
func NewAmmPosition(tokenA, tokenB TokenWithAmount, pool common.Address) *AmmPosition {
	return &AmmPosition{
		TokenA: tokenA,
		TokenB: tokenB,
		Pool:   pool,
	}
}

// GetRatio returns the ratio of tokenA to tokenB
func (p *AmmPosition) GetRatio() *big.Float {
	if p.TokenB.Amount.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}

	tokenAFloat := new(big.Float).SetInt(p.TokenA.Amount)
	tokenBFloat := new(big.Float).SetInt(p.TokenB.Amount)

	// Adjust for decimals
	adjustmentA := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.TokenA.Decimals)), nil))
	adjustmentB := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.TokenB.Decimals)), nil))

	normalizedA := new(big.Float).Quo(tokenAFloat, adjustmentA)
	normalizedB := new(big.Float).Quo(tokenBFloat, adjustmentB)

	return new(big.Float).Quo(normalizedA, normalizedB)
}

// GetSpotPrice returns the current spot price of tokenA in terms of tokenB
func (p *AmmPosition) GetSpotPrice() *big.Float {
	return new(big.Float).Quo(p.TokenB.PriceInUSD, p.TokenA.PriceInUSD)
}

// GetValueInUSD returns the total value of the position in USD
func (p *AmmPosition) GetValueInUSD() *big.Float {
	valueA := p.GetTokenValueInUSD(p.TokenA)
	valueB := p.GetTokenValueInUSD(p.TokenB)

	return new(big.Float).Add(valueA, valueB)
}

// GetTokenValueInUSD returns the USD value of a token amount
func (p *AmmPosition) GetTokenValueInUSD(token TokenWithAmount) *big.Float {
	tokenFloat := new(big.Float).SetInt(token.Amount)
	adjustment := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token.Decimals)), nil))
	normalizedAmount := new(big.Float).Quo(tokenFloat, adjustment)

	return new(big.Float).Mul(normalizedAmount, token.PriceInUSD)
}

// GetImpermanentLoss calculates the impermanent loss compared to holding
// Takes the initial price ratio and current price ratio as parameters
func (p *AmmPosition) GetImpermanentLoss(initialPriceRatio, currentPriceRatio *big.Float) *big.Float {
	// Calculate price ratio change (k)
	priceRatioChange := new(big.Float).Quo(currentPriceRatio, initialPriceRatio)

	// Calculate sqrt(k)
	squareRoot := new(big.Float).Sqrt(priceRatioChange)

	// Calculate 2 * sqrt(k) / (1 + k)
	numerator := new(big.Float).Mul(big.NewFloat(2), squareRoot)
	denominator := new(big.Float).Add(big.NewFloat(1), priceRatioChange)
	fraction := new(big.Float).Quo(numerator, denominator)

	// IL = 1 - 2 * sqrt(k) / (1 + k)
	return new(big.Float).Sub(big.NewFloat(1), fraction)
}

// PrintPositionInfo prints detailed information about the AMM position
func (p *AmmPosition) PrintPositionInfo() {
	fmt.Printf("Pool Address: %s\n", p.Pool.Hex())
	fmt.Printf("Token A: %s, Amount: %s, Decimals: %d\n", p.TokenA.Token.Hex(), p.TokenA.Amount.String(), p.TokenA.Decimals)
	fmt.Printf("Token B: %s, Amount: %s, Decimals: %d\n", p.TokenB.Token.Hex(), p.TokenB.Amount.String(), p.TokenB.Decimals)

	ratio := p.GetRatio()
	fmt.Printf("Token Ratio (A/B): %s\n", ratio.Text('f', 8))

	spotPrice := p.GetSpotPrice()
	fmt.Printf("Spot Price (B per A): %s\n", spotPrice.Text('f', 8))

	valueUSD := p.GetValueInUSD()
	fmt.Printf("Total Position Value: $%s\n", valueUSD.Text('f', 2))
}

// CalculateOptimalAmounts calculates optimal amounts to deposit to match the current pool ratio
func (p *AmmPosition) CalculateOptimalAmounts(availableA, availableB *big.Int) (TokenWithAmount, TokenWithAmount) {
	poolRatio := p.GetRatio()

	availableAFloat := new(big.Float).SetInt(availableA)
	availableBFloat := new(big.Float).SetInt(availableB)

	// Normalize by decimals
	adjustmentA := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.TokenA.Decimals)), nil))
	adjustmentB := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(p.TokenB.Decimals)), nil))

	normalizedAvailableA := new(big.Float).Quo(availableAFloat, adjustmentA)
	normalizedAvailableB := new(big.Float).Quo(availableBFloat, adjustmentB)

	// Calculate how much B would be needed for all of A
	neededB := new(big.Float).Quo(normalizedAvailableA, poolRatio)

	// Calculate how much A would be needed for all of B
	neededA := new(big.Float).Mul(normalizedAvailableB, poolRatio)

	var optimalA, optimalB *big.Int

	if normalizedAvailableB.Cmp(neededB) >= 0 {
		// We have enough B, use all A
		optimalA = availableA

		// Calculate B needed (with decimals)
		bNeeded := new(big.Float).Mul(neededB, adjustmentB)
		optimalB = new(big.Int)
		bNeeded.Int(optimalB)
	} else {
		// We have enough A, use all B
		optimalB = availableB

		// Calculate A needed (with decimals)
		aNeeded := new(big.Float).Mul(neededA, adjustmentA)
		optimalA = new(big.Int)
		aNeeded.Int(optimalA)
	}

	return TokenWithAmount{
			Token:      p.TokenA.Token,
			Amount:     optimalA,
			PriceInUSD: p.TokenA.PriceInUSD,
			Decimals:   p.TokenA.Decimals,
		}, TokenWithAmount{
			Token:      p.TokenB.Token,
			Amount:     optimalB,
			PriceInUSD: p.TokenB.PriceInUSD,
			Decimals:   p.TokenB.Decimals,
		}
}

// CalculateTokensOutForTokensIn calculates expected tokens output for a swap
// based on the AMM's constant product formula
func (p *AmmPosition) CalculateTokensOutForTokensIn(tokenIn TokenWithAmount, isTokenA bool) *big.Int {
	var reserveIn, reserveOut *big.Int
	var decimalsIn, decimalsOut uint8

	if isTokenA {
		reserveIn = p.TokenA.Amount
		reserveOut = p.TokenB.Amount
		decimalsIn = p.TokenA.Decimals
		decimalsOut = p.TokenB.Decimals
	} else {
		reserveIn = p.TokenB.Amount
		reserveOut = p.TokenA.Amount
		decimalsIn = p.TokenB.Decimals
		decimalsOut = p.TokenA.Decimals
	}

	// Normalize token amounts to account for decimals if they differ
	if decimalsIn != decimalsOut {
		// We'll work with normalized (decimal-adjusted) values using big.Float for precision
		reserveInFloat := new(big.Float).SetInt(reserveIn)
		reserveOutFloat := new(big.Float).SetInt(reserveOut)
		amountInFloat := new(big.Float).SetInt(tokenIn.Amount)

		// Calculate decimal adjustment factors
		inAdjustment := new(big.Float).SetFloat64(math.Pow(10, float64(decimalsIn)))
		outAdjustment := new(big.Float).SetFloat64(math.Pow(10, float64(decimalsOut)))

		// Normalize to a common decimal base (e.g., 18 decimals)
		normalizedReserveIn := new(big.Float).Quo(reserveInFloat, inAdjustment)
		normalizedReserveOut := new(big.Float).Quo(reserveOutFloat, outAdjustment)
		normalizedAmountIn := new(big.Float).Quo(amountInFloat, inAdjustment)

		// Apply constant product formula with floating point math
		// amountOut = reserveOut * amountIn / (reserveIn + amountIn)
		numerator := new(big.Float).Mul(normalizedReserveOut, normalizedAmountIn)
		denominator := new(big.Float).Add(normalizedReserveIn, normalizedAmountIn)
		normalizedAmountOut := new(big.Float).Quo(numerator, denominator)

		// Apply fee (0.3%)
		fee := new(big.Float).Mul(normalizedAmountOut, new(big.Float).SetInt(p.FeeRatio)) // 0.3% fee
		normalizedAmountOut = new(big.Float).Sub(normalizedAmountOut, fee)

		// Convert back to token decimals
		actualAmountOut := new(big.Float).Mul(normalizedAmountOut, outAdjustment)

		// Convert to big.Int
		amountOut := new(big.Int)
		actualAmountOut.Int(amountOut)
		return amountOut
	}

	// If decimals are the same, use the more efficient integer math
	// Apply constant product formula: x * y = k
	// (reserveIn + amountIn) * (reserveOut - amountOut) = reserveIn * reserveOut
	// amountOut = reserveOut - (reserveIn * reserveOut) / (reserveIn + amountIn)

	// Calculate numerator and denominator with big.Int to avoid overflow
	numerator := new(big.Int).Mul(reserveIn, reserveOut)
	denominator := new(big.Int).Add(reserveIn, tokenIn.Amount)

	// Calculate intermediate result
	intermediate := new(big.Int).Div(numerator, denominator)

	// Calculate amountOut
	amountOut := new(big.Int).Sub(reserveOut, intermediate)

	// Apply fee (typically 0.3% for Uniswap V2 style pools)
	fee := new(big.Int).Div(amountOut, p.FeeRatio) // Approximates 0.3%
	amountOut = new(big.Int).Sub(amountOut, fee)

	return amountOut
}

// CalculateLiquidityShare calculates the share of the total pool the position represents
func (p *AmmPosition) CalculateLiquidityShare(totalSupply *big.Int, lpTokens *big.Int) *big.Float {
	if totalSupply.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}

	lpTokensFloat := new(big.Float).SetInt(lpTokens)
	totalSupplyFloat := new(big.Float).SetInt(totalSupply)

	share := new(big.Float).Quo(lpTokensFloat, totalSupplyFloat)
	return share
}
