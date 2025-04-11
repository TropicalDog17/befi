// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RewardVaultFactoryMetaData contains all meta data concerning the RewardVaultFactory contract.
var RewardVaultFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountLessThanMinIncentiveRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRecoverIncentiveToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRecoverRewardToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRecoverStakingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommissionChangeAlreadyQueued\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommissionNotQueuedOrDelayNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositNotMultipleOfGwei\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositValueTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DonateAmountLessThanPayoutAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncentiveRateTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsolventReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientDelegateStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientIncentiveTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientSelfStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidActivateBoostDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBaseRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBoostMultiplier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommissionChangeDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCommissionValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredentialsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDistribution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDropBoostDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxIncentiveTokensCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinBoostedRewardRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPubKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardAllocationWeights\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardClaimDelay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardConvexity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStartBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidateDefaultRewardAllocation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvariantCheckFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxNumWeightsPerRewardAllocationIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinIncentiveRateIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBGT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBlockRewardController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotDistributor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughBoostedBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFactoryVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotFeeCollector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotIncentiveManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNewOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelistedVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayoutAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardAllocationAlreadyQueued\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardAllocationBlockDelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardCycleNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardInactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardsDurationIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimestampAlreadyProcessed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyWhitelistedOrLimitReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyWeights\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalSupplyOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroOperatorOnFirstDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPercentageWeight\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBGTIncentiveDistributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldBGTIncentiveDistributor\",\"type\":\"address\"}],\"name\":\"BGTIncentiveDistributorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allVaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allVaultsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconDepositContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bgt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bgtIncentiveDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"createRewardVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bgt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_beaconDepositContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingToken\",\"type\":\"address\"}],\"name\":\"predictRewardVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bgtIncentiveDistributor\",\"type\":\"address\"}],\"name\":\"setBGTIncentiveDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RewardVaultFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardVaultFactoryMetaData.ABI instead.
var RewardVaultFactoryABI = RewardVaultFactoryMetaData.ABI

// RewardVaultFactory is an auto generated Go binding around an Ethereum contract.
type RewardVaultFactory struct {
	RewardVaultFactoryCaller     // Read-only binding to the contract
	RewardVaultFactoryTransactor // Write-only binding to the contract
	RewardVaultFactoryFilterer   // Log filterer for contract events
}

// RewardVaultFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardVaultFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardVaultFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardVaultFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardVaultFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardVaultFactorySession struct {
	Contract     *RewardVaultFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewardVaultFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardVaultFactoryCallerSession struct {
	Contract *RewardVaultFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RewardVaultFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardVaultFactoryTransactorSession struct {
	Contract     *RewardVaultFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RewardVaultFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardVaultFactoryRaw struct {
	Contract *RewardVaultFactory // Generic contract binding to access the raw methods on
}

// RewardVaultFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardVaultFactoryCallerRaw struct {
	Contract *RewardVaultFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RewardVaultFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardVaultFactoryTransactorRaw struct {
	Contract *RewardVaultFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardVaultFactory creates a new instance of RewardVaultFactory, bound to a specific deployed contract.
func NewRewardVaultFactory(address common.Address, backend bind.ContractBackend) (*RewardVaultFactory, error) {
	contract, err := bindRewardVaultFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactory{RewardVaultFactoryCaller: RewardVaultFactoryCaller{contract: contract}, RewardVaultFactoryTransactor: RewardVaultFactoryTransactor{contract: contract}, RewardVaultFactoryFilterer: RewardVaultFactoryFilterer{contract: contract}}, nil
}

// NewRewardVaultFactoryCaller creates a new read-only instance of RewardVaultFactory, bound to a specific deployed contract.
func NewRewardVaultFactoryCaller(address common.Address, caller bind.ContractCaller) (*RewardVaultFactoryCaller, error) {
	contract, err := bindRewardVaultFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryCaller{contract: contract}, nil
}

// NewRewardVaultFactoryTransactor creates a new write-only instance of RewardVaultFactory, bound to a specific deployed contract.
func NewRewardVaultFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardVaultFactoryTransactor, error) {
	contract, err := bindRewardVaultFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryTransactor{contract: contract}, nil
}

// NewRewardVaultFactoryFilterer creates a new log filterer instance of RewardVaultFactory, bound to a specific deployed contract.
func NewRewardVaultFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardVaultFactoryFilterer, error) {
	contract, err := bindRewardVaultFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryFilterer{contract: contract}, nil
}

// bindRewardVaultFactory binds a generic wrapper to an already deployed contract.
func bindRewardVaultFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardVaultFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardVaultFactory *RewardVaultFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardVaultFactory.Contract.RewardVaultFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardVaultFactory *RewardVaultFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RewardVaultFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardVaultFactory *RewardVaultFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RewardVaultFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardVaultFactory *RewardVaultFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardVaultFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardVaultFactory *RewardVaultFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardVaultFactory *RewardVaultFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.DEFAULTADMINROLE(&_RewardVaultFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.DEFAULTADMINROLE(&_RewardVaultFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RewardVaultFactory *RewardVaultFactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RewardVaultFactory *RewardVaultFactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RewardVaultFactory.Contract.UPGRADEINTERFACEVERSION(&_RewardVaultFactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RewardVaultFactory.Contract.UPGRADEINTERFACEVERSION(&_RewardVaultFactory.CallOpts)
}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCaller) VAULTMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "VAULT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactorySession) VAULTMANAGERROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.VAULTMANAGERROLE(&_RewardVaultFactory.CallOpts)
}

// VAULTMANAGERROLE is a free data retrieval call binding the contract method 0x5e5a24a4.
//
// Solidity: function VAULT_MANAGER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) VAULTMANAGERROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.VAULTMANAGERROLE(&_RewardVaultFactory.CallOpts)
}

