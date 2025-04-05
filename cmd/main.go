package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TropicalDog17/befi"
	"github.com/TropicalDog17/befi/cache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Global variables to track resources for graceful shutdown
var (
	mainVaultCache        *cache.VaultCache
	mainWhitelistedWorker *befi.WhitelistedVaultWorker
)

// setupGracefulShutdown sets up signal handling for graceful shutdown
func setupGracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received interrupt signal, performing graceful shutdown...")

		// Stop the worker
		if mainWhitelistedWorker != nil {
			mainWhitelistedWorker.Stop()
		}

		// Save the cache on interrupt
		if mainVaultCache != nil {
			if err := mainVaultCache.SaveToDisk(); err != nil {
				log.Printf("Error saving cache before exit: %v", err)
			} else {
				log.Printf("Successfully saved cache before exit")
			}
		}

		// Save the whitelisted vaults to a file
		if mainWhitelistedWorker != nil {
			whitelistedVaults := mainWhitelistedWorker.GetWhitelistedVaults()
			saveWhitelistedVaultsToFile(whitelistedVaults, "whitelisted_vaults.json")
		}

		os.Exit(0)
	}()
}

// saveWhitelistedVaultsToFile saves the current whitelisted vaults to a JSON file
func saveWhitelistedVaultsToFile(vaults []befi.WhitelistedVault, filename string) {
	// Convert to a more JSON-friendly format
	type jsonVault struct {
		Address  string `json:"address"`
		Metadata string `json:"metadata"`
	}

	jsonVaults := make([]jsonVault, len(vaults))
	for i, vault := range vaults {
		jsonVaults[i] = jsonVault{
			Address:  vault.Address.Hex(),
			Metadata: vault.Metadata,
		}
	}

	// Marshal to JSON
	data, err := json.MarshalIndent(jsonVaults, "", "  ")
	if err != nil {
		log.Printf("Error marshaling whitelisted vaults: %v", err)
		return
	}

	// Write to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Printf("Error writing whitelisted vaults to file: %v", err)
		return
	}

	log.Printf("Successfully saved %d whitelisted vaults to %s", len(vaults), filename)
}

func main() {
	// Set up signal handling for graceful shutdown
	setupGracefulShutdown()

	// Connect to an Ethereum node
	rpcURL := "https://berachain-rpc.publicnode.com"
	log.Printf("Connecting to %s...", rpcURL)
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Printf("Connected to Ethereum node")

	// Set the contract addresses
	factoryAddress := common.HexToAddress("0x94Ad6Ac84f6C6FbA8b8CCbD71d9f4f101def52a8")
	beraChefAddress := common.HexToAddress("0xYourBeraChefContractAddress") // Replace with actual address

	// Create contract instances
	log.Printf("Creating contract instances...")
	factory, err := befi.NewVaultFactory(factoryAddress, client)
	if err != nil {
		log.Fatalf("Failed to create vault factory instance: %v", err)
	}

	beraChef, err := befi.NewBeraChef(beraChefAddress, client)
	if err != nil {
		log.Fatalf("Failed to create BeraChef instance: %v", err)
	}
	log.Printf("Contract instances created successfully")

	// Network identifier (for caching purposes)
	networkID := "berachain"

	// Create a cache instance with persistence
	log.Printf("Initializing vault cache...")
	vaultCache := cache.NewVaultCache(networkID, factoryAddress)
	mainVaultCache = vaultCache // Set global for shutdown handler

	// Load existing cache from disk
	if err := vaultCache.LoadFromDisk(); err != nil {
		log.Printf("Warning: Could not load cache from disk: %v", err)
	} else {
		log.Printf("Vault cache loaded successfully")
	}

	// Create and start the whitelisted vault worker
	log.Printf("Starting WhitelistedVaultWorker...")
	worker := befi.NewWhitelistedVaultWorker(
		client,
		beraChef,
		factory,
		vaultCache,
		5*time.Minute, // Update every 5 minutes
	)
	mainWhitelistedWorker = worker // Set global for shutdown handler

	// Start the worker
	worker.Start()

	// Force initial update and get the whitelisted vaults
	log.Printf("Performing initial whitelisted vault check...")
	if err := worker.ForceUpdate(); err != nil {
		log.Printf("Warning: Initial update failed: %v", err)
	}

	// Get and display initial whitelisted vaults
	whitelistedVaults := worker.GetWhitelistedVaults()
	log.Printf("Found %d whitelisted vaults initially", len(whitelistedVaults))

	// Print details about each whitelisted vault
	fmt.Println("Initial Whitelisted Vaults:")
	for i, vault := range whitelistedVaults {
		fmt.Printf("%d. Address: %s\n   Metadata: %s\n",
			i+1, vault.Address.Hex(), vault.Metadata)
	}

	// Start a simple monitor to periodically report status
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for {
			<-ticker.C
			vaults := worker.GetWhitelistedVaults()
			log.Printf("Status update: %d whitelisted vaults currently known", len(vaults))
		}
	}()

	log.Printf("Application running. Press Ctrl+C to exit.")

	// Block indefinitely (until signal handler executes)
	select {}
}
