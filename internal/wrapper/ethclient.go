package wrapper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/TropicalDog17/befi/internal/constants"
	"github.com/TropicalDog17/befi/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Wrapper around the ethereum client
type EthClient struct {
	// Ethereum client
	clients   []*ethclient.Client
	endpoints []string
	counter   int
	mu        sync.Mutex

	failureTracker map[int]time.Time // Tracks failed clients and their cooldown periods
	cooldown       time.Duration     // Cooldown period for retrying a failed client

}

// NewEthClient creates a new Ethereum client
func NewEthClient(rpcEndpoints []string) (*EthClient, error) {
	if len(rpcEndpoints) == 0 {
		return nil, fmt.Errorf("no RPC endpoints provided")
	}

	clients := make([]*ethclient.Client, len(rpcEndpoints))
	for i, endpoint := range rpcEndpoints {
		client, err := ethclient.Dial(endpoint)
		if err != nil {
			return nil, err
		}
		clients[i] = client
	}

	return &EthClient{
		clients:   clients,
		endpoints: rpcEndpoints,
		counter:   0,
		// failureTracker: make(map[int]time.Time),
		// cooldown:       constants.EthClientCooldown,
	}, nil
}

func (c *EthClient) GetClient() *ethclient.Client {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Track the index of the last client
	lastClientIndex := len(c.clients) - 1

	for {
		clientIndex := c.counter
		c.counter = (c.counter + 1) % len(c.clients)

		// Check if the client is in cooldown
		if cooldownEnd, failed := c.failureTracker[clientIndex]; failed {
			if time.Now().Before(cooldownEnd) && clientIndex != lastClientIndex {
				fmt.Printf("Skipping client %s due to recent failure (cooldown active)", c.endpoints[clientIndex])
				continue
			}
			// Remove from failure tracker if cooldown has expired
			delete(c.failureTracker, clientIndex)
		}

		// Return the client if not in cooldown or it's the last client
		return c.clients[clientIndex]
	}
}

func (c *EthClient) GetAuth(ctx context.Context,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
	client *ethclient.Client) (*bind.TransactOpts, error) {

	if privateKey == nil || chainID == nil || client == nil {
		return nil, fmt.Errorf("invalid parameters: privateKey, chainID, and client must not be nil")
	}

	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Fetch pending nonce
	pendingNonce, err := client.PendingNonceAt(ctx, fromAddress) // Highest pending nonce
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	bufferedGasPrice, err := utils.CalculateBufferedGasPrice(gasPrice, constants.GasPriceMultiplier)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate buffered gas price: %w", err)
	}

	// Create transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Set transaction options
	auth.Nonce = big.NewInt(int64(pendingNonce))
	auth.Value = big.NewInt(0) // 0 wei, since we're not sending Ether
	auth.GasPrice = bufferedGasPrice

	// logger.GetLogger().Infof("Using nonce %d for address %s", pendingNonce, fromAddress.Hex())
	return auth, nil

}

// executeWithRetry executes a function and retries with the next client on failure
func (c *EthClient) executeWithRetry(
	fn func(client *ethclient.Client) (any, error),
) (any, error) {
	var lastErr error
	baseDelay := constants.RetryDelay

	for range c.clients {
		client := c.GetClient()
		clientIndex := (c.counter - 1 + len(c.clients)) % len(c.clients) // Get the last used client index

		fmt.Printf("Using RPC endpoint: %s", c.endpoints[clientIndex])

		// Retry the current client up to 3 times before switching
		for attempt := 1; attempt <= 3; attempt++ {
			result, err := fn(client)
			if err == nil {
				return result, nil // Success
			}

			lastErr = err
			fmt.Printf("RPC call failed on endpoint %s (attempt %d/3): %v", c.endpoints[clientIndex], attempt, err)
			time.Sleep(baseDelay * (1 << (attempt - 1))) // Exponential backoff

			// If this was the last attempt for this client, mark it as failed
			if attempt == 3 {
				c.mu.Lock()
				c.failureTracker[clientIndex] = time.Now().Add(c.cooldown)
				c.mu.Unlock()
				fmt.Printf("Marking RPC endpoint %s as temporarily failed (cooldown active)", c.endpoints[clientIndex])
			}
		}
	}

	// Try the last client as a fallback
	lastClient := c.clients[len(c.clients)-1]
	lastClientIndex := len(c.clients) - 1
	fmt.Printf("Falling back to last RPC endpoint: %s", c.endpoints[lastClientIndex])

	for attempt := 1; attempt <= 3; attempt++ {
		result, err := fn(lastClient)
		if err == nil {
			return result, nil // Success
		}
		lastErr = err
		fmt.Printf("Fallback client failed on attempt %d/3: %v", attempt, err)
		time.Sleep(baseDelay * (1 << (attempt - 1))) // Exponential backoff
	}

	// Return the last encountered error if all retries fail
	return nil, fmt.Errorf("all RPC clients failed after retries, including fallback: %w", lastErr)
}

// SuggestGasTipCap retrieves the suggested gas tip cap for dynamic fee transactions using round-robin retry logic.
func (c *EthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	// Use executeWithRetry to simplify retry logic
	result, err := c.executeWithRetry(func(client *ethclient.Client) (any, error) {
		// Use the client's SuggestGasTipCap method to fetch the tip cap
		gasTipCap, err := client.SuggestGasTipCap(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas tip cap: %w", err)
		}
		return gasTipCap, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve suggested gas tip cap after retries: %w", err)
	}

	// Cast the result to *big.Int and return
	return result.(*big.Int), nil
}
