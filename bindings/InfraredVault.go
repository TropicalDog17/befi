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

// IInfraredVaultUserReward is an auto generated low-level Go binding around an user-defined struct.
type IInfraredVaultUserReward struct {
	Token  common.Address
	Amount *big.Int
}

// InfraredVaultMetaData contains all meta data concerning the InfraredVault contract.
var InfraredVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxNumberOfRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"RewardStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NUM_REWARD_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getAllRewardsForUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIInfraredVault.UserReward[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRewardForUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infrared\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardResidual\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsVault\",\"outputs\":[{\"internalType\":\"contractIRewardVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"updateRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InfraredVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use InfraredVaultMetaData.ABI instead.
var InfraredVaultABI = InfraredVaultMetaData.ABI

// InfraredVault is an auto generated Go binding around an Ethereum contract.
type InfraredVault struct {
	InfraredVaultCaller     // Read-only binding to the contract
	InfraredVaultTransactor // Write-only binding to the contract
	InfraredVaultFilterer   // Log filterer for contract events
}

// InfraredVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfraredVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfraredVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfraredVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfraredVaultSession struct {
	Contract     *InfraredVault    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfraredVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfraredVaultCallerSession struct {
	Contract *InfraredVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InfraredVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfraredVaultTransactorSession struct {
	Contract     *InfraredVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InfraredVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfraredVaultRaw struct {
	Contract *InfraredVault // Generic contract binding to access the raw methods on
}

// InfraredVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfraredVaultCallerRaw struct {
	Contract *InfraredVaultCaller // Generic read-only contract binding to access the raw methods on
}

// InfraredVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfraredVaultTransactorRaw struct {
	Contract *InfraredVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfraredVault creates a new instance of InfraredVault, bound to a specific deployed contract.
func NewInfraredVault(address common.Address, backend bind.ContractBackend) (*InfraredVault, error) {
	contract, err := bindInfraredVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfraredVault{InfraredVaultCaller: InfraredVaultCaller{contract: contract}, InfraredVaultTransactor: InfraredVaultTransactor{contract: contract}, InfraredVaultFilterer: InfraredVaultFilterer{contract: contract}}, nil
}

// NewInfraredVaultCaller creates a new read-only instance of InfraredVault, bound to a specific deployed contract.
func NewInfraredVaultCaller(address common.Address, caller bind.ContractCaller) (*InfraredVaultCaller, error) {
	contract, err := bindInfraredVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultCaller{contract: contract}, nil
}

// NewInfraredVaultTransactor creates a new write-only instance of InfraredVault, bound to a specific deployed contract.
func NewInfraredVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*InfraredVaultTransactor, error) {
	contract, err := bindInfraredVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultTransactor{contract: contract}, nil
}

// NewInfraredVaultFilterer creates a new log filterer instance of InfraredVault, bound to a specific deployed contract.
func NewInfraredVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*InfraredVaultFilterer, error) {
	contract, err := bindInfraredVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultFilterer{contract: contract}, nil
}

