package befi

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/TropicalDog17/befi/cache"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// WhitelistedVaultCache provides caching for whitelisted vault addresses
type WhitelistedVaultCache struct {
	vaults        map[common.Address]WhitelistedVault
	mu            sync.RWMutex
	lastUpdated   time.Time
	updateCounter int
}

// NewWhitelistedVaultCache creates a new cache for whitelisted vaults
func NewWhitelistedVaultCache() *WhitelistedVaultCache {
	return &WhitelistedVaultCache{
		vaults:        make(map[common.Address]WhitelistedVault),
		mu:            sync.RWMutex{},
		lastUpdated:   time.Time{},
		updateCounter: 0,
	}
}

// Get returns all whitelisted vaults from the cache
func (wvc *WhitelistedVaultCache) Get() []WhitelistedVault {
	wvc.mu.RLock()
	defer wvc.mu.RUnlock()

	result := make([]WhitelistedVault, 0, len(wvc.vaults))
	for _, vault := range wvc.vaults {
		result = append(result, vault)
	}
	return result
}

// Update updates the cache with the latest whitelisted vaults
func (wvc *WhitelistedVaultCache) Update(vaults []WhitelistedVault) {
	wvc.mu.Lock()
	defer wvc.mu.Unlock()

	// Clear the existing cache and update with new data
	wvc.vaults = make(map[common.Address]WhitelistedVault, len(vaults))
	for _, vault := range vaults {
		wvc.vaults[vault.Address] = vault
	}

	wvc.lastUpdated = time.Now()
	wvc.updateCounter++

	log.Printf("Updated whitelisted vaults cache with %d vaults (update #%d)",
		len(vaults), wvc.updateCounter)
}

// GetLastUpdated returns the time when the cache was last updated
func (wvc *WhitelistedVaultCache) GetLastUpdated() time.Time {
	wvc.mu.RLock()
	defer wvc.mu.RUnlock()
	return wvc.lastUpdated
}

// WhitelistedVaultWorker periodically fetches and updates whitelisted vaults
type WhitelistedVaultWorker struct {
	client           *ethclient.Client
	beraChef         *BeraChef
	vaultFactory     *VaultFactory
	vaultCache       *cache.VaultCache
	whitelistedCache *WhitelistedVaultCache
	updateInterval   time.Duration
	done             chan bool
	wg               sync.WaitGroup
}

// NewWhitelistedVaultWorker creates a new worker to update whitelisted vaults
func NewWhitelistedVaultWorker(
	client *ethclient.Client,
	beraChef *BeraChef,
	vaultFactory *VaultFactory,
	vaultCache *cache.VaultCache,
	updateInterval time.Duration,
) *WhitelistedVaultWorker {
	return &WhitelistedVaultWorker{
		client:           client,
		beraChef:         beraChef,
		vaultFactory:     vaultFactory,
		vaultCache:       vaultCache,
		whitelistedCache: NewWhitelistedVaultCache(),
		updateInterval:   updateInterval,
		done:             make(chan bool),
	}
}

// Start begins the worker's update loop
func (wvw *WhitelistedVaultWorker) Start() {
	wvw.wg.Add(1)
	go wvw.updateLoop()
	log.Printf("WhitelistedVaultWorker started with %s update interval", wvw.updateInterval)
}

// Stop signals the worker to stop and waits for it to finish
func (wvw *WhitelistedVaultWorker) Stop() {
	log.Println("Stopping WhitelistedVaultWorker...")
	wvw.done <- true
	wvw.wg.Wait()
	log.Println("WhitelistedVaultWorker stopped")
}

// GetWhitelistedVaults returns the currently cached whitelisted vaults
func (wvw *WhitelistedVaultWorker) GetWhitelistedVaults() []WhitelistedVault {
	return wvw.whitelistedCache.Get()
}

// ForceUpdate forces an immediate update of the whitelisted vaults
func (wvw *WhitelistedVaultWorker) ForceUpdate() error {
	return wvw.updateWhitelistedVaults()
}

// updateLoop is the main worker loop that periodically updates whitelisted vaults
func (wvw *WhitelistedVaultWorker) updateLoop() {
	defer wvw.wg.Done()

	// Perform initial update
	if err := wvw.updateWhitelistedVaults(); err != nil {
		log.Printf("Initial whitelisted vaults update failed: %v", err)
	}

	// Set up ticker for periodic updates
	ticker := time.NewTicker(wvw.updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := wvw.updateWhitelistedVaults(); err != nil {
				log.Printf("Whitelisted vaults update failed: %v", err)
			}
		case <-wvw.done:
			log.Println("Update loop stopping")
			return
		}
	}
}

// updateWhitelistedVaults fetches the latest whitelisted vaults from the chain
func (wvw *WhitelistedVaultWorker) updateWhitelistedVaults() error {
	// Fetch all vaults using the vault cache
	vaults, err := wvw.vaultCache.GetVaults(wvw.client, wvw.vaultFactory.Address())
	if err != nil {
		return fmt.Errorf("failed to get all vaults: %v", err)
	}

	log.Printf("Checking whitelisted status for %d vaults", len(vaults))

	// Check which vaults are whitelisted
	whitelistedVaults, err := wvw.beraChef.GetWhitelistedVaults(vaults)
	if err != nil {
		return fmt.Errorf("failed to get whitelisted vaults: %v", err)
	}

	// Update the cache with the latest information
	wvw.whitelistedCache.Update(whitelistedVaults)

	log.Printf("Found %d whitelisted vaults out of %d total vaults",
		len(whitelistedVaults), len(vaults))

	return nil
}

// Example usage in main function:
/*
func main() {
	// Connect to Ethereum node
	client, err := ethclient.Dial("https://berachain-rpc.publicnode.com")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Set contract addresses
	factoryAddress := common.HexToAddress("0x94Ad6Ac84f6C6FbA8b8CCbD71d9f4f101def52a8")
	beraChefAddress := common.HexToAddress("0xYourBeraChefContractAddress")

	// Create contract instances
	factory, err := NewVaultFactory(factoryAddress, client)
	if err != nil {
		log.Fatalf("Failed to create vault factory instance: %v", err)
	}

	beraChef, err := NewBeraChef(beraChefAddress, client)
	if err != nil {
		log.Fatalf("Failed to create BeraChef instance: %v", err)
	}

	// Create cache for vaults
	vaultCache := NewVaultCache("berachain", factoryAddress)
	if err := vaultCache.LoadFromDisk(); err != nil {
		log.Printf("Warning: Could not load vault cache: %v", err)
	}

	// Create and start the whitelisted vault worker
	worker := NewWhitelistedVaultWorker(
		client,
		beraChef,
		factory,
		vaultCache,
		5*time.Minute, // Update every 5 minutes
	)

	// Register worker for graceful shutdown
	// Add to existing signal handler...

	// Start the worker
	worker.Start()

	// Use the worker to get whitelisted vaults
	whitelistedVaults := worker.GetWhitelistedVaults()
	fmt.Printf("Current whitelisted vaults: %d\n", len(whitelistedVaults))

	// ... rest of your application code ...

	// When shutting down:
	worker.Stop()
}
*/