// VAULTPAUSERROLE is a free data retrieval call binding the contract method 0xc209bca2.
//
// Solidity: function VAULT_PAUSER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCaller) VAULTPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "VAULT_PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VAULTPAUSERROLE is a free data retrieval call binding the contract method 0xc209bca2.
//
// Solidity: function VAULT_PAUSER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactorySession) VAULTPAUSERROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.VAULTPAUSERROLE(&_RewardVaultFactory.CallOpts)
}

// VAULTPAUSERROLE is a free data retrieval call binding the contract method 0xc209bca2.
//
// Solidity: function VAULT_PAUSER_ROLE() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) VAULTPAUSERROLE() ([32]byte, error) {
	return _RewardVaultFactory.Contract.VAULTPAUSERROLE(&_RewardVaultFactory.CallOpts)
}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) AllVaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "allVaults", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) AllVaults(arg0 *big.Int) (common.Address, error) {
	return _RewardVaultFactory.Contract.AllVaults(&_RewardVaultFactory.CallOpts, arg0)
}

// AllVaults is a free data retrieval call binding the contract method 0x9094a91e.
//
// Solidity: function allVaults(uint256 ) view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) AllVaults(arg0 *big.Int) (common.Address, error) {
	return _RewardVaultFactory.Contract.AllVaults(&_RewardVaultFactory.CallOpts, arg0)
}

// AllVaultsLength is a free data retrieval call binding the contract method 0x4cd18577.
//
// Solidity: function allVaultsLength() view returns(uint256)
func (_RewardVaultFactory *RewardVaultFactoryCaller) AllVaultsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "allVaultsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllVaultsLength is a free data retrieval call binding the contract method 0x4cd18577.
//
// Solidity: function allVaultsLength() view returns(uint256)
func (_RewardVaultFactory *RewardVaultFactorySession) AllVaultsLength() (*big.Int, error) {
	return _RewardVaultFactory.Contract.AllVaultsLength(&_RewardVaultFactory.CallOpts)
}

// AllVaultsLength is a free data retrieval call binding the contract method 0x4cd18577.
//
// Solidity: function allVaultsLength() view returns(uint256)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) AllVaultsLength() (*big.Int, error) {
	return _RewardVaultFactory.Contract.AllVaultsLength(&_RewardVaultFactory.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) Beacon() (common.Address, error) {
	return _RewardVaultFactory.Contract.Beacon(&_RewardVaultFactory.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) Beacon() (common.Address, error) {
	return _RewardVaultFactory.Contract.Beacon(&_RewardVaultFactory.CallOpts)
}

// BeaconDepositContract is a free data retrieval call binding the contract method 0xeafe0f4c.
//
// Solidity: function beaconDepositContract() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) BeaconDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "beaconDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconDepositContract is a free data retrieval call binding the contract method 0xeafe0f4c.
//
// Solidity: function beaconDepositContract() view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) BeaconDepositContract() (common.Address, error) {
	return _RewardVaultFactory.Contract.BeaconDepositContract(&_RewardVaultFactory.CallOpts)
}

