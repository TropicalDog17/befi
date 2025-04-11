package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/TropicalDog17/befi/bindings"
	"github.com/TropicalDog17/befi/internal/constants"
	"github.com/TropicalDog17/befi/internal/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/cenkalti/backoff/v4"
)

type InfraredService struct {
	// Infrared service dependencies
	Infrared       *bindings.Infrared
	InfraredVaults map[common.Address]*bindings.InfraredVault
	Multicall3     *bindings.Multicall3
	ChainID        *big.Int
	EthClient      *wrapper.EthClient
	Keypair        *wrapper.KeyPair
}

func NewInfraredService(infrared *bindings.Infrared, infraredVaultAddresses []common.Address, multicall3 *bindings.Multicall3, backend *wrapper.EthClient, keypair *wrapper.KeyPair) *InfraredService {
	// Create a map to hold the InfraredVault instances
	infraredVaults := make(map[common.Address]*bindings.InfraredVault)
	// Iterate through the provided vault addresses and create InfraredVault instances
	for _, vaultAddress := range infraredVaultAddresses {
		vault, err := bindings.NewInfraredVault(vaultAddress, backend.GetClient())
		if err != nil {
			log.Fatalf("Failed to create InfraredVault contract instance: %v", err)
		}
		infraredVaults[vaultAddress] = vault
	}
	return &InfraredService{
		Infrared:       infrared,
		InfraredVaults: infraredVaults,
		Multicall3:     multicall3,
		ChainID:        big.NewInt(80094),
		EthClient:      backend,
		Keypair:        keypair,
	}
}

// GetAccumulatedReward returns the current accumulated rewards of a given address for a specific vault
func (s *InfraredService) GetAccumulatedReward(vaultAddress *common.Address, userAddress *common.Address) (amount *big.Int, err error) {
	infraredVault, ok := s.InfraredVaults[*vaultAddress]
	if !ok {
		return nil, fmt.Errorf("vault not found: %s", vaultAddress.Hex())
	}
	amount, err = infraredVault.Earned(nil, *userAddress, constants.IBGTAddressMainnet)
	fmt.Println("Accumulated reward for vault", vaultAddress.Hex(), ":", amount.String())
	if err != nil {
		return nil, err
	}

	return amount, nil
}

// GetAccumulatedRewards returns the total accumulated rewards across all vaults for a given user
func (s *InfraredService) GetAccumulatedRewards(userAddress *common.Address) (amount *big.Int, err error) {
	// Initialize the total amount to zero
	totalAmount := big.NewInt(0)

	// Iterate through the vault addresses and sum the rewards
	for vaultAddress := range s.InfraredVaults {
		vaultAddr := vaultAddress // Create a copy to avoid issues with referencing loop variable
		vaultAmount, err := s.GetAccumulatedReward(&vaultAddr, userAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get accumulated reward for vault %s: %w", vaultAddr.Hex(), err)
		}
		totalAmount.Add(totalAmount, vaultAmount)
	}

	return totalAmount, nil
}

func (s *InfraredService) GetVaultAddresses() []common.Address {
	vaultAddresses := make([]common.Address, 0, len(s.InfraredVaults))
	for vaultAddress := range s.InfraredVaults {
		vaultAddresses = append(vaultAddresses, vaultAddress)
	}
	return vaultAddresses
}

func (s *InfraredService) ClaimRewards() (amount *big.Int, err error) {
	// Get the list of vault addresses
	vaultAddresses := s.GetVaultAddresses()
	calldatas := make([]bindings.Multicall3Call3, len(vaultAddresses))

	totalAmount, err := s.GetAccumulatedRewards(&constants.DefaultUserAddress)
	fmt.Println("Total amount before claiming:", totalAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to get accumulated rewards: %w", err)
	}

	// Iterate through the vault addresses and prepare the calldata for each vault
	for i, vaultAddress := range vaultAddresses {
		_, ok := s.InfraredVaults[vaultAddress]
		if !ok {
			return nil, fmt.Errorf("vault not found: %s", vaultAddress.Hex())
		}

		// Get the ABI of the InfraredVault contract
		infraredVaultABI, err := bindings.InfraredVaultMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("failed to get InfraredVault ABI: %w", err)
		}

		// Pack the function call data
		calldata, err := infraredVaultABI.Pack("getRewardForUser", constants.DefaultUserAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack getRewardForUser function call: %w", err)
		}

		calldatas[i] = bindings.Multicall3Call3{
			Target:       vaultAddress,
			CallData:     calldata,
			AllowFailure: false,
		}
	}

	// Create a signer using the private key
	auth, err := s.EthClient.GetAuth(context.Background(), s.Keypair.PrivateKey, s.ChainID, s.EthClient.GetClient())
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	// Call the Multicall3 contract to execute the getReward function on all vaults
	tx, err := s.Multicall3.Aggregate3(auth, calldatas)

	if err != nil {
		return nil, fmt.Errorf("failed to execute multicall: %w", err)
	}

	// Wait for the transaction to be mined with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), constants.TransactionTimeout*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, (s.EthClient).GetClient(), tx)
	if err != nil {
		return nil, fmt.Errorf("failed waiting for transaction to be mined (timeout: %v): %w", constants.TransactionTimeout, err)
	}

	if receipt.Status == 0 {
		return nil, fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}

	return totalAmount, nil
}

func (s *InfraredService) ClaimRewardsWithRetry() (amount *big.Int, err error) {
	// Retry the claim rewards operation with exponential backoff
	maxRetries := 5
	retryDelay := time.Second
	err = backoff.Retry(func() error {
		amount, err = s.ClaimRewards()
		if err != nil {
			return fmt.Errorf("failed to claim rewards: %w", err)
		}
		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(backoff.WithInitialInterval(retryDelay)), uint64(maxRetries)))
	if err != nil {
		return nil, fmt.Errorf("failed to claim rewards after %d attempts: %w", maxRetries, err)
	}

	return amount, nil
}
