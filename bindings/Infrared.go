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

// DataTypesToken is an auto generated low-level Go binding around an user-defined struct.
type DataTypesToken struct {
	TokenAddress common.Address
	Amount       *big.Int
}

// IBeraChefWeight is an auto generated low-level Go binding around an user-defined struct.
type IBeraChefWeight struct {
	Receiver            common.Address
	PercentageNumerator *big.Int
}

// ValidatorTypesValidator is an auto generated low-level Go binding around an user-defined struct.
type ValidatorTypesValidator struct {
	Pubkey []byte
	Addr   common.Address
}

// InfraredMetaData contains all meta data concerning the Infrared contract.
var InfraredMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRewardsVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardTokenNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardTokenNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensReservedForProtocolFees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultAlreadyUpToDate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"ActivatedBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bgtAmt\",\"type\":\"uint256\"}],\"name\":\"BaseHarvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"BribeSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amtWiberaVault\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amtIbgtVault\",\"type\":\"uint256\"}],\"name\":\"BribesCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"CancelDropBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"CancelledBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"DroppedBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumConfigTypes.FeeType\",\"name\":\"_feeType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_ibgt\",\"type\":\"address\"}],\"name\":\"IBGTSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_IR\",\"type\":\"address\"}],\"name\":\"IRSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldWeight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newWeight\",\"type\":\"uint256\"}],\"name\":\"InfraredBERABribeSplitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ibgtAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_iredAmt\",\"type\":\"uint256\"}],\"name\":\"InfraredBGTSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldIbgt\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newIbgt\",\"type\":\"address\"}],\"name\":\"InfraredBGTUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldIbgtVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newIbgtVault\",\"type\":\"address\"}],\"name\":\"InfraredBGTVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"NewVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_ibera\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"OperatorRewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_voterAmt\",\"type\":\"uint256\"}],\"name\":\"ProtocolFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"QueueDropBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"QueuedBoosts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_from\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"RewardSupplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"RewardTokenNotSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"UpdatedIRMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"commissionRate\",\"type\":\"uint96\"}],\"name\":\"ValidatorCommissionActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"commissionRate\",\"type\":\"uint96\"}],\"name\":\"ValidatorCommissionQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"_validator\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structDataTypes.Token[]\",\"name\":\"_rewards\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bgtAmt\",\"type\":\"uint256\"}],\"name\":\"ValidatorHarvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_current\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_new\",\"type\":\"bytes\"}],\"name\":\"ValidatorReplaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structValidatorTypes.Validator[]\",\"name\":\"_validators\",\"type\":\"tuple[]\"}],\"name\":\"ValidatorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"ValidatorsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bgtAmt\",\"type\":\"uint256\"}],\"name\":\"VaultHarvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"VaultMigrated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"VaultRegistrationPauseStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"VoterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_wasWhitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isWhitelisted\",\"type\":\"bool\"}],\"name\":\"WhiteListedRewardTokensUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEEPER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VAULT_STORAGE_LOCATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"activateBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"activateQueuedValCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addIncentives\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structValidatorTypes.Validator[]\",\"name\":\"_validators\",\"type\":\"tuple[]\"}],\"name\":\"addValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"cancelBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"cancelDropBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeProtocol\",\"type\":\"uint256\"}],\"name\":\"chargedFeesOnRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amtRecipient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amtVoter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amtProtocol\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chef\",\"outputs\":[{\"internalType\":\"contractIBeraChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"claimLostRewardsOnVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"claimProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"collectBribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collector\",\"outputs\":[{\"internalType\":\"contractIBribeCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatee\",\"type\":\"address\"}],\"name\":\"delegateBGT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"contractIInfraredDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"dropBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"t\",\"type\":\"uint256\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBGTBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestBoostRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"harvestBribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"harvestOldVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestOperatorRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"harvestVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"honey\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibera\",\"outputs\":[{\"internalType\":\"contractIInfraredBERA\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgt\",\"outputs\":[{\"internalType\":\"contractInfraredBGT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgtVault\",\"outputs\":[{\"internalType\":\"contractIInfraredVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infrared\",\"outputs\":[{\"internalType\":\"contractIInfrared\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infraredValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"internalType\":\"structValidatorTypes.Validator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakingTokens\",\"type\":\"address[]\"}],\"name\":\"initializeV1_2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ir\",\"outputs\":[{\"internalType\":\"contractIInfraredGovernanceToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"isInfraredValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"versionToUpgradeTo\",\"type\":\"uint8\"}],\"name\":\"migrateVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numInfraredValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"pauseOldStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"pauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"protocolFeeAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"queueBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint128[]\",\"name\":\"_amts\",\"type\":\"uint128[]\"}],\"name\":\"queueDropBoosts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint96\",\"name\":\"_commissionRate\",\"type\":\"uint96\"}],\"name\":\"queueMultipleValCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_startBlock\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"percentageNumerator\",\"type\":\"uint96\"}],\"internalType\":\"structIBeraChef.Weight[]\",\"name\":\"_weights\",\"type\":\"tuple[]\"}],\"name\":\"queueNewCuttingBoard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"_commissionRate\",\"type\":\"uint96\"}],\"name\":\"queueValCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20FromOldVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20FromVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"registerVault\",\"outputs\":[{\"internalType\":\"contractIInfraredVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"removeValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_current\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_new\",\"type\":\"bytes\"}],\"name\":\"replaceValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsFactory\",\"outputs\":[{\"internalType\":\"contractIRewardVaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ir\",\"type\":\"address\"}],\"name\":\"setIR\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"name\":\"setVaultRegistrationPauseStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"unpauseOldStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"unpauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumConfigTypes.FeeType\",\"name\":\"_t\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_irMintRate\",\"type\":\"uint256\"}],\"name\":\"updateIRMintRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"updateInfraredBERABribeSplit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"updateRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"updateRewardsDurationForVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_whitelisted\",\"type\":\"bool\"}],\"name\":\"updateWhiteListedRewardTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"name\":\"vaultRegistry\",\"outputs\":[{\"internalType\":\"contractIInfraredVault\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voter\",\"outputs\":[{\"internalType\":\"contractIVoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wbera\",\"outputs\":[{\"internalType\":\"contractIWBERA\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"whitelistedRewardTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// InfraredABI is the input ABI used to generate the binding from.
// Deprecated: Use InfraredMetaData.ABI instead.
var InfraredABI = InfraredMetaData.ABI

// Infrared is an auto generated Go binding around an Ethereum contract.
type Infrared struct {
	InfraredCaller     // Read-only binding to the contract
	InfraredTransactor // Write-only binding to the contract
	InfraredFilterer   // Log filterer for contract events
}

// InfraredCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfraredCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfraredTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfraredFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfraredSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfraredSession struct {
	Contract     *Infrared         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfraredCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfraredCallerSession struct {
	Contract *InfraredCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// InfraredTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfraredTransactorSession struct {
	Contract     *InfraredTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InfraredRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfraredRaw struct {
	Contract *Infrared // Generic contract binding to access the raw methods on
}

// InfraredCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfraredCallerRaw struct {
	Contract *InfraredCaller // Generic read-only contract binding to access the raw methods on
}

// InfraredTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfraredTransactorRaw struct {
	Contract *InfraredTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfrared creates a new instance of Infrared, bound to a specific deployed contract.