// BeaconDepositContract is a free data retrieval call binding the contract method 0xeafe0f4c.
//
// Solidity: function beaconDepositContract() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) BeaconDepositContract() (common.Address, error) {
	return _RewardVaultFactory.Contract.BeaconDepositContract(&_RewardVaultFactory.CallOpts)
}

// Bgt is a free data retrieval call binding the contract method 0x9a6c31c4.
//
// Solidity: function bgt() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) Bgt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "bgt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bgt is a free data retrieval call binding the contract method 0x9a6c31c4.
//
// Solidity: function bgt() view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) Bgt() (common.Address, error) {
	return _RewardVaultFactory.Contract.Bgt(&_RewardVaultFactory.CallOpts)
}

// Bgt is a free data retrieval call binding the contract method 0x9a6c31c4.
//
// Solidity: function bgt() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) Bgt() (common.Address, error) {
	return _RewardVaultFactory.Contract.Bgt(&_RewardVaultFactory.CallOpts)
}

// BgtIncentiveDistributor is a free data retrieval call binding the contract method 0x6cb9f230.
//
// Solidity: function bgtIncentiveDistributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) BgtIncentiveDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "bgtIncentiveDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BgtIncentiveDistributor is a free data retrieval call binding the contract method 0x6cb9f230.
//
// Solidity: function bgtIncentiveDistributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) BgtIncentiveDistributor() (common.Address, error) {
	return _RewardVaultFactory.Contract.BgtIncentiveDistributor(&_RewardVaultFactory.CallOpts)
}

