package services_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/TropicalDog17/befi/bindings"
	"github.com/TropicalDog17/befi/internal/constants"
	"github.com/TropicalDog17/befi/internal/services"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func GetMockInfraredService() *services.InfraredService {
	client, err := ethclient.Dial("https://berachain-rpc.publicnode.com")
	if err != nil {
		panic("Failed to connect to Ethereum client")
	}

	infraredContract, err := bindings.NewInfrared(
		constants.InfaredCoreMainnet,
		client,
	)
	if err != nil {
		panic("Failed to create Infrared contract instance")
	}

	infraredVaultContract, err := bindings.NewInfraredVault(
		constants.NAVBeraInfraredVaultMainnet,
		client,
	)
	if err != nil {
		panic("Failed to create InfraredVault contract instance")
	}

	// Mock Infrared service dependencies
	infrared := &services.InfraredService{
		Infrared: infraredContract, // Mock Infrared contract
		InfraredVaults: map[common.Address]*bindings.InfraredVault{
			constants.NAVBeraInfraredVaultMainnet: infraredVaultContract,
		},
		Multicall3: nil, // Mock Multicall3 contract
	}

	return infrared
}

func TestGetReward(t *testing.T) {
	infraredService := GetMockInfraredService()

	rewardAmount, err := infraredService.GetAccumulatedRewards(
		&constants.DefaultUserAddress,
	)

	require.NoError(t, err)
	require.NotNil(t, rewardAmount)
	require.Greater(t, rewardAmount.Cmp(big.NewInt(0)), 0, "Reward amount should be greater than 0")

	fmt.Printf("Reward amount	: %s iBGT\n", new(big.Float).Quo(new(big.Float).SetInt(rewardAmount), big.NewFloat(1e18)).String())
}