func NewInfrared(address common.Address, backend bind.ContractBackend) (*Infrared, error) {
	contract, err := bindInfrared(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Infrared{InfraredCaller: InfraredCaller{contract: contract}, InfraredTransactor: InfraredTransactor{contract: contract}, InfraredFilterer: InfraredFilterer{contract: contract}}, nil
}

// NewInfraredCaller creates a new read-only instance of Infrared, bound to a specific deployed contract.
func NewInfraredCaller(address common.Address, caller bind.ContractCaller) (*InfraredCaller, error) {
	contract, err := bindInfrared(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfraredCaller{contract: contract}, nil
}

// NewInfraredTransactor creates a new write-only instance of Infrared, bound to a specific deployed contract.
func NewInfraredTransactor(address common.Address, transactor bind.ContractTransactor) (*InfraredTransactor, error) {
	contract, err := bindInfrared(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfraredTransactor{contract: contract}, nil
}

// NewInfraredFilterer creates a new log filterer instance of Infrared, bound to a specific deployed contract.
func NewInfraredFilterer(address common.Address, filterer bind.ContractFilterer) (*InfraredFilterer, error) {
	contract, err := bindInfrared(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfraredFilterer{contract: contract}, nil
}

// bindInfrared binds a generic wrapper to an already deployed contract.
func bindInfrared(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfraredMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Infrared *InfraredRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Infrared.Contract.InfraredCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Infrared *InfraredRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.Contract.InfraredTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Infrared *InfraredRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Infrared.Contract.InfraredTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Infrared *InfraredCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Infrared.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Infrared *InfraredTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Infrared *InfraredTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Infrared.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Infrared *InfraredCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Infrared *InfraredSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Infrared.Contract.DEFAULTADMINROLE(&_Infrared.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Infrared *InfraredCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Infrared.Contract.DEFAULTADMINROLE(&_Infrared.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Infrared *InfraredCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Infrared *InfraredSession) GOVERNANCEROLE() ([32]byte, error) {
	return _Infrared.Contract.GOVERNANCEROLE(&_Infrared.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_Infrared *InfraredCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _Infrared.Contract.GOVERNANCEROLE(&_Infrared.CallOpts)
}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Infrared *InfraredCaller) KEEPERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "KEEPER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Infrared *InfraredSession) KEEPERROLE() ([32]byte, error) {
	return _Infrared.Contract.KEEPERROLE(&_Infrared.CallOpts)
}

// KEEPERROLE is a free data retrieval call binding the contract method 0x364bc15a.
//
// Solidity: function KEEPER_ROLE() view returns(bytes32)
func (_Infrared *InfraredCallerSession) KEEPERROLE() ([32]byte, error) {
	return _Infrared.Contract.KEEPERROLE(&_Infrared.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Infrared *InfraredCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Infrared *InfraredSession) PAUSERROLE() ([32]byte, error) {
	return _Infrared.Contract.PAUSERROLE(&_Infrared.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Infrared *InfraredCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Infrared.Contract.PAUSERROLE(&_Infrared.CallOpts)
}

// REWARDSSTORAGELOCATION is a free data retrieval call binding the contract method 0xdefe0f03.
//
// Solidity: function REWARDS_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCaller) REWARDSSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "REWARDS_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDSSTORAGELOCATION is a free data retrieval call binding the contract method 0xdefe0f03.
//
// Solidity: function REWARDS_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredSession) REWARDSSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.REWARDSSTORAGELOCATION(&_Infrared.CallOpts)
}

// REWARDSSTORAGELOCATION is a free data retrieval call binding the contract method 0xdefe0f03.
//
// Solidity: function REWARDS_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCallerSession) REWARDSSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.REWARDSSTORAGELOCATION(&_Infrared.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Infrared *InfraredCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Infrared *InfraredSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Infrared.Contract.UPGRADEINTERFACEVERSION(&_Infrared.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Infrared *InfraredCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Infrared.Contract.UPGRADEINTERFACEVERSION(&_Infrared.CallOpts)
}

// VALIDATORSTORAGELOCATION is a free data retrieval call binding the contract method 0x0c0c2ca0.
//
// Solidity: function VALIDATOR_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCaller) VALIDATORSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "VALIDATOR_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATORSTORAGELOCATION is a free data retrieval call binding the contract method 0x0c0c2ca0.
//
// Solidity: function VALIDATOR_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredSession) VALIDATORSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.VALIDATORSTORAGELOCATION(&_Infrared.CallOpts)
}

// VALIDATORSTORAGELOCATION is a free data retrieval call binding the contract method 0x0c0c2ca0.
//
// Solidity: function VALIDATOR_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCallerSession) VALIDATORSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.VALIDATORSTORAGELOCATION(&_Infrared.CallOpts)
}

// VAULTSTORAGELOCATION is a free data retrieval call binding the contract method 0x032a3398.
//
// Solidity: function VAULT_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCaller) VAULTSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "VAULT_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VAULTSTORAGELOCATION is a free data retrieval call binding the contract method 0x032a3398.
//
// Solidity: function VAULT_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredSession) VAULTSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.VAULTSTORAGELOCATION(&_Infrared.CallOpts)
}

// VAULTSTORAGELOCATION is a free data retrieval call binding the contract method 0x032a3398.
//
// Solidity: function VAULT_STORAGE_LOCATION() view returns(bytes32)
func (_Infrared *InfraredCallerSession) VAULTSTORAGELOCATION() ([32]byte, error) {
	return _Infrared.Contract.VAULTSTORAGELOCATION(&_Infrared.CallOpts)
}

// ChargedFeesOnRewards is a free data retrieval call binding the contract method 0x178b4094.
//
// Solidity: function chargedFeesOnRewards(uint256 _amt, uint256 _feeTotal, uint256 _feeProtocol) pure returns(uint256 amtRecipient, uint256 amtVoter, uint256 amtProtocol)
func (_Infrared *InfraredCaller) ChargedFeesOnRewards(opts *bind.CallOpts, _amt *big.Int, _feeTotal *big.Int, _feeProtocol *big.Int) (struct {
	AmtRecipient *big.Int
	AmtVoter     *big.Int
	AmtProtocol  *big.Int
}, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "chargedFeesOnRewards", _amt, _feeTotal, _feeProtocol)

	outstruct := new(struct {
		AmtRecipient *big.Int
		AmtVoter     *big.Int
		AmtProtocol  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmtRecipient = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmtVoter = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmtProtocol = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChargedFeesOnRewards is a free data retrieval call binding the contract method 0x178b4094.
//
// Solidity: function chargedFeesOnRewards(uint256 _amt, uint256 _feeTotal, uint256 _feeProtocol) pure returns(uint256 amtRecipient, uint256 amtVoter, uint256 amtProtocol)
func (_Infrared *InfraredSession) ChargedFeesOnRewards(_amt *big.Int, _feeTotal *big.Int, _feeProtocol *big.Int) (struct {
	AmtRecipient *big.Int
	AmtVoter     *big.Int
	AmtProtocol  *big.Int
}, error) {
	return _Infrared.Contract.ChargedFeesOnRewards(&_Infrared.CallOpts, _amt, _feeTotal, _feeProtocol)
}

// ChargedFeesOnRewards is a free data retrieval call binding the contract method 0x178b4094.
//
// Solidity: function chargedFeesOnRewards(uint256 _amt, uint256 _feeTotal, uint256 _feeProtocol) pure returns(uint256 amtRecipient, uint256 amtVoter, uint256 amtProtocol)
func (_Infrared *InfraredCallerSession) ChargedFeesOnRewards(_amt *big.Int, _feeTotal *big.Int, _feeProtocol *big.Int) (struct {
	AmtRecipient *big.Int
	AmtVoter     *big.Int
	AmtProtocol  *big.Int
}, error) {
	return _Infrared.Contract.ChargedFeesOnRewards(&_Infrared.CallOpts, _amt, _feeTotal, _feeProtocol)
}

// Chef is a free data retrieval call binding the contract method 0x1fc8bc5d.
//
// Solidity: function chef() view returns(address)
func (_Infrared *InfraredCaller) Chef(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "chef")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chef is a free data retrieval call binding the contract method 0x1fc8bc5d.
//
// Solidity: function chef() view returns(address)
func (_Infrared *InfraredSession) Chef() (common.Address, error) {
	return _Infrared.Contract.Chef(&_Infrared.CallOpts)
}

// Chef is a free data retrieval call binding the contract method 0x1fc8bc5d.
//
// Solidity: function chef() view returns(address)
func (_Infrared *InfraredCallerSession) Chef() (common.Address, error) {
	return _Infrared.Contract.Chef(&_Infrared.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Infrared *InfraredCaller) Collector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "collector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Infrared *InfraredSession) Collector() (common.Address, error) {
	return _Infrared.Contract.Collector(&_Infrared.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Infrared *InfraredCallerSession) Collector() (common.Address, error) {
	return _Infrared.Contract.Collector(&_Infrared.CallOpts)
}

// CurrentImplementation is a free data retrieval call binding the contract method 0xd8bd5c29.
//
// Solidity: function currentImplementation() view returns(address)
func (_Infrared *InfraredCaller) CurrentImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "currentImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentImplementation is a free data retrieval call binding the contract method 0xd8bd5c29.
//
// Solidity: function currentImplementation() view returns(address)
func (_Infrared *InfraredSession) CurrentImplementation() (common.Address, error) {
	return _Infrared.Contract.CurrentImplementation(&_Infrared.CallOpts)
}

// CurrentImplementation is a free data retrieval call binding the contract method 0xd8bd5c29.
//
// Solidity: function currentImplementation() view returns(address)
func (_Infrared *InfraredCallerSession) CurrentImplementation() (common.Address, error) {
	return _Infrared.Contract.CurrentImplementation(&_Infrared.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Infrared *InfraredCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Infrared *InfraredSession) Distributor() (common.Address, error) {
	return _Infrared.Contract.Distributor(&_Infrared.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Infrared *InfraredCallerSession) Distributor() (common.Address, error) {
	return _Infrared.Contract.Distributor(&_Infrared.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 t) view returns(uint256)
func (_Infrared *InfraredCaller) Fees(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "fees", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 t) view returns(uint256)
func (_Infrared *InfraredSession) Fees(t *big.Int) (*big.Int, error) {
	return _Infrared.Contract.Fees(&_Infrared.CallOpts, t)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 t) view returns(uint256)
func (_Infrared *InfraredCallerSession) Fees(t *big.Int) (*big.Int, error) {
	return _Infrared.Contract.Fees(&_Infrared.CallOpts, t)
}

// GetBGTBalance is a free data retrieval call binding the contract method 0xf0d68244.
//
// Solidity: function getBGTBalance() view returns(uint256)
func (_Infrared *InfraredCaller) GetBGTBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "getBGTBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBGTBalance is a free data retrieval call binding the contract method 0xf0d68244.
//
// Solidity: function getBGTBalance() view returns(uint256)
func (_Infrared *InfraredSession) GetBGTBalance() (*big.Int, error) {
	return _Infrared.Contract.GetBGTBalance(&_Infrared.CallOpts)
}

// GetBGTBalance is a free data retrieval call binding the contract method 0xf0d68244.
//
// Solidity: function getBGTBalance() view returns(uint256)
func (_Infrared *InfraredCallerSession) GetBGTBalance() (*big.Int, error) {
	return _Infrared.Contract.GetBGTBalance(&_Infrared.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Infrared *InfraredCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Infrared *InfraredSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Infrared.Contract.GetRoleAdmin(&_Infrared.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Infrared *InfraredCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Infrared.Contract.GetRoleAdmin(&_Infrared.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Infrared *InfraredCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Infrared *InfraredSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Infrared.Contract.HasRole(&_Infrared.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Infrared *InfraredCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Infrared.Contract.HasRole(&_Infrared.CallOpts, role, account)
}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_Infrared *InfraredCaller) Honey(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "honey")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_Infrared *InfraredSession) Honey() (common.Address, error) {
	return _Infrared.Contract.Honey(&_Infrared.CallOpts)
}

// Honey is a free data retrieval call binding the contract method 0x36b2c4b2.
//
// Solidity: function honey() view returns(address)
func (_Infrared *InfraredCallerSession) Honey() (common.Address, error) {
	return _Infrared.Contract.Honey(&_Infrared.CallOpts)
}

// Ibera is a free data retrieval call binding the contract method 0x5c9750ee.
//
// Solidity: function ibera() view returns(address)
func (_Infrared *InfraredCaller) Ibera(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "ibera")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ibera is a free data retrieval call binding the contract method 0x5c9750ee.
//
// Solidity: function ibera() view returns(address)
func (_Infrared *InfraredSession) Ibera() (common.Address, error) {
	return _Infrared.Contract.Ibera(&_Infrared.CallOpts)
}

// Ibera is a free data retrieval call binding the contract method 0x5c9750ee.
//
// Solidity: function ibera() view returns(address)
func (_Infrared *InfraredCallerSession) Ibera() (common.Address, error) {
	return _Infrared.Contract.Ibera(&_Infrared.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_Infrared *InfraredCaller) Ibgt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "ibgt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_Infrared *InfraredSession) Ibgt() (common.Address, error) {
	return _Infrared.Contract.Ibgt(&_Infrared.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_Infrared *InfraredCallerSession) Ibgt() (common.Address, error) {
	return _Infrared.Contract.Ibgt(&_Infrared.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_Infrared *InfraredCaller) IbgtVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "ibgtVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_Infrared *InfraredSession) IbgtVault() (common.Address, error) {
	return _Infrared.Contract.IbgtVault(&_Infrared.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_Infrared *InfraredCallerSession) IbgtVault() (common.Address, error) {
	return _Infrared.Contract.IbgtVault(&_Infrared.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Infrared *InfraredCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Infrared *InfraredSession) Implementation() (common.Address, error) {
	return _Infrared.Contract.Implementation(&_Infrared.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Infrared *InfraredCallerSession) Implementation() (common.Address, error) {
	return _Infrared.Contract.Implementation(&_Infrared.CallOpts)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Infrared *InfraredCaller) Infrared(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "infrared")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Infrared *InfraredSession) Infrared() (common.Address, error) {
	return _Infrared.Contract.Infrared(&_Infrared.CallOpts)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Infrared *InfraredCallerSession) Infrared() (common.Address, error) {
	return _Infrared.Contract.Infrared(&_Infrared.CallOpts)
}

// InfraredValidators is a free data retrieval call binding the contract method 0xadc51dcb.
//
// Solidity: function infraredValidators() view returns((bytes,address)[] validators)
func (_Infrared *InfraredCaller) InfraredValidators(opts *bind.CallOpts) ([]ValidatorTypesValidator, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "infraredValidators")

	if err != nil {
		return *new([]ValidatorTypesValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]ValidatorTypesValidator)).(*[]ValidatorTypesValidator)

	return out0, err

}

// InfraredValidators is a free data retrieval call binding the contract method 0xadc51dcb.
//
// Solidity: function infraredValidators() view returns((bytes,address)[] validators)
func (_Infrared *InfraredSession) InfraredValidators() ([]ValidatorTypesValidator, error) {
	return _Infrared.Contract.InfraredValidators(&_Infrared.CallOpts)
}

// InfraredValidators is a free data retrieval call binding the contract method 0xadc51dcb.
//
// Solidity: function infraredValidators() view returns((bytes,address)[] validators)
func (_Infrared *InfraredCallerSession) InfraredValidators() ([]ValidatorTypesValidator, error) {
	return _Infrared.Contract.InfraredValidators(&_Infrared.CallOpts)
}

// Ir is a free data retrieval call binding the contract method 0x5db4cd21.
//
// Solidity: function ir() view returns(address)
func (_Infrared *InfraredCaller) Ir(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "ir")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ir is a free data retrieval call binding the contract method 0x5db4cd21.
//
// Solidity: function ir() view returns(address)
func (_Infrared *InfraredSession) Ir() (common.Address, error) {
	return _Infrared.Contract.Ir(&_Infrared.CallOpts)
}

// Ir is a free data retrieval call binding the contract method 0x5db4cd21.
//
// Solidity: function ir() view returns(address)
func (_Infrared *InfraredCallerSession) Ir() (common.Address, error) {
	return _Infrared.Contract.Ir(&_Infrared.CallOpts)
}

// IsInfraredValidator is a free data retrieval call binding the contract method 0xbab97c7b.
//
// Solidity: function isInfraredValidator(bytes _pubkey) view returns(bool)
func (_Infrared *InfraredCaller) IsInfraredValidator(opts *bind.CallOpts, _pubkey []byte) (bool, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "isInfraredValidator", _pubkey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInfraredValidator is a free data retrieval call binding the contract method 0xbab97c7b.
//
// Solidity: function isInfraredValidator(bytes _pubkey) view returns(bool)
func (_Infrared *InfraredSession) IsInfraredValidator(_pubkey []byte) (bool, error) {
	return _Infrared.Contract.IsInfraredValidator(&_Infrared.CallOpts, _pubkey)
}

// IsInfraredValidator is a free data retrieval call binding the contract method 0xbab97c7b.
//
// Solidity: function isInfraredValidator(bytes _pubkey) view returns(bool)
func (_Infrared *InfraredCallerSession) IsInfraredValidator(_pubkey []byte) (bool, error) {
	return _Infrared.Contract.IsInfraredValidator(&_Infrared.CallOpts, _pubkey)
}

// NumInfraredValidators is a free data retrieval call binding the contract method 0xaf34596b.
//
// Solidity: function numInfraredValidators() view returns(uint256)
func (_Infrared *InfraredCaller) NumInfraredValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "numInfraredValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumInfraredValidators is a free data retrieval call binding the contract method 0xaf34596b.
//
// Solidity: function numInfraredValidators() view returns(uint256)
func (_Infrared *InfraredSession) NumInfraredValidators() (*big.Int, error) {
	return _Infrared.Contract.NumInfraredValidators(&_Infrared.CallOpts)
}

// NumInfraredValidators is a free data retrieval call binding the contract method 0xaf34596b.
//
// Solidity: function numInfraredValidators() view returns(uint256)
func (_Infrared *InfraredCallerSession) NumInfraredValidators() (*big.Int, error) {
	return _Infrared.Contract.NumInfraredValidators(&_Infrared.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Infrared *InfraredCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Infrared *InfraredSession) Paused() (bool, error) {
	return _Infrared.Contract.Paused(&_Infrared.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Infrared *InfraredCallerSession) Paused() (bool, error) {
	return _Infrared.Contract.Paused(&_Infrared.CallOpts)
}

// ProtocolFeeAmounts is a free data retrieval call binding the contract method 0x1f89b683.
//
// Solidity: function protocolFeeAmounts(address _token) view returns(uint256)
func (_Infrared *InfraredCaller) ProtocolFeeAmounts(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "protocolFeeAmounts", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeAmounts is a free data retrieval call binding the contract method 0x1f89b683.
//
// Solidity: function protocolFeeAmounts(address _token) view returns(uint256)
func (_Infrared *InfraredSession) ProtocolFeeAmounts(_token common.Address) (*big.Int, error) {
	return _Infrared.Contract.ProtocolFeeAmounts(&_Infrared.CallOpts, _token)
}

// ProtocolFeeAmounts is a free data retrieval call binding the contract method 0x1f89b683.
//
// Solidity: function protocolFeeAmounts(address _token) view returns(uint256)
func (_Infrared *InfraredCallerSession) ProtocolFeeAmounts(_token common.Address) (*big.Int, error) {
	return _Infrared.Contract.ProtocolFeeAmounts(&_Infrared.CallOpts, _token)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Infrared *InfraredCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Infrared *InfraredSession) ProxiableUUID() ([32]byte, error) {
	return _Infrared.Contract.ProxiableUUID(&_Infrared.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Infrared *InfraredCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Infrared.Contract.ProxiableUUID(&_Infrared.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256 duration)
func (_Infrared *InfraredCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256 duration)
func (_Infrared *InfraredSession) RewardsDuration() (*big.Int, error) {
	return _Infrared.Contract.RewardsDuration(&_Infrared.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256 duration)
func (_Infrared *InfraredCallerSession) RewardsDuration() (*big.Int, error) {
	return _Infrared.Contract.RewardsDuration(&_Infrared.CallOpts)
}

// RewardsFactory is a free data retrieval call binding the contract method 0x51d08944.
//
// Solidity: function rewardsFactory() view returns(address)
func (_Infrared *InfraredCaller) RewardsFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "rewardsFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsFactory is a free data retrieval call binding the contract method 0x51d08944.
//
// Solidity: function rewardsFactory() view returns(address)
func (_Infrared *InfraredSession) RewardsFactory() (common.Address, error) {
	return _Infrared.Contract.RewardsFactory(&_Infrared.CallOpts)
}

// RewardsFactory is a free data retrieval call binding the contract method 0x51d08944.
//
// Solidity: function rewardsFactory() view returns(address)
func (_Infrared *InfraredCallerSession) RewardsFactory() (common.Address, error) {
	return _Infrared.Contract.RewardsFactory(&_Infrared.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Infrared *InfraredCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Infrared *InfraredSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Infrared.Contract.SupportsInterface(&_Infrared.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Infrared *InfraredCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Infrared.Contract.SupportsInterface(&_Infrared.CallOpts, interfaceId)
}

// VaultRegistry is a free data retrieval call binding the contract method 0x5487beb6.
//
// Solidity: function vaultRegistry(address _stakingToken) view returns(address vault)
func (_Infrared *InfraredCaller) VaultRegistry(opts *bind.CallOpts, _stakingToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "vaultRegistry", _stakingToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultRegistry is a free data retrieval call binding the contract method 0x5487beb6.
//
// Solidity: function vaultRegistry(address _stakingToken) view returns(address vault)
func (_Infrared *InfraredSession) VaultRegistry(_stakingToken common.Address) (common.Address, error) {
	return _Infrared.Contract.VaultRegistry(&_Infrared.CallOpts, _stakingToken)
}

// VaultRegistry is a free data retrieval call binding the contract method 0x5487beb6.
//
// Solidity: function vaultRegistry(address _stakingToken) view returns(address vault)
func (_Infrared *InfraredCallerSession) VaultRegistry(_stakingToken common.Address) (common.Address, error) {
	return _Infrared.Contract.VaultRegistry(&_Infrared.CallOpts, _stakingToken)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Infrared *InfraredCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Infrared *InfraredSession) Voter() (common.Address, error) {
	return _Infrared.Contract.Voter(&_Infrared.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_Infrared *InfraredCallerSession) Voter() (common.Address, error) {
	return _Infrared.Contract.Voter(&_Infrared.CallOpts)
}

// Wbera is a free data retrieval call binding the contract method 0x31f41a33.
//
// Solidity: function wbera() view returns(address)
func (_Infrared *InfraredCaller) Wbera(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "wbera")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wbera is a free data retrieval call binding the contract method 0x31f41a33.
//
// Solidity: function wbera() view returns(address)
func (_Infrared *InfraredSession) Wbera() (common.Address, error) {
	return _Infrared.Contract.Wbera(&_Infrared.CallOpts)
}

// Wbera is a free data retrieval call binding the contract method 0x31f41a33.
//
// Solidity: function wbera() view returns(address)
func (_Infrared *InfraredCallerSession) Wbera() (common.Address, error) {
	return _Infrared.Contract.Wbera(&_Infrared.CallOpts)
}

// WhitelistedRewardTokens is a free data retrieval call binding the contract method 0x5225f987.
//
// Solidity: function whitelistedRewardTokens(address token) view returns(bool)
func (_Infrared *InfraredCaller) WhitelistedRewardTokens(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Infrared.contract.Call(opts, &out, "whitelistedRewardTokens", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedRewardTokens is a free data retrieval call binding the contract method 0x5225f987.
//
// Solidity: function whitelistedRewardTokens(address token) view returns(bool)
func (_Infrared *InfraredSession) WhitelistedRewardTokens(token common.Address) (bool, error) {
	return _Infrared.Contract.WhitelistedRewardTokens(&_Infrared.CallOpts, token)
}

// WhitelistedRewardTokens is a free data retrieval call binding the contract method 0x5225f987.
//
// Solidity: function whitelistedRewardTokens(address token) view returns(bool)
func (_Infrared *InfraredCallerSession) WhitelistedRewardTokens(token common.Address) (bool, error) {
	return _Infrared.Contract.WhitelistedRewardTokens(&_Infrared.CallOpts, token)
}

// ActivateBoosts is a paid mutator transaction binding the contract method 0x93857df5.
//
// Solidity: function activateBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactor) ActivateBoosts(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "activateBoosts", _pubkeys)
}

// ActivateBoosts is a paid mutator transaction binding the contract method 0x93857df5.
//
// Solidity: function activateBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredSession) ActivateBoosts(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.ActivateBoosts(&_Infrared.TransactOpts, _pubkeys)
}

// ActivateBoosts is a paid mutator transaction binding the contract method 0x93857df5.
//
// Solidity: function activateBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactorSession) ActivateBoosts(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.ActivateBoosts(&_Infrared.TransactOpts, _pubkeys)
}

// ActivateQueuedValCommission is a paid mutator transaction binding the contract method 0x37e4c5bf.
//
// Solidity: function activateQueuedValCommission(bytes _pubkey) returns()
func (_Infrared *InfraredTransactor) ActivateQueuedValCommission(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "activateQueuedValCommission", _pubkey)
}

// ActivateQueuedValCommission is a paid mutator transaction binding the contract method 0x37e4c5bf.
//
// Solidity: function activateQueuedValCommission(bytes _pubkey) returns()
func (_Infrared *InfraredSession) ActivateQueuedValCommission(_pubkey []byte) (*types.Transaction, error) {
	return _Infrared.Contract.ActivateQueuedValCommission(&_Infrared.TransactOpts, _pubkey)
}

// ActivateQueuedValCommission is a paid mutator transaction binding the contract method 0x37e4c5bf.
//
// Solidity: function activateQueuedValCommission(bytes _pubkey) returns()
func (_Infrared *InfraredTransactorSession) ActivateQueuedValCommission(_pubkey []byte) (*types.Transaction, error) {
	return _Infrared.Contract.ActivateQueuedValCommission(&_Infrared.TransactOpts, _pubkey)
}

// AddIncentives is a paid mutator transaction binding the contract method 0x1b75e798.
//
// Solidity: function addIncentives(address _stakingToken, address _rewardsToken, uint256 _amount) returns()
func (_Infrared *InfraredTransactor) AddIncentives(opts *bind.TransactOpts, _stakingToken common.Address, _rewardsToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "addIncentives", _stakingToken, _rewardsToken, _amount)
}

// AddIncentives is a paid mutator transaction binding the contract method 0x1b75e798.
//
// Solidity: function addIncentives(address _stakingToken, address _rewardsToken, uint256 _amount) returns()
func (_Infrared *InfraredSession) AddIncentives(_stakingToken common.Address, _rewardsToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.AddIncentives(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _amount)
}

// AddIncentives is a paid mutator transaction binding the contract method 0x1b75e798.
//
// Solidity: function addIncentives(address _stakingToken, address _rewardsToken, uint256 _amount) returns()
func (_Infrared *InfraredTransactorSession) AddIncentives(_stakingToken common.Address, _rewardsToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.AddIncentives(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _amount)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactor) AddReward(opts *bind.TransactOpts, _stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "addReward", _stakingToken, _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredSession) AddReward(_stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.AddReward(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0xd0ed26ae.
//
// Solidity: function addReward(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactorSession) AddReward(_stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.AddReward(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _rewardsDuration)
}

// AddValidators is a paid mutator transaction binding the contract method 0x062b79fb.
//
// Solidity: function addValidators((bytes,address)[] _validators) returns()
func (_Infrared *InfraredTransactor) AddValidators(opts *bind.TransactOpts, _validators []ValidatorTypesValidator) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "addValidators", _validators)
}

// AddValidators is a paid mutator transaction binding the contract method 0x062b79fb.
//
// Solidity: function addValidators((bytes,address)[] _validators) returns()
func (_Infrared *InfraredSession) AddValidators(_validators []ValidatorTypesValidator) (*types.Transaction, error) {
	return _Infrared.Contract.AddValidators(&_Infrared.TransactOpts, _validators)
}

// AddValidators is a paid mutator transaction binding the contract method 0x062b79fb.
//
// Solidity: function addValidators((bytes,address)[] _validators) returns()
func (_Infrared *InfraredTransactorSession) AddValidators(_validators []ValidatorTypesValidator) (*types.Transaction, error) {
	return _Infrared.Contract.AddValidators(&_Infrared.TransactOpts, _validators)
}

// CancelBoosts is a paid mutator transaction binding the contract method 0xc277a2d6.
//
// Solidity: function cancelBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactor) CancelBoosts(opts *bind.TransactOpts, _pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "cancelBoosts", _pubkeys, _amts)
}

// CancelBoosts is a paid mutator transaction binding the contract method 0xc277a2d6.
//
// Solidity: function cancelBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredSession) CancelBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CancelBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// CancelBoosts is a paid mutator transaction binding the contract method 0xc277a2d6.
//
// Solidity: function cancelBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactorSession) CancelBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CancelBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// CancelDropBoosts is a paid mutator transaction binding the contract method 0xf667b703.
//
// Solidity: function cancelDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactor) CancelDropBoosts(opts *bind.TransactOpts, _pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "cancelDropBoosts", _pubkeys, _amts)
}

// CancelDropBoosts is a paid mutator transaction binding the contract method 0xf667b703.
//
// Solidity: function cancelDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredSession) CancelDropBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CancelDropBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// CancelDropBoosts is a paid mutator transaction binding the contract method 0xf667b703.
//
// Solidity: function cancelDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactorSession) CancelDropBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CancelDropBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// ClaimLostRewardsOnVault is a paid mutator transaction binding the contract method 0xdcaae5c7.
//
// Solidity: function claimLostRewardsOnVault(address _asset) returns()
func (_Infrared *InfraredTransactor) ClaimLostRewardsOnVault(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "claimLostRewardsOnVault", _asset)
}

// ClaimLostRewardsOnVault is a paid mutator transaction binding the contract method 0xdcaae5c7.
//
// Solidity: function claimLostRewardsOnVault(address _asset) returns()
func (_Infrared *InfraredSession) ClaimLostRewardsOnVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.ClaimLostRewardsOnVault(&_Infrared.TransactOpts, _asset)
}

// ClaimLostRewardsOnVault is a paid mutator transaction binding the contract method 0xdcaae5c7.
//
// Solidity: function claimLostRewardsOnVault(address _asset) returns()
func (_Infrared *InfraredTransactorSession) ClaimLostRewardsOnVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.ClaimLostRewardsOnVault(&_Infrared.TransactOpts, _asset)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x77c654a3.
//
// Solidity: function claimProtocolFees(address _to, address _token) returns()
func (_Infrared *InfraredTransactor) ClaimProtocolFees(opts *bind.TransactOpts, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "claimProtocolFees", _to, _token)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x77c654a3.
//
// Solidity: function claimProtocolFees(address _to, address _token) returns()
func (_Infrared *InfraredSession) ClaimProtocolFees(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.ClaimProtocolFees(&_Infrared.TransactOpts, _to, _token)
}

// ClaimProtocolFees is a paid mutator transaction binding the contract method 0x77c654a3.
//
// Solidity: function claimProtocolFees(address _to, address _token) returns()
func (_Infrared *InfraredTransactorSession) ClaimProtocolFees(_to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.ClaimProtocolFees(&_Infrared.TransactOpts, _to, _token)
}

// CollectBribes is a paid mutator transaction binding the contract method 0xd9ea2d7d.
//
// Solidity: function collectBribes(address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactor) CollectBribes(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "collectBribes", _token, _amount)
}

// CollectBribes is a paid mutator transaction binding the contract method 0xd9ea2d7d.
//
// Solidity: function collectBribes(address _token, uint256 _amount) returns()
func (_Infrared *InfraredSession) CollectBribes(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CollectBribes(&_Infrared.TransactOpts, _token, _amount)
}

// CollectBribes is a paid mutator transaction binding the contract method 0xd9ea2d7d.
//
// Solidity: function collectBribes(address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactorSession) CollectBribes(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.CollectBribes(&_Infrared.TransactOpts, _token, _amount)
}

// DelegateBGT is a paid mutator transaction binding the contract method 0x5137bbdf.
//
// Solidity: function delegateBGT(address _delegatee) returns()
func (_Infrared *InfraredTransactor) DelegateBGT(opts *bind.TransactOpts, _delegatee common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "delegateBGT", _delegatee)
}

// DelegateBGT is a paid mutator transaction binding the contract method 0x5137bbdf.
//
// Solidity: function delegateBGT(address _delegatee) returns()
func (_Infrared *InfraredSession) DelegateBGT(_delegatee common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.DelegateBGT(&_Infrared.TransactOpts, _delegatee)
}

// DelegateBGT is a paid mutator transaction binding the contract method 0x5137bbdf.
//
// Solidity: function delegateBGT(address _delegatee) returns()
func (_Infrared *InfraredTransactorSession) DelegateBGT(_delegatee common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.DelegateBGT(&_Infrared.TransactOpts, _delegatee)
}

// DropBoosts is a paid mutator transaction binding the contract method 0x3f1b754e.
//
// Solidity: function dropBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactor) DropBoosts(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "dropBoosts", _pubkeys)
}

// DropBoosts is a paid mutator transaction binding the contract method 0x3f1b754e.
//
// Solidity: function dropBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredSession) DropBoosts(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.DropBoosts(&_Infrared.TransactOpts, _pubkeys)
}

// DropBoosts is a paid mutator transaction binding the contract method 0x3f1b754e.
//
// Solidity: function dropBoosts(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactorSession) DropBoosts(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.DropBoosts(&_Infrared.TransactOpts, _pubkeys)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Infrared *InfraredTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Infrared *InfraredSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.GrantRole(&_Infrared.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Infrared *InfraredTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.GrantRole(&_Infrared.TransactOpts, role, account)
}

// HarvestBase is a paid mutator transaction binding the contract method 0xa3aad78f.
//
// Solidity: function harvestBase() returns()
func (_Infrared *InfraredTransactor) HarvestBase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestBase")
}

// HarvestBase is a paid mutator transaction binding the contract method 0xa3aad78f.
//
// Solidity: function harvestBase() returns()
func (_Infrared *InfraredSession) HarvestBase() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBase(&_Infrared.TransactOpts)
}

// HarvestBase is a paid mutator transaction binding the contract method 0xa3aad78f.
//
// Solidity: function harvestBase() returns()
func (_Infrared *InfraredTransactorSession) HarvestBase() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBase(&_Infrared.TransactOpts)
}

// HarvestBoostRewards is a paid mutator transaction binding the contract method 0xaf1c921d.
//
// Solidity: function harvestBoostRewards() returns()
func (_Infrared *InfraredTransactor) HarvestBoostRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestBoostRewards")
}

// HarvestBoostRewards is a paid mutator transaction binding the contract method 0xaf1c921d.
//
// Solidity: function harvestBoostRewards() returns()
func (_Infrared *InfraredSession) HarvestBoostRewards() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBoostRewards(&_Infrared.TransactOpts)
}

// HarvestBoostRewards is a paid mutator transaction binding the contract method 0xaf1c921d.
//
// Solidity: function harvestBoostRewards() returns()
func (_Infrared *InfraredTransactorSession) HarvestBoostRewards() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBoostRewards(&_Infrared.TransactOpts)
}

// HarvestBribes is a paid mutator transaction binding the contract method 0xe3a65fc1.
//
// Solidity: function harvestBribes(address[] _tokens) returns()
func (_Infrared *InfraredTransactor) HarvestBribes(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestBribes", _tokens)
}

// HarvestBribes is a paid mutator transaction binding the contract method 0xe3a65fc1.
//
// Solidity: function harvestBribes(address[] _tokens) returns()
func (_Infrared *InfraredSession) HarvestBribes(_tokens []common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBribes(&_Infrared.TransactOpts, _tokens)
}

// HarvestBribes is a paid mutator transaction binding the contract method 0xe3a65fc1.
//
// Solidity: function harvestBribes(address[] _tokens) returns()
func (_Infrared *InfraredTransactorSession) HarvestBribes(_tokens []common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestBribes(&_Infrared.TransactOpts, _tokens)
}

// HarvestOldVault is a paid mutator transaction binding the contract method 0x02f7ecb3.
//
// Solidity: function harvestOldVault(address _vault, address _asset) returns()
func (_Infrared *InfraredTransactor) HarvestOldVault(opts *bind.TransactOpts, _vault common.Address, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestOldVault", _vault, _asset)
}

// HarvestOldVault is a paid mutator transaction binding the contract method 0x02f7ecb3.
//
// Solidity: function harvestOldVault(address _vault, address _asset) returns()
func (_Infrared *InfraredSession) HarvestOldVault(_vault common.Address, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestOldVault(&_Infrared.TransactOpts, _vault, _asset)
}

// HarvestOldVault is a paid mutator transaction binding the contract method 0x02f7ecb3.
//
// Solidity: function harvestOldVault(address _vault, address _asset) returns()
func (_Infrared *InfraredTransactorSession) HarvestOldVault(_vault common.Address, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestOldVault(&_Infrared.TransactOpts, _vault, _asset)
}

// HarvestOperatorRewards is a paid mutator transaction binding the contract method 0x549e71d0.
//
// Solidity: function harvestOperatorRewards() returns()
func (_Infrared *InfraredTransactor) HarvestOperatorRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestOperatorRewards")
}

// HarvestOperatorRewards is a paid mutator transaction binding the contract method 0x549e71d0.
//
// Solidity: function harvestOperatorRewards() returns()
func (_Infrared *InfraredSession) HarvestOperatorRewards() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestOperatorRewards(&_Infrared.TransactOpts)
}

