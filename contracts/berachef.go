package befi

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/TropicalDog17/befi/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
type WhitelistedVault = types.WhitelistedVault

// BeraChef is the main client for interacting with the BeraChef contract
type BeraChef struct {
	client  *ethclient.Client
	address common.Address
	*BeraChefCaller
}

// BeraChefCaller contains read-only methods
type BeraChefCaller struct {
	contract *bind.BoundContract
}

// NewBeraChef creates a new client instance for the BeraChef contract
func NewBeraChef(address common.Address, client *ethclient.Client) (*BeraChef, error) {
	contractABI, err := loadContractABI()
	if err != nil {
		return nil, fmt.Errorf("failed to load ABI: %v", err)
	}

	contract := bind.NewBoundContract(address, contractABI, client, client, client)

	return &BeraChef{
		client:         client,
		address:        address,
		BeraChefCaller: &BeraChefCaller{contract: contract},
	}, nil
}

// NewBeraChefFromURL creates a new client instance for the BeraChef contract from an RPC URL
func NewBeraChefFromURL(rpcURL, contractAddress string) (*BeraChef, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	return NewBeraChef(common.HexToAddress(contractAddress), client)
}

// loadContractABI loads the contract ABI from file or embedded JSON
func loadContractABI() (abi.ABI, error) {
	// Try loading from file first
	abiFile := "abi/BeraChef.json"
	if _, err := os.Stat(abiFile); err == nil {
		jsonFile, err := os.ReadFile(abiFile)
		if err == nil {
			return abi.JSON(strings.NewReader(string(jsonFile)))
		}
	}

	return abi.ABI{}, fmt.Errorf("failed to load ABI from file")
}

// IsWhitelistedVault checks if an address is a whitelisted vault
func (c *BeraChefCaller) IsWhitelistedVault(opts *bind.CallOpts, vaultAddress common.Address) (bool, error) {
	var out []any
	err := c.contract.Call(opts, &out, "isWhitelistedVault", vaultAddress)
	if err != nil {
		return false, fmt.Errorf("contract call failed: %v", err)
	}

	return *abi.ConvertType(out[0], new(bool)).(*bool), nil
}

// GetDefaultRewardAllocation retrieves the default reward allocation
func (c *BeraChefCaller) GetDefaultRewardAllocation(opts *bind.CallOpts) (*RewardAllocation, error) {
	var out []any
	err := c.contract.Call(opts, &out, "getDefaultRewardAllocation")
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %v", err)
	}

	// Process the returned data
	startBlock := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	rawWeights := *abi.ConvertType(out[1], new([]struct {
		Receiver            common.Address
		PercentageNumerator *big.Int
	})).(*[]struct {
		Receiver            common.Address
		PercentageNumerator *big.Int
	})

	// Convert to our RewardAllocation type
	weights := make([]Weight, len(rawWeights))
	for i, w := range rawWeights {
		weights[i] = Weight{
			Receiver:            w.Receiver,
			PercentageNumerator: w.PercentageNumerator,
		}
	}

	return &RewardAllocation{
		StartBlock: startBlock,
		Weights:    weights,
	}, nil
}

// GetActiveRewardAllocation retrieves the active reward allocation for a validator
func (c *BeraChefCaller) GetActiveRewardAllocation(opts *bind.CallOpts, validatorPubKey []byte) (*RewardAllocation, error) {
	var out []any
	err := c.contract.Call(opts, &out, "getActiveRewardAllocation", validatorPubKey)
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %v", err)
	}

	// Process the returned data
	startBlock := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	rawWeights := *abi.ConvertType(out[1], new([]struct {
		Receiver            common.Address
		PercentageNumerator *big.Int
	})).(*[]struct {
		Receiver            common.Address
		PercentageNumerator *big.Int
	})

	// Convert to our RewardAllocation type
	weights := make([]Weight, len(rawWeights))
	for i, w := range rawWeights {
		weights[i] = Weight{
			Receiver:            w.Receiver,
			PercentageNumerator: w.PercentageNumerator,
		}
	}

	return &RewardAllocation{
		StartBlock: startBlock,
		Weights:    weights,
	}, nil
}

// GetValCommissionOnIncentiveTokens retrieves commission rate for a validator
func (c *BeraChefCaller) GetValCommissionOnIncentiveTokens(opts *bind.CallOpts, validatorPubKey []byte) (*big.Int, error) {
	var out []any
	err := c.contract.Call(opts, &out, "getValCommissionOnIncentiveTokens", validatorPubKey)
	if err != nil {
		return nil, fmt.Errorf("contract call failed: %v", err)
	}

	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// GetWhitelistedVaults retrieves all whitelisted vaults from a list of candidate addresses
func (c *BeraChef) GetWhitelistedVaults(candidateAddresses []common.Address) ([]WhitelistedVault, error) {
	whitelistedVaults := make([]WhitelistedVault, 0)
	opts := &bind.CallOpts{}

	for _, addr := range candidateAddresses {
		isWhitelisted, err := c.IsWhitelistedVault(opts, addr)
		if err != nil {
			return nil, fmt.Errorf("failed to check if vault is whitelisted: %v", err)
		}

		if isWhitelisted {
			whitelistedVaults = append(whitelistedVaults, WhitelistedVault{
				Address:  addr,
				Metadata: "", // Note: To get metadata, you would need to query events
			})
		}
	}

	return whitelistedVaults, nil
}

// // FindWhitelistedVaultsFromEvents retrieves whitelisted vaults by filtering events
// func (c *BeraChef) FindWhitelistedVaultsFromEvents(fromBlock, toBlock *big.Int) ([]WhitelistedVault, error) {
// 	// Create a filter query for VaultWhitelistedStatusUpdated events
// 	query := ethereum.FilterQuery{
// 		FromBlock: fromBlock,
// 		ToBlock:   toBlock,
// 		Addresses: []common.Address{c.address},
// 		Topics: [][]common.Hash{
// 			{crypto.Keccak256Hash([]byte("VaultWhitelistedStatusUpdated(address,bool,string)"))},
// 		},
// 	}

// 	logs, err := c.client.FilterLogs(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to filter logs: %v", err)
// 	}

// 	contractABI, err := loadContractABI()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to load ABI: %v", err)
// 	}

// 	vaults := make(map[common.Address]WhitelistedVault)
// 	for _, log := range logs {
// 		// Parse the log data
// 		event := struct {
// 			Receiver      common.Address
// 			IsWhitelisted bool
// 			Metadata      string
// 		}{}

// 		err = contractABI.UnpackIntoInterface(&event, "VaultWhitelistedStatusUpdated", log.Data)
// 		if err != nil {
// 			continue // Skip entries that cannot be parsed
// 		}

// 		// Extract receiver address from the indexed parameter (first topic)
// 		if len(log.Topics) >= 2 {
// 			receiver := common.BytesToAddress(log.Topics[1].Bytes())
// 			isWhitelisted := false
// 			if len(log.Topics) >= 3 {
// 				isWhitelisted = log.Topics[2].Big().Uint64() == 1
// 			}

// 			if isWhitelisted {
// 				vaults[receiver] = WhitelistedVault{
// 					Address:  receiver,
// 					Metadata: event.Metadata,
// 				}
// 			} else {
// 				delete(vaults, receiver)
// 			}
// 		}
// 	}

// 	// Convert map to slice
// 	result := make([]WhitelistedVault, 0, len(vaults))
// 	for _, vault := range vaults {
// 		result = append(result, vault)
// 	}

// 	return result, nil
// }
