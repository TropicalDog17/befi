package befi

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// VaultFactory is a simplified interface for the contract
type VaultFactory struct {
	client  *ethclient.Client
	address common.Address
	*VaultFactoryCaller
}

// VaultFactoryCaller contains read-only methods
type VaultFactoryCaller struct {
	contract *bind.BoundContract
}

func NewVaultFactory(address common.Address, client *ethclient.Client) (*VaultFactory, error) {
	jsonFile, err := os.Open("abi/RewardVaultFactory.json")
	if err != nil {
		return nil, err
	}

	parsedABI, err := abi.JSON(jsonFile)
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(address, parsedABI, client, client, client)

	return &VaultFactory{
		client:             client,
		address:            address,
		VaultFactoryCaller: &VaultFactoryCaller{contract: contract},
	}, nil

}

func (c *VaultFactory) Address() common.Address {
	return c.address
}

// AllVaultsLength retrieves the number of vaults
func (c *VaultFactoryCaller) AllVaultsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []any
	err := c.contract.Call(opts, &out, "allVaultsLength")
	if err != nil {
		return nil, err
	}
	return *abi.ConvertType(out[0], new(*big.Int)).(**big.Int), nil
}

// AllVaults retrieves a vault by index
func (c *VaultFactoryCaller) AllVaults(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []any
	err := c.contract.Call(opts, &out, "allVaults", index)
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// GetVault retrieves a vault by staking token
func (c *VaultFactoryCaller) GetVault(opts *bind.CallOpts, stakingToken common.Address) (common.Address, error) {
	var out []any
	err := c.contract.Call(opts, &out, "getVault", stakingToken)
	if err != nil {
		return common.Address{}, err
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// GetAllVaults retrieves all vaults by using allVaultsLength and iterating through indices
func GetAllVaults(client *ethclient.Client, factoryAddress common.Address) ([]common.Address, error) {
	factory, err := NewVaultFactory(factoryAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create factory instance: %v", err)
	}

	// Get the total number of vaults
	length, err := factory.AllVaultsLength(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults length: %v", err)
	}

	vaults := make([]common.Address, 0, length.Int64())

	// Iterate through all vaults
	for i := int64(0); i < length.Int64(); i++ {
		vault, err := factory.AllVaults(&bind.CallOpts{}, big.NewInt(i))
		if err != nil {
			return nil, fmt.Errorf("failed to get vault at index %d: %v", i, err)
		}

		vaults = append(vaults, vault)
	}

	return vaults, nil
}