// bindInfraredVault binds a generic wrapper to an already deployed contract.
func bindInfraredVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfraredVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfraredVault *InfraredVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfraredVault.Contract.InfraredVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfraredVault *InfraredVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.Contract.InfraredVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfraredVault *InfraredVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfraredVault.Contract.InfraredVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfraredVault *InfraredVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfraredVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfraredVault *InfraredVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfraredVault *InfraredVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfraredVault.Contract.contract.Transact(opts, method, params...)
}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) MAXNUMREWARDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "MAX_NUM_REWARD_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_InfraredVault *InfraredVaultSession) MAXNUMREWARDTOKENS() (*big.Int, error) {
	return _InfraredVault.Contract.MAXNUMREWARDTOKENS(&_InfraredVault.CallOpts)
}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) MAXNUMREWARDTOKENS() (*big.Int, error) {
	return _InfraredVault.Contract.MAXNUMREWARDTOKENS(&_InfraredVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_InfraredVault *InfraredVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_InfraredVault *InfraredVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.BalanceOf(&_InfraredVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_InfraredVault *InfraredVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.BalanceOf(&_InfraredVault.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) Earned(opts *bind.CallOpts, account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "earned", account, _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) Earned(account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.Earned(&_InfraredVault.CallOpts, account, _rewardsToken)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) Earned(account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.Earned(&_InfraredVault.CallOpts, account, _rewardsToken)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_InfraredVault *InfraredVaultCaller) GetAllRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "getAllRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_InfraredVault *InfraredVaultSession) GetAllRewardTokens() ([]common.Address, error) {
	return _InfraredVault.Contract.GetAllRewardTokens(&_InfraredVault.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_InfraredVault *InfraredVaultCallerSession) GetAllRewardTokens() ([]common.Address, error) {
	return _InfraredVault.Contract.GetAllRewardTokens(&_InfraredVault.CallOpts)
}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_InfraredVault *InfraredVaultCaller) GetAllRewardsForUser(opts *bind.CallOpts, _user common.Address) ([]IInfraredVaultUserReward, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "getAllRewardsForUser", _user)

	if err != nil {
		return *new([]IInfraredVaultUserReward), err
	}

	out0 := *abi.ConvertType(out[0], new([]IInfraredVaultUserReward)).(*[]IInfraredVaultUserReward)

	return out0, err

}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_InfraredVault *InfraredVaultSession) GetAllRewardsForUser(_user common.Address) ([]IInfraredVaultUserReward, error) {
	return _InfraredVault.Contract.GetAllRewardsForUser(&_InfraredVault.CallOpts, _user)
}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_InfraredVault *InfraredVaultCallerSession) GetAllRewardsForUser(_user common.Address) ([]IInfraredVaultUserReward, error) {
	return _InfraredVault.Contract.GetAllRewardsForUser(&_InfraredVault.CallOpts, _user)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) GetRewardForDuration(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "getRewardForDuration", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.GetRewardForDuration(&_InfraredVault.CallOpts, _rewardsToken)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.GetRewardForDuration(&_InfraredVault.CallOpts, _rewardsToken)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_InfraredVault *InfraredVaultCaller) Infrared(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "infrared")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_InfraredVault *InfraredVaultSession) Infrared() (common.Address, error) {
	return _InfraredVault.Contract.Infrared(&_InfraredVault.CallOpts)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_InfraredVault *InfraredVaultCallerSession) Infrared() (common.Address, error) {
	return _InfraredVault.Contract.Infrared(&_InfraredVault.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) LastTimeRewardApplicable(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "lastTimeRewardApplicable", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.LastTimeRewardApplicable(&_InfraredVault.CallOpts, _rewardsToken)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.LastTimeRewardApplicable(&_InfraredVault.CallOpts, _rewardsToken)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfraredVault *InfraredVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfraredVault *InfraredVaultSession) Paused() (bool, error) {
	return _InfraredVault.Contract.Paused(&_InfraredVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InfraredVault *InfraredVaultCallerSession) Paused() (bool, error) {
	return _InfraredVault.Contract.Paused(&_InfraredVault.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_InfraredVault *InfraredVaultCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "rewardData", arg0)

	outstruct := new(struct {
		RewardsDistributor   common.Address
		RewardsDuration      *big.Int
		PeriodFinish         *big.Int
		RewardRate           *big.Int
		LastUpdateTime       *big.Int
		RewardPerTokenStored *big.Int
		RewardResidual       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsDistributor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardsDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PeriodFinish = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardResidual = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_InfraredVault *InfraredVaultSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	return _InfraredVault.Contract.RewardData(&_InfraredVault.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_InfraredVault *InfraredVaultCallerSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	return _InfraredVault.Contract.RewardData(&_InfraredVault.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) RewardPerToken(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "rewardPerToken", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.RewardPerToken(&_InfraredVault.CallOpts, _rewardsToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.RewardPerToken(&_InfraredVault.CallOpts, _rewardsToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_InfraredVault *InfraredVaultCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_InfraredVault *InfraredVaultSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _InfraredVault.Contract.RewardTokens(&_InfraredVault.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_InfraredVault *InfraredVaultCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _InfraredVault.Contract.RewardTokens(&_InfraredVault.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) Rewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "rewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.Rewards(&_InfraredVault.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.Rewards(&_InfraredVault.CallOpts, arg0, arg1)
}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_InfraredVault *InfraredVaultCaller) RewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "rewardsVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_InfraredVault *InfraredVaultSession) RewardsVault() (common.Address, error) {
	return _InfraredVault.Contract.RewardsVault(&_InfraredVault.CallOpts)
}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_InfraredVault *InfraredVaultCallerSession) RewardsVault() (common.Address, error) {
	return _InfraredVault.Contract.RewardsVault(&_InfraredVault.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InfraredVault *InfraredVaultCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InfraredVault *InfraredVaultSession) StakingToken() (common.Address, error) {
	return _InfraredVault.Contract.StakingToken(&_InfraredVault.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InfraredVault *InfraredVaultCallerSession) StakingToken() (common.Address, error) {
	return _InfraredVault.Contract.StakingToken(&_InfraredVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfraredVault *InfraredVaultSession) TotalSupply() (*big.Int, error) {
	return _InfraredVault.Contract.TotalSupply(&_InfraredVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _InfraredVault.Contract.TotalSupply(&_InfraredVault.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfraredVault.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.UserRewardPerTokenPaid(&_InfraredVault.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_InfraredVault *InfraredVaultCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _InfraredVault.Contract.UserRewardPerTokenPaid(&_InfraredVault.CallOpts, arg0, arg1)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultTransactor) AddReward(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "addReward", _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultSession) AddReward(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.AddReward(&_InfraredVault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultTransactorSession) AddReward(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.AddReward(&_InfraredVault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InfraredVault *InfraredVaultTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InfraredVault *InfraredVaultSession) Exit() (*types.Transaction, error) {
	return _InfraredVault.Contract.Exit(&_InfraredVault.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InfraredVault *InfraredVaultTransactorSession) Exit() (*types.Transaction, error) {
	return _InfraredVault.Contract.Exit(&_InfraredVault.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InfraredVault *InfraredVaultTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InfraredVault *InfraredVaultSession) GetReward() (*types.Transaction, error) {
	return _InfraredVault.Contract.GetReward(&_InfraredVault.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InfraredVault *InfraredVaultTransactorSession) GetReward() (*types.Transaction, error) {
	return _InfraredVault.Contract.GetReward(&_InfraredVault.TransactOpts)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_InfraredVault *InfraredVaultTransactor) GetRewardForUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "getRewardForUser", _user)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_InfraredVault *InfraredVaultSession) GetRewardForUser(_user common.Address) (*types.Transaction, error) {
	return _InfraredVault.Contract.GetRewardForUser(&_InfraredVault.TransactOpts, _user)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_InfraredVault *InfraredVaultTransactorSession) GetRewardForUser(_user common.Address) (*types.Transaction, error) {
	return _InfraredVault.Contract.GetRewardForUser(&_InfraredVault.TransactOpts, _user)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_InfraredVault *InfraredVaultTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "notifyRewardAmount", _rewardToken, _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_InfraredVault *InfraredVaultSession) NotifyRewardAmount(_rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.NotifyRewardAmount(&_InfraredVault.TransactOpts, _rewardToken, _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_InfraredVault *InfraredVaultTransactorSession) NotifyRewardAmount(_rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.NotifyRewardAmount(&_InfraredVault.TransactOpts, _rewardToken, _reward)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_InfraredVault *InfraredVaultTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_InfraredVault *InfraredVaultSession) PauseStaking() (*types.Transaction, error) {
	return _InfraredVault.Contract.PauseStaking(&_InfraredVault.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_InfraredVault *InfraredVaultTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _InfraredVault.Contract.PauseStaking(&_InfraredVault.TransactOpts)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_InfraredVault *InfraredVaultTransactor) RecoverERC20(opts *bind.TransactOpts, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "recoverERC20", _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_InfraredVault *InfraredVaultSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.RecoverERC20(&_InfraredVault.TransactOpts, _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_InfraredVault *InfraredVaultTransactorSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.RecoverERC20(&_InfraredVault.TransactOpts, _to, _token, _amount)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_InfraredVault *InfraredVaultTransactor) RemoveReward(opts *bind.TransactOpts, _rewardsToken common.Address) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "removeReward", _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_InfraredVault *InfraredVaultSession) RemoveReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _InfraredVault.Contract.RemoveReward(&_InfraredVault.TransactOpts, _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_InfraredVault *InfraredVaultTransactorSession) RemoveReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _InfraredVault.Contract.RemoveReward(&_InfraredVault.TransactOpts, _rewardsToken)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InfraredVault *InfraredVaultTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InfraredVault *InfraredVaultSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.Stake(&_InfraredVault.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InfraredVault *InfraredVaultTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.Stake(&_InfraredVault.TransactOpts, amount)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_InfraredVault *InfraredVaultTransactor) UnpauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "unpauseStaking")
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_InfraredVault *InfraredVaultSession) UnpauseStaking() (*types.Transaction, error) {
	return _InfraredVault.Contract.UnpauseStaking(&_InfraredVault.TransactOpts)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_InfraredVault *InfraredVaultTransactorSession) UnpauseStaking() (*types.Transaction, error) {
	return _InfraredVault.Contract.UnpauseStaking(&_InfraredVault.TransactOpts)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultTransactor) UpdateRewardsDuration(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "updateRewardsDuration", _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultSession) UpdateRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.UpdateRewardsDuration(&_InfraredVault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_InfraredVault *InfraredVaultTransactorSession) UpdateRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.UpdateRewardsDuration(&_InfraredVault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InfraredVault *InfraredVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InfraredVault *InfraredVaultSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.Withdraw(&_InfraredVault.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InfraredVault *InfraredVaultTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InfraredVault.Contract.Withdraw(&_InfraredVault.TransactOpts, amount)
}

// InfraredVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the InfraredVault contract.
type InfraredVaultPausedIterator struct {
	Event *InfraredVaultPaused // Event containing the contract specifics and raw log

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
func (it *InfraredVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultPaused)
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
		it.Event = new(InfraredVaultPaused)
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
func (it *InfraredVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultPaused represents a Paused event raised by the InfraredVault contract.
type InfraredVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfraredVault *InfraredVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*InfraredVaultPausedIterator, error) {

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultPausedIterator{contract: _InfraredVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfraredVault *InfraredVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InfraredVaultPaused) (event.Subscription, error) {

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultPaused)
				if err := _InfraredVault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_InfraredVault *InfraredVaultFilterer) ParsePaused(log types.Log) (*InfraredVaultPaused, error) {
	event := new(InfraredVaultPaused)
	if err := _InfraredVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the InfraredVault contract.
type InfraredVaultRecoveredIterator struct {
	Event *InfraredVaultRecovered // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRecovered)
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
		it.Event = new(InfraredVaultRecovered)
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
func (it *InfraredVaultRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRecovered represents a Recovered event raised by the InfraredVault contract.
type InfraredVaultRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) FilterRecovered(opts *bind.FilterOpts) (*InfraredVaultRecoveredIterator, error) {

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRecoveredIterator{contract: _InfraredVault.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *InfraredVaultRecovered) (event.Subscription, error) {

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRecovered)
				if err := _InfraredVault.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) ParseRecovered(log types.Log) (*InfraredVaultRecovered, error) {
	event := new(InfraredVaultRecovered)
	if err := _InfraredVault.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the InfraredVault contract.
type InfraredVaultRewardAddedIterator struct {
	Event *InfraredVaultRewardAdded // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardAdded)
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
		it.Event = new(InfraredVaultRewardAdded)
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
func (it *InfraredVaultRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardAdded represents a RewardAdded event raised by the InfraredVault contract.
type InfraredVaultRewardAdded struct {
	RewardsToken common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardAdded(opts *bind.FilterOpts, rewardsToken []common.Address) (*InfraredVaultRewardAddedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardAdded", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardAddedIterator{contract: _InfraredVault.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardAdded, rewardsToken []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardAdded", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardAdded)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardAdded(log types.Log) (*InfraredVaultRewardAdded, error) {
	event := new(InfraredVaultRewardAdded)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the InfraredVault contract.
type InfraredVaultRewardPaidIterator struct {
	Event *InfraredVaultRewardPaid // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardPaid)
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
		it.Event = new(InfraredVaultRewardPaid)
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
func (it *InfraredVaultRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardPaid represents a RewardPaid event raised by the InfraredVault contract.
type InfraredVaultRewardPaid struct {
	User         common.Address
	RewardsToken common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address, rewardsToken []common.Address) (*InfraredVaultRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardPaid", userRule, rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardPaidIterator{contract: _InfraredVault.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardPaid, user []common.Address, rewardsToken []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardPaid", userRule, rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardPaid)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardPaid(log types.Log) (*InfraredVaultRewardPaid, error) {
	event := new(InfraredVaultRewardPaid)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the InfraredVault contract.
type InfraredVaultRewardRemovedIterator struct {
	Event *InfraredVaultRewardRemoved // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardRemoved)
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
		it.Event = new(InfraredVaultRewardRemoved)
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
func (it *InfraredVaultRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardRemoved represents a RewardRemoved event raised by the InfraredVault contract.
type InfraredVaultRewardRemoved struct {
	RewardsToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardRemoved(opts *bind.FilterOpts, rewardsToken []common.Address) (*InfraredVaultRewardRemovedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardRemoved", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardRemovedIterator{contract: _InfraredVault.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardRemoved, rewardsToken []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardRemoved", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardRemoved)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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

// ParseRewardRemoved is a log parse operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardRemoved(log types.Log) (*InfraredVaultRewardRemoved, error) {
	event := new(InfraredVaultRewardRemoved)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardStoredIterator is returned from FilterRewardStored and is used to iterate over the raw logs and unpacked data for RewardStored events raised by the InfraredVault contract.
type InfraredVaultRewardStoredIterator struct {
	Event *InfraredVaultRewardStored // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardStored)
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
		it.Event = new(InfraredVaultRewardStored)
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
func (it *InfraredVaultRewardStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardStored represents a RewardStored event raised by the InfraredVault contract.
type InfraredVaultRewardStored struct {
	RewardsToken    common.Address
	RewardsDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardStored is a free log retrieval operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardStored(opts *bind.FilterOpts) (*InfraredVaultRewardStoredIterator, error) {

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardStored")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardStoredIterator{contract: _InfraredVault.contract, event: "RewardStored", logs: logs, sub: sub}, nil
}

// WatchRewardStored is a free log subscription operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardStored(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardStored) (event.Subscription, error) {

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardStored)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardStored", log); err != nil {
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

// ParseRewardStored is a log parse operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardStored(log types.Log) (*InfraredVaultRewardStored, error) {
	event := new(InfraredVaultRewardStored)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardsDistributorUpdatedIterator is returned from FilterRewardsDistributorUpdated and is used to iterate over the raw logs and unpacked data for RewardsDistributorUpdated events raised by the InfraredVault contract.
type InfraredVaultRewardsDistributorUpdatedIterator struct {
	Event *InfraredVaultRewardsDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardsDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardsDistributorUpdated)
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
		it.Event = new(InfraredVaultRewardsDistributorUpdated)
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
func (it *InfraredVaultRewardsDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardsDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardsDistributorUpdated represents a RewardsDistributorUpdated event raised by the InfraredVault contract.
type InfraredVaultRewardsDistributorUpdated struct {
	RewardsToken   common.Address
	NewDistributor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorUpdated is a free log retrieval operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardsDistributorUpdated(opts *bind.FilterOpts, rewardsToken []common.Address, newDistributor []common.Address) (*InfraredVaultRewardsDistributorUpdatedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardsDistributorUpdated", rewardsTokenRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardsDistributorUpdatedIterator{contract: _InfraredVault.contract, event: "RewardsDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorUpdated is a free log subscription operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardsDistributorUpdated(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardsDistributorUpdated, rewardsToken []common.Address, newDistributor []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardsDistributorUpdated", rewardsTokenRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardsDistributorUpdated)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardsDistributorUpdated", log); err != nil {
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

// ParseRewardsDistributorUpdated is a log parse operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardsDistributorUpdated(log types.Log) (*InfraredVaultRewardsDistributorUpdated, error) {
	event := new(InfraredVaultRewardsDistributorUpdated)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardsDistributorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the InfraredVault contract.
type InfraredVaultRewardsDurationUpdatedIterator struct {
	Event *InfraredVaultRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRewardsDurationUpdated)
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
		it.Event = new(InfraredVaultRewardsDurationUpdated)
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
func (it *InfraredVaultRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the InfraredVault contract.
type InfraredVaultRewardsDurationUpdated struct {
	Token       common.Address
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_InfraredVault *InfraredVaultFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*InfraredVaultRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRewardsDurationUpdatedIterator{contract: _InfraredVault.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_InfraredVault *InfraredVaultFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *InfraredVaultRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRewardsDurationUpdated)
				if err := _InfraredVault.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_InfraredVault *InfraredVaultFilterer) ParseRewardsDurationUpdated(log types.Log) (*InfraredVaultRewardsDurationUpdated, error) {
	event := new(InfraredVaultRewardsDurationUpdated)
	if err := _InfraredVault.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the InfraredVault contract.
type InfraredVaultStakedIterator struct {
	Event *InfraredVaultStaked // Event containing the contract specifics and raw log

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
func (it *InfraredVaultStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultStaked)
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
		it.Event = new(InfraredVaultStaked)
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
func (it *InfraredVaultStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultStaked represents a Staked event raised by the InfraredVault contract.
type InfraredVaultStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*InfraredVaultStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultStakedIterator{contract: _InfraredVault.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *InfraredVaultStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultStaked)
				if err := _InfraredVault.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) ParseStaked(log types.Log) (*InfraredVaultStaked, error) {
	event := new(InfraredVaultStaked)
	if err := _InfraredVault.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the InfraredVault contract.
type InfraredVaultUnpausedIterator struct {
	Event *InfraredVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *InfraredVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultUnpaused)
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
		it.Event = new(InfraredVaultUnpaused)
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
func (it *InfraredVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultUnpaused represents a Unpaused event raised by the InfraredVault contract.
type InfraredVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfraredVault *InfraredVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InfraredVaultUnpausedIterator, error) {

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultUnpausedIterator{contract: _InfraredVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfraredVault *InfraredVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InfraredVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultUnpaused)
				if err := _InfraredVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_InfraredVault *InfraredVaultFilterer) ParseUnpaused(log types.Log) (*InfraredVaultUnpaused, error) {
	event := new(InfraredVaultUnpaused)
	if err := _InfraredVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the InfraredVault contract.
type InfraredVaultWithdrawnIterator struct {
	Event *InfraredVaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *InfraredVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultWithdrawn)
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
		it.Event = new(InfraredVaultWithdrawn)
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
func (it *InfraredVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultWithdrawn represents a Withdrawn event raised by the InfraredVault contract.
type InfraredVaultWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*InfraredVaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InfraredVault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultWithdrawnIterator{contract: _InfraredVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *InfraredVaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InfraredVault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultWithdrawn)
				if err := _InfraredVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InfraredVault *InfraredVaultFilterer) ParseWithdrawn(log types.Log) (*InfraredVaultWithdrawn, error) {
	event := new(InfraredVaultWithdrawn)
	if err := _InfraredVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
