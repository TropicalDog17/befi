package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/TropicalDog17/befi/bindings"
	"github.com/TropicalDog17/befi/internal/constants"
	"github.com/TropicalDog17/befi/internal/types"
	"github.com/TropicalDog17/befi/internal/wrapper"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type SwapService struct {
	client  *wrapper.EthClient
	router  *bindings.KodiakSwapRouter
	keypair *wrapper.KeyPair
}

func NewSwapService(client *wrapper.EthClient, router *bindings.KodiakSwapRouter, keypair *wrapper.KeyPair) *SwapService {
	return &SwapService{
		client:  client,
		router:  router,
		keypair: keypair,
	}
}

// SwapIbgtToLp performs a token swap on the Berachain blockchain.
// It swaps all IBGT tokens of sender for a token using the provided parameters.
func (s *SwapService) SwapIbgtToToken(
	tokenOut common.Address,
	slippageTolerance float64, // e.g.,
) (*ethtypes.Transaction, error) {
	auth, err := s.client.GetAuth(context.Background(), s.keypair.PrivateKey, constants.ChainID, s.client.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction options: %v", err)
	}

	// Get the balance of IBGT tokens for the sender
	// extract this logic to a separate service
	ibgtErc20, err := bindings.NewERC20(constants.IBGTAddressMainnet, s.client.GetClient())
	if err != nil {
		return nil, fmt.Errorf("error when initiate erc20 binding", err)
	}

	userBalance, err := ibgtErc20.BalanceOf(nil, auth.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get user balance: %v", err)
	}
	// Get the amount of IBGT tokens to swap
	amountIn := new(big.Int).Set(userBalance)
	fmt.Println("Amount to swap:", amountIn.String())
	// Perform the swap using the router contract, retry if it fails
	var tx *ethtypes.Transaction
	backoff.Retry(
		func() error {
			tx, err = swapWithSlippage(
				s.router,
				auth,
				constants.IBGTAddressMainnet,
				tokenOut,
				amountIn,
				slippageTolerance,
			)
			if err != nil {
				return fmt.Errorf("swap failed: %v", err)
			}

			// Wait for the transaction to be mined
			receipt, err := bind.WaitMined(context.Background(), s.client.GetClient(), tx)
			if err != nil {
				return fmt.Errorf("failed to wait for transaction receipt: %v", err)
			}

			if receipt.Status != 1 {
				return fmt.Errorf("transaction failed: %v", receipt.Status)
			}

			return nil
		}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3),
	)
	// If the swap is successful, return the transaction
	if err != nil {
		return nil, fmt.Errorf("swap failed after retries: %v", err)
	}
	return tx, nil
}

func swapWithSlippage(
	router *bindings.KodiakSwapRouter,
	auth *bind.TransactOpts,
	tokenIn common.Address,
	tokenOut common.Address,
	amountIn *big.Int,
	slippageTolerance float64, // e.g., 0.01 for 1%
) (*ethtypes.Transaction, error) {
	// 1. Define the path for the swap
	// 2. Fetch the quote from the API
	quoteResp, err := GetQuote(types.GetQuoteParams{
		Protocols:       "v2,v3,mixed",
		TokenInAddress:  tokenIn.Hex(),
		TokenInChainId:  "80094",
		TokenOutAddress: tokenOut.Hex(),
		TokenOutChainId: "80094",
		Amount:          amountIn.String(),
		Type:            "exactIn",
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %v", err)
	}
	expectedOutput := new(big.Int).Set(quoteResp.Route[0][0].AmountOut.Int) // Assuming the first route is the best one

	fmt.Printf("Expected output amount: %s\n", expectedOutput.String())
	// Check if the expected output is zero
	if expectedOutput.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("expected output amount is zero or negative: %s", expectedOutput.String())
	}
	// Calculate the minimum output amount based on slippage tolerance
	// Convert slippage percentage to a big.Float representation
	slippageFloat := big.NewFloat(slippageTolerance)

	// Create a big.Float for the expected output
	expectedOutputFloat := new(big.Float).SetInt(expectedOutput)

	// Calculate slippage amount
	slippageAmount := new(big.Float).Mul(expectedOutputFloat, slippageFloat)

	// Subtract slippage amount from expected output
	minOutputFloat := new(big.Float).Sub(expectedOutputFloat, slippageAmount)

	// Convert back to big.Int (truncating any fractional component)
	minOutput := new(big.Int)
	minOutputFloat.Int(minOutput)
	// Ensure minOutput is not zero
	if minOutput.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("minimum output amount is zero or negative: %s", minOutput.String())
	}
	fmt.Printf("Minimum output amount: %s\n", minOutput.String())

	// 3. Approve the router to spend the tokenIn amount
	// This step is necessary to allow the router to transfer tokens on behalf of the user

	// tokenInContract, err := bindings.NewToken(tokenIn, auth.Client)
	// if err != nil {
	// 	return fmt.Errorf("failed to create token contract instance: %v", err)
	// }

	// tx, err := tokenInContract.Approve(auth, routerAddress, amountIn)
	// if err != nil {
	// 	return fmt.Errorf("failed to approve token transfer: %v", err)
	// }

	// fmt.Printf("Approval transaction sent: %s\n", tx.Hash().Hex())
	// 4. Execute the swap using SwapExactTokensForTokens
	tx, err := router.ExactInputSingle(
		auth,
		bindings.IV3SwapRouterExactInputSingleParams{
			TokenIn:           tokenIn,
			TokenOut:          tokenOut,
			Fee:               big.NewInt(3000), // 0.3% fee
			Recipient:         auth.From,        // recipient
			AmountIn:          amountIn,
			AmountOutMinimum:  minOutput,
			SqrtPriceLimitX96: big.NewInt(0),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("swap failed: %v", err)
	}

	fmt.Printf("Swap transaction sent: %s\n", tx.Hash().Hex())
	return tx, nil
}

// GetQuote fetches a quote for swapping tokens
func GetQuote(params types.GetQuoteParams) (*types.QuoteResponse, error) {
	url := fmt.Sprintf("https://ebey72gfe6.execute-api.us-east-1.amazonaws.com/prod/quote?protocols=%s&tokenInAddress=%s&tokenInChainId=%s&tokenOutAddress=%s&tokenOutChainId=%s&amount=%s&type=%s",
		params.Protocols,
		params.TokenInAddress,
		params.TokenInChainId,
		params.TokenOutAddress,
		params.TokenOutChainId,
		params.Amount,
		params.Type)

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("origin", "https://app.kodiak.finance")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://app.kodiak.finance/")
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	// Create HTTP client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	// Parse JSON response
	var quoteResp types.QuoteResponse
	if err := json.Unmarshal(body, &quoteResp); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w - Raw: %s", err, string(body))
	}

	return &quoteResp, nil
}
