package types_test

import (
	"encoding/json"
	"testing"

	"github.com/TropicalDog17/befi/internal/types"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalQuoteResponse(t *testing.T) {
	// Read from quote_resp.json
	// Unmarshal the JSON into a QuoteResponse struct
	// Check if the unmarshalled data matches the expected values

	quoteJson := `
	{
    "blockNumber": "3528856",
    "amount": "7167650998764773",
    "amountDecimals": "0.007167650998764773",
    "quote": "12124076429550338",
    "quoteDecimals": "0.012124076429550338",
    "quoteGasAdjusted": "12123990945502338",
    "quoteGasAdjustedDecimals": "0.012123990945502338",
    "gasUseEstimateQuote": "85484048000",
    "gasUseEstimateQuoteDecimals": "0.000000085484048",
    "gasUseEstimate": "113000",
    "gasUseEstimateUSD": "0",
    "gasPriceWei": "756496",
    "route": [
        [
            {
                "type": "v3-pool",
                "address": "0x12bf773F18cEC56F14e7cb91d82984eF5A3148EE",
                "tokenIn": {
                    "chainId": 80094,
                    "decimals": "18",
                    "address": "0xac03CABA51e17c86c921E1f6CBFBdC91F8BB2E6b",
                    "symbol": "iBGT"
                },
                "tokenOut": {
                    "chainId": 80094,
                    "decimals": "18",
                    "address": "0x6969696969696969696969696969696969696969",
                    "symbol": "WBERA"
                },
                "fee": "3000",
                "liquidity": "3923957943664711481868144",
                "sqrtRatioX96": "60826281702219052508389780701",
                "tickCurrent": "-5287",
                "amountIn": "7167650998764773",
                "amountOut": "12124076429550338"
            }
        ]
    ],
    "routeString": "[V3] 100.00% = iBGT -- 0.3% [0x12bf773F18cEC56F14e7cb91d82984eF5A3148EE] --> WBERA",
    "quoteId": "499d8"
}
	`

	// Unmarshal the JSON into a QuoteResponse struct
	var quoteResponse types.QuoteResponse
	err := json.Unmarshal([]byte(quoteJson), &quoteResponse)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if the unmarshalled data matches the expected values
	if quoteResponse.BlockNumber != "3528856" {
		t.Errorf("Expected block number 3528856, got %s", quoteResponse.BlockNumber)
	}

	require.True(t, quoteResponse.Route[0][0].AmountOut.Int64() > 0)

}