// BgtIncentiveDistributor is a free data retrieval call binding the contract method 0x6cb9f230.
//
// Solidity: function bgtIncentiveDistributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) BgtIncentiveDistributor() (common.Address, error) {
	return _RewardVaultFactory.Contract.BgtIncentiveDistributor(&_RewardVaultFactory.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) Distributor() (common.Address, error) {
	return _RewardVaultFactory.Contract.Distributor(&_RewardVaultFactory.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) Distributor() (common.Address, error) {
	return _RewardVaultFactory.Contract.Distributor(&_RewardVaultFactory.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardVaultFactory.Contract.GetRoleAdmin(&_RewardVaultFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RewardVaultFactory.Contract.GetRoleAdmin(&_RewardVaultFactory.CallOpts, role)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address stakingToken) view returns(address vault)
func (_RewardVaultFactory *RewardVaultFactoryCaller) GetVault(opts *bind.CallOpts, stakingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "getVault", stakingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address stakingToken) view returns(address vault)
func (_RewardVaultFactory *RewardVaultFactorySession) GetVault(stakingToken common.Address) (common.Address, error) {
	return _RewardVaultFactory.Contract.GetVault(&_RewardVaultFactory.CallOpts, stakingToken)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address stakingToken) view returns(address vault)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) GetVault(stakingToken common.Address) (common.Address, error) {
	return _RewardVaultFactory.Contract.GetVault(&_RewardVaultFactory.CallOpts, stakingToken)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardVaultFactory.Contract.HasRole(&_RewardVaultFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RewardVaultFactory.Contract.HasRole(&_RewardVaultFactory.CallOpts, role, account)
}

// PredictRewardVaultAddress is a free data retrieval call binding the contract method 0x5b3b9a18.
//
// Solidity: function predictRewardVaultAddress(address stakingToken) view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCaller) PredictRewardVaultAddress(opts *bind.CallOpts, stakingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "predictRewardVaultAddress", stakingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictRewardVaultAddress is a free data retrieval call binding the contract method 0x5b3b9a18.
//
// Solidity: function predictRewardVaultAddress(address stakingToken) view returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) PredictRewardVaultAddress(stakingToken common.Address) (common.Address, error) {
	return _RewardVaultFactory.Contract.PredictRewardVaultAddress(&_RewardVaultFactory.CallOpts, stakingToken)
}

// PredictRewardVaultAddress is a free data retrieval call binding the contract method 0x5b3b9a18.
//
// Solidity: function predictRewardVaultAddress(address stakingToken) view returns(address)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) PredictRewardVaultAddress(stakingToken common.Address) (common.Address, error) {
	return _RewardVaultFactory.Contract.PredictRewardVaultAddress(&_RewardVaultFactory.CallOpts, stakingToken)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactorySession) ProxiableUUID() ([32]byte, error) {
	return _RewardVaultFactory.Contract.ProxiableUUID(&_RewardVaultFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RewardVaultFactory.Contract.ProxiableUUID(&_RewardVaultFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RewardVaultFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardVaultFactory.Contract.SupportsInterface(&_RewardVaultFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RewardVaultFactory *RewardVaultFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RewardVaultFactory.Contract.SupportsInterface(&_RewardVaultFactory.CallOpts, interfaceId)
}

// CreateRewardVault is a paid mutator transaction binding the contract method 0x577ee5c7.
//
// Solidity: function createRewardVault(address stakingToken) returns(address)
func (_RewardVaultFactory *RewardVaultFactoryTransactor) CreateRewardVault(opts *bind.TransactOpts, stakingToken common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "createRewardVault", stakingToken)
}

// CreateRewardVault is a paid mutator transaction binding the contract method 0x577ee5c7.
//
// Solidity: function createRewardVault(address stakingToken) returns(address)
func (_RewardVaultFactory *RewardVaultFactorySession) CreateRewardVault(stakingToken common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.CreateRewardVault(&_RewardVaultFactory.TransactOpts, stakingToken)
}

// CreateRewardVault is a paid mutator transaction binding the contract method 0x577ee5c7.
//
// Solidity: function createRewardVault(address stakingToken) returns(address)
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) CreateRewardVault(stakingToken common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.CreateRewardVault(&_RewardVaultFactory.TransactOpts, stakingToken)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.GrantRole(&_RewardVaultFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.GrantRole(&_RewardVaultFactory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _bgt, address _distributor, address _beaconDepositContract, address _governance, address _vaultImpl) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) Initialize(opts *bind.TransactOpts, _bgt common.Address, _distributor common.Address, _beaconDepositContract common.Address, _governance common.Address, _vaultImpl common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "initialize", _bgt, _distributor, _beaconDepositContract, _governance, _vaultImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _bgt, address _distributor, address _beaconDepositContract, address _governance, address _vaultImpl) returns()
func (_RewardVaultFactory *RewardVaultFactorySession) Initialize(_bgt common.Address, _distributor common.Address, _beaconDepositContract common.Address, _governance common.Address, _vaultImpl common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.Initialize(&_RewardVaultFactory.TransactOpts, _bgt, _distributor, _beaconDepositContract, _governance, _vaultImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _bgt, address _distributor, address _beaconDepositContract, address _governance, address _vaultImpl) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) Initialize(_bgt common.Address, _distributor common.Address, _beaconDepositContract common.Address, _governance common.Address, _vaultImpl common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.Initialize(&_RewardVaultFactory.TransactOpts, _bgt, _distributor, _beaconDepositContract, _governance, _vaultImpl)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RewardVaultFactory *RewardVaultFactorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RenounceRole(&_RewardVaultFactory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RenounceRole(&_RewardVaultFactory.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RevokeRole(&_RewardVaultFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.RevokeRole(&_RewardVaultFactory.TransactOpts, role, account)
}

// SetBGTIncentiveDistributor is a paid mutator transaction binding the contract method 0x349de329.
//
// Solidity: function setBGTIncentiveDistributor(address _bgtIncentiveDistributor) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) SetBGTIncentiveDistributor(opts *bind.TransactOpts, _bgtIncentiveDistributor common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "setBGTIncentiveDistributor", _bgtIncentiveDistributor)
}

// SetBGTIncentiveDistributor is a paid mutator transaction binding the contract method 0x349de329.
//
// Solidity: function setBGTIncentiveDistributor(address _bgtIncentiveDistributor) returns()
func (_RewardVaultFactory *RewardVaultFactorySession) SetBGTIncentiveDistributor(_bgtIncentiveDistributor common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.SetBGTIncentiveDistributor(&_RewardVaultFactory.TransactOpts, _bgtIncentiveDistributor)
}

// SetBGTIncentiveDistributor is a paid mutator transaction binding the contract method 0x349de329.
//
// Solidity: function setBGTIncentiveDistributor(address _bgtIncentiveDistributor) returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) SetBGTIncentiveDistributor(_bgtIncentiveDistributor common.Address) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.SetBGTIncentiveDistributor(&_RewardVaultFactory.TransactOpts, _bgtIncentiveDistributor)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RewardVaultFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RewardVaultFactory *RewardVaultFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.UpgradeToAndCall(&_RewardVaultFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RewardVaultFactory *RewardVaultFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RewardVaultFactory.Contract.UpgradeToAndCall(&_RewardVaultFactory.TransactOpts, newImplementation, data)
}

// RewardVaultFactoryBGTIncentiveDistributorSetIterator is returned from FilterBGTIncentiveDistributorSet and is used to iterate over the raw logs and unpacked data for BGTIncentiveDistributorSet events raised by the RewardVaultFactory contract.
type RewardVaultFactoryBGTIncentiveDistributorSetIterator struct {
	Event *RewardVaultFactoryBGTIncentiveDistributorSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryBGTIncentiveDistributorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryBGTIncentiveDistributorSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryBGTIncentiveDistributorSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryBGTIncentiveDistributorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryBGTIncentiveDistributorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryBGTIncentiveDistributorSet represents a BGTIncentiveDistributorSet event raised by the RewardVaultFactory contract.
type RewardVaultFactoryBGTIncentiveDistributorSet struct {
	NewBGTIncentiveDistributor common.Address
	OldBGTIncentiveDistributor common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterBGTIncentiveDistributorSet is a free log retrieval operation binding the contract event 0x38bffcf8509d63333ad84013d433268a1d044bf5f68fdc7ffea76c73b737eda0.
//
// Solidity: event BGTIncentiveDistributorSet(address indexed newBGTIncentiveDistributor, address indexed oldBGTIncentiveDistributor)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterBGTIncentiveDistributorSet(opts *bind.FilterOpts, newBGTIncentiveDistributor []common.Address, oldBGTIncentiveDistributor []common.Address) (*RewardVaultFactoryBGTIncentiveDistributorSetIterator, error) {

	var newBGTIncentiveDistributorRule []interface{}
	for _, newBGTIncentiveDistributorItem := range newBGTIncentiveDistributor {
		newBGTIncentiveDistributorRule = append(newBGTIncentiveDistributorRule, newBGTIncentiveDistributorItem)
	}
	var oldBGTIncentiveDistributorRule []interface{}
	for _, oldBGTIncentiveDistributorItem := range oldBGTIncentiveDistributor {
		oldBGTIncentiveDistributorRule = append(oldBGTIncentiveDistributorRule, oldBGTIncentiveDistributorItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "BGTIncentiveDistributorSet", newBGTIncentiveDistributorRule, oldBGTIncentiveDistributorRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryBGTIncentiveDistributorSetIterator{contract: _RewardVaultFactory.contract, event: "BGTIncentiveDistributorSet", logs: logs, sub: sub}, nil
}

// WatchBGTIncentiveDistributorSet is a free log subscription operation binding the contract event 0x38bffcf8509d63333ad84013d433268a1d044bf5f68fdc7ffea76c73b737eda0.
//
// Solidity: event BGTIncentiveDistributorSet(address indexed newBGTIncentiveDistributor, address indexed oldBGTIncentiveDistributor)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchBGTIncentiveDistributorSet(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryBGTIncentiveDistributorSet, newBGTIncentiveDistributor []common.Address, oldBGTIncentiveDistributor []common.Address) (event.Subscription, error) {

	var newBGTIncentiveDistributorRule []interface{}
	for _, newBGTIncentiveDistributorItem := range newBGTIncentiveDistributor {
		newBGTIncentiveDistributorRule = append(newBGTIncentiveDistributorRule, newBGTIncentiveDistributorItem)
	}
	var oldBGTIncentiveDistributorRule []interface{}
	for _, oldBGTIncentiveDistributorItem := range oldBGTIncentiveDistributor {
		oldBGTIncentiveDistributorRule = append(oldBGTIncentiveDistributorRule, oldBGTIncentiveDistributorItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "BGTIncentiveDistributorSet", newBGTIncentiveDistributorRule, oldBGTIncentiveDistributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryBGTIncentiveDistributorSet)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "BGTIncentiveDistributorSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBGTIncentiveDistributorSet is a log parse operation binding the contract event 0x38bffcf8509d63333ad84013d433268a1d044bf5f68fdc7ffea76c73b737eda0.
//
// Solidity: event BGTIncentiveDistributorSet(address indexed newBGTIncentiveDistributor, address indexed oldBGTIncentiveDistributor)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseBGTIncentiveDistributorSet(log types.Log) (*RewardVaultFactoryBGTIncentiveDistributorSet, error) {
	event := new(RewardVaultFactoryBGTIncentiveDistributorSet)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "BGTIncentiveDistributorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardVaultFactory contract.
type RewardVaultFactoryInitializedIterator struct {
	Event *RewardVaultFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryInitialized represents a Initialized event raised by the RewardVaultFactory contract.
type RewardVaultFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardVaultFactoryInitializedIterator, error) {

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryInitializedIterator{contract: _RewardVaultFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryInitialized)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseInitialized(log types.Log) (*RewardVaultFactoryInitialized, error) {
	event := new(RewardVaultFactoryInitialized)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleAdminChangedIterator struct {
	Event *RewardVaultFactoryRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RewardVaultFactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryRoleAdminChangedIterator{contract: _RewardVaultFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryRoleAdminChanged)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*RewardVaultFactoryRoleAdminChanged, error) {
	event := new(RewardVaultFactoryRoleAdminChanged)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleGrantedIterator struct {
	Event *RewardVaultFactoryRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryRoleGranted represents a RoleGranted event raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardVaultFactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryRoleGrantedIterator{contract: _RewardVaultFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryRoleGranted)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseRoleGranted(log types.Log) (*RewardVaultFactoryRoleGranted, error) {
	event := new(RewardVaultFactoryRoleGranted)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleRevokedIterator struct {
	Event *RewardVaultFactoryRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryRoleRevoked represents a RoleRevoked event raised by the RewardVaultFactory contract.
type RewardVaultFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RewardVaultFactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryRoleRevokedIterator{contract: _RewardVaultFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryRoleRevoked)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseRoleRevoked(log types.Log) (*RewardVaultFactoryRoleRevoked, error) {
	event := new(RewardVaultFactoryRoleRevoked)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RewardVaultFactory contract.
type RewardVaultFactoryUpgradedIterator struct {
	Event *RewardVaultFactoryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryUpgraded represents a Upgraded event raised by the RewardVaultFactory contract.
type RewardVaultFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RewardVaultFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryUpgradedIterator{contract: _RewardVaultFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryUpgraded)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseUpgraded(log types.Log) (*RewardVaultFactoryUpgraded, error) {
	event := new(RewardVaultFactoryUpgraded)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardVaultFactoryVaultCreatedIterator is returned from FilterVaultCreated and is used to iterate over the raw logs and unpacked data for VaultCreated events raised by the RewardVaultFactory contract.
type RewardVaultFactoryVaultCreatedIterator struct {
	Event *RewardVaultFactoryVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RewardVaultFactoryVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardVaultFactoryVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RewardVaultFactoryVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RewardVaultFactoryVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardVaultFactoryVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardVaultFactoryVaultCreated represents a VaultCreated event raised by the RewardVaultFactory contract.
type RewardVaultFactoryVaultCreated struct {
	StakingToken common.Address
	Vault        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVaultCreated is a free log retrieval operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed stakingToken, address indexed vault)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) FilterVaultCreated(opts *bind.FilterOpts, stakingToken []common.Address, vault []common.Address) (*RewardVaultFactoryVaultCreatedIterator, error) {

	var stakingTokenRule []interface{}
	for _, stakingTokenItem := range stakingToken {
		stakingTokenRule = append(stakingTokenRule, stakingTokenItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.FilterLogs(opts, "VaultCreated", stakingTokenRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &RewardVaultFactoryVaultCreatedIterator{contract: _RewardVaultFactory.contract, event: "VaultCreated", logs: logs, sub: sub}, nil
}

// WatchVaultCreated is a free log subscription operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed stakingToken, address indexed vault)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) WatchVaultCreated(opts *bind.WatchOpts, sink chan<- *RewardVaultFactoryVaultCreated, stakingToken []common.Address, vault []common.Address) (event.Subscription, error) {

	var stakingTokenRule []interface{}
	for _, stakingTokenItem := range stakingToken {
		stakingTokenRule = append(stakingTokenRule, stakingTokenItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _RewardVaultFactory.contract.WatchLogs(opts, "VaultCreated", stakingTokenRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardVaultFactoryVaultCreated)
				if err := _RewardVaultFactory.contract.UnpackLog(event, "VaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultCreated is a log parse operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed stakingToken, address indexed vault)
func (_RewardVaultFactory *RewardVaultFactoryFilterer) ParseVaultCreated(log types.Log) (*RewardVaultFactoryVaultCreated, error) {
	event := new(RewardVaultFactoryVaultCreated)
	if err := _RewardVaultFactory.contract.UnpackLog(event, "VaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
