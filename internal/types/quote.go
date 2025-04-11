package types

import (
	"encoding/json"
	"math/big"
	"strconv"
)

// BigInt is a wrapper around big.Int that implements json unmarshaling
type BigInt struct {
	*big.Int
}

// NewBigInt creates a new BigInt from a string
func NewBigInt(s string) (*BigInt, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return nil, &strconv.NumError{Func: "NewBigInt", Num: s, Err: strconv.ErrSyntax}
	}
	return &BigInt{n}, nil
}

// UnmarshalJSON custom unmarshaler for BigInt
func (b *BigInt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return &strconv.NumError{Func: "UnmarshalJSON(BigInt)", Num: s, Err: strconv.ErrSyntax}
	}

	b.Int = n
	return nil
}

// MarshalJSON custom marshaler for BigInt
func (b *BigInt) MarshalJSON() ([]byte, error) {
	if b == nil || b.Int == nil {
		return []byte("null"), nil
	}
	return json.Marshal(b.String())
}

// Token represents a cryptocurrency token
type Token struct {
	ChainID  int64  `json:"chainId"`
	Decimals string `json:"decimals"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
}

// V3Pool represents a V3 liquidity pool
type V3Pool struct {
	Type         string  `json:"type"`
	Address      string  `json:"address"`
	TokenIn      Token   `json:"tokenIn"`
	TokenOut     Token   `json:"tokenOut"`
	Fee          string  `json:"fee"`
	Liquidity    *BigInt `json:"liquidity"`
	SqrtRatioX96 *BigInt `json:"sqrtRatioX96"`
	TickCurrent  string  `json:"tickCurrent"`
	AmountIn     *BigInt `json:"amountIn"`
	AmountOut    *BigInt `json:"amountOut"`
}

// Route represents a swap route (array of arrays of pools)
type Route [][]V3Pool

// QuoteResponse represents the API response
type QuoteResponse struct {
	BlockNumber                 string  `json:"blockNumber"`
	Amount                      *BigInt `json:"amount"`
	AmountDecimals              string  `json:"amountDecimals"`
	Quote                       *BigInt `json:"quote"`
	QuoteDecimals               string  `json:"quoteDecimals"`
	QuoteGasAdjusted            *BigInt `json:"quoteGasAdjusted"`
	QuoteGasAdjustedDecimals    string  `json:"quoteGasAdjustedDecimals"`
	GasUseEstimateQuote         *BigInt `json:"gasUseEstimateQuote"`
	GasUseEstimateQuoteDecimals string  `json:"gasUseEstimateQuoteDecimals"`
	GasUseEstimate              string  `json:"gasUseEstimate"`
	GasUseEstimateUSD           string  `json:"gasUseEstimateUSD"`
	GasPriceWei                 string  `json:"gasPriceWei"`
	Route                       Route   `json:"route"`
	RouteString                 string  `json:"routeString"`
	QuoteId                     string  `json:"quoteId"`
}

// Raw quote response for unmarshaling
type rawQuoteResponse struct {
	BlockNumber                 string          `json:"blockNumber"`
	Amount                      string          `json:"amount"`
	AmountDecimals              string          `json:"amountDecimals"`
	Quote                       string          `json:"quote"`
	QuoteDecimals               string          `json:"quoteDecimals"`
	QuoteGasAdjusted            string          `json:"quoteGasAdjusted"`
	QuoteGasAdjustedDecimals    string          `json:"quoteGasAdjustedDecimals"`
	GasUseEstimateQuote         string          `json:"gasUseEstimateQuote"`
	GasUseEstimateQuoteDecimals string          `json:"gasUseEstimateQuoteDecimals"`
	GasUseEstimate              string          `json:"gasUseEstimate"`
	GasUseEstimateUSD           string          `json:"gasUseEstimateUSD"`
	GasPriceWei                 string          `json:"gasPriceWei"`
	Route                       json.RawMessage `json:"route"`
	RouteString                 string          `json:"routeString"`
	QuoteId                     string          `json:"quoteId"`
}

// Raw V3Pool for unmarshaling
type rawV3Pool struct {
	Type         string          `json:"type"`
	Address      string          `json:"address"`
	TokenIn      json.RawMessage `json:"tokenIn"`
	TokenOut     json.RawMessage `json:"tokenOut"`
	Fee          string          `json:"fee"`
	Liquidity    string          `json:"liquidity"`
	SqrtRatioX96 string          `json:"sqrtRatioX96"`
	TickCurrent  string          `json:"tickCurrent"`
	AmountIn     string          `json:"amountIn"`
	AmountOut    string          `json:"amountOut"`
}

// UnmarshalJSON for QuoteResponse
func (q *QuoteResponse) UnmarshalJSON(data []byte) error {
	var raw rawQuoteResponse
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Parse the basic string fields
	q.BlockNumber = raw.BlockNumber
	q.AmountDecimals = raw.AmountDecimals
	q.QuoteDecimals = raw.QuoteDecimals
	q.QuoteGasAdjustedDecimals = raw.QuoteGasAdjustedDecimals
	q.GasUseEstimateQuoteDecimals = raw.GasUseEstimateQuoteDecimals
	q.GasUseEstimate = raw.GasUseEstimate
	q.GasUseEstimateUSD = raw.GasUseEstimateUSD
	q.GasPriceWei = raw.GasPriceWei
	q.RouteString = raw.RouteString
	q.QuoteId = raw.QuoteId

	// Parse the BigInt fields
	amount, err := NewBigInt(raw.Amount)
	if err != nil {
		return err
	}
	q.Amount = amount

	quote, err := NewBigInt(raw.Quote)
	if err != nil {
		return err
	}
	q.Quote = quote

	quoteGasAdjusted, err := NewBigInt(raw.QuoteGasAdjusted)
	if err != nil {
		return err
	}
	q.QuoteGasAdjusted = quoteGasAdjusted

	gasUseEstimateQuote, err := NewBigInt(raw.GasUseEstimateQuote)
	if err != nil {
		return err
	}
	q.GasUseEstimateQuote = gasUseEstimateQuote

	// Parse the route (array of arrays of pools)
	var rawRoutes [][]json.RawMessage
	if err := json.Unmarshal(raw.Route, &rawRoutes); err != nil {
		return err
	}

	routes := make([][]V3Pool, len(rawRoutes))
	for i, rawRoute := range rawRoutes {
		route := make([]V3Pool, len(rawRoute))
		for j, rawPoolJson := range rawRoute {
			var rawPool rawV3Pool
			if err := json.Unmarshal(rawPoolJson, &rawPool); err != nil {
				return err
			}

			var tokenIn, tokenOut Token
			if err := json.Unmarshal(rawPool.TokenIn, &tokenIn); err != nil {
				return err
			}
			if err := json.Unmarshal(rawPool.TokenOut, &tokenOut); err != nil {
				return err
			}

			liquidity, err := NewBigInt(rawPool.Liquidity)
			if err != nil {
				return err
			}

			sqrtRatioX96, err := NewBigInt(rawPool.SqrtRatioX96)
			if err != nil {
				return err
			}

			amountIn, err := NewBigInt(rawPool.AmountIn)
			if err != nil {
				return err
			}

			amountOut, err := NewBigInt(rawPool.AmountOut)
			if err != nil {
				return err
			}

			route[j] = V3Pool{
				Type:         rawPool.Type,
				Address:      rawPool.Address,
				TokenIn:      tokenIn,
				TokenOut:     tokenOut,
				Fee:          rawPool.Fee,
				Liquidity:    liquidity,
				SqrtRatioX96: sqrtRatioX96,
				TickCurrent:  rawPool.TickCurrent,
				AmountIn:     amountIn,
				AmountOut:    amountOut,
			}
		}
		routes[i] = route
	}
	q.Route = routes

	return nil
}

// GetQuoteParams holds parameters for the quote request
type GetQuoteParams struct {
	Protocols       string
	TokenInAddress  string
	TokenInChainId  string
	TokenOutAddress string
	TokenOutChainId string
	Amount          string
	Type            string // "exactIn" or "exactOut"
}
