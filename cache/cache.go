package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/TropicalDog17/befi"
	"github.com/ethereum/go-ethereum/common"
)

// WhitelistCacheData represents the data structure for serializing to disk
type WhitelistCacheData struct {
	LastUpdated time.Time                      `json:"last_updated"`
	NetworkID   string                         `json:"network_id"`
	ChefAddress string                         `json:"chef_address"`
	Vaults      map[string]SerializedWhitelist `json:"vaults"`
}

// SerializedWhitelist represents a serializable version of WhitelistedVault
type SerializedWhitelist struct {
	Address     string    `json:"address"`
	Metadata    string    `json:"metadata"`
	AddedAt     time.Time `json:"added_at"`
	LastChecked time.Time `json:"last_checked"`
}

// PersistentWhitelistCache manages whitelisted vault information with disk persistence
type PersistentWhitelistCache struct {
	vaults      map[common.Address]WhitelistedVaultInfo
	mu          sync.RWMutex
	cachePath   string
	networkID   string
	chefAddress string
	lastUpdated time.Time
}

// WhitelistedVaultInfo extends WhitelistedVault with additional metadata
type WhitelistedVaultInfo struct {
	befi.WhitelistedVault
	AddedAt     time.Time
	LastChecked time.Time
}

// NewPersistentWhitelistCache creates a new whitelist cache with persistence
func NewPersistentWhitelistCache(networkID string, chefAddress common.Address) *PersistentWhitelistCache {
	// Create cache directory if it doesn't exist
	cacheDir := filepath.Join(".cache", "befi-cache")
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		log.Printf("Warning: Could not create cache directory: %v", err)
	}

	// Create cache file path based on network and chef address
	cachePath := filepath.Join(cacheDir, fmt.Sprintf("whitelist_%s_%s.json", networkID, chefAddress.Hex()))

	return &PersistentWhitelistCache{
		vaults:      make(map[common.Address]WhitelistedVaultInfo),
		mu:          sync.RWMutex{},
		cachePath:   cachePath,
		networkID:   networkID,
		chefAddress: chefAddress.Hex(),
		lastUpdated: time.Time{},
	}
}

// LoadFromDisk loads the whitelist cache from disk if available
func (wc *PersistentWhitelistCache) LoadFromDisk() error {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	// Check if cache file exists
	if _, err := os.Stat(wc.cachePath); os.IsNotExist(err) {
		// No cache file yet, that's fine
		return nil
	}

	// Read cache file
	data, err := os.ReadFile(wc.cachePath)
	if err != nil {
		return fmt.Errorf("failed to read cache file: %w", err)
	}

	// Parse JSON
	var cacheData WhitelistCacheData
	if err := json.Unmarshal(data, &cacheData); err != nil {
		return fmt.Errorf("failed to parse cache file: %w", err)
	}

	// Verify the cache is for the correct network and chef
	if cacheData.NetworkID != wc.networkID || cacheData.ChefAddress != wc.chefAddress {
		return fmt.Errorf("cache mismatch: expected network=%s chef=%s, got network=%s chef=%s",
			wc.networkID, wc.chefAddress, cacheData.NetworkID, cacheData.ChefAddress)
	}

	// Load vault addresses into cache
	wc.vaults = make(map[common.Address]WhitelistedVaultInfo, len(cacheData.Vaults))
	for addrStr, serialized := range cacheData.Vaults {
		addr := common.HexToAddress(addrStr)
		wc.vaults[addr] = WhitelistedVaultInfo{
			WhitelistedVault: befi.WhitelistedVault{
				Address:  addr,
				Metadata: serialized.Metadata,
			},
			AddedAt:     serialized.AddedAt,
			LastChecked: serialized.LastChecked,
		}
	}

	wc.lastUpdated = cacheData.LastUpdated
	log.Printf("Loaded %d whitelisted vaults from disk cache (last updated: %s)",
		len(wc.vaults), wc.lastUpdated.Format(time.RFC3339))

	return nil
}

// SaveToDisk saves the whitelist cache to disk
func (wc *PersistentWhitelistCache) SaveToDisk() error {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	// Convert vault data for serialization
	serializedVaults := make(map[string]SerializedWhitelist, len(wc.vaults))
	for addr, info := range wc.vaults {
		serializedVaults[addr.Hex()] = SerializedWhitelist{
			Address:     addr.Hex(),
			Metadata:    info.Metadata,
			AddedAt:     info.AddedAt,
			LastChecked: info.LastChecked,
		}
	}

	// Prepare data for persistence
	cacheData := WhitelistCacheData{
		LastUpdated: wc.lastUpdated,
		NetworkID:   wc.networkID,
		ChefAddress: wc.chefAddress,
		Vaults:      serializedVaults,
	}

	// Convert to JSON
	data, err := json.MarshalIndent(cacheData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %w", err)
	}

	// Write to file (use WriteFile for atomic write)
	if err := os.WriteFile(wc.cachePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	log.Printf("Saved %d whitelisted vaults to disk cache", len(wc.vaults))
	return nil
}

// Update adds or updates a whitelisted vault
func (wc *PersistentWhitelistCache) Update(vault befi.WhitelistedVault) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	now := time.Now()

	// Check if vault already exists to preserve the added timestamp
	if existing, exists := wc.vaults[vault.Address]; exists {
		wc.vaults[vault.Address] = WhitelistedVaultInfo{
			WhitelistedVault: vault,
			AddedAt:          existing.AddedAt,
			LastChecked:      now,
		}
	} else {
		wc.vaults[vault.Address] = WhitelistedVaultInfo{
			WhitelistedVault: vault,
			AddedAt:          now,
			LastChecked:      now,
		}
	}

	wc.lastUpdated = now
}

// Remove removes a vault from the whitelist cache
func (wc *PersistentWhitelistCache) Remove(address common.Address) {
	wc.mu.Lock()
	defer wc.mu.Unlock()

	delete(wc.vaults, address)
	wc.lastUpdated = time.Now()
}

// Get returns a whitelisted vault by address
func (wc *PersistentWhitelistCache) Get(address common.Address) (befi.WhitelistedVault, bool) {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	info, exists := wc.vaults[address]
	if !exists {
		return befi.WhitelistedVault{}, false
	}

	return info.WhitelistedVault, true
}

// GetAll returns all whitelisted vaults
func (wc *PersistentWhitelistCache) GetAll() []befi.WhitelistedVault {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	vaults := make([]befi.WhitelistedVault, 0, len(wc.vaults))
	for _, info := range wc.vaults {
		vaults = append(vaults, info.WhitelistedVault)
	}

	return vaults
}

// GetAllWithMetadata returns all whitelisted vaults with metadata
func (wc *PersistentWhitelistCache) GetAllWithMetadata() []WhitelistedVaultInfo {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	vaults := make([]WhitelistedVaultInfo, 0, len(wc.vaults))
	for _, info := range wc.vaults {
		vaults = append(vaults, info)
	}

	return vaults
}

// GetLastUpdated returns the timestamp of the last update
func (wc *PersistentWhitelistCache) GetLastUpdated() time.Time {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	return wc.lastUpdated
}

// Count returns the number of whitelisted vaults
func (wc *PersistentWhitelistCache) Count() int {
	wc.mu.RLock()
	defer wc.mu.RUnlock()

	return len(wc.vaults)
}