// HarvestOperatorRewards is a paid mutator transaction binding the contract method 0x549e71d0.
//
// Solidity: function harvestOperatorRewards() returns()
func (_Infrared *InfraredTransactorSession) HarvestOperatorRewards() (*types.Transaction, error) {
	return _Infrared.Contract.HarvestOperatorRewards(&_Infrared.TransactOpts)
}

// HarvestVault is a paid mutator transaction binding the contract method 0x0a2f023e.
//
// Solidity: function harvestVault(address _asset) returns()
func (_Infrared *InfraredTransactor) HarvestVault(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "harvestVault", _asset)
}

// HarvestVault is a paid mutator transaction binding the contract method 0x0a2f023e.
//
// Solidity: function harvestVault(address _asset) returns()
func (_Infrared *InfraredSession) HarvestVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestVault(&_Infrared.TransactOpts, _asset)
}

// HarvestVault is a paid mutator transaction binding the contract method 0x0a2f023e.
//
// Solidity: function harvestVault(address _asset) returns()
func (_Infrared *InfraredTransactorSession) HarvestVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.HarvestVault(&_Infrared.TransactOpts, _asset)
}

// InitializeV12 is a paid mutator transaction binding the contract method 0x317e7dfe.
//
// Solidity: function initializeV1_2(address[] _stakingTokens) returns()
func (_Infrared *InfraredTransactor) InitializeV12(opts *bind.TransactOpts, _stakingTokens []common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "initializeV1_2", _stakingTokens)
}

// InitializeV12 is a paid mutator transaction binding the contract method 0x317e7dfe.
//
// Solidity: function initializeV1_2(address[] _stakingTokens) returns()
func (_Infrared *InfraredSession) InitializeV12(_stakingTokens []common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.InitializeV12(&_Infrared.TransactOpts, _stakingTokens)
}

// InitializeV12 is a paid mutator transaction binding the contract method 0x317e7dfe.
//
// Solidity: function initializeV1_2(address[] _stakingTokens) returns()
func (_Infrared *InfraredTransactorSession) InitializeV12(_stakingTokens []common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.InitializeV12(&_Infrared.TransactOpts, _stakingTokens)
}

// MigrateVault is a paid mutator transaction binding the contract method 0xd060af2b.
//
// Solidity: function migrateVault(address _asset, uint8 versionToUpgradeTo) returns(address newVault)
func (_Infrared *InfraredTransactor) MigrateVault(opts *bind.TransactOpts, _asset common.Address, versionToUpgradeTo uint8) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "migrateVault", _asset, versionToUpgradeTo)
}

// MigrateVault is a paid mutator transaction binding the contract method 0xd060af2b.
//
// Solidity: function migrateVault(address _asset, uint8 versionToUpgradeTo) returns(address newVault)
func (_Infrared *InfraredSession) MigrateVault(_asset common.Address, versionToUpgradeTo uint8) (*types.Transaction, error) {
	return _Infrared.Contract.MigrateVault(&_Infrared.TransactOpts, _asset, versionToUpgradeTo)
}

// MigrateVault is a paid mutator transaction binding the contract method 0xd060af2b.
//
// Solidity: function migrateVault(address _asset, uint8 versionToUpgradeTo) returns(address newVault)
func (_Infrared *InfraredTransactorSession) MigrateVault(_asset common.Address, versionToUpgradeTo uint8) (*types.Transaction, error) {
	return _Infrared.Contract.MigrateVault(&_Infrared.TransactOpts, _asset, versionToUpgradeTo)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Infrared *InfraredTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Infrared *InfraredSession) Pause() (*types.Transaction, error) {
	return _Infrared.Contract.Pause(&_Infrared.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Infrared *InfraredTransactorSession) Pause() (*types.Transaction, error) {
	return _Infrared.Contract.Pause(&_Infrared.TransactOpts)
}

// PauseOldStaking is a paid mutator transaction binding the contract method 0x247e92ab.
//
// Solidity: function pauseOldStaking(address _vault) returns()
func (_Infrared *InfraredTransactor) PauseOldStaking(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "pauseOldStaking", _vault)
}

// PauseOldStaking is a paid mutator transaction binding the contract method 0x247e92ab.
//
// Solidity: function pauseOldStaking(address _vault) returns()
func (_Infrared *InfraredSession) PauseOldStaking(_vault common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.PauseOldStaking(&_Infrared.TransactOpts, _vault)
}

// PauseOldStaking is a paid mutator transaction binding the contract method 0x247e92ab.
//
// Solidity: function pauseOldStaking(address _vault) returns()
func (_Infrared *InfraredTransactorSession) PauseOldStaking(_vault common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.PauseOldStaking(&_Infrared.TransactOpts, _vault)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xef8cf92d.
//
// Solidity: function pauseStaking(address _asset) returns()
func (_Infrared *InfraredTransactor) PauseStaking(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "pauseStaking", _asset)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xef8cf92d.
//
// Solidity: function pauseStaking(address _asset) returns()
func (_Infrared *InfraredSession) PauseStaking(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.PauseStaking(&_Infrared.TransactOpts, _asset)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xef8cf92d.
//
// Solidity: function pauseStaking(address _asset) returns()
func (_Infrared *InfraredTransactorSession) PauseStaking(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.PauseStaking(&_Infrared.TransactOpts, _asset)
}

// QueueBoosts is a paid mutator transaction binding the contract method 0x60930183.
//
// Solidity: function queueBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactor) QueueBoosts(opts *bind.TransactOpts, _pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "queueBoosts", _pubkeys, _amts)
}

// QueueBoosts is a paid mutator transaction binding the contract method 0x60930183.
//
// Solidity: function queueBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredSession) QueueBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// QueueBoosts is a paid mutator transaction binding the contract method 0x60930183.
//
// Solidity: function queueBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactorSession) QueueBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// QueueDropBoosts is a paid mutator transaction binding the contract method 0x694be773.
//
// Solidity: function queueDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactor) QueueDropBoosts(opts *bind.TransactOpts, _pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "queueDropBoosts", _pubkeys, _amts)
}

// QueueDropBoosts is a paid mutator transaction binding the contract method 0x694be773.
//
// Solidity: function queueDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredSession) QueueDropBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueDropBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// QueueDropBoosts is a paid mutator transaction binding the contract method 0x694be773.
//
// Solidity: function queueDropBoosts(bytes[] _pubkeys, uint128[] _amts) returns()
func (_Infrared *InfraredTransactorSession) QueueDropBoosts(_pubkeys [][]byte, _amts []*big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueDropBoosts(&_Infrared.TransactOpts, _pubkeys, _amts)
}

// QueueMultipleValCommissions is a paid mutator transaction binding the contract method 0x268cadbe.
//
// Solidity: function queueMultipleValCommissions(bytes[] _pubkeys, uint96 _commissionRate) returns()
func (_Infrared *InfraredTransactor) QueueMultipleValCommissions(opts *bind.TransactOpts, _pubkeys [][]byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "queueMultipleValCommissions", _pubkeys, _commissionRate)
}

// QueueMultipleValCommissions is a paid mutator transaction binding the contract method 0x268cadbe.
//
// Solidity: function queueMultipleValCommissions(bytes[] _pubkeys, uint96 _commissionRate) returns()
func (_Infrared *InfraredSession) QueueMultipleValCommissions(_pubkeys [][]byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueMultipleValCommissions(&_Infrared.TransactOpts, _pubkeys, _commissionRate)
}

// QueueMultipleValCommissions is a paid mutator transaction binding the contract method 0x268cadbe.
//
// Solidity: function queueMultipleValCommissions(bytes[] _pubkeys, uint96 _commissionRate) returns()
func (_Infrared *InfraredTransactorSession) QueueMultipleValCommissions(_pubkeys [][]byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueMultipleValCommissions(&_Infrared.TransactOpts, _pubkeys, _commissionRate)
}

// QueueNewCuttingBoard is a paid mutator transaction binding the contract method 0xa66ce82d.
//
// Solidity: function queueNewCuttingBoard(bytes _pubkey, uint64 _startBlock, (address,uint96)[] _weights) returns()
func (_Infrared *InfraredTransactor) QueueNewCuttingBoard(opts *bind.TransactOpts, _pubkey []byte, _startBlock uint64, _weights []IBeraChefWeight) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "queueNewCuttingBoard", _pubkey, _startBlock, _weights)
}

// QueueNewCuttingBoard is a paid mutator transaction binding the contract method 0xa66ce82d.
//
// Solidity: function queueNewCuttingBoard(bytes _pubkey, uint64 _startBlock, (address,uint96)[] _weights) returns()
func (_Infrared *InfraredSession) QueueNewCuttingBoard(_pubkey []byte, _startBlock uint64, _weights []IBeraChefWeight) (*types.Transaction, error) {
	return _Infrared.Contract.QueueNewCuttingBoard(&_Infrared.TransactOpts, _pubkey, _startBlock, _weights)
}

// QueueNewCuttingBoard is a paid mutator transaction binding the contract method 0xa66ce82d.
//
// Solidity: function queueNewCuttingBoard(bytes _pubkey, uint64 _startBlock, (address,uint96)[] _weights) returns()
func (_Infrared *InfraredTransactorSession) QueueNewCuttingBoard(_pubkey []byte, _startBlock uint64, _weights []IBeraChefWeight) (*types.Transaction, error) {
	return _Infrared.Contract.QueueNewCuttingBoard(&_Infrared.TransactOpts, _pubkey, _startBlock, _weights)
}

// QueueValCommission is a paid mutator transaction binding the contract method 0x8f7e0f0b.
//
// Solidity: function queueValCommission(bytes _pubkey, uint96 _commissionRate) returns()
func (_Infrared *InfraredTransactor) QueueValCommission(opts *bind.TransactOpts, _pubkey []byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "queueValCommission", _pubkey, _commissionRate)
}

// QueueValCommission is a paid mutator transaction binding the contract method 0x8f7e0f0b.
//
// Solidity: function queueValCommission(bytes _pubkey, uint96 _commissionRate) returns()
func (_Infrared *InfraredSession) QueueValCommission(_pubkey []byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueValCommission(&_Infrared.TransactOpts, _pubkey, _commissionRate)
}

// QueueValCommission is a paid mutator transaction binding the contract method 0x8f7e0f0b.
//
// Solidity: function queueValCommission(bytes _pubkey, uint96 _commissionRate) returns()
func (_Infrared *InfraredTransactorSession) QueueValCommission(_pubkey []byte, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.QueueValCommission(&_Infrared.TransactOpts, _pubkey, _commissionRate)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactor) RecoverERC20(opts *bind.TransactOpts, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "recoverERC20", _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20(&_Infrared.TransactOpts, _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactorSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20(&_Infrared.TransactOpts, _to, _token, _amount)
}

// RecoverERC20FromOldVault is a paid mutator transaction binding the contract method 0x4f2bf2f1.
//
// Solidity: function recoverERC20FromOldVault(address _vault, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactor) RecoverERC20FromOldVault(opts *bind.TransactOpts, _vault common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "recoverERC20FromOldVault", _vault, _to, _token, _amount)
}

// RecoverERC20FromOldVault is a paid mutator transaction binding the contract method 0x4f2bf2f1.
//
// Solidity: function recoverERC20FromOldVault(address _vault, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredSession) RecoverERC20FromOldVault(_vault common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20FromOldVault(&_Infrared.TransactOpts, _vault, _to, _token, _amount)
}

// RecoverERC20FromOldVault is a paid mutator transaction binding the contract method 0x4f2bf2f1.
//
// Solidity: function recoverERC20FromOldVault(address _vault, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactorSession) RecoverERC20FromOldVault(_vault common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20FromOldVault(&_Infrared.TransactOpts, _vault, _to, _token, _amount)
}

// RecoverERC20FromVault is a paid mutator transaction binding the contract method 0xe18cb2c0.
//
// Solidity: function recoverERC20FromVault(address _asset, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactor) RecoverERC20FromVault(opts *bind.TransactOpts, _asset common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "recoverERC20FromVault", _asset, _to, _token, _amount)
}

// RecoverERC20FromVault is a paid mutator transaction binding the contract method 0xe18cb2c0.
//
// Solidity: function recoverERC20FromVault(address _asset, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredSession) RecoverERC20FromVault(_asset common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20FromVault(&_Infrared.TransactOpts, _asset, _to, _token, _amount)
}

// RecoverERC20FromVault is a paid mutator transaction binding the contract method 0xe18cb2c0.
//
// Solidity: function recoverERC20FromVault(address _asset, address _to, address _token, uint256 _amount) returns()
func (_Infrared *InfraredTransactorSession) RecoverERC20FromVault(_asset common.Address, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.RecoverERC20FromVault(&_Infrared.TransactOpts, _asset, _to, _token, _amount)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x67ccd1cb.
//
// Solidity: function registerVault(address _asset) returns(address vault)
func (_Infrared *InfraredTransactor) RegisterVault(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "registerVault", _asset)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x67ccd1cb.
//
// Solidity: function registerVault(address _asset) returns(address vault)
func (_Infrared *InfraredSession) RegisterVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RegisterVault(&_Infrared.TransactOpts, _asset)
}

