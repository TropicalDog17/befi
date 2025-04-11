package main

import (
	"fmt"
	"math/big"

	"github.com/TropicalDog17/befi/bindings"
	"github.com/TropicalDog17/befi/internal/constants"
	"github.com/TropicalDog17/befi/internal/services"
	"github.com/TropicalDog17/befi/internal/wrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	keypair, err := wrapper.NewKeyPair()
	if err != nil {
		panic("Failed to create key pair")
	}
	client, err := wrapper.NewEthClient([]string{"https://rpc.berachain.com", "https://berachain-rpc.publicnode.com", "https://berachain.blockpi.network/v1/rpc/public", "https://rpc.berachain-apis.com"})
	if err != nil {
		panic("Failed to create Ethereum client")
	}

	innerClient := client.GetClient()

	infraredContract, err := bindings.NewInfrared(
		constants.InfaredCoreMainnet,
		innerClient,
	)
	if err != nil {
		panic("Failed to create Infrared contract instance")
	}

	multicall3Contract, err := bindings.NewMulticall3(
		constants.Multicall3AddressMainnet,
		innerClient)

	if err != nil {
		panic("Failed to create Multicall3 contract instance")
	}

	routerContract, err := bindings.NewKodiakSwapRouter(
		constants.KodiakSwapRouterAddressMainnet,
		innerClient,
	)
	if err != nil {
		panic("Failed to create KodiakSwapRouter contract instance")
	}

	infrared := services.NewInfraredService(infraredContract, []common.Address{constants.NAVBeraInfraredVaultMainnet, constants.BIXBTBeraInfraredVaultMainnet, constants.BMBeraInfraredVaultMainnet, constants.HOLDBeraInfraredVaultMainnet}, multicall3Contract, client, keypair)

	swap := services.NewSwapService(client, routerContract, keypair)

	totalAmount, err := infrared.ClaimRewardsWithRetry()
	if err != nil {
		fmt.Printf("Failed to claim rewards: %s", err)
		return
	}

	if totalAmount.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("No rewards to claim")
		return
	}
	fmt.Println("Total amount claimed:", totalAmount.String())

	// Try to swap the claimed rewards to WBera token
	tx, err := swap.SwapIbgtToToken(constants.BeraAddressMainnet, 0.05)
	if err != nil {
		panic("Failed to swap IBGT to LP: " + err.Error())
	}
	// Print the transaction hash
	println("Transaction hash:", tx.Hash().Hex())
}
