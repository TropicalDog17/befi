package cache

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"sync"

	"github.com/TropicalDog17/befi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// VaultCache provides caching for vault addresses to reduce RPC calls
// with persistence to local storage
type VaultCache struct {
	vaults         map[common.Address]bool // Using a map for fast lookup
	mu             sync.RWMutex
	cachePath      string
	networkID      string
	factoryAddress string
}

// NewVaultCache creates a new vault cache with persistence
func NewVaultCache(networkID string, factoryAddress common.Address) *VaultCache {
	// Create cache directory if it doesn't exist
	cacheDir := filepath.Join(os.TempDir(), "befi-cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		log.Printf("Warning: Could not create cache directory: %v", err)
	}

	// Create cache file path based on network and factory address
	cachePath := filepath.Join(cacheDir, fmt.Sprintf("vaults_%s_%s.json", networkID, factoryAddress.Hex()))

	return &VaultCache{
		vaults:         make(map[common.Address]bool),
		mu:             sync.RWMutex{},
		cachePath:      cachePath,
		networkID:      networkID,
		factoryAddress: factoryAddress.Hex(),
	}
}

// persistentVaultData is the structure we'll save to disk
type persistentVaultData struct {
	NetworkID      string   `json:"network_id"`
	FactoryAddress string   `json:"factory_address"`
	VaultAddresses []string `json:"vault_addresses"`
}

// LoadFromDisk loads the vault cache from disk if available
func (vc *VaultCache) LoadFromDisk() error {
	vc.mu.Lock()
	defer vc.mu.Unlock()

	// Check if cache file exists
	if _, err := os.Stat(vc.cachePath); os.IsNotExist(err) {
		// No cache file yet, that's fine
		return nil
	}

	// Read cache file
	data, err := ioutil.ReadFile(vc.cachePath)
	if err != nil {
		return fmt.Errorf("failed to read cache file: %w", err)
	}

	// Parse JSON
	var persistentData persistentVaultData
	if err := json.Unmarshal(data, &persistentData); err != nil {
		return fmt.Errorf("failed to parse cache file: %w", err)
	}

	// Verify the cache is for the correct network and factory
	if persistentData.NetworkID != vc.networkID || persistentData.FactoryAddress != vc.factoryAddress {
		return fmt.Errorf("cache mismatch: expected network=%s factory=%s, got network=%s factory=%s",
			vc.networkID, vc.factoryAddress, persistentData.NetworkID, persistentData.FactoryAddress)
	}

	// Load vault addresses into cache
	vc.vaults = make(map[common.Address]bool, len(persistentData.VaultAddresses))
	for _, addrStr := range persistentData.VaultAddresses {
		vc.vaults[common.HexToAddress(addrStr)] = true
	}

	log.Printf("Loaded %d vaults from disk cache", len(vc.vaults))
	return nil
}

// SaveToDisk saves the vault cache to disk
func (vc *VaultCache) SaveToDisk() error {
	vc.mu.RLock()
	defer vc.mu.RUnlock()

	// Convert vault addresses to strings
	vaultAddresses := make([]string, 0, len(vc.vaults))
	for addr := range vc.vaults {
		vaultAddresses = append(vaultAddresses, addr.Hex())
	}

	// Prepare data for persistence
	persistentData := persistentVaultData{
		NetworkID:      vc.networkID,
		FactoryAddress: vc.factoryAddress,
		VaultAddresses: vaultAddresses,
	}

	// Convert to JSON
	data, err := json.MarshalIndent(persistentData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %w", err)
	}

	// Write to file (use WriteFile for atomic write)
	if err := ioutil.WriteFile(vc.cachePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	log.Printf("Saved %d vaults to disk cache", len(vc.vaults))
	return nil
}

// GetVaults returns all vaults, fetching only new ones if needed
func (vc *VaultCache) GetVaults(client *ethclient.Client, factoryAddress common.Address) ([]common.Address, error) {
	// First check if we have any cached vaults
	vc.mu.RLock()
	cachedCount := len(vc.vaults)
	vc.mu.RUnlock()

	// Create a factory instance
	factory, err := befi.NewVaultFactory(factoryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory instance: %v", err)
	}

	// Get the current total vault count
	length, err := factory.AllVaultsLength(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults length: %v", err)
	}

	totalVaultCount := int(length.Int64())

	// If we have all vaults already, return the cached ones
	if cachedCount >= totalVaultCount {
		log.Printf("Using %d cached vaults (no new vaults found)", cachedCount)
		return vc.getCachedVaults(), nil
	}

	log.Printf("Found %d new vaults (had %d cached)", totalVaultCount-cachedCount, cachedCount)

	// We need to fetch new vaults
	vaults, err := vc.fetchNewVaults(client, factory, cachedCount, totalVaultCount)
	if err != nil {
		return nil, err
	}

	// Save updated cache to disk
	if err := vc.SaveToDisk(); err != nil {
		log.Printf("Warning: Could not save cache to disk: %v", err)
	}

	return vaults, nil
}

// getCachedVaults returns the currently cached vault addresses
func (vc *VaultCache) getCachedVaults() []common.Address {
	vc.mu.RLock()
	defer vc.mu.RUnlock()

	vaults := make([]common.Address, 0, len(vc.vaults))
	for vault := range vc.vaults {
		vaults = append(vaults, vault)
	}
	return vaults
}

// fetchNewVaults fetches only the new vaults that aren't in cache yet
// Saves cache incrementally to avoid data loss if interrupted
func (vc *VaultCache) fetchNewVaults(client *ethclient.Client, factory *befi.VaultFactory,
	cachedCount, totalVaultCount int) ([]common.Address, error) {

	// Only fetch the new vaults (from cachedCount to totalVaultCount-1)
	newVaults := make([]common.Address, 0, totalVaultCount-cachedCount)

	// Save cache after every batch of vaults to prevent data loss on interruption
	batchSize := 10
	saveThreshold := cachedCount + batchSize

	for i := cachedCount; i < totalVaultCount; i++ {
		vault, err := factory.AllVaults(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			return nil, fmt.Errorf("failed to get vault at index %d: %v", i, err)
		}

		// Add to new vaults list
		newVaults = append(newVaults, vault)

		// Update the cache immediately
		vc.mu.Lock()
		vc.vaults[vault] = true
		vc.mu.Unlock()

		// Save incrementally after each batch
		if i+1 >= saveThreshold || i+1 == totalVaultCount {
			if err := vc.SaveToDisk(); err != nil {
				log.Printf("Warning: Could not save cache checkpoint to disk: %v", err)
			} else {
				log.Printf("Saved checkpoint with %d/%d vaults", i+1, totalVaultCount)
			}
			saveThreshold = i + 1 + batchSize
		}
	}

	// Return all vaults
	return vc.getCachedVaults(), nil
}