// RegisterVault is a paid mutator transaction binding the contract method 0x67ccd1cb.
//
// Solidity: function registerVault(address _asset) returns(address vault)
func (_Infrared *InfraredTransactorSession) RegisterVault(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RegisterVault(&_Infrared.TransactOpts, _asset)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x165d21c9.
//
// Solidity: function removeReward(address _stakingToken, address _rewardsToken) returns()
func (_Infrared *InfraredTransactor) RemoveReward(opts *bind.TransactOpts, _stakingToken common.Address, _rewardsToken common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "removeReward", _stakingToken, _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x165d21c9.
//
// Solidity: function removeReward(address _stakingToken, address _rewardsToken) returns()
func (_Infrared *InfraredSession) RemoveReward(_stakingToken common.Address, _rewardsToken common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RemoveReward(&_Infrared.TransactOpts, _stakingToken, _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x165d21c9.
//
// Solidity: function removeReward(address _stakingToken, address _rewardsToken) returns()
func (_Infrared *InfraredTransactorSession) RemoveReward(_stakingToken common.Address, _rewardsToken common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RemoveReward(&_Infrared.TransactOpts, _stakingToken, _rewardsToken)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xd6fc45fd.
//
// Solidity: function removeValidators(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactor) RemoveValidators(opts *bind.TransactOpts, _pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "removeValidators", _pubkeys)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xd6fc45fd.
//
// Solidity: function removeValidators(bytes[] _pubkeys) returns()
func (_Infrared *InfraredSession) RemoveValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.RemoveValidators(&_Infrared.TransactOpts, _pubkeys)
}

// RemoveValidators is a paid mutator transaction binding the contract method 0xd6fc45fd.
//
// Solidity: function removeValidators(bytes[] _pubkeys) returns()
func (_Infrared *InfraredTransactorSession) RemoveValidators(_pubkeys [][]byte) (*types.Transaction, error) {
	return _Infrared.Contract.RemoveValidators(&_Infrared.TransactOpts, _pubkeys)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Infrared *InfraredTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Infrared *InfraredSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RenounceRole(&_Infrared.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Infrared *InfraredTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RenounceRole(&_Infrared.TransactOpts, role, callerConfirmation)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x37f6cfe7.
//
// Solidity: function replaceValidator(bytes _current, bytes _new) returns()
func (_Infrared *InfraredTransactor) ReplaceValidator(opts *bind.TransactOpts, _current []byte, _new []byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "replaceValidator", _current, _new)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x37f6cfe7.
//
// Solidity: function replaceValidator(bytes _current, bytes _new) returns()
func (_Infrared *InfraredSession) ReplaceValidator(_current []byte, _new []byte) (*types.Transaction, error) {
	return _Infrared.Contract.ReplaceValidator(&_Infrared.TransactOpts, _current, _new)
}

// ReplaceValidator is a paid mutator transaction binding the contract method 0x37f6cfe7.
//
// Solidity: function replaceValidator(bytes _current, bytes _new) returns()
func (_Infrared *InfraredTransactorSession) ReplaceValidator(_current []byte, _new []byte) (*types.Transaction, error) {
	return _Infrared.Contract.ReplaceValidator(&_Infrared.TransactOpts, _current, _new)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Infrared *InfraredTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Infrared *InfraredSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RevokeRole(&_Infrared.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Infrared *InfraredTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.RevokeRole(&_Infrared.TransactOpts, role, account)
}

// SetIR is a paid mutator transaction binding the contract method 0x013c5e52.
//
// Solidity: function setIR(address _ir) returns()
func (_Infrared *InfraredTransactor) SetIR(opts *bind.TransactOpts, _ir common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "setIR", _ir)
}

// SetIR is a paid mutator transaction binding the contract method 0x013c5e52.
//
// Solidity: function setIR(address _ir) returns()
func (_Infrared *InfraredSession) SetIR(_ir common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.SetIR(&_Infrared.TransactOpts, _ir)
}

// SetIR is a paid mutator transaction binding the contract method 0x013c5e52.
//
// Solidity: function setIR(address _ir) returns()
func (_Infrared *InfraredTransactorSession) SetIR(_ir common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.SetIR(&_Infrared.TransactOpts, _ir)
}

// SetVaultRegistrationPauseStatus is a paid mutator transaction binding the contract method 0x270ecc22.
//
// Solidity: function setVaultRegistrationPauseStatus(bool pause) returns()
func (_Infrared *InfraredTransactor) SetVaultRegistrationPauseStatus(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "setVaultRegistrationPauseStatus", pause)
}

// SetVaultRegistrationPauseStatus is a paid mutator transaction binding the contract method 0x270ecc22.
//
// Solidity: function setVaultRegistrationPauseStatus(bool pause) returns()
func (_Infrared *InfraredSession) SetVaultRegistrationPauseStatus(pause bool) (*types.Transaction, error) {
	return _Infrared.Contract.SetVaultRegistrationPauseStatus(&_Infrared.TransactOpts, pause)
}

// SetVaultRegistrationPauseStatus is a paid mutator transaction binding the contract method 0x270ecc22.
//
// Solidity: function setVaultRegistrationPauseStatus(bool pause) returns()
func (_Infrared *InfraredTransactorSession) SetVaultRegistrationPauseStatus(pause bool) (*types.Transaction, error) {
	return _Infrared.Contract.SetVaultRegistrationPauseStatus(&_Infrared.TransactOpts, pause)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Infrared *InfraredTransactor) SetVoter(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "setVoter", _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Infrared *InfraredSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.SetVoter(&_Infrared.TransactOpts, _voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address _voter) returns()
func (_Infrared *InfraredTransactorSession) SetVoter(_voter common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.SetVoter(&_Infrared.TransactOpts, _voter)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Infrared *InfraredTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Infrared *InfraredSession) Unpause() (*types.Transaction, error) {
	return _Infrared.Contract.Unpause(&_Infrared.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Infrared *InfraredTransactorSession) Unpause() (*types.Transaction, error) {
	return _Infrared.Contract.Unpause(&_Infrared.TransactOpts)
}

// UnpauseOldStaking is a paid mutator transaction binding the contract method 0xe0de574f.
//
// Solidity: function unpauseOldStaking(address _vault) returns()
func (_Infrared *InfraredTransactor) UnpauseOldStaking(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "unpauseOldStaking", _vault)
}

// UnpauseOldStaking is a paid mutator transaction binding the contract method 0xe0de574f.
//
// Solidity: function unpauseOldStaking(address _vault) returns()
func (_Infrared *InfraredSession) UnpauseOldStaking(_vault common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.UnpauseOldStaking(&_Infrared.TransactOpts, _vault)
}

// UnpauseOldStaking is a paid mutator transaction binding the contract method 0xe0de574f.
//
// Solidity: function unpauseOldStaking(address _vault) returns()
func (_Infrared *InfraredTransactorSession) UnpauseOldStaking(_vault common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.UnpauseOldStaking(&_Infrared.TransactOpts, _vault)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0xc4911bd2.
//
// Solidity: function unpauseStaking(address _asset) returns()
func (_Infrared *InfraredTransactor) UnpauseStaking(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "unpauseStaking", _asset)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0xc4911bd2.
//
// Solidity: function unpauseStaking(address _asset) returns()
func (_Infrared *InfraredSession) UnpauseStaking(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.UnpauseStaking(&_Infrared.TransactOpts, _asset)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0xc4911bd2.
//
// Solidity: function unpauseStaking(address _asset) returns()
func (_Infrared *InfraredTransactorSession) UnpauseStaking(_asset common.Address) (*types.Transaction, error) {
	return _Infrared.Contract.UnpauseStaking(&_Infrared.TransactOpts, _asset)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x7052a5e5.
//
// Solidity: function updateFee(uint8 _t, uint256 _fee) returns()
func (_Infrared *InfraredTransactor) UpdateFee(opts *bind.TransactOpts, _t uint8, _fee *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateFee", _t, _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x7052a5e5.
//
// Solidity: function updateFee(uint8 _t, uint256 _fee) returns()
func (_Infrared *InfraredSession) UpdateFee(_t uint8, _fee *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateFee(&_Infrared.TransactOpts, _t, _fee)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x7052a5e5.
//
// Solidity: function updateFee(uint8 _t, uint256 _fee) returns()
func (_Infrared *InfraredTransactorSession) UpdateFee(_t uint8, _fee *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateFee(&_Infrared.TransactOpts, _t, _fee)
}

// UpdateIRMintRate is a paid mutator transaction binding the contract method 0x243bbfe6.
//
// Solidity: function updateIRMintRate(uint256 _irMintRate) returns()
func (_Infrared *InfraredTransactor) UpdateIRMintRate(opts *bind.TransactOpts, _irMintRate *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateIRMintRate", _irMintRate)
}

// UpdateIRMintRate is a paid mutator transaction binding the contract method 0x243bbfe6.
//
// Solidity: function updateIRMintRate(uint256 _irMintRate) returns()
func (_Infrared *InfraredSession) UpdateIRMintRate(_irMintRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateIRMintRate(&_Infrared.TransactOpts, _irMintRate)
}

// UpdateIRMintRate is a paid mutator transaction binding the contract method 0x243bbfe6.
//
// Solidity: function updateIRMintRate(uint256 _irMintRate) returns()
func (_Infrared *InfraredTransactorSession) UpdateIRMintRate(_irMintRate *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateIRMintRate(&_Infrared.TransactOpts, _irMintRate)
}

// UpdateInfraredBERABribeSplit is a paid mutator transaction binding the contract method 0xd080edc9.
//
// Solidity: function updateInfraredBERABribeSplit(uint256 _weight) returns()
func (_Infrared *InfraredTransactor) UpdateInfraredBERABribeSplit(opts *bind.TransactOpts, _weight *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateInfraredBERABribeSplit", _weight)
}

// UpdateInfraredBERABribeSplit is a paid mutator transaction binding the contract method 0xd080edc9.
//
// Solidity: function updateInfraredBERABribeSplit(uint256 _weight) returns()
func (_Infrared *InfraredSession) UpdateInfraredBERABribeSplit(_weight *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateInfraredBERABribeSplit(&_Infrared.TransactOpts, _weight)
}

// UpdateInfraredBERABribeSplit is a paid mutator transaction binding the contract method 0xd080edc9.
//
// Solidity: function updateInfraredBERABribeSplit(uint256 _weight) returns()
func (_Infrared *InfraredTransactorSession) UpdateInfraredBERABribeSplit(_weight *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateInfraredBERABribeSplit(&_Infrared.TransactOpts, _weight)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xd94ef29d.
//
// Solidity: function updateRewardsDuration(uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactor) UpdateRewardsDuration(opts *bind.TransactOpts, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateRewardsDuration", _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xd94ef29d.
//
// Solidity: function updateRewardsDuration(uint256 _rewardsDuration) returns()
func (_Infrared *InfraredSession) UpdateRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateRewardsDuration(&_Infrared.TransactOpts, _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xd94ef29d.
//
// Solidity: function updateRewardsDuration(uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactorSession) UpdateRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateRewardsDuration(&_Infrared.TransactOpts, _rewardsDuration)
}

// UpdateRewardsDurationForVault is a paid mutator transaction binding the contract method 0x792868bb.
//
// Solidity: function updateRewardsDurationForVault(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactor) UpdateRewardsDurationForVault(opts *bind.TransactOpts, _stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateRewardsDurationForVault", _stakingToken, _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDurationForVault is a paid mutator transaction binding the contract method 0x792868bb.
//
// Solidity: function updateRewardsDurationForVault(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredSession) UpdateRewardsDurationForVault(_stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateRewardsDurationForVault(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDurationForVault is a paid mutator transaction binding the contract method 0x792868bb.
//
// Solidity: function updateRewardsDurationForVault(address _stakingToken, address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Infrared *InfraredTransactorSession) UpdateRewardsDurationForVault(_stakingToken common.Address, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateRewardsDurationForVault(&_Infrared.TransactOpts, _stakingToken, _rewardsToken, _rewardsDuration)
}

// UpdateWhiteListedRewardTokens is a paid mutator transaction binding the contract method 0x5787077d.
//
// Solidity: function updateWhiteListedRewardTokens(address _token, bool _whitelisted) returns()
func (_Infrared *InfraredTransactor) UpdateWhiteListedRewardTokens(opts *bind.TransactOpts, _token common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "updateWhiteListedRewardTokens", _token, _whitelisted)
}

// UpdateWhiteListedRewardTokens is a paid mutator transaction binding the contract method 0x5787077d.
//
// Solidity: function updateWhiteListedRewardTokens(address _token, bool _whitelisted) returns()
func (_Infrared *InfraredSession) UpdateWhiteListedRewardTokens(_token common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateWhiteListedRewardTokens(&_Infrared.TransactOpts, _token, _whitelisted)
}

// UpdateWhiteListedRewardTokens is a paid mutator transaction binding the contract method 0x5787077d.
//
// Solidity: function updateWhiteListedRewardTokens(address _token, bool _whitelisted) returns()
func (_Infrared *InfraredTransactorSession) UpdateWhiteListedRewardTokens(_token common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Infrared.Contract.UpdateWhiteListedRewardTokens(&_Infrared.TransactOpts, _token, _whitelisted)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Infrared *InfraredTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Infrared.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Infrared *InfraredSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Infrared.Contract.UpgradeToAndCall(&_Infrared.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Infrared *InfraredTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Infrared.Contract.UpgradeToAndCall(&_Infrared.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Infrared *InfraredTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Infrared.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Infrared *InfraredSession) Receive() (*types.Transaction, error) {
	return _Infrared.Contract.Receive(&_Infrared.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Infrared *InfraredTransactorSession) Receive() (*types.Transaction, error) {
	return _Infrared.Contract.Receive(&_Infrared.TransactOpts)
}

// InfraredActivatedBoostsIterator is returned from FilterActivatedBoosts and is used to iterate over the raw logs and unpacked data for ActivatedBoosts events raised by the Infrared contract.
type InfraredActivatedBoostsIterator struct {
	Event *InfraredActivatedBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredActivatedBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredActivatedBoosts)
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
		it.Event = new(InfraredActivatedBoosts)
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
func (it *InfraredActivatedBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredActivatedBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredActivatedBoosts represents a ActivatedBoosts event raised by the Infrared contract.
type InfraredActivatedBoosts struct {
	Sender  common.Address
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterActivatedBoosts is a free log retrieval operation binding the contract event 0x369eb69e56e3a11fe35cac617eb5c87feeb40cfc71402e8d0c43b323d88d838a.
//
// Solidity: event ActivatedBoosts(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) FilterActivatedBoosts(opts *bind.FilterOpts) (*InfraredActivatedBoostsIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ActivatedBoosts")
	if err != nil {
		return nil, err
	}
	return &InfraredActivatedBoostsIterator{contract: _Infrared.contract, event: "ActivatedBoosts", logs: logs, sub: sub}, nil
}

// WatchActivatedBoosts is a free log subscription operation binding the contract event 0x369eb69e56e3a11fe35cac617eb5c87feeb40cfc71402e8d0c43b323d88d838a.
//
// Solidity: event ActivatedBoosts(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) WatchActivatedBoosts(opts *bind.WatchOpts, sink chan<- *InfraredActivatedBoosts) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ActivatedBoosts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredActivatedBoosts)
				if err := _Infrared.contract.UnpackLog(event, "ActivatedBoosts", log); err != nil {
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

// ParseActivatedBoosts is a log parse operation binding the contract event 0x369eb69e56e3a11fe35cac617eb5c87feeb40cfc71402e8d0c43b323d88d838a.
//
// Solidity: event ActivatedBoosts(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) ParseActivatedBoosts(log types.Log) (*InfraredActivatedBoosts, error) {
	event := new(InfraredActivatedBoosts)
	if err := _Infrared.contract.UnpackLog(event, "ActivatedBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredBaseHarvestedIterator is returned from FilterBaseHarvested and is used to iterate over the raw logs and unpacked data for BaseHarvested events raised by the Infrared contract.
type InfraredBaseHarvestedIterator struct {
	Event *InfraredBaseHarvested // Event containing the contract specifics and raw log

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
func (it *InfraredBaseHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredBaseHarvested)
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
		it.Event = new(InfraredBaseHarvested)
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
func (it *InfraredBaseHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredBaseHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredBaseHarvested represents a BaseHarvested event raised by the Infrared contract.
type InfraredBaseHarvested struct {
	Sender common.Address
	BgtAmt *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBaseHarvested is a free log retrieval operation binding the contract event 0x631fea7b21a7883939599d132c75a0b4bcb42f2ade466a484cdab4909e1cc46c.
//
// Solidity: event BaseHarvested(address _sender, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) FilterBaseHarvested(opts *bind.FilterOpts) (*InfraredBaseHarvestedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "BaseHarvested")
	if err != nil {
		return nil, err
	}
	return &InfraredBaseHarvestedIterator{contract: _Infrared.contract, event: "BaseHarvested", logs: logs, sub: sub}, nil
}

// WatchBaseHarvested is a free log subscription operation binding the contract event 0x631fea7b21a7883939599d132c75a0b4bcb42f2ade466a484cdab4909e1cc46c.
//
// Solidity: event BaseHarvested(address _sender, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) WatchBaseHarvested(opts *bind.WatchOpts, sink chan<- *InfraredBaseHarvested) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "BaseHarvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredBaseHarvested)
				if err := _Infrared.contract.UnpackLog(event, "BaseHarvested", log); err != nil {
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

// ParseBaseHarvested is a log parse operation binding the contract event 0x631fea7b21a7883939599d132c75a0b4bcb42f2ade466a484cdab4909e1cc46c.
//
// Solidity: event BaseHarvested(address _sender, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) ParseBaseHarvested(log types.Log) (*InfraredBaseHarvested, error) {
	event := new(InfraredBaseHarvested)
	if err := _Infrared.contract.UnpackLog(event, "BaseHarvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredBribeSuppliedIterator is returned from FilterBribeSupplied and is used to iterate over the raw logs and unpacked data for BribeSupplied events raised by the Infrared contract.
type InfraredBribeSuppliedIterator struct {
	Event *InfraredBribeSupplied // Event containing the contract specifics and raw log

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
func (it *InfraredBribeSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredBribeSupplied)
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
		it.Event = new(InfraredBribeSupplied)
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
func (it *InfraredBribeSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredBribeSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredBribeSupplied represents a BribeSupplied event raised by the Infrared contract.
type InfraredBribeSupplied struct {
	Recipient common.Address
	Token     common.Address
	Amt       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBribeSupplied is a free log retrieval operation binding the contract event 0xff3ccffdf851bb4b5a34b7eb3895ff9a15395bd707a9fcc66963b78540eda2dc.
//
// Solidity: event BribeSupplied(address indexed _recipient, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) FilterBribeSupplied(opts *bind.FilterOpts, _recipient []common.Address, _token []common.Address) (*InfraredBribeSuppliedIterator, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "BribeSupplied", _recipientRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredBribeSuppliedIterator{contract: _Infrared.contract, event: "BribeSupplied", logs: logs, sub: sub}, nil
}

// WatchBribeSupplied is a free log subscription operation binding the contract event 0xff3ccffdf851bb4b5a34b7eb3895ff9a15395bd707a9fcc66963b78540eda2dc.
//
// Solidity: event BribeSupplied(address indexed _recipient, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) WatchBribeSupplied(opts *bind.WatchOpts, sink chan<- *InfraredBribeSupplied, _recipient []common.Address, _token []common.Address) (event.Subscription, error) {

	var _recipientRule []interface{}
	for _, _recipientItem := range _recipient {
		_recipientRule = append(_recipientRule, _recipientItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "BribeSupplied", _recipientRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredBribeSupplied)
				if err := _Infrared.contract.UnpackLog(event, "BribeSupplied", log); err != nil {
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

// ParseBribeSupplied is a log parse operation binding the contract event 0xff3ccffdf851bb4b5a34b7eb3895ff9a15395bd707a9fcc66963b78540eda2dc.
//
// Solidity: event BribeSupplied(address indexed _recipient, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) ParseBribeSupplied(log types.Log) (*InfraredBribeSupplied, error) {
	event := new(InfraredBribeSupplied)
	if err := _Infrared.contract.UnpackLog(event, "BribeSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredBribesCollectedIterator is returned from FilterBribesCollected and is used to iterate over the raw logs and unpacked data for BribesCollected events raised by the Infrared contract.
type InfraredBribesCollectedIterator struct {
	Event *InfraredBribesCollected // Event containing the contract specifics and raw log

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
func (it *InfraredBribesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredBribesCollected)
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
		it.Event = new(InfraredBribesCollected)
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
func (it *InfraredBribesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredBribesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredBribesCollected represents a BribesCollected event raised by the Infrared contract.
type InfraredBribesCollected struct {
	Sender         common.Address
	Token          common.Address
	AmtWiberaVault *big.Int
	AmtIbgtVault   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBribesCollected is a free log retrieval operation binding the contract event 0xfe24b5984a4e07d71fc93f1efdc6bd554af25e0d759757a0f71f90e2af67e2bd.
//
// Solidity: event BribesCollected(address _sender, address _token, uint256 _amtWiberaVault, uint256 _amtIbgtVault)
func (_Infrared *InfraredFilterer) FilterBribesCollected(opts *bind.FilterOpts) (*InfraredBribesCollectedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "BribesCollected")
	if err != nil {
		return nil, err
	}
	return &InfraredBribesCollectedIterator{contract: _Infrared.contract, event: "BribesCollected", logs: logs, sub: sub}, nil
}

// WatchBribesCollected is a free log subscription operation binding the contract event 0xfe24b5984a4e07d71fc93f1efdc6bd554af25e0d759757a0f71f90e2af67e2bd.
//
// Solidity: event BribesCollected(address _sender, address _token, uint256 _amtWiberaVault, uint256 _amtIbgtVault)
func (_Infrared *InfraredFilterer) WatchBribesCollected(opts *bind.WatchOpts, sink chan<- *InfraredBribesCollected) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "BribesCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredBribesCollected)
				if err := _Infrared.contract.UnpackLog(event, "BribesCollected", log); err != nil {
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

// ParseBribesCollected is a log parse operation binding the contract event 0xfe24b5984a4e07d71fc93f1efdc6bd554af25e0d759757a0f71f90e2af67e2bd.
//
// Solidity: event BribesCollected(address _sender, address _token, uint256 _amtWiberaVault, uint256 _amtIbgtVault)
func (_Infrared *InfraredFilterer) ParseBribesCollected(log types.Log) (*InfraredBribesCollected, error) {
	event := new(InfraredBribesCollected)
	if err := _Infrared.contract.UnpackLog(event, "BribesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredCancelDropBoostsIterator is returned from FilterCancelDropBoosts and is used to iterate over the raw logs and unpacked data for CancelDropBoosts events raised by the Infrared contract.
type InfraredCancelDropBoostsIterator struct {
	Event *InfraredCancelDropBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredCancelDropBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredCancelDropBoosts)
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
		it.Event = new(InfraredCancelDropBoosts)
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
func (it *InfraredCancelDropBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredCancelDropBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredCancelDropBoosts represents a CancelDropBoosts event raised by the Infrared contract.
type InfraredCancelDropBoosts struct {
	User    common.Address
	Pubkeys [][]byte
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelDropBoosts is a free log retrieval operation binding the contract event 0xe407367666886c3ff8f6ae8d450e5abdac1688f7b06e1b22c1b6be93d112f2ee.
//
// Solidity: event CancelDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) FilterCancelDropBoosts(opts *bind.FilterOpts, user []common.Address, pubkeys [][][]byte) (*InfraredCancelDropBoostsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pubkeysRule []interface{}
	for _, pubkeysItem := range pubkeys {
		pubkeysRule = append(pubkeysRule, pubkeysItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "CancelDropBoosts", userRule, pubkeysRule)
	if err != nil {
		return nil, err
	}
	return &InfraredCancelDropBoostsIterator{contract: _Infrared.contract, event: "CancelDropBoosts", logs: logs, sub: sub}, nil
}

// WatchCancelDropBoosts is a free log subscription operation binding the contract event 0xe407367666886c3ff8f6ae8d450e5abdac1688f7b06e1b22c1b6be93d112f2ee.
//
// Solidity: event CancelDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) WatchCancelDropBoosts(opts *bind.WatchOpts, sink chan<- *InfraredCancelDropBoosts, user []common.Address, pubkeys [][][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pubkeysRule []interface{}
	for _, pubkeysItem := range pubkeys {
		pubkeysRule = append(pubkeysRule, pubkeysItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "CancelDropBoosts", userRule, pubkeysRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredCancelDropBoosts)
				if err := _Infrared.contract.UnpackLog(event, "CancelDropBoosts", log); err != nil {
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

// ParseCancelDropBoosts is a log parse operation binding the contract event 0xe407367666886c3ff8f6ae8d450e5abdac1688f7b06e1b22c1b6be93d112f2ee.
//
// Solidity: event CancelDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) ParseCancelDropBoosts(log types.Log) (*InfraredCancelDropBoosts, error) {
	event := new(InfraredCancelDropBoosts)
	if err := _Infrared.contract.UnpackLog(event, "CancelDropBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredCancelledBoostsIterator is returned from FilterCancelledBoosts and is used to iterate over the raw logs and unpacked data for CancelledBoosts events raised by the Infrared contract.
type InfraredCancelledBoostsIterator struct {
	Event *InfraredCancelledBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredCancelledBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredCancelledBoosts)
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
		it.Event = new(InfraredCancelledBoosts)
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
func (it *InfraredCancelledBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredCancelledBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredCancelledBoosts represents a CancelledBoosts event raised by the Infrared contract.
type InfraredCancelledBoosts struct {
	Sender  common.Address
	Pubkeys [][]byte
	Amts    []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelledBoosts is a free log retrieval operation binding the contract event 0x277aebb8a2c70752d1c2f967a481e28346237d7d182ff6f8e1992901c0034f67.
//
// Solidity: event CancelledBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) FilterCancelledBoosts(opts *bind.FilterOpts) (*InfraredCancelledBoostsIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "CancelledBoosts")
	if err != nil {
		return nil, err
	}
	return &InfraredCancelledBoostsIterator{contract: _Infrared.contract, event: "CancelledBoosts", logs: logs, sub: sub}, nil
}

// WatchCancelledBoosts is a free log subscription operation binding the contract event 0x277aebb8a2c70752d1c2f967a481e28346237d7d182ff6f8e1992901c0034f67.
//
// Solidity: event CancelledBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) WatchCancelledBoosts(opts *bind.WatchOpts, sink chan<- *InfraredCancelledBoosts) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "CancelledBoosts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredCancelledBoosts)
				if err := _Infrared.contract.UnpackLog(event, "CancelledBoosts", log); err != nil {
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

// ParseCancelledBoosts is a log parse operation binding the contract event 0x277aebb8a2c70752d1c2f967a481e28346237d7d182ff6f8e1992901c0034f67.
//
// Solidity: event CancelledBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) ParseCancelledBoosts(log types.Log) (*InfraredCancelledBoosts, error) {
	event := new(InfraredCancelledBoosts)
	if err := _Infrared.contract.UnpackLog(event, "CancelledBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredDroppedBoostsIterator is returned from FilterDroppedBoosts and is used to iterate over the raw logs and unpacked data for DroppedBoosts events raised by the Infrared contract.
type InfraredDroppedBoostsIterator struct {
	Event *InfraredDroppedBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredDroppedBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredDroppedBoosts)
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
		it.Event = new(InfraredDroppedBoosts)
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
func (it *InfraredDroppedBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredDroppedBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredDroppedBoosts represents a DroppedBoosts event raised by the Infrared contract.
type InfraredDroppedBoosts struct {
	Sender  common.Address
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDroppedBoosts is a free log retrieval operation binding the contract event 0x00cdcd3acc123d91be1f19ef5d8f8be0aaf6730b054ffeb0cdd45e4c2e504055.
//
// Solidity: event DroppedBoosts(address indexed _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) FilterDroppedBoosts(opts *bind.FilterOpts, _sender []common.Address) (*InfraredDroppedBoostsIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "DroppedBoosts", _senderRule)
	if err != nil {
		return nil, err
	}
	return &InfraredDroppedBoostsIterator{contract: _Infrared.contract, event: "DroppedBoosts", logs: logs, sub: sub}, nil
}

// WatchDroppedBoosts is a free log subscription operation binding the contract event 0x00cdcd3acc123d91be1f19ef5d8f8be0aaf6730b054ffeb0cdd45e4c2e504055.
//
// Solidity: event DroppedBoosts(address indexed _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) WatchDroppedBoosts(opts *bind.WatchOpts, sink chan<- *InfraredDroppedBoosts, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "DroppedBoosts", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredDroppedBoosts)
				if err := _Infrared.contract.UnpackLog(event, "DroppedBoosts", log); err != nil {
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

// ParseDroppedBoosts is a log parse operation binding the contract event 0x00cdcd3acc123d91be1f19ef5d8f8be0aaf6730b054ffeb0cdd45e4c2e504055.
//
// Solidity: event DroppedBoosts(address indexed _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) ParseDroppedBoosts(log types.Log) (*InfraredDroppedBoosts, error) {
	event := new(InfraredDroppedBoosts)
	if err := _Infrared.contract.UnpackLog(event, "DroppedBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the Infrared contract.
type InfraredFeeUpdatedIterator struct {
	Event *InfraredFeeUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredFeeUpdated)
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
		it.Event = new(InfraredFeeUpdated)
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
func (it *InfraredFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredFeeUpdated represents a FeeUpdated event raised by the Infrared contract.
type InfraredFeeUpdated struct {
	Sender     common.Address
	FeeType    uint8
	OldFeeRate *big.Int
	NewFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x5332bf595be156bc0b759aebeff52c12f0c2656e3beea83e881c3b0dacea7655.
//
// Solidity: event FeeUpdated(address _sender, uint8 _feeType, uint256 _oldFeeRate, uint256 _newFeeRate)
func (_Infrared *InfraredFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*InfraredFeeUpdatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredFeeUpdatedIterator{contract: _Infrared.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x5332bf595be156bc0b759aebeff52c12f0c2656e3beea83e881c3b0dacea7655.
//
// Solidity: event FeeUpdated(address _sender, uint8 _feeType, uint256 _oldFeeRate, uint256 _newFeeRate)
func (_Infrared *InfraredFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *InfraredFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredFeeUpdated)
				if err := _Infrared.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x5332bf595be156bc0b759aebeff52c12f0c2656e3beea83e881c3b0dacea7655.
//
// Solidity: event FeeUpdated(address _sender, uint8 _feeType, uint256 _oldFeeRate, uint256 _newFeeRate)
func (_Infrared *InfraredFilterer) ParseFeeUpdated(log types.Log) (*InfraredFeeUpdated, error) {
	event := new(InfraredFeeUpdated)
	if err := _Infrared.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredIBGTSetIterator is returned from FilterIBGTSet and is used to iterate over the raw logs and unpacked data for IBGTSet events raised by the Infrared contract.
type InfraredIBGTSetIterator struct {
	Event *InfraredIBGTSet // Event containing the contract specifics and raw log

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
func (it *InfraredIBGTSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredIBGTSet)
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
		it.Event = new(InfraredIBGTSet)
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
func (it *InfraredIBGTSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredIBGTSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredIBGTSet represents a IBGTSet event raised by the Infrared contract.
type InfraredIBGTSet struct {
	Sender common.Address
	Ibgt   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIBGTSet is a free log retrieval operation binding the contract event 0x01fb3d0e426042c8588b7aec3073a19eeb8b8554fb8eeae7e4266eaaca696cde.
//
// Solidity: event IBGTSet(address _sender, address _ibgt)
func (_Infrared *InfraredFilterer) FilterIBGTSet(opts *bind.FilterOpts) (*InfraredIBGTSetIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "IBGTSet")
	if err != nil {
		return nil, err
	}
	return &InfraredIBGTSetIterator{contract: _Infrared.contract, event: "IBGTSet", logs: logs, sub: sub}, nil
}

// WatchIBGTSet is a free log subscription operation binding the contract event 0x01fb3d0e426042c8588b7aec3073a19eeb8b8554fb8eeae7e4266eaaca696cde.
//
// Solidity: event IBGTSet(address _sender, address _ibgt)
func (_Infrared *InfraredFilterer) WatchIBGTSet(opts *bind.WatchOpts, sink chan<- *InfraredIBGTSet) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "IBGTSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredIBGTSet)
				if err := _Infrared.contract.UnpackLog(event, "IBGTSet", log); err != nil {
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

// ParseIBGTSet is a log parse operation binding the contract event 0x01fb3d0e426042c8588b7aec3073a19eeb8b8554fb8eeae7e4266eaaca696cde.
//
// Solidity: event IBGTSet(address _sender, address _ibgt)
func (_Infrared *InfraredFilterer) ParseIBGTSet(log types.Log) (*InfraredIBGTSet, error) {
	event := new(InfraredIBGTSet)
	if err := _Infrared.contract.UnpackLog(event, "IBGTSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredIRSetIterator is returned from FilterIRSet and is used to iterate over the raw logs and unpacked data for IRSet events raised by the Infrared contract.
type InfraredIRSetIterator struct {
	Event *InfraredIRSet // Event containing the contract specifics and raw log

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
func (it *InfraredIRSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredIRSet)
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
		it.Event = new(InfraredIRSet)
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
func (it *InfraredIRSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredIRSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredIRSet represents a IRSet event raised by the Infrared contract.
type InfraredIRSet struct {
	Sender common.Address
	IR     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIRSet is a free log retrieval operation binding the contract event 0x479870560c8b27c5b7f3ba9bfdea7a215d7fa90fa99ac5f6f4130e83219fb396.
//
// Solidity: event IRSet(address _sender, address _IR)
func (_Infrared *InfraredFilterer) FilterIRSet(opts *bind.FilterOpts) (*InfraredIRSetIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "IRSet")
	if err != nil {
		return nil, err
	}
	return &InfraredIRSetIterator{contract: _Infrared.contract, event: "IRSet", logs: logs, sub: sub}, nil
}

// WatchIRSet is a free log subscription operation binding the contract event 0x479870560c8b27c5b7f3ba9bfdea7a215d7fa90fa99ac5f6f4130e83219fb396.
//
// Solidity: event IRSet(address _sender, address _IR)
func (_Infrared *InfraredFilterer) WatchIRSet(opts *bind.WatchOpts, sink chan<- *InfraredIRSet) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "IRSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredIRSet)
				if err := _Infrared.contract.UnpackLog(event, "IRSet", log); err != nil {
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

// ParseIRSet is a log parse operation binding the contract event 0x479870560c8b27c5b7f3ba9bfdea7a215d7fa90fa99ac5f6f4130e83219fb396.
//
// Solidity: event IRSet(address _sender, address _IR)
func (_Infrared *InfraredFilterer) ParseIRSet(log types.Log) (*InfraredIRSet, error) {
	event := new(InfraredIRSet)
	if err := _Infrared.contract.UnpackLog(event, "IRSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredInfraredBERABribeSplitUpdatedIterator is returned from FilterInfraredBERABribeSplitUpdated and is used to iterate over the raw logs and unpacked data for InfraredBERABribeSplitUpdated events raised by the Infrared contract.
type InfraredInfraredBERABribeSplitUpdatedIterator struct {
	Event *InfraredInfraredBERABribeSplitUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredInfraredBERABribeSplitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredInfraredBERABribeSplitUpdated)
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
		it.Event = new(InfraredInfraredBERABribeSplitUpdated)
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
func (it *InfraredInfraredBERABribeSplitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredInfraredBERABribeSplitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredInfraredBERABribeSplitUpdated represents a InfraredBERABribeSplitUpdated event raised by the Infrared contract.
type InfraredInfraredBERABribeSplitUpdated struct {
	Sender    common.Address
	OldWeight *big.Int
	NewWeight *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInfraredBERABribeSplitUpdated is a free log retrieval operation binding the contract event 0x0fd67e3c16d445befd821bc8cb4a2ab0c3eedd866f8089ceb99b19ced15fdb88.
//
// Solidity: event InfraredBERABribeSplitUpdated(address _sender, uint256 _oldWeight, uint256 _newWeight)
func (_Infrared *InfraredFilterer) FilterInfraredBERABribeSplitUpdated(opts *bind.FilterOpts) (*InfraredInfraredBERABribeSplitUpdatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "InfraredBERABribeSplitUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredInfraredBERABribeSplitUpdatedIterator{contract: _Infrared.contract, event: "InfraredBERABribeSplitUpdated", logs: logs, sub: sub}, nil
}

// WatchInfraredBERABribeSplitUpdated is a free log subscription operation binding the contract event 0x0fd67e3c16d445befd821bc8cb4a2ab0c3eedd866f8089ceb99b19ced15fdb88.
//
// Solidity: event InfraredBERABribeSplitUpdated(address _sender, uint256 _oldWeight, uint256 _newWeight)
func (_Infrared *InfraredFilterer) WatchInfraredBERABribeSplitUpdated(opts *bind.WatchOpts, sink chan<- *InfraredInfraredBERABribeSplitUpdated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "InfraredBERABribeSplitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredInfraredBERABribeSplitUpdated)
				if err := _Infrared.contract.UnpackLog(event, "InfraredBERABribeSplitUpdated", log); err != nil {
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

// ParseInfraredBERABribeSplitUpdated is a log parse operation binding the contract event 0x0fd67e3c16d445befd821bc8cb4a2ab0c3eedd866f8089ceb99b19ced15fdb88.
//
// Solidity: event InfraredBERABribeSplitUpdated(address _sender, uint256 _oldWeight, uint256 _newWeight)
func (_Infrared *InfraredFilterer) ParseInfraredBERABribeSplitUpdated(log types.Log) (*InfraredInfraredBERABribeSplitUpdated, error) {
	event := new(InfraredInfraredBERABribeSplitUpdated)
	if err := _Infrared.contract.UnpackLog(event, "InfraredBERABribeSplitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredInfraredBGTSuppliedIterator is returned from FilterInfraredBGTSupplied and is used to iterate over the raw logs and unpacked data for InfraredBGTSupplied events raised by the Infrared contract.
type InfraredInfraredBGTSuppliedIterator struct {
	Event *InfraredInfraredBGTSupplied // Event containing the contract specifics and raw log

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
func (it *InfraredInfraredBGTSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredInfraredBGTSupplied)
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
		it.Event = new(InfraredInfraredBGTSupplied)
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
func (it *InfraredInfraredBGTSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredInfraredBGTSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredInfraredBGTSupplied represents a InfraredBGTSupplied event raised by the Infrared contract.
type InfraredInfraredBGTSupplied struct {
	Vault   common.Address
	IbgtAmt *big.Int
	IredAmt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInfraredBGTSupplied is a free log retrieval operation binding the contract event 0x594615b4057efb5d66d20575d74bb57176dd99f876b613531b9d6b539af56b19.
//
// Solidity: event InfraredBGTSupplied(address indexed _vault, uint256 _ibgtAmt, uint256 _iredAmt)
func (_Infrared *InfraredFilterer) FilterInfraredBGTSupplied(opts *bind.FilterOpts, _vault []common.Address) (*InfraredInfraredBGTSuppliedIterator, error) {

	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "InfraredBGTSupplied", _vaultRule)
	if err != nil {
		return nil, err
	}
	return &InfraredInfraredBGTSuppliedIterator{contract: _Infrared.contract, event: "InfraredBGTSupplied", logs: logs, sub: sub}, nil
}

// WatchInfraredBGTSupplied is a free log subscription operation binding the contract event 0x594615b4057efb5d66d20575d74bb57176dd99f876b613531b9d6b539af56b19.
//
// Solidity: event InfraredBGTSupplied(address indexed _vault, uint256 _ibgtAmt, uint256 _iredAmt)
func (_Infrared *InfraredFilterer) WatchInfraredBGTSupplied(opts *bind.WatchOpts, sink chan<- *InfraredInfraredBGTSupplied, _vault []common.Address) (event.Subscription, error) {

	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "InfraredBGTSupplied", _vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredInfraredBGTSupplied)
				if err := _Infrared.contract.UnpackLog(event, "InfraredBGTSupplied", log); err != nil {
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

// ParseInfraredBGTSupplied is a log parse operation binding the contract event 0x594615b4057efb5d66d20575d74bb57176dd99f876b613531b9d6b539af56b19.
//
// Solidity: event InfraredBGTSupplied(address indexed _vault, uint256 _ibgtAmt, uint256 _iredAmt)
func (_Infrared *InfraredFilterer) ParseInfraredBGTSupplied(log types.Log) (*InfraredInfraredBGTSupplied, error) {
	event := new(InfraredInfraredBGTSupplied)
	if err := _Infrared.contract.UnpackLog(event, "InfraredBGTSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredInfraredBGTUpdatedIterator is returned from FilterInfraredBGTUpdated and is used to iterate over the raw logs and unpacked data for InfraredBGTUpdated events raised by the Infrared contract.
type InfraredInfraredBGTUpdatedIterator struct {
	Event *InfraredInfraredBGTUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredInfraredBGTUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredInfraredBGTUpdated)
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
		it.Event = new(InfraredInfraredBGTUpdated)
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
func (it *InfraredInfraredBGTUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredInfraredBGTUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredInfraredBGTUpdated represents a InfraredBGTUpdated event raised by the Infrared contract.
type InfraredInfraredBGTUpdated struct {
	Sender  common.Address
	OldIbgt common.Address
	NewIbgt common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInfraredBGTUpdated is a free log retrieval operation binding the contract event 0x82328691e54c53dc3eb5ee58e6f153738b2fea020058273ed5c05e1bcd75db95.
//
// Solidity: event InfraredBGTUpdated(address _sender, address _oldIbgt, address _newIbgt)
func (_Infrared *InfraredFilterer) FilterInfraredBGTUpdated(opts *bind.FilterOpts) (*InfraredInfraredBGTUpdatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "InfraredBGTUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredInfraredBGTUpdatedIterator{contract: _Infrared.contract, event: "InfraredBGTUpdated", logs: logs, sub: sub}, nil
}

// WatchInfraredBGTUpdated is a free log subscription operation binding the contract event 0x82328691e54c53dc3eb5ee58e6f153738b2fea020058273ed5c05e1bcd75db95.
//
// Solidity: event InfraredBGTUpdated(address _sender, address _oldIbgt, address _newIbgt)
func (_Infrared *InfraredFilterer) WatchInfraredBGTUpdated(opts *bind.WatchOpts, sink chan<- *InfraredInfraredBGTUpdated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "InfraredBGTUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredInfraredBGTUpdated)
				if err := _Infrared.contract.UnpackLog(event, "InfraredBGTUpdated", log); err != nil {
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

// ParseInfraredBGTUpdated is a log parse operation binding the contract event 0x82328691e54c53dc3eb5ee58e6f153738b2fea020058273ed5c05e1bcd75db95.
//
// Solidity: event InfraredBGTUpdated(address _sender, address _oldIbgt, address _newIbgt)
func (_Infrared *InfraredFilterer) ParseInfraredBGTUpdated(log types.Log) (*InfraredInfraredBGTUpdated, error) {
	event := new(InfraredInfraredBGTUpdated)
	if err := _Infrared.contract.UnpackLog(event, "InfraredBGTUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredInfraredBGTVaultUpdatedIterator is returned from FilterInfraredBGTVaultUpdated and is used to iterate over the raw logs and unpacked data for InfraredBGTVaultUpdated events raised by the Infrared contract.
type InfraredInfraredBGTVaultUpdatedIterator struct {
	Event *InfraredInfraredBGTVaultUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredInfraredBGTVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredInfraredBGTVaultUpdated)
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
		it.Event = new(InfraredInfraredBGTVaultUpdated)
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
func (it *InfraredInfraredBGTVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredInfraredBGTVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredInfraredBGTVaultUpdated represents a InfraredBGTVaultUpdated event raised by the Infrared contract.
type InfraredInfraredBGTVaultUpdated struct {
	Sender       common.Address
	OldIbgtVault common.Address
	NewIbgtVault common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInfraredBGTVaultUpdated is a free log retrieval operation binding the contract event 0x884b83a6c9e565ae8ed1c43bd13563d349f3b8adec010f4e2adb485c66086808.
//
// Solidity: event InfraredBGTVaultUpdated(address _sender, address _oldIbgtVault, address _newIbgtVault)
func (_Infrared *InfraredFilterer) FilterInfraredBGTVaultUpdated(opts *bind.FilterOpts) (*InfraredInfraredBGTVaultUpdatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "InfraredBGTVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredInfraredBGTVaultUpdatedIterator{contract: _Infrared.contract, event: "InfraredBGTVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchInfraredBGTVaultUpdated is a free log subscription operation binding the contract event 0x884b83a6c9e565ae8ed1c43bd13563d349f3b8adec010f4e2adb485c66086808.
//
// Solidity: event InfraredBGTVaultUpdated(address _sender, address _oldIbgtVault, address _newIbgtVault)
func (_Infrared *InfraredFilterer) WatchInfraredBGTVaultUpdated(opts *bind.WatchOpts, sink chan<- *InfraredInfraredBGTVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "InfraredBGTVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredInfraredBGTVaultUpdated)
				if err := _Infrared.contract.UnpackLog(event, "InfraredBGTVaultUpdated", log); err != nil {
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

// ParseInfraredBGTVaultUpdated is a log parse operation binding the contract event 0x884b83a6c9e565ae8ed1c43bd13563d349f3b8adec010f4e2adb485c66086808.
//
// Solidity: event InfraredBGTVaultUpdated(address _sender, address _oldIbgtVault, address _newIbgtVault)
func (_Infrared *InfraredFilterer) ParseInfraredBGTVaultUpdated(log types.Log) (*InfraredInfraredBGTVaultUpdated, error) {
	event := new(InfraredInfraredBGTVaultUpdated)
	if err := _Infrared.contract.UnpackLog(event, "InfraredBGTVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Infrared contract.
type InfraredInitializedIterator struct {
	Event *InfraredInitialized // Event containing the contract specifics and raw log

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
func (it *InfraredInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredInitialized)
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
		it.Event = new(InfraredInitialized)
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
func (it *InfraredInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredInitialized represents a Initialized event raised by the Infrared contract.
type InfraredInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Infrared *InfraredFilterer) FilterInitialized(opts *bind.FilterOpts) (*InfraredInitializedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InfraredInitializedIterator{contract: _Infrared.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Infrared *InfraredFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InfraredInitialized) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredInitialized)
				if err := _Infrared.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseInitialized(log types.Log) (*InfraredInitialized, error) {
	event := new(InfraredInitialized)
	if err := _Infrared.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredNewVaultIterator is returned from FilterNewVault and is used to iterate over the raw logs and unpacked data for NewVault events raised by the Infrared contract.
type InfraredNewVaultIterator struct {
	Event *InfraredNewVault // Event containing the contract specifics and raw log

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
func (it *InfraredNewVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredNewVault)
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
		it.Event = new(InfraredNewVault)
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
func (it *InfraredNewVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredNewVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredNewVault represents a NewVault event raised by the Infrared contract.
type InfraredNewVault struct {
	Sender common.Address
	Asset  common.Address
	Vault  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewVault is a free log retrieval operation binding the contract event 0x891f008b8c9cadc4c9114ed37ce718f739b28049d58e50e35a02d94c4e9b06ff.
//
// Solidity: event NewVault(address _sender, address indexed _asset, address indexed _vault)
func (_Infrared *InfraredFilterer) FilterNewVault(opts *bind.FilterOpts, _asset []common.Address, _vault []common.Address) (*InfraredNewVaultIterator, error) {

	var _assetRule []interface{}
	for _, _assetItem := range _asset {
		_assetRule = append(_assetRule, _assetItem)
	}
	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "NewVault", _assetRule, _vaultRule)
	if err != nil {
		return nil, err
	}
	return &InfraredNewVaultIterator{contract: _Infrared.contract, event: "NewVault", logs: logs, sub: sub}, nil
}

// WatchNewVault is a free log subscription operation binding the contract event 0x891f008b8c9cadc4c9114ed37ce718f739b28049d58e50e35a02d94c4e9b06ff.
//
// Solidity: event NewVault(address _sender, address indexed _asset, address indexed _vault)
func (_Infrared *InfraredFilterer) WatchNewVault(opts *bind.WatchOpts, sink chan<- *InfraredNewVault, _asset []common.Address, _vault []common.Address) (event.Subscription, error) {

	var _assetRule []interface{}
	for _, _assetItem := range _asset {
		_assetRule = append(_assetRule, _assetItem)
	}
	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "NewVault", _assetRule, _vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredNewVault)
				if err := _Infrared.contract.UnpackLog(event, "NewVault", log); err != nil {
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

// ParseNewVault is a log parse operation binding the contract event 0x891f008b8c9cadc4c9114ed37ce718f739b28049d58e50e35a02d94c4e9b06ff.
//
// Solidity: event NewVault(address _sender, address indexed _asset, address indexed _vault)
func (_Infrared *InfraredFilterer) ParseNewVault(log types.Log) (*InfraredNewVault, error) {
	event := new(InfraredNewVault)
	if err := _Infrared.contract.UnpackLog(event, "NewVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredOperatorRewardsDistributedIterator is returned from FilterOperatorRewardsDistributed and is used to iterate over the raw logs and unpacked data for OperatorRewardsDistributed events raised by the Infrared contract.
type InfraredOperatorRewardsDistributedIterator struct {
	Event *InfraredOperatorRewardsDistributed // Event containing the contract specifics and raw log

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
func (it *InfraredOperatorRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredOperatorRewardsDistributed)
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
		it.Event = new(InfraredOperatorRewardsDistributed)
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
func (it *InfraredOperatorRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredOperatorRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredOperatorRewardsDistributed represents a OperatorRewardsDistributed event raised by the Infrared contract.
type InfraredOperatorRewardsDistributed struct {
	Ibera       common.Address
	Distributor common.Address
	Amt         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorRewardsDistributed is a free log retrieval operation binding the contract event 0x2e42554d495020b729a23098e4bcbddb25d4a71511cc6b03c7faa5ab51f0570d.
//
// Solidity: event OperatorRewardsDistributed(address indexed _ibera, address indexed _distributor, uint256 _amt)
func (_Infrared *InfraredFilterer) FilterOperatorRewardsDistributed(opts *bind.FilterOpts, _ibera []common.Address, _distributor []common.Address) (*InfraredOperatorRewardsDistributedIterator, error) {

	var _iberaRule []interface{}
	for _, _iberaItem := range _ibera {
		_iberaRule = append(_iberaRule, _iberaItem)
	}
	var _distributorRule []interface{}
	for _, _distributorItem := range _distributor {
		_distributorRule = append(_distributorRule, _distributorItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "OperatorRewardsDistributed", _iberaRule, _distributorRule)
	if err != nil {
		return nil, err
	}
	return &InfraredOperatorRewardsDistributedIterator{contract: _Infrared.contract, event: "OperatorRewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchOperatorRewardsDistributed is a free log subscription operation binding the contract event 0x2e42554d495020b729a23098e4bcbddb25d4a71511cc6b03c7faa5ab51f0570d.
//
// Solidity: event OperatorRewardsDistributed(address indexed _ibera, address indexed _distributor, uint256 _amt)
func (_Infrared *InfraredFilterer) WatchOperatorRewardsDistributed(opts *bind.WatchOpts, sink chan<- *InfraredOperatorRewardsDistributed, _ibera []common.Address, _distributor []common.Address) (event.Subscription, error) {

	var _iberaRule []interface{}
	for _, _iberaItem := range _ibera {
		_iberaRule = append(_iberaRule, _iberaItem)
	}
	var _distributorRule []interface{}
	for _, _distributorItem := range _distributor {
		_distributorRule = append(_distributorRule, _distributorItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "OperatorRewardsDistributed", _iberaRule, _distributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredOperatorRewardsDistributed)
				if err := _Infrared.contract.UnpackLog(event, "OperatorRewardsDistributed", log); err != nil {
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

// ParseOperatorRewardsDistributed is a log parse operation binding the contract event 0x2e42554d495020b729a23098e4bcbddb25d4a71511cc6b03c7faa5ab51f0570d.
//
// Solidity: event OperatorRewardsDistributed(address indexed _ibera, address indexed _distributor, uint256 _amt)
func (_Infrared *InfraredFilterer) ParseOperatorRewardsDistributed(log types.Log) (*InfraredOperatorRewardsDistributed, error) {
	event := new(InfraredOperatorRewardsDistributed)
	if err := _Infrared.contract.UnpackLog(event, "OperatorRewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Infrared contract.
type InfraredPausedIterator struct {
	Event *InfraredPaused // Event containing the contract specifics and raw log

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
func (it *InfraredPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredPaused)
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
		it.Event = new(InfraredPaused)
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
func (it *InfraredPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredPaused represents a Paused event raised by the Infrared contract.
type InfraredPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Infrared *InfraredFilterer) FilterPaused(opts *bind.FilterOpts) (*InfraredPausedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InfraredPausedIterator{contract: _Infrared.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Infrared *InfraredFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InfraredPaused) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredPaused)
				if err := _Infrared.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParsePaused(log types.Log) (*InfraredPaused, error) {
	event := new(InfraredPaused)
	if err := _Infrared.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredProtocolFeesIterator is returned from FilterProtocolFees and is used to iterate over the raw logs and unpacked data for ProtocolFees events raised by the Infrared contract.
type InfraredProtocolFeesIterator struct {
	Event *InfraredProtocolFees // Event containing the contract specifics and raw log

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
func (it *InfraredProtocolFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredProtocolFees)
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
		it.Event = new(InfraredProtocolFees)
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
func (it *InfraredProtocolFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredProtocolFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredProtocolFees represents a ProtocolFees event raised by the Infrared contract.
type InfraredProtocolFees struct {
	Token    common.Address
	Amt      *big.Int
	VoterAmt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFees is a free log retrieval operation binding the contract event 0x6236815025b2e02f65f95d322e662ef46303653aec5edefdd706cf84e0d73223.
//
// Solidity: event ProtocolFees(address indexed _token, uint256 _amt, uint256 _voterAmt)
func (_Infrared *InfraredFilterer) FilterProtocolFees(opts *bind.FilterOpts, _token []common.Address) (*InfraredProtocolFeesIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ProtocolFees", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredProtocolFeesIterator{contract: _Infrared.contract, event: "ProtocolFees", logs: logs, sub: sub}, nil
}

// WatchProtocolFees is a free log subscription operation binding the contract event 0x6236815025b2e02f65f95d322e662ef46303653aec5edefdd706cf84e0d73223.
//
// Solidity: event ProtocolFees(address indexed _token, uint256 _amt, uint256 _voterAmt)
func (_Infrared *InfraredFilterer) WatchProtocolFees(opts *bind.WatchOpts, sink chan<- *InfraredProtocolFees, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ProtocolFees", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredProtocolFees)
				if err := _Infrared.contract.UnpackLog(event, "ProtocolFees", log); err != nil {
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

// ParseProtocolFees is a log parse operation binding the contract event 0x6236815025b2e02f65f95d322e662ef46303653aec5edefdd706cf84e0d73223.
//
// Solidity: event ProtocolFees(address indexed _token, uint256 _amt, uint256 _voterAmt)
func (_Infrared *InfraredFilterer) ParseProtocolFees(log types.Log) (*InfraredProtocolFees, error) {
	event := new(InfraredProtocolFees)
	if err := _Infrared.contract.UnpackLog(event, "ProtocolFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredProtocolFeesClaimedIterator is returned from FilterProtocolFeesClaimed and is used to iterate over the raw logs and unpacked data for ProtocolFeesClaimed events raised by the Infrared contract.
type InfraredProtocolFeesClaimedIterator struct {
	Event *InfraredProtocolFeesClaimed // Event containing the contract specifics and raw log

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
func (it *InfraredProtocolFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredProtocolFeesClaimed)
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
		it.Event = new(InfraredProtocolFeesClaimed)
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
func (it *InfraredProtocolFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredProtocolFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredProtocolFeesClaimed represents a ProtocolFeesClaimed event raised by the Infrared contract.
type InfraredProtocolFeesClaimed struct {
	Sender common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeesClaimed is a free log retrieval operation binding the contract event 0x347de86417f108df3ab27b51ed2076c0aa2c9cb0c4743ecefaacad8eab02c0f0.
//
// Solidity: event ProtocolFeesClaimed(address _sender, address _to, address _token, uint256 _amount)
func (_Infrared *InfraredFilterer) FilterProtocolFeesClaimed(opts *bind.FilterOpts) (*InfraredProtocolFeesClaimedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ProtocolFeesClaimed")
	if err != nil {
		return nil, err
	}
	return &InfraredProtocolFeesClaimedIterator{contract: _Infrared.contract, event: "ProtocolFeesClaimed", logs: logs, sub: sub}, nil
}

// WatchProtocolFeesClaimed is a free log subscription operation binding the contract event 0x347de86417f108df3ab27b51ed2076c0aa2c9cb0c4743ecefaacad8eab02c0f0.
//
// Solidity: event ProtocolFeesClaimed(address _sender, address _to, address _token, uint256 _amount)
func (_Infrared *InfraredFilterer) WatchProtocolFeesClaimed(opts *bind.WatchOpts, sink chan<- *InfraredProtocolFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ProtocolFeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredProtocolFeesClaimed)
				if err := _Infrared.contract.UnpackLog(event, "ProtocolFeesClaimed", log); err != nil {
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

// ParseProtocolFeesClaimed is a log parse operation binding the contract event 0x347de86417f108df3ab27b51ed2076c0aa2c9cb0c4743ecefaacad8eab02c0f0.
//
// Solidity: event ProtocolFeesClaimed(address _sender, address _to, address _token, uint256 _amount)
func (_Infrared *InfraredFilterer) ParseProtocolFeesClaimed(log types.Log) (*InfraredProtocolFeesClaimed, error) {
	event := new(InfraredProtocolFeesClaimed)
	if err := _Infrared.contract.UnpackLog(event, "ProtocolFeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredQueueDropBoostsIterator is returned from FilterQueueDropBoosts and is used to iterate over the raw logs and unpacked data for QueueDropBoosts events raised by the Infrared contract.
type InfraredQueueDropBoostsIterator struct {
	Event *InfraredQueueDropBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredQueueDropBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredQueueDropBoosts)
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
		it.Event = new(InfraredQueueDropBoosts)
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
func (it *InfraredQueueDropBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredQueueDropBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredQueueDropBoosts represents a QueueDropBoosts event raised by the Infrared contract.
type InfraredQueueDropBoosts struct {
	User    common.Address
	Pubkeys [][]byte
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterQueueDropBoosts is a free log retrieval operation binding the contract event 0x86e1fbae4fac9b3d74ffd63a3ae649b9e1348f1492710f2dc4c74583007cd64b.
//
// Solidity: event QueueDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) FilterQueueDropBoosts(opts *bind.FilterOpts, user []common.Address, pubkeys [][][]byte) (*InfraredQueueDropBoostsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pubkeysRule []interface{}
	for _, pubkeysItem := range pubkeys {
		pubkeysRule = append(pubkeysRule, pubkeysItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "QueueDropBoosts", userRule, pubkeysRule)
	if err != nil {
		return nil, err
	}
	return &InfraredQueueDropBoostsIterator{contract: _Infrared.contract, event: "QueueDropBoosts", logs: logs, sub: sub}, nil
}

// WatchQueueDropBoosts is a free log subscription operation binding the contract event 0x86e1fbae4fac9b3d74ffd63a3ae649b9e1348f1492710f2dc4c74583007cd64b.
//
// Solidity: event QueueDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) WatchQueueDropBoosts(opts *bind.WatchOpts, sink chan<- *InfraredQueueDropBoosts, user []common.Address, pubkeys [][][]byte) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pubkeysRule []interface{}
	for _, pubkeysItem := range pubkeys {
		pubkeysRule = append(pubkeysRule, pubkeysItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "QueueDropBoosts", userRule, pubkeysRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredQueueDropBoosts)
				if err := _Infrared.contract.UnpackLog(event, "QueueDropBoosts", log); err != nil {
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

// ParseQueueDropBoosts is a log parse operation binding the contract event 0x86e1fbae4fac9b3d74ffd63a3ae649b9e1348f1492710f2dc4c74583007cd64b.
//
// Solidity: event QueueDropBoosts(address indexed user, bytes[] indexed pubkeys, uint128[] amounts)
func (_Infrared *InfraredFilterer) ParseQueueDropBoosts(log types.Log) (*InfraredQueueDropBoosts, error) {
	event := new(InfraredQueueDropBoosts)
	if err := _Infrared.contract.UnpackLog(event, "QueueDropBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredQueuedBoostsIterator is returned from FilterQueuedBoosts and is used to iterate over the raw logs and unpacked data for QueuedBoosts events raised by the Infrared contract.
type InfraredQueuedBoostsIterator struct {
	Event *InfraredQueuedBoosts // Event containing the contract specifics and raw log

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
func (it *InfraredQueuedBoostsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredQueuedBoosts)
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
		it.Event = new(InfraredQueuedBoosts)
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
func (it *InfraredQueuedBoostsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredQueuedBoostsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredQueuedBoosts represents a QueuedBoosts event raised by the Infrared contract.
type InfraredQueuedBoosts struct {
	Sender  common.Address
	Pubkeys [][]byte
	Amts    []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterQueuedBoosts is a free log retrieval operation binding the contract event 0xd4834079da4b292baebce77b556196b77c9cf5fafd01cfcaa97b3cb7a05310fd.
//
// Solidity: event QueuedBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) FilterQueuedBoosts(opts *bind.FilterOpts) (*InfraredQueuedBoostsIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "QueuedBoosts")
	if err != nil {
		return nil, err
	}
	return &InfraredQueuedBoostsIterator{contract: _Infrared.contract, event: "QueuedBoosts", logs: logs, sub: sub}, nil
}

// WatchQueuedBoosts is a free log subscription operation binding the contract event 0xd4834079da4b292baebce77b556196b77c9cf5fafd01cfcaa97b3cb7a05310fd.
//
// Solidity: event QueuedBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) WatchQueuedBoosts(opts *bind.WatchOpts, sink chan<- *InfraredQueuedBoosts) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "QueuedBoosts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredQueuedBoosts)
				if err := _Infrared.contract.UnpackLog(event, "QueuedBoosts", log); err != nil {
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

// ParseQueuedBoosts is a log parse operation binding the contract event 0xd4834079da4b292baebce77b556196b77c9cf5fafd01cfcaa97b3cb7a05310fd.
//
// Solidity: event QueuedBoosts(address _sender, bytes[] _pubkeys, uint128[] _amts)
func (_Infrared *InfraredFilterer) ParseQueuedBoosts(log types.Log) (*InfraredQueuedBoosts, error) {
	event := new(InfraredQueuedBoosts)
	if err := _Infrared.contract.UnpackLog(event, "QueuedBoosts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Infrared contract.
type InfraredRecoveredIterator struct {
	Event *InfraredRecovered // Event containing the contract specifics and raw log

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
func (it *InfraredRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRecovered)
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
		it.Event = new(InfraredRecovered)
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
func (it *InfraredRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRecovered represents a Recovered event raised by the Infrared contract.
type InfraredRecovered struct {
	Sender common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0xfff3b3844276f57024e0b42afec1a37f75db36511e43819a4f2a63ab7862b648.
//
// Solidity: event Recovered(address _sender, address indexed _token, uint256 _amount)
func (_Infrared *InfraredFilterer) FilterRecovered(opts *bind.FilterOpts, _token []common.Address) (*InfraredRecoveredIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Recovered", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredRecoveredIterator{contract: _Infrared.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0xfff3b3844276f57024e0b42afec1a37f75db36511e43819a4f2a63ab7862b648.
//
// Solidity: event Recovered(address _sender, address indexed _token, uint256 _amount)
func (_Infrared *InfraredFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *InfraredRecovered, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Recovered", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRecovered)
				if err := _Infrared.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0xfff3b3844276f57024e0b42afec1a37f75db36511e43819a4f2a63ab7862b648.
//
// Solidity: event Recovered(address _sender, address indexed _token, uint256 _amount)
func (_Infrared *InfraredFilterer) ParseRecovered(log types.Log) (*InfraredRecovered, error) {
	event := new(InfraredRecovered)
	if err := _Infrared.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the Infrared contract.
type InfraredRedelegatedIterator struct {
	Event *InfraredRedelegated // Event containing the contract specifics and raw log

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
func (it *InfraredRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRedelegated)
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
		it.Event = new(InfraredRedelegated)
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
func (it *InfraredRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRedelegated represents a Redelegated event raised by the Infrared contract.
type InfraredRedelegated struct {
	Sender common.Address
	From   []byte
	To     []byte
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0x54a6a94d88b70488019f8f04499f6c3929cbeea46f70cccd6ab5139d28bd701b.
//
// Solidity: event Redelegated(address _sender, bytes _from, bytes _to, uint256 _amt)
func (_Infrared *InfraredFilterer) FilterRedelegated(opts *bind.FilterOpts) (*InfraredRedelegatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Redelegated")
	if err != nil {
		return nil, err
	}
	return &InfraredRedelegatedIterator{contract: _Infrared.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0x54a6a94d88b70488019f8f04499f6c3929cbeea46f70cccd6ab5139d28bd701b.
//
// Solidity: event Redelegated(address _sender, bytes _from, bytes _to, uint256 _amt)
func (_Infrared *InfraredFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *InfraredRedelegated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Redelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRedelegated)
				if err := _Infrared.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0x54a6a94d88b70488019f8f04499f6c3929cbeea46f70cccd6ab5139d28bd701b.
//
// Solidity: event Redelegated(address _sender, bytes _from, bytes _to, uint256 _amt)
func (_Infrared *InfraredFilterer) ParseRedelegated(log types.Log) (*InfraredRedelegated, error) {
	event := new(InfraredRedelegated)
	if err := _Infrared.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRewardSuppliedIterator is returned from FilterRewardSupplied and is used to iterate over the raw logs and unpacked data for RewardSupplied events raised by the Infrared contract.
type InfraredRewardSuppliedIterator struct {
	Event *InfraredRewardSupplied // Event containing the contract specifics and raw log

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
func (it *InfraredRewardSuppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRewardSupplied)
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
		it.Event = new(InfraredRewardSupplied)
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
func (it *InfraredRewardSuppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRewardSuppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRewardSupplied represents a RewardSupplied event raised by the Infrared contract.
type InfraredRewardSupplied struct {
	Vault common.Address
	Token common.Address
	Amt   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardSupplied is a free log retrieval operation binding the contract event 0x75f899a4c93beaef4ceef04dcf09b52b95c2f6d76cc2f366b059f3f6ab165909.
//
// Solidity: event RewardSupplied(address indexed _vault, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) FilterRewardSupplied(opts *bind.FilterOpts, _vault []common.Address, _token []common.Address) (*InfraredRewardSuppliedIterator, error) {

	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RewardSupplied", _vaultRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredRewardSuppliedIterator{contract: _Infrared.contract, event: "RewardSupplied", logs: logs, sub: sub}, nil
}

// WatchRewardSupplied is a free log subscription operation binding the contract event 0x75f899a4c93beaef4ceef04dcf09b52b95c2f6d76cc2f366b059f3f6ab165909.
//
// Solidity: event RewardSupplied(address indexed _vault, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) WatchRewardSupplied(opts *bind.WatchOpts, sink chan<- *InfraredRewardSupplied, _vault []common.Address, _token []common.Address) (event.Subscription, error) {

	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RewardSupplied", _vaultRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRewardSupplied)
				if err := _Infrared.contract.UnpackLog(event, "RewardSupplied", log); err != nil {
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

// ParseRewardSupplied is a log parse operation binding the contract event 0x75f899a4c93beaef4ceef04dcf09b52b95c2f6d76cc2f366b059f3f6ab165909.
//
// Solidity: event RewardSupplied(address indexed _vault, address indexed _token, uint256 _amt)
func (_Infrared *InfraredFilterer) ParseRewardSupplied(log types.Log) (*InfraredRewardSupplied, error) {
	event := new(InfraredRewardSupplied)
	if err := _Infrared.contract.UnpackLog(event, "RewardSupplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRewardTokenNotSupportedIterator is returned from FilterRewardTokenNotSupported and is used to iterate over the raw logs and unpacked data for RewardTokenNotSupported events raised by the Infrared contract.
type InfraredRewardTokenNotSupportedIterator struct {
	Event *InfraredRewardTokenNotSupported // Event containing the contract specifics and raw log

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
func (it *InfraredRewardTokenNotSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRewardTokenNotSupported)
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
		it.Event = new(InfraredRewardTokenNotSupported)
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
func (it *InfraredRewardTokenNotSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRewardTokenNotSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRewardTokenNotSupported represents a RewardTokenNotSupported event raised by the Infrared contract.
type InfraredRewardTokenNotSupported struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardTokenNotSupported is a free log retrieval operation binding the contract event 0x97ae829a24746c7eadb50a0698ad54133086676ad079ea8a06dee26fc2ef29a3.
//
// Solidity: event RewardTokenNotSupported(address _token)
func (_Infrared *InfraredFilterer) FilterRewardTokenNotSupported(opts *bind.FilterOpts) (*InfraredRewardTokenNotSupportedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RewardTokenNotSupported")
	if err != nil {
		return nil, err
	}
	return &InfraredRewardTokenNotSupportedIterator{contract: _Infrared.contract, event: "RewardTokenNotSupported", logs: logs, sub: sub}, nil
}

// WatchRewardTokenNotSupported is a free log subscription operation binding the contract event 0x97ae829a24746c7eadb50a0698ad54133086676ad079ea8a06dee26fc2ef29a3.
//
// Solidity: event RewardTokenNotSupported(address _token)
func (_Infrared *InfraredFilterer) WatchRewardTokenNotSupported(opts *bind.WatchOpts, sink chan<- *InfraredRewardTokenNotSupported) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RewardTokenNotSupported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRewardTokenNotSupported)
				if err := _Infrared.contract.UnpackLog(event, "RewardTokenNotSupported", log); err != nil {
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

// ParseRewardTokenNotSupported is a log parse operation binding the contract event 0x97ae829a24746c7eadb50a0698ad54133086676ad079ea8a06dee26fc2ef29a3.
//
// Solidity: event RewardTokenNotSupported(address _token)
func (_Infrared *InfraredFilterer) ParseRewardTokenNotSupported(log types.Log) (*InfraredRewardTokenNotSupported, error) {
	event := new(InfraredRewardTokenNotSupported)
	if err := _Infrared.contract.UnpackLog(event, "RewardTokenNotSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the Infrared contract.
type InfraredRewardsDurationUpdatedIterator struct {
	Event *InfraredRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRewardsDurationUpdated)
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
		it.Event = new(InfraredRewardsDurationUpdated)
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
func (it *InfraredRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the Infrared contract.
type InfraredRewardsDurationUpdated struct {
	Sender      common.Address
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0x5010042fbbde5830368b15f9875f904d7611dcd3f3efaa3f4469398250c84aaa.
//
// Solidity: event RewardsDurationUpdated(address _sender, uint256 _oldDuration, uint256 _newDuration)
func (_Infrared *InfraredFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*InfraredRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &InfraredRewardsDurationUpdatedIterator{contract: _Infrared.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0x5010042fbbde5830368b15f9875f904d7611dcd3f3efaa3f4469398250c84aaa.
//
// Solidity: event RewardsDurationUpdated(address _sender, uint256 _oldDuration, uint256 _newDuration)
func (_Infrared *InfraredFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *InfraredRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRewardsDurationUpdated)
				if err := _Infrared.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0x5010042fbbde5830368b15f9875f904d7611dcd3f3efaa3f4469398250c84aaa.
//
// Solidity: event RewardsDurationUpdated(address _sender, uint256 _oldDuration, uint256 _newDuration)
func (_Infrared *InfraredFilterer) ParseRewardsDurationUpdated(log types.Log) (*InfraredRewardsDurationUpdated, error) {
	event := new(InfraredRewardsDurationUpdated)
	if err := _Infrared.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Infrared contract.
type InfraredRoleAdminChangedIterator struct {
	Event *InfraredRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *InfraredRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRoleAdminChanged)
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
		it.Event = new(InfraredRoleAdminChanged)
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
func (it *InfraredRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRoleAdminChanged represents a RoleAdminChanged event raised by the Infrared contract.
type InfraredRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Infrared *InfraredFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*InfraredRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &InfraredRoleAdminChangedIterator{contract: _Infrared.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Infrared *InfraredFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *InfraredRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRoleAdminChanged)
				if err := _Infrared.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseRoleAdminChanged(log types.Log) (*InfraredRoleAdminChanged, error) {
	event := new(InfraredRoleAdminChanged)
	if err := _Infrared.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Infrared contract.
type InfraredRoleGrantedIterator struct {
	Event *InfraredRoleGranted // Event containing the contract specifics and raw log

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
func (it *InfraredRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRoleGranted)
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
		it.Event = new(InfraredRoleGranted)
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
func (it *InfraredRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRoleGranted represents a RoleGranted event raised by the Infrared contract.
type InfraredRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Infrared *InfraredFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InfraredRoleGrantedIterator, error) {

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

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InfraredRoleGrantedIterator{contract: _Infrared.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Infrared *InfraredFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *InfraredRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRoleGranted)
				if err := _Infrared.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseRoleGranted(log types.Log) (*InfraredRoleGranted, error) {
	event := new(InfraredRoleGranted)
	if err := _Infrared.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Infrared contract.
type InfraredRoleRevokedIterator struct {
	Event *InfraredRoleRevoked // Event containing the contract specifics and raw log

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
func (it *InfraredRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredRoleRevoked)
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
		it.Event = new(InfraredRoleRevoked)
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
func (it *InfraredRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredRoleRevoked represents a RoleRevoked event raised by the Infrared contract.
type InfraredRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Infrared *InfraredFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*InfraredRoleRevokedIterator, error) {

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

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &InfraredRoleRevokedIterator{contract: _Infrared.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Infrared *InfraredFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *InfraredRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredRoleRevoked)
				if err := _Infrared.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseRoleRevoked(log types.Log) (*InfraredRoleRevoked, error) {
	event := new(InfraredRoleRevoked)
	if err := _Infrared.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Infrared contract.
type InfraredUndelegatedIterator struct {
	Event *InfraredUndelegated // Event containing the contract specifics and raw log

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
func (it *InfraredUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredUndelegated)
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
		it.Event = new(InfraredUndelegated)
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
func (it *InfraredUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredUndelegated represents a Undelegated event raised by the Infrared contract.
type InfraredUndelegated struct {
	Sender common.Address
	Pubkey []byte
	Amt    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x8fcff9315ad35ed2f36722045fc75aea491dce9380125ff8855990ef7bb47053.
//
// Solidity: event Undelegated(address _sender, bytes _pubkey, uint256 _amt)
func (_Infrared *InfraredFilterer) FilterUndelegated(opts *bind.FilterOpts) (*InfraredUndelegatedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Undelegated")
	if err != nil {
		return nil, err
	}
	return &InfraredUndelegatedIterator{contract: _Infrared.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x8fcff9315ad35ed2f36722045fc75aea491dce9380125ff8855990ef7bb47053.
//
// Solidity: event Undelegated(address _sender, bytes _pubkey, uint256 _amt)
func (_Infrared *InfraredFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *InfraredUndelegated) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Undelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredUndelegated)
				if err := _Infrared.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x8fcff9315ad35ed2f36722045fc75aea491dce9380125ff8855990ef7bb47053.
//
// Solidity: event Undelegated(address _sender, bytes _pubkey, uint256 _amt)
func (_Infrared *InfraredFilterer) ParseUndelegated(log types.Log) (*InfraredUndelegated, error) {
	event := new(InfraredUndelegated)
	if err := _Infrared.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Infrared contract.
type InfraredUnpausedIterator struct {
	Event *InfraredUnpaused // Event containing the contract specifics and raw log

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
func (it *InfraredUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredUnpaused)
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
		it.Event = new(InfraredUnpaused)
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
func (it *InfraredUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredUnpaused represents a Unpaused event raised by the Infrared contract.
type InfraredUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Infrared *InfraredFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InfraredUnpausedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InfraredUnpausedIterator{contract: _Infrared.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Infrared *InfraredFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InfraredUnpaused) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredUnpaused)
				if err := _Infrared.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseUnpaused(log types.Log) (*InfraredUnpaused, error) {
	event := new(InfraredUnpaused)
	if err := _Infrared.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredUpdatedIRMintRateIterator is returned from FilterUpdatedIRMintRate and is used to iterate over the raw logs and unpacked data for UpdatedIRMintRate events raised by the Infrared contract.
type InfraredUpdatedIRMintRateIterator struct {
	Event *InfraredUpdatedIRMintRate // Event containing the contract specifics and raw log

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
func (it *InfraredUpdatedIRMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredUpdatedIRMintRate)
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
		it.Event = new(InfraredUpdatedIRMintRate)
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
func (it *InfraredUpdatedIRMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredUpdatedIRMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredUpdatedIRMintRate represents a UpdatedIRMintRate event raised by the Infrared contract.
type InfraredUpdatedIRMintRate struct {
	OldMintRate *big.Int
	NewMintRate *big.Int
	Sender      common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatedIRMintRate is a free log retrieval operation binding the contract event 0xbfdc5161a1732f5628c29f6d7a72159b2f58074d77b24384d9d97a9dd7baebe5.
//
// Solidity: event UpdatedIRMintRate(uint256 oldMintRate, uint256 newMintRate, address sender)
func (_Infrared *InfraredFilterer) FilterUpdatedIRMintRate(opts *bind.FilterOpts) (*InfraredUpdatedIRMintRateIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "UpdatedIRMintRate")
	if err != nil {
		return nil, err
	}
	return &InfraredUpdatedIRMintRateIterator{contract: _Infrared.contract, event: "UpdatedIRMintRate", logs: logs, sub: sub}, nil
}

// WatchUpdatedIRMintRate is a free log subscription operation binding the contract event 0xbfdc5161a1732f5628c29f6d7a72159b2f58074d77b24384d9d97a9dd7baebe5.
//
// Solidity: event UpdatedIRMintRate(uint256 oldMintRate, uint256 newMintRate, address sender)
func (_Infrared *InfraredFilterer) WatchUpdatedIRMintRate(opts *bind.WatchOpts, sink chan<- *InfraredUpdatedIRMintRate) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "UpdatedIRMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredUpdatedIRMintRate)
				if err := _Infrared.contract.UnpackLog(event, "UpdatedIRMintRate", log); err != nil {
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

// ParseUpdatedIRMintRate is a log parse operation binding the contract event 0xbfdc5161a1732f5628c29f6d7a72159b2f58074d77b24384d9d97a9dd7baebe5.
//
// Solidity: event UpdatedIRMintRate(uint256 oldMintRate, uint256 newMintRate, address sender)
func (_Infrared *InfraredFilterer) ParseUpdatedIRMintRate(log types.Log) (*InfraredUpdatedIRMintRate, error) {
	event := new(InfraredUpdatedIRMintRate)
	if err := _Infrared.contract.UnpackLog(event, "UpdatedIRMintRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Infrared contract.
type InfraredUpgradedIterator struct {
	Event *InfraredUpgraded // Event containing the contract specifics and raw log

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
func (it *InfraredUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredUpgraded)
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
		it.Event = new(InfraredUpgraded)
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
func (it *InfraredUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredUpgraded represents a Upgraded event raised by the Infrared contract.
type InfraredUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Infrared *InfraredFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*InfraredUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &InfraredUpgradedIterator{contract: _Infrared.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Infrared *InfraredFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *InfraredUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredUpgraded)
				if err := _Infrared.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Infrared *InfraredFilterer) ParseUpgraded(log types.Log) (*InfraredUpgraded, error) {
	event := new(InfraredUpgraded)
	if err := _Infrared.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorCommissionActivatedIterator is returned from FilterValidatorCommissionActivated and is used to iterate over the raw logs and unpacked data for ValidatorCommissionActivated events raised by the Infrared contract.
type InfraredValidatorCommissionActivatedIterator struct {
	Event *InfraredValidatorCommissionActivated // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorCommissionActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorCommissionActivated)
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
		it.Event = new(InfraredValidatorCommissionActivated)
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
func (it *InfraredValidatorCommissionActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorCommissionActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorCommissionActivated represents a ValidatorCommissionActivated event raised by the Infrared contract.
type InfraredValidatorCommissionActivated struct {
	Operator       common.Address
	Pubkey         []byte
	CommissionRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorCommissionActivated is a free log retrieval operation binding the contract event 0x3bdc55fdb9ac22163caf219f1a92b71fb34b10f66d829f61aa440dbf533e2182.
//
// Solidity: event ValidatorCommissionActivated(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) FilterValidatorCommissionActivated(opts *bind.FilterOpts, operator []common.Address) (*InfraredValidatorCommissionActivatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorCommissionActivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorCommissionActivatedIterator{contract: _Infrared.contract, event: "ValidatorCommissionActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorCommissionActivated is a free log subscription operation binding the contract event 0x3bdc55fdb9ac22163caf219f1a92b71fb34b10f66d829f61aa440dbf533e2182.
//
// Solidity: event ValidatorCommissionActivated(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) WatchValidatorCommissionActivated(opts *bind.WatchOpts, sink chan<- *InfraredValidatorCommissionActivated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorCommissionActivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorCommissionActivated)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorCommissionActivated", log); err != nil {
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

// ParseValidatorCommissionActivated is a log parse operation binding the contract event 0x3bdc55fdb9ac22163caf219f1a92b71fb34b10f66d829f61aa440dbf533e2182.
//
// Solidity: event ValidatorCommissionActivated(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) ParseValidatorCommissionActivated(log types.Log) (*InfraredValidatorCommissionActivated, error) {
	event := new(InfraredValidatorCommissionActivated)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorCommissionActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorCommissionQueuedIterator is returned from FilterValidatorCommissionQueued and is used to iterate over the raw logs and unpacked data for ValidatorCommissionQueued events raised by the Infrared contract.
type InfraredValidatorCommissionQueuedIterator struct {
	Event *InfraredValidatorCommissionQueued // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorCommissionQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorCommissionQueued)
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
		it.Event = new(InfraredValidatorCommissionQueued)
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
func (it *InfraredValidatorCommissionQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorCommissionQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorCommissionQueued represents a ValidatorCommissionQueued event raised by the Infrared contract.
type InfraredValidatorCommissionQueued struct {
	Operator       common.Address
	Pubkey         []byte
	CommissionRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorCommissionQueued is a free log retrieval operation binding the contract event 0xab57e59e4883cf7d90a19fd4e915ed2951ec032d1e09cecf5b7cb9d20e8da925.
//
// Solidity: event ValidatorCommissionQueued(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) FilterValidatorCommissionQueued(opts *bind.FilterOpts, operator []common.Address) (*InfraredValidatorCommissionQueuedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorCommissionQueued", operatorRule)
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorCommissionQueuedIterator{contract: _Infrared.contract, event: "ValidatorCommissionQueued", logs: logs, sub: sub}, nil
}

// WatchValidatorCommissionQueued is a free log subscription operation binding the contract event 0xab57e59e4883cf7d90a19fd4e915ed2951ec032d1e09cecf5b7cb9d20e8da925.
//
// Solidity: event ValidatorCommissionQueued(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) WatchValidatorCommissionQueued(opts *bind.WatchOpts, sink chan<- *InfraredValidatorCommissionQueued, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorCommissionQueued", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorCommissionQueued)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorCommissionQueued", log); err != nil {
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

// ParseValidatorCommissionQueued is a log parse operation binding the contract event 0xab57e59e4883cf7d90a19fd4e915ed2951ec032d1e09cecf5b7cb9d20e8da925.
//
// Solidity: event ValidatorCommissionQueued(address indexed operator, bytes pubkey, uint96 commissionRate)
func (_Infrared *InfraredFilterer) ParseValidatorCommissionQueued(log types.Log) (*InfraredValidatorCommissionQueued, error) {
	event := new(InfraredValidatorCommissionQueued)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorCommissionQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorHarvestedIterator is returned from FilterValidatorHarvested and is used to iterate over the raw logs and unpacked data for ValidatorHarvested events raised by the Infrared contract.
type InfraredValidatorHarvestedIterator struct {
	Event *InfraredValidatorHarvested // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorHarvested)
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
		it.Event = new(InfraredValidatorHarvested)
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
func (it *InfraredValidatorHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorHarvested represents a ValidatorHarvested event raised by the Infrared contract.
type InfraredValidatorHarvested struct {
	Sender    common.Address
	Validator common.Hash
	Rewards   []DataTypesToken
	BgtAmt    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorHarvested is a free log retrieval operation binding the contract event 0x4134d1d53bea75c703a5d3b3a30b4d8b91e84d9b03335766e30674cd1e45d4cb.
//
// Solidity: event ValidatorHarvested(address _sender, bytes indexed _validator, (address,uint256)[] _rewards, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) FilterValidatorHarvested(opts *bind.FilterOpts, _validator [][]byte) (*InfraredValidatorHarvestedIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorHarvested", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorHarvestedIterator{contract: _Infrared.contract, event: "ValidatorHarvested", logs: logs, sub: sub}, nil
}

// WatchValidatorHarvested is a free log subscription operation binding the contract event 0x4134d1d53bea75c703a5d3b3a30b4d8b91e84d9b03335766e30674cd1e45d4cb.
//
// Solidity: event ValidatorHarvested(address _sender, bytes indexed _validator, (address,uint256)[] _rewards, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) WatchValidatorHarvested(opts *bind.WatchOpts, sink chan<- *InfraredValidatorHarvested, _validator [][]byte) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorHarvested", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorHarvested)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorHarvested", log); err != nil {
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

// ParseValidatorHarvested is a log parse operation binding the contract event 0x4134d1d53bea75c703a5d3b3a30b4d8b91e84d9b03335766e30674cd1e45d4cb.
//
// Solidity: event ValidatorHarvested(address _sender, bytes indexed _validator, (address,uint256)[] _rewards, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) ParseValidatorHarvested(log types.Log) (*InfraredValidatorHarvested, error) {
	event := new(InfraredValidatorHarvested)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorHarvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorReplacedIterator is returned from FilterValidatorReplaced and is used to iterate over the raw logs and unpacked data for ValidatorReplaced events raised by the Infrared contract.
type InfraredValidatorReplacedIterator struct {
	Event *InfraredValidatorReplaced // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorReplaced)
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
		it.Event = new(InfraredValidatorReplaced)
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
func (it *InfraredValidatorReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorReplaced represents a ValidatorReplaced event raised by the Infrared contract.
type InfraredValidatorReplaced struct {
	Sender  common.Address
	Current []byte
	New     []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorReplaced is a free log retrieval operation binding the contract event 0xb21ae6014605260f5065cfb6310215dc49bedceadf75b84b71bf28279eddca4a.
//
// Solidity: event ValidatorReplaced(address _sender, bytes _current, bytes _new)
func (_Infrared *InfraredFilterer) FilterValidatorReplaced(opts *bind.FilterOpts) (*InfraredValidatorReplacedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorReplaced")
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorReplacedIterator{contract: _Infrared.contract, event: "ValidatorReplaced", logs: logs, sub: sub}, nil
}

// WatchValidatorReplaced is a free log subscription operation binding the contract event 0xb21ae6014605260f5065cfb6310215dc49bedceadf75b84b71bf28279eddca4a.
//
// Solidity: event ValidatorReplaced(address _sender, bytes _current, bytes _new)
func (_Infrared *InfraredFilterer) WatchValidatorReplaced(opts *bind.WatchOpts, sink chan<- *InfraredValidatorReplaced) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorReplaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorReplaced)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorReplaced", log); err != nil {
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

// ParseValidatorReplaced is a log parse operation binding the contract event 0xb21ae6014605260f5065cfb6310215dc49bedceadf75b84b71bf28279eddca4a.
//
// Solidity: event ValidatorReplaced(address _sender, bytes _current, bytes _new)
func (_Infrared *InfraredFilterer) ParseValidatorReplaced(log types.Log) (*InfraredValidatorReplaced, error) {
	event := new(InfraredValidatorReplaced)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorReplaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorsAddedIterator is returned from FilterValidatorsAdded and is used to iterate over the raw logs and unpacked data for ValidatorsAdded events raised by the Infrared contract.
type InfraredValidatorsAddedIterator struct {
	Event *InfraredValidatorsAdded // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorsAdded)
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
		it.Event = new(InfraredValidatorsAdded)
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
func (it *InfraredValidatorsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorsAdded represents a ValidatorsAdded event raised by the Infrared contract.
type InfraredValidatorsAdded struct {
	Sender     common.Address
	Validators []ValidatorTypesValidator
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsAdded is a free log retrieval operation binding the contract event 0x5e67f9e05470203c71c0840375c93153a9fb6c1fd825aae5706b3d08c1dedddb.
//
// Solidity: event ValidatorsAdded(address _sender, (bytes,address)[] _validators)
func (_Infrared *InfraredFilterer) FilterValidatorsAdded(opts *bind.FilterOpts) (*InfraredValidatorsAddedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorsAdded")
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorsAddedIterator{contract: _Infrared.contract, event: "ValidatorsAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorsAdded is a free log subscription operation binding the contract event 0x5e67f9e05470203c71c0840375c93153a9fb6c1fd825aae5706b3d08c1dedddb.
//
// Solidity: event ValidatorsAdded(address _sender, (bytes,address)[] _validators)
func (_Infrared *InfraredFilterer) WatchValidatorsAdded(opts *bind.WatchOpts, sink chan<- *InfraredValidatorsAdded) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorsAdded)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorsAdded", log); err != nil {
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

// ParseValidatorsAdded is a log parse operation binding the contract event 0x5e67f9e05470203c71c0840375c93153a9fb6c1fd825aae5706b3d08c1dedddb.
//
// Solidity: event ValidatorsAdded(address _sender, (bytes,address)[] _validators)
func (_Infrared *InfraredFilterer) ParseValidatorsAdded(log types.Log) (*InfraredValidatorsAdded, error) {
	event := new(InfraredValidatorsAdded)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredValidatorsRemovedIterator is returned from FilterValidatorsRemoved and is used to iterate over the raw logs and unpacked data for ValidatorsRemoved events raised by the Infrared contract.
type InfraredValidatorsRemovedIterator struct {
	Event *InfraredValidatorsRemoved // Event containing the contract specifics and raw log

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
func (it *InfraredValidatorsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredValidatorsRemoved)
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
		it.Event = new(InfraredValidatorsRemoved)
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
func (it *InfraredValidatorsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredValidatorsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredValidatorsRemoved represents a ValidatorsRemoved event raised by the Infrared contract.
type InfraredValidatorsRemoved struct {
	Sender  common.Address
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorsRemoved is a free log retrieval operation binding the contract event 0x297b9ce1c3391a04fb73f5f0378f876fcebcd2c70a74cfdd7c00f82847e9194f.
//
// Solidity: event ValidatorsRemoved(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) FilterValidatorsRemoved(opts *bind.FilterOpts) (*InfraredValidatorsRemovedIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return &InfraredValidatorsRemovedIterator{contract: _Infrared.contract, event: "ValidatorsRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorsRemoved is a free log subscription operation binding the contract event 0x297b9ce1c3391a04fb73f5f0378f876fcebcd2c70a74cfdd7c00f82847e9194f.
//
// Solidity: event ValidatorsRemoved(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) WatchValidatorsRemoved(opts *bind.WatchOpts, sink chan<- *InfraredValidatorsRemoved) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "ValidatorsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredValidatorsRemoved)
				if err := _Infrared.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
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

// ParseValidatorsRemoved is a log parse operation binding the contract event 0x297b9ce1c3391a04fb73f5f0378f876fcebcd2c70a74cfdd7c00f82847e9194f.
//
// Solidity: event ValidatorsRemoved(address _sender, bytes[] _pubkeys)
func (_Infrared *InfraredFilterer) ParseValidatorsRemoved(log types.Log) (*InfraredValidatorsRemoved, error) {
	event := new(InfraredValidatorsRemoved)
	if err := _Infrared.contract.UnpackLog(event, "ValidatorsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultHarvestedIterator is returned from FilterVaultHarvested and is used to iterate over the raw logs and unpacked data for VaultHarvested events raised by the Infrared contract.
type InfraredVaultHarvestedIterator struct {
	Event *InfraredVaultHarvested // Event containing the contract specifics and raw log

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
func (it *InfraredVaultHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultHarvested)
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
		it.Event = new(InfraredVaultHarvested)
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
func (it *InfraredVaultHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultHarvested represents a VaultHarvested event raised by the Infrared contract.
type InfraredVaultHarvested struct {
	Sender common.Address
	Asset  common.Address
	Vault  common.Address
	BgtAmt *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultHarvested is a free log retrieval operation binding the contract event 0x546659bfb7e6ea18193d3711c72bc25c799442cc60749bb9d98bd4c56ca36f27.
//
// Solidity: event VaultHarvested(address _sender, address indexed _asset, address indexed _vault, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) FilterVaultHarvested(opts *bind.FilterOpts, _asset []common.Address, _vault []common.Address) (*InfraredVaultHarvestedIterator, error) {

	var _assetRule []interface{}
	for _, _assetItem := range _asset {
		_assetRule = append(_assetRule, _assetItem)
	}
	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "VaultHarvested", _assetRule, _vaultRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultHarvestedIterator{contract: _Infrared.contract, event: "VaultHarvested", logs: logs, sub: sub}, nil
}

// WatchVaultHarvested is a free log subscription operation binding the contract event 0x546659bfb7e6ea18193d3711c72bc25c799442cc60749bb9d98bd4c56ca36f27.
//
// Solidity: event VaultHarvested(address _sender, address indexed _asset, address indexed _vault, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) WatchVaultHarvested(opts *bind.WatchOpts, sink chan<- *InfraredVaultHarvested, _asset []common.Address, _vault []common.Address) (event.Subscription, error) {

	var _assetRule []interface{}
	for _, _assetItem := range _asset {
		_assetRule = append(_assetRule, _assetItem)
	}
	var _vaultRule []interface{}
	for _, _vaultItem := range _vault {
		_vaultRule = append(_vaultRule, _vaultItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "VaultHarvested", _assetRule, _vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultHarvested)
				if err := _Infrared.contract.UnpackLog(event, "VaultHarvested", log); err != nil {
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

// ParseVaultHarvested is a log parse operation binding the contract event 0x546659bfb7e6ea18193d3711c72bc25c799442cc60749bb9d98bd4c56ca36f27.
//
// Solidity: event VaultHarvested(address _sender, address indexed _asset, address indexed _vault, uint256 _bgtAmt)
func (_Infrared *InfraredFilterer) ParseVaultHarvested(log types.Log) (*InfraredVaultHarvested, error) {
	event := new(InfraredVaultHarvested)
	if err := _Infrared.contract.UnpackLog(event, "VaultHarvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultMigratedIterator is returned from FilterVaultMigrated and is used to iterate over the raw logs and unpacked data for VaultMigrated events raised by the Infrared contract.
type InfraredVaultMigratedIterator struct {
	Event *InfraredVaultMigrated // Event containing the contract specifics and raw log

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
func (it *InfraredVaultMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultMigrated)
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
		it.Event = new(InfraredVaultMigrated)
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
func (it *InfraredVaultMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultMigrated represents a VaultMigrated event raised by the Infrared contract.
type InfraredVaultMigrated struct {
	Sender   common.Address
	Asset    common.Address
	OldVault common.Address
	NewVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVaultMigrated is a free log retrieval operation binding the contract event 0xb18bfcac0b82d1210cde442dd427cf1b8e75f4ba5d76fe35c52398bf500373a5.
//
// Solidity: event VaultMigrated(address indexed sender, address indexed asset, address indexed oldVault, address newVault)
func (_Infrared *InfraredFilterer) FilterVaultMigrated(opts *bind.FilterOpts, sender []common.Address, asset []common.Address, oldVault []common.Address) (*InfraredVaultMigratedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var oldVaultRule []interface{}
	for _, oldVaultItem := range oldVault {
		oldVaultRule = append(oldVaultRule, oldVaultItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "VaultMigrated", senderRule, assetRule, oldVaultRule)
	if err != nil {
		return nil, err
	}
	return &InfraredVaultMigratedIterator{contract: _Infrared.contract, event: "VaultMigrated", logs: logs, sub: sub}, nil
}

// WatchVaultMigrated is a free log subscription operation binding the contract event 0xb18bfcac0b82d1210cde442dd427cf1b8e75f4ba5d76fe35c52398bf500373a5.
//
// Solidity: event VaultMigrated(address indexed sender, address indexed asset, address indexed oldVault, address newVault)
func (_Infrared *InfraredFilterer) WatchVaultMigrated(opts *bind.WatchOpts, sink chan<- *InfraredVaultMigrated, sender []common.Address, asset []common.Address, oldVault []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var oldVaultRule []interface{}
	for _, oldVaultItem := range oldVault {
		oldVaultRule = append(oldVaultRule, oldVaultItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "VaultMigrated", senderRule, assetRule, oldVaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultMigrated)
				if err := _Infrared.contract.UnpackLog(event, "VaultMigrated", log); err != nil {
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

// ParseVaultMigrated is a log parse operation binding the contract event 0xb18bfcac0b82d1210cde442dd427cf1b8e75f4ba5d76fe35c52398bf500373a5.
//
// Solidity: event VaultMigrated(address indexed sender, address indexed asset, address indexed oldVault, address newVault)
func (_Infrared *InfraredFilterer) ParseVaultMigrated(log types.Log) (*InfraredVaultMigrated, error) {
	event := new(InfraredVaultMigrated)
	if err := _Infrared.contract.UnpackLog(event, "VaultMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVaultRegistrationPauseStatusIterator is returned from FilterVaultRegistrationPauseStatus and is used to iterate over the raw logs and unpacked data for VaultRegistrationPauseStatus events raised by the Infrared contract.
type InfraredVaultRegistrationPauseStatusIterator struct {
	Event *InfraredVaultRegistrationPauseStatus // Event containing the contract specifics and raw log

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
func (it *InfraredVaultRegistrationPauseStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVaultRegistrationPauseStatus)
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
		it.Event = new(InfraredVaultRegistrationPauseStatus)
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
func (it *InfraredVaultRegistrationPauseStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVaultRegistrationPauseStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVaultRegistrationPauseStatus represents a VaultRegistrationPauseStatus event raised by the Infrared contract.
type InfraredVaultRegistrationPauseStatus struct {
	Pause bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultRegistrationPauseStatus is a free log retrieval operation binding the contract event 0x7446b242925aecc762aac930585d1135bece13513207e064626068285095c89e.
//
// Solidity: event VaultRegistrationPauseStatus(bool pause)
func (_Infrared *InfraredFilterer) FilterVaultRegistrationPauseStatus(opts *bind.FilterOpts) (*InfraredVaultRegistrationPauseStatusIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "VaultRegistrationPauseStatus")
	if err != nil {
		return nil, err
	}
	return &InfraredVaultRegistrationPauseStatusIterator{contract: _Infrared.contract, event: "VaultRegistrationPauseStatus", logs: logs, sub: sub}, nil
}

// WatchVaultRegistrationPauseStatus is a free log subscription operation binding the contract event 0x7446b242925aecc762aac930585d1135bece13513207e064626068285095c89e.
//
// Solidity: event VaultRegistrationPauseStatus(bool pause)
func (_Infrared *InfraredFilterer) WatchVaultRegistrationPauseStatus(opts *bind.WatchOpts, sink chan<- *InfraredVaultRegistrationPauseStatus) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "VaultRegistrationPauseStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVaultRegistrationPauseStatus)
				if err := _Infrared.contract.UnpackLog(event, "VaultRegistrationPauseStatus", log); err != nil {
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

// ParseVaultRegistrationPauseStatus is a log parse operation binding the contract event 0x7446b242925aecc762aac930585d1135bece13513207e064626068285095c89e.
//
// Solidity: event VaultRegistrationPauseStatus(bool pause)
func (_Infrared *InfraredFilterer) ParseVaultRegistrationPauseStatus(log types.Log) (*InfraredVaultRegistrationPauseStatus, error) {
	event := new(InfraredVaultRegistrationPauseStatus)
	if err := _Infrared.contract.UnpackLog(event, "VaultRegistrationPauseStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredVoterSetIterator is returned from FilterVoterSet and is used to iterate over the raw logs and unpacked data for VoterSet events raised by the Infrared contract.
type InfraredVoterSetIterator struct {
	Event *InfraredVoterSet // Event containing the contract specifics and raw log

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
func (it *InfraredVoterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredVoterSet)
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
		it.Event = new(InfraredVoterSet)
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
func (it *InfraredVoterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredVoterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredVoterSet represents a VoterSet event raised by the Infrared contract.
type InfraredVoterSet struct {
	Sender common.Address
	Voter  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVoterSet is a free log retrieval operation binding the contract event 0xb32f3288ab299698a7e6e1d7a3dc2cecab02b83346c3745708bd2121c9adb589.
//
// Solidity: event VoterSet(address _sender, address _voter)
func (_Infrared *InfraredFilterer) FilterVoterSet(opts *bind.FilterOpts) (*InfraredVoterSetIterator, error) {

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "VoterSet")
	if err != nil {
		return nil, err
	}
	return &InfraredVoterSetIterator{contract: _Infrared.contract, event: "VoterSet", logs: logs, sub: sub}, nil
}

// WatchVoterSet is a free log subscription operation binding the contract event 0xb32f3288ab299698a7e6e1d7a3dc2cecab02b83346c3745708bd2121c9adb589.
//
// Solidity: event VoterSet(address _sender, address _voter)
func (_Infrared *InfraredFilterer) WatchVoterSet(opts *bind.WatchOpts, sink chan<- *InfraredVoterSet) (event.Subscription, error) {

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "VoterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredVoterSet)
				if err := _Infrared.contract.UnpackLog(event, "VoterSet", log); err != nil {
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

// ParseVoterSet is a log parse operation binding the contract event 0xb32f3288ab299698a7e6e1d7a3dc2cecab02b83346c3745708bd2121c9adb589.
//
// Solidity: event VoterSet(address _sender, address _voter)
func (_Infrared *InfraredFilterer) ParseVoterSet(log types.Log) (*InfraredVoterSet, error) {
	event := new(InfraredVoterSet)
	if err := _Infrared.contract.UnpackLog(event, "VoterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfraredWhiteListedRewardTokensUpdatedIterator is returned from FilterWhiteListedRewardTokensUpdated and is used to iterate over the raw logs and unpacked data for WhiteListedRewardTokensUpdated events raised by the Infrared contract.
type InfraredWhiteListedRewardTokensUpdatedIterator struct {
	Event *InfraredWhiteListedRewardTokensUpdated // Event containing the contract specifics and raw log

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
func (it *InfraredWhiteListedRewardTokensUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfraredWhiteListedRewardTokensUpdated)
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
		it.Event = new(InfraredWhiteListedRewardTokensUpdated)
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
func (it *InfraredWhiteListedRewardTokensUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfraredWhiteListedRewardTokensUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfraredWhiteListedRewardTokensUpdated represents a WhiteListedRewardTokensUpdated event raised by the Infrared contract.
type InfraredWhiteListedRewardTokensUpdated struct {
	Sender         common.Address
	Token          common.Address
	WasWhitelisted bool
	IsWhitelisted  bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWhiteListedRewardTokensUpdated is a free log retrieval operation binding the contract event 0xdd01adaecf9dc9676b6096c7f319855230d8501da40b7b53c83f3c23d976f67f.
//
// Solidity: event WhiteListedRewardTokensUpdated(address _sender, address indexed _token, bool _wasWhitelisted, bool _isWhitelisted)
func (_Infrared *InfraredFilterer) FilterWhiteListedRewardTokensUpdated(opts *bind.FilterOpts, _token []common.Address) (*InfraredWhiteListedRewardTokensUpdatedIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.FilterLogs(opts, "WhiteListedRewardTokensUpdated", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &InfraredWhiteListedRewardTokensUpdatedIterator{contract: _Infrared.contract, event: "WhiteListedRewardTokensUpdated", logs: logs, sub: sub}, nil
}

// WatchWhiteListedRewardTokensUpdated is a free log subscription operation binding the contract event 0xdd01adaecf9dc9676b6096c7f319855230d8501da40b7b53c83f3c23d976f67f.
//
// Solidity: event WhiteListedRewardTokensUpdated(address _sender, address indexed _token, bool _wasWhitelisted, bool _isWhitelisted)
func (_Infrared *InfraredFilterer) WatchWhiteListedRewardTokensUpdated(opts *bind.WatchOpts, sink chan<- *InfraredWhiteListedRewardTokensUpdated, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Infrared.contract.WatchLogs(opts, "WhiteListedRewardTokensUpdated", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfraredWhiteListedRewardTokensUpdated)
				if err := _Infrared.contract.UnpackLog(event, "WhiteListedRewardTokensUpdated", log); err != nil {
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

// ParseWhiteListedRewardTokensUpdated is a log parse operation binding the contract event 0xdd01adaecf9dc9676b6096c7f319855230d8501da40b7b53c83f3c23d976f67f.
//
// Solidity: event WhiteListedRewardTokensUpdated(address _sender, address indexed _token, bool _wasWhitelisted, bool _isWhitelisted)
func (_Infrared *InfraredFilterer) ParseWhiteListedRewardTokensUpdated(log types.Log) (*InfraredWhiteListedRewardTokensUpdated, error) {
	event := new(InfraredWhiteListedRewardTokensUpdated)
	if err := _Infrared.contract.UnpackLog(event, "WhiteListedRewardTokensUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
