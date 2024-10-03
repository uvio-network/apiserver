// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package claimscontract

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

// ClaimsContractBindingV050MetaData contains all meta data concerning the ClaimsContractBindingV050 contract.
var ClaimsContractBindingV050MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unx\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Mapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"BalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"ClaimUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"DisputeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ProposeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"ProposeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ResolveCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROPOSER\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_TOTAL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_S\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_P\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_R\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_T\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_A\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_B\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_C\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_D\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MID_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_S\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"con\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"createDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"con\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"tok\",\"type\":\"address[]\"}],\"name\":\"createPropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ind\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"createResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationBasis\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMax\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMin\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"}],\"name\":\"searchBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchContent\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchExpired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchHistory\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchIndices\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchLatest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchPropose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"ind\",\"type\":\"uint8\"}],\"name\":\"searchResolve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchResults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchSamples\",\"outputs\":[{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"}],\"name\":\"updatePropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"}],\"name\":\"updateResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ClaimsContractBindingV050ABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimsContractBindingV050MetaData.ABI instead.
var ClaimsContractBindingV050ABI = ClaimsContractBindingV050MetaData.ABI

// ClaimsContractBindingV050 is an auto generated Go binding around an Ethereum contract.
type ClaimsContractBindingV050 struct {
	ClaimsContractBindingV050Caller     // Read-only binding to the contract
	ClaimsContractBindingV050Transactor // Write-only binding to the contract
	ClaimsContractBindingV050Filterer   // Log filterer for contract events
}

// ClaimsContractBindingV050Caller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimsContractBindingV050Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingV050Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimsContractBindingV050Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingV050Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimsContractBindingV050Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingV050Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimsContractBindingV050Session struct {
	Contract     *ClaimsContractBindingV050 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ClaimsContractBindingV050CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimsContractBindingV050CallerSession struct {
	Contract *ClaimsContractBindingV050Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ClaimsContractBindingV050TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimsContractBindingV050TransactorSession struct {
	Contract     *ClaimsContractBindingV050Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ClaimsContractBindingV050Raw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimsContractBindingV050Raw struct {
	Contract *ClaimsContractBindingV050 // Generic contract binding to access the raw methods on
}

// ClaimsContractBindingV050CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimsContractBindingV050CallerRaw struct {
	Contract *ClaimsContractBindingV050Caller // Generic read-only contract binding to access the raw methods on
}

// ClaimsContractBindingV050TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimsContractBindingV050TransactorRaw struct {
	Contract *ClaimsContractBindingV050Transactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimsContractBindingV050 creates a new instance of ClaimsContractBindingV050, bound to a specific deployed contract.
func NewClaimsContractBindingV050(address common.Address, backend bind.ContractBackend) (*ClaimsContractBindingV050, error) {
	contract, err := bindClaimsContractBindingV050(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050{ClaimsContractBindingV050Caller: ClaimsContractBindingV050Caller{contract: contract}, ClaimsContractBindingV050Transactor: ClaimsContractBindingV050Transactor{contract: contract}, ClaimsContractBindingV050Filterer: ClaimsContractBindingV050Filterer{contract: contract}}, nil
}

// NewClaimsContractBindingV050Caller creates a new read-only instance of ClaimsContractBindingV050, bound to a specific deployed contract.
func NewClaimsContractBindingV050Caller(address common.Address, caller bind.ContractCaller) (*ClaimsContractBindingV050Caller, error) {
	contract, err := bindClaimsContractBindingV050(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050Caller{contract: contract}, nil
}

// NewClaimsContractBindingV050Transactor creates a new write-only instance of ClaimsContractBindingV050, bound to a specific deployed contract.
func NewClaimsContractBindingV050Transactor(address common.Address, transactor bind.ContractTransactor) (*ClaimsContractBindingV050Transactor, error) {
	contract, err := bindClaimsContractBindingV050(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050Transactor{contract: contract}, nil
}

// NewClaimsContractBindingV050Filterer creates a new log filterer instance of ClaimsContractBindingV050, bound to a specific deployed contract.
func NewClaimsContractBindingV050Filterer(address common.Address, filterer bind.ContractFilterer) (*ClaimsContractBindingV050Filterer, error) {
	contract, err := bindClaimsContractBindingV050(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050Filterer{contract: contract}, nil
}

// bindClaimsContractBindingV050 binds a generic wrapper to an already deployed contract.
func bindClaimsContractBindingV050(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimsContractBindingV050MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimsContractBindingV050.Contract.ClaimsContractBindingV050Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.ClaimsContractBindingV050Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.ClaimsContractBindingV050Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimsContractBindingV050.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) ADDRESSSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "ADDRESS_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) ADDRESSSTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.ADDRESSSTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) ADDRESSSTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.ADDRESSSTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) ADDRESSSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "ADDRESS_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) ADDRESSSTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.ADDRESSSTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) ADDRESSSTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.ADDRESSSTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) BASISFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "BASIS_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) BASISFEE() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISFEE(&_ClaimsContractBindingV050.CallOpts)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) BASISFEE() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISFEE(&_ClaimsContractBindingV050.CallOpts)
}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) BASISPROPOSER(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "BASIS_PROPOSER")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) BASISPROPOSER() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISPROPOSER(&_ClaimsContractBindingV050.CallOpts)
}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) BASISPROPOSER() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISPROPOSER(&_ClaimsContractBindingV050.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) BASISPROTOCOL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "BASIS_PROTOCOL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) BASISPROTOCOL() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISPROTOCOL(&_ClaimsContractBindingV050.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) BASISPROTOCOL() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISPROTOCOL(&_ClaimsContractBindingV050.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) BASISTOTAL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "BASIS_TOTAL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) BASISTOTAL() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISTOTAL(&_ClaimsContractBindingV050.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) BASISTOTAL() (uint16, error) {
	return _ClaimsContractBindingV050.Contract.BASISTOTAL(&_ClaimsContractBindingV050.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) BOTROLE() ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.BOTROLE(&_ClaimsContractBindingV050.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) BOTROLE() ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.BOTROLE(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMADDRESSN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_ADDRESS_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMADDRESSN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMADDRESSN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMADDRESSN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMADDRESSN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMADDRESSY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_ADDRESS_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMADDRESSY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMADDRESSY(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMADDRESSY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMADDRESSY(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMBALANCES(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_BALANCE_S")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMBALANCES() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMBALANCES(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMBALANCES() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMBALANCES(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMBALANCEV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_BALANCE_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMBALANCEV() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMBALANCEV(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMBALANCEV() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMBALANCEV(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMEXPIRYP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_EXPIRY_P")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMEXPIRYP() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYP(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMEXPIRYP() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYP(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMEXPIRYR(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_EXPIRY_R")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMEXPIRYR() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYR(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMEXPIRYR() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYR(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMEXPIRYT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_EXPIRY_T")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMEXPIRYT() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYT(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMEXPIRYT() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMEXPIRYT(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKEA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_A")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKEA() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEA(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKEA() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEA(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKEB(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_B")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKEB() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEB(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKEB() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEB(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKEC(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_C")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKEC() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEC(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKEC() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEC(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_D")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKED() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKED(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKED() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKED(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMSTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMSTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMSTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMTRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMTRUTHN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMTRUTHN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMTRUTHN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMTRUTHN(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) CLAIMTRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "CLAIM_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CLAIMTRUTHY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMTRUTHY(&_ClaimsContractBindingV050.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) CLAIMTRUTHY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.CLAIMTRUTHY(&_ClaimsContractBindingV050.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.DEFAULTADMINROLE(&_ClaimsContractBindingV050.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.DEFAULTADMINROLE(&_ClaimsContractBindingV050.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) MAXUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "MAX_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) MAXUINT256() (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.MAXUINT256(&_ClaimsContractBindingV050.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) MAXUINT256() (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.MAXUINT256(&_ClaimsContractBindingV050.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) MIDUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "MID_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) MIDUINT256() (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.MIDUINT256(&_ClaimsContractBindingV050.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) MIDUINT256() (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.MIDUINT256(&_ClaimsContractBindingV050.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VERSION() (string, error) {
	return _ClaimsContractBindingV050.Contract.VERSION(&_ClaimsContractBindingV050.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VERSION() (string, error) {
	return _ClaimsContractBindingV050.Contract.VERSION(&_ClaimsContractBindingV050.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTESTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTESTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTESTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTESTAKEN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTESTAKEN(&_ClaimsContractBindingV050.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTESTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTESTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTESTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTESTAKEY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTESTAKEY(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTETRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTETRUTHN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHN(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTETRUTHN() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHN(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTETRUTHS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_TRUTH_S")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTETRUTHS() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHS(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTETRUTHS() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHS(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTETRUTHV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_TRUTH_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTETRUTHV() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHV(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTETRUTHV() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHV(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) VOTETRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "VOTE_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) VOTETRUTHY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHY(&_ClaimsContractBindingV050.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) VOTETRUTHY() (uint8, error) {
	return _ClaimsContractBindingV050.Contract.VOTETRUTHY(&_ClaimsContractBindingV050.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) DurationBasis(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "durationBasis")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) DurationBasis() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationBasis(&_ClaimsContractBindingV050.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) DurationBasis() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationBasis(&_ClaimsContractBindingV050.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) DurationMax(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "durationMax")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) DurationMax() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationMax(&_ClaimsContractBindingV050.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) DurationMax() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationMax(&_ClaimsContractBindingV050.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) DurationMin(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "durationMin")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) DurationMin() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationMin(&_ClaimsContractBindingV050.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) DurationMin() (uint64, error) {
	return _ClaimsContractBindingV050.Contract.DurationMin(&_ClaimsContractBindingV050.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleAdmin(&_ClaimsContractBindingV050.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleAdmin(&_ClaimsContractBindingV050.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleMember(&_ClaimsContractBindingV050.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleMember(&_ClaimsContractBindingV050.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleMemberCount(&_ClaimsContractBindingV050.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.GetRoleMemberCount(&_ClaimsContractBindingV050.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ClaimsContractBindingV050.Contract.HasRole(&_ClaimsContractBindingV050.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ClaimsContractBindingV050.Contract.HasRole(&_ClaimsContractBindingV050.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) Owner() (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.Owner(&_ClaimsContractBindingV050.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) Owner() (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.Owner(&_ClaimsContractBindingV050.CallOpts)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchBalance(opts *bind.CallOpts, use common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchBalance", use)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchBalance(&_ClaimsContractBindingV050.CallOpts, use)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchBalance(&_ClaimsContractBindingV050.CallOpts, use)
}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchContent(opts *bind.CallOpts, pod *big.Int) (string, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchContent", pod)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchContent(pod *big.Int) (string, error) {
	return _ClaimsContractBindingV050.Contract.SearchContent(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchContent(pod *big.Int) (string, error) {
	return _ClaimsContractBindingV050.Contract.SearchContent(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchExpired(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchExpired", pod)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchExpired(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchExpired(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchExpired(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchExpired(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchHistory(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchHistory", pod, lef, rig)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchHistory(pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchHistory(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchHistory(pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchHistory(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256[8])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchIndices(opts *bind.CallOpts, pod *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchIndices", pod)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256[8])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchIndices(pod *big.Int) ([8]*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchIndices(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256[8])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchIndices(pod *big.Int) ([8]*big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchIndices(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchLatest is a free data retrieval call binding the contract method 0x5d0d720f.
//
// Solidity: function searchLatest(uint256 pod) view returns(uint8, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchLatest(opts *bind.CallOpts, pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchLatest", pod)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// SearchLatest is a free data retrieval call binding the contract method 0x5d0d720f.
//
// Solidity: function searchLatest(uint256 pod) view returns(uint8, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchLatest(pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchLatest(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchLatest is a free data retrieval call binding the contract method 0x5d0d720f.
//
// Solidity: function searchLatest(uint256 pod) view returns(uint8, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchLatest(pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchLatest(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pod) view returns(uint256, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchPropose(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchPropose", pod)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pod) view returns(uint256, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchPropose(pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchPropose(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pod) view returns(uint256, uint256, uint256)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchPropose(pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBindingV050.Contract.SearchPropose(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchResolve(opts *bind.CallOpts, pod *big.Int, ind uint8) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchResolve", pod, ind)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchResolve(pod *big.Int, ind uint8) (bool, error) {
	return _ClaimsContractBindingV050.Contract.SearchResolve(&_ClaimsContractBindingV050.CallOpts, pod, ind)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchResolve(pod *big.Int, ind uint8) (bool, error) {
	return _ClaimsContractBindingV050.Contract.SearchResolve(&_ClaimsContractBindingV050.CallOpts, pod, ind)
}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool, bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchResults(opts *bind.CallOpts, pod *big.Int) (bool, bool, bool, bool, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchResults", pod)

	if err != nil {
		return *new(bool), *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)

	return out0, out1, out2, out3, err

}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool, bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchResults(pod *big.Int) (bool, bool, bool, bool, error) {
	return _ClaimsContractBindingV050.Contract.SearchResults(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool, bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchResults(pod *big.Int) (bool, bool, bool, bool, error) {
	return _ClaimsContractBindingV050.Contract.SearchResults(&_ClaimsContractBindingV050.CallOpts, pod)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(uint8[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchSamples(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]uint8, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchSamples", pod, lef, rig)

	if err != nil {
		return *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint8)).(*[]uint8)

	return out0, err

}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(uint8[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchSamples(pod *big.Int, lef *big.Int, rig *big.Int) ([]uint8, error) {
	return _ClaimsContractBindingV050.Contract.SearchSamples(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(uint8[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchSamples(pod *big.Int, lef *big.Int, rig *big.Int) ([]uint8, error) {
	return _ClaimsContractBindingV050.Contract.SearchSamples(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchStakers(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchStakers", pod, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchStakers(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchStakers(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchStakers(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchStakers(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchToken(opts *bind.CallOpts, pro *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchToken", pro)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchToken(&_ClaimsContractBindingV050.CallOpts, pro)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchToken(&_ClaimsContractBindingV050.CallOpts, pro)
}

// SearchVoters is a free data retrieval call binding the contract method 0x0867d8fc.
//
// Solidity: function searchVoters(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SearchVoters(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "searchVoters", pod, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchVoters is a free data retrieval call binding the contract method 0x0867d8fc.
//
// Solidity: function searchVoters(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SearchVoters(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchVoters(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SearchVoters is a free data retrieval call binding the contract method 0x0867d8fc.
//
// Solidity: function searchVoters(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SearchVoters(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBindingV050.Contract.SearchVoters(&_ClaimsContractBindingV050.CallOpts, pod, lef, rig)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ClaimsContractBindingV050.Contract.SupportsInterface(&_ClaimsContractBindingV050.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ClaimsContractBindingV050.Contract.SupportsInterface(&_ClaimsContractBindingV050.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBindingV050.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) Token() (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.Token(&_ClaimsContractBindingV050.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050CallerSession) Token() (common.Address, error) {
	return _ClaimsContractBindingV050.Contract.Token(&_ClaimsContractBindingV050.CallOpts)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) CreateDispute(opts *bind.TransactOpts, dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "createDispute", dis, bal, vot, exp, con, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreateDispute(&_ClaimsContractBindingV050.TransactOpts, dis, bal, vot, exp, con, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreateDispute(&_ClaimsContractBindingV050.TransactOpts, dis, bal, vot, exp, con, pro)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) CreatePropose(opts *bind.TransactOpts, pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "createPropose", pro, bal, vot, exp, con, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreatePropose(&_ClaimsContractBindingV050.TransactOpts, pro, bal, vot, exp, con, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreatePropose(&_ClaimsContractBindingV050.TransactOpts, pro, bal, vot, exp, con, tok)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) CreateResolve(opts *bind.TransactOpts, pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "createResolve", pod, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) CreateResolve(pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreateResolve(&_ClaimsContractBindingV050.TransactOpts, pod, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) CreateResolve(pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.CreateResolve(&_ClaimsContractBindingV050.TransactOpts, pod, ind, exp)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.GrantRole(&_ClaimsContractBindingV050.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.GrantRole(&_ClaimsContractBindingV050.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.RenounceRole(&_ClaimsContractBindingV050.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.RenounceRole(&_ClaimsContractBindingV050.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.RevokeRole(&_ClaimsContractBindingV050.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.RevokeRole(&_ClaimsContractBindingV050.TransactOpts, role, account)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) UpdateBalance(opts *bind.TransactOpts, pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "updateBalance", pod, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) UpdateBalance(pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateBalance(&_ClaimsContractBindingV050.TransactOpts, pod, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) UpdateBalance(pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateBalance(&_ClaimsContractBindingV050.TransactOpts, pod, max)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) UpdateDuration(opts *bind.TransactOpts, bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "updateDuration", bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateDuration(&_ClaimsContractBindingV050.TransactOpts, bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateDuration(&_ClaimsContractBindingV050.TransactOpts, bas, max, min)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) UpdateOwner(opts *bind.TransactOpts, own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "updateOwner", own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateOwner(&_ClaimsContractBindingV050.TransactOpts, own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateOwner(&_ClaimsContractBindingV050.TransactOpts, own)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) UpdatePropose(opts *bind.TransactOpts, pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "updatePropose", pod, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) UpdatePropose(pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdatePropose(&_ClaimsContractBindingV050.TransactOpts, pod, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) UpdatePropose(pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdatePropose(&_ClaimsContractBindingV050.TransactOpts, pod, bal, vot, tok)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) UpdateResolve(opts *bind.TransactOpts, pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "updateResolve", pod, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) UpdateResolve(pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateResolve(&_ClaimsContractBindingV050.TransactOpts, pod, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) UpdateResolve(pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.UpdateResolve(&_ClaimsContractBindingV050.TransactOpts, pod, vot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) Withdraw(opts *bind.TransactOpts, bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.Transact(opts, "withdraw", bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.Withdraw(&_ClaimsContractBindingV050.TransactOpts, bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.Withdraw(&_ClaimsContractBindingV050.TransactOpts, bal)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBindingV050.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Session) Receive() (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.Receive(&_ClaimsContractBindingV050.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050TransactorSession) Receive() (*types.Transaction, error) {
	return _ClaimsContractBindingV050.Contract.Receive(&_ClaimsContractBindingV050.TransactOpts)
}

// ClaimsContractBindingV050BalanceUpdatedIterator is returned from FilterBalanceUpdated and is used to iterate over the raw logs and unpacked data for BalanceUpdated events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050BalanceUpdatedIterator struct {
	Event *ClaimsContractBindingV050BalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050BalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050BalanceUpdated)
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
		it.Event = new(ClaimsContractBindingV050BalanceUpdated)
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
func (it *ClaimsContractBindingV050BalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050BalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050BalanceUpdated represents a BalanceUpdated event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050BalanceUpdated struct {
	Pod *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBalanceUpdated is a free log retrieval operation binding the contract event 0x04be25a4c0fe4c39136ca055daaac08ac6b3041b2e1f2a1b698bac12f1cbdc4a.
//
// Solidity: event BalanceUpdated(uint256 pod)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterBalanceUpdated(opts *bind.FilterOpts) (*ClaimsContractBindingV050BalanceUpdatedIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "BalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050BalanceUpdatedIterator{contract: _ClaimsContractBindingV050.contract, event: "BalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceUpdated is a free log subscription operation binding the contract event 0x04be25a4c0fe4c39136ca055daaac08ac6b3041b2e1f2a1b698bac12f1cbdc4a.
//
// Solidity: event BalanceUpdated(uint256 pod)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050BalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "BalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050BalanceUpdated)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
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

// ParseBalanceUpdated is a log parse operation binding the contract event 0x04be25a4c0fe4c39136ca055daaac08ac6b3041b2e1f2a1b698bac12f1cbdc4a.
//
// Solidity: event BalanceUpdated(uint256 pod)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseBalanceUpdated(log types.Log) (*ClaimsContractBindingV050BalanceUpdated, error) {
	event := new(ClaimsContractBindingV050BalanceUpdated)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050ClaimUpdatedIterator is returned from FilterClaimUpdated and is used to iterate over the raw logs and unpacked data for ClaimUpdated events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ClaimUpdatedIterator struct {
	Event *ClaimsContractBindingV050ClaimUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050ClaimUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050ClaimUpdated)
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
		it.Event = new(ClaimsContractBindingV050ClaimUpdated)
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
func (it *ClaimsContractBindingV050ClaimUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050ClaimUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050ClaimUpdated represents a ClaimUpdated event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ClaimUpdated struct {
	Pod *big.Int
	Use common.Address
	Bal *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimUpdated is a free log retrieval operation binding the contract event 0x456e7c3ba307b1c39b4aa965dad80444be2ddffd40b0d6e930507c5858faf33e.
//
// Solidity: event ClaimUpdated(uint256 pod, address use, uint256 bal)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterClaimUpdated(opts *bind.FilterOpts) (*ClaimsContractBindingV050ClaimUpdatedIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "ClaimUpdated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050ClaimUpdatedIterator{contract: _ClaimsContractBindingV050.contract, event: "ClaimUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimUpdated is a free log subscription operation binding the contract event 0x456e7c3ba307b1c39b4aa965dad80444be2ddffd40b0d6e930507c5858faf33e.
//
// Solidity: event ClaimUpdated(uint256 pod, address use, uint256 bal)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchClaimUpdated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050ClaimUpdated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "ClaimUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050ClaimUpdated)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
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

// ParseClaimUpdated is a log parse operation binding the contract event 0x456e7c3ba307b1c39b4aa965dad80444be2ddffd40b0d6e930507c5858faf33e.
//
// Solidity: event ClaimUpdated(uint256 pod, address use, uint256 bal)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseClaimUpdated(log types.Log) (*ClaimsContractBindingV050ClaimUpdated, error) {
	event := new(ClaimsContractBindingV050ClaimUpdated)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050DisputeCreatedIterator is returned from FilterDisputeCreated and is used to iterate over the raw logs and unpacked data for DisputeCreated events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050DisputeCreatedIterator struct {
	Event *ClaimsContractBindingV050DisputeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050DisputeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050DisputeCreated)
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
		it.Event = new(ClaimsContractBindingV050DisputeCreated)
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
func (it *ClaimsContractBindingV050DisputeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050DisputeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050DisputeCreated represents a DisputeCreated event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050DisputeCreated struct {
	Dis *big.Int
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisputeCreated is a free log retrieval operation binding the contract event 0x848ff8058c972bb6cd367b414b3cd1799ef42a941e7800185f0e580bda9c3cc5.
//
// Solidity: event DisputeCreated(uint256 dis, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterDisputeCreated(opts *bind.FilterOpts) (*ClaimsContractBindingV050DisputeCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050DisputeCreatedIterator{contract: _ClaimsContractBindingV050.contract, event: "DisputeCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCreated is a free log subscription operation binding the contract event 0x848ff8058c972bb6cd367b414b3cd1799ef42a941e7800185f0e580bda9c3cc5.
//
// Solidity: event DisputeCreated(uint256 dis, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchDisputeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050DisputeCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050DisputeCreated)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
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

// ParseDisputeCreated is a log parse operation binding the contract event 0x848ff8058c972bb6cd367b414b3cd1799ef42a941e7800185f0e580bda9c3cc5.
//
// Solidity: event DisputeCreated(uint256 dis, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseDisputeCreated(log types.Log) (*ClaimsContractBindingV050DisputeCreated, error) {
	event := new(ClaimsContractBindingV050DisputeCreated)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050DisputeSettledIterator is returned from FilterDisputeSettled and is used to iterate over the raw logs and unpacked data for DisputeSettled events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050DisputeSettledIterator struct {
	Event *ClaimsContractBindingV050DisputeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050DisputeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050DisputeSettled)
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
		it.Event = new(ClaimsContractBindingV050DisputeSettled)
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
func (it *ClaimsContractBindingV050DisputeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050DisputeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050DisputeSettled represents a DisputeSettled event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050DisputeSettled struct {
	Dis *big.Int
	All *big.Int
	Yay *big.Int
	Nah *big.Int
	Tot *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisputeSettled is a free log retrieval operation binding the contract event 0x885d609f8d9d1fc5b04cdedac72087d79d587398e79a263f148e64f9275fc929.
//
// Solidity: event DisputeSettled(uint256 dis, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterDisputeSettled(opts *bind.FilterOpts) (*ClaimsContractBindingV050DisputeSettledIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050DisputeSettledIterator{contract: _ClaimsContractBindingV050.contract, event: "DisputeSettled", logs: logs, sub: sub}, nil
}

// WatchDisputeSettled is a free log subscription operation binding the contract event 0x885d609f8d9d1fc5b04cdedac72087d79d587398e79a263f148e64f9275fc929.
//
// Solidity: event DisputeSettled(uint256 dis, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchDisputeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050DisputeSettled) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050DisputeSettled)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
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

// ParseDisputeSettled is a log parse operation binding the contract event 0x885d609f8d9d1fc5b04cdedac72087d79d587398e79a263f148e64f9275fc929.
//
// Solidity: event DisputeSettled(uint256 dis, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseDisputeSettled(log types.Log) (*ClaimsContractBindingV050DisputeSettled, error) {
	event := new(ClaimsContractBindingV050DisputeSettled)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050ProposeCreatedIterator is returned from FilterProposeCreated and is used to iterate over the raw logs and unpacked data for ProposeCreated events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ProposeCreatedIterator struct {
	Event *ClaimsContractBindingV050ProposeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050ProposeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050ProposeCreated)
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
		it.Event = new(ClaimsContractBindingV050ProposeCreated)
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
func (it *ClaimsContractBindingV050ProposeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050ProposeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050ProposeCreated represents a ProposeCreated event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ProposeCreated struct {
	Pro *big.Int
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposeCreated is a free log retrieval operation binding the contract event 0x560897c1e1894395c161cb6ee32a257c436c4f5e19d550b5b80a6043fd166c78.
//
// Solidity: event ProposeCreated(uint256 pro, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterProposeCreated(opts *bind.FilterOpts) (*ClaimsContractBindingV050ProposeCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050ProposeCreatedIterator{contract: _ClaimsContractBindingV050.contract, event: "ProposeCreated", logs: logs, sub: sub}, nil
}

// WatchProposeCreated is a free log subscription operation binding the contract event 0x560897c1e1894395c161cb6ee32a257c436c4f5e19d550b5b80a6043fd166c78.
//
// Solidity: event ProposeCreated(uint256 pro, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchProposeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050ProposeCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050ProposeCreated)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
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

// ParseProposeCreated is a log parse operation binding the contract event 0x560897c1e1894395c161cb6ee32a257c436c4f5e19d550b5b80a6043fd166c78.
//
// Solidity: event ProposeCreated(uint256 pro, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseProposeCreated(log types.Log) (*ClaimsContractBindingV050ProposeCreated, error) {
	event := new(ClaimsContractBindingV050ProposeCreated)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050ProposeSettledIterator is returned from FilterProposeSettled and is used to iterate over the raw logs and unpacked data for ProposeSettled events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ProposeSettledIterator struct {
	Event *ClaimsContractBindingV050ProposeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050ProposeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050ProposeSettled)
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
		it.Event = new(ClaimsContractBindingV050ProposeSettled)
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
func (it *ClaimsContractBindingV050ProposeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050ProposeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050ProposeSettled represents a ProposeSettled event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ProposeSettled struct {
	Pro *big.Int
	All *big.Int
	Yay *big.Int
	Nah *big.Int
	Tot *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposeSettled is a free log retrieval operation binding the contract event 0x02b6b8b1e0e988010b10993e9424cd6b9a8f599a653942c3fec3756d53f14e68.
//
// Solidity: event ProposeSettled(uint256 pro, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterProposeSettled(opts *bind.FilterOpts) (*ClaimsContractBindingV050ProposeSettledIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050ProposeSettledIterator{contract: _ClaimsContractBindingV050.contract, event: "ProposeSettled", logs: logs, sub: sub}, nil
}

// WatchProposeSettled is a free log subscription operation binding the contract event 0x02b6b8b1e0e988010b10993e9424cd6b9a8f599a653942c3fec3756d53f14e68.
//
// Solidity: event ProposeSettled(uint256 pro, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchProposeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050ProposeSettled) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050ProposeSettled)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
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

// ParseProposeSettled is a log parse operation binding the contract event 0x02b6b8b1e0e988010b10993e9424cd6b9a8f599a653942c3fec3756d53f14e68.
//
// Solidity: event ProposeSettled(uint256 pro, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseProposeSettled(log types.Log) (*ClaimsContractBindingV050ProposeSettled, error) {
	event := new(ClaimsContractBindingV050ProposeSettled)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050ResolveCreatedIterator is returned from FilterResolveCreated and is used to iterate over the raw logs and unpacked data for ResolveCreated events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ResolveCreatedIterator struct {
	Event *ClaimsContractBindingV050ResolveCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050ResolveCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050ResolveCreated)
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
		it.Event = new(ClaimsContractBindingV050ResolveCreated)
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
func (it *ClaimsContractBindingV050ResolveCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050ResolveCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050ResolveCreated represents a ResolveCreated event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050ResolveCreated struct {
	Pod *big.Int
	Use common.Address
	Len *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResolveCreated is a free log retrieval operation binding the contract event 0xe3c5a61bbb499af1d75605281b4c8c6e4635121595fe0cdc3ef366e28204dbef.
//
// Solidity: event ResolveCreated(uint256 pod, address use, uint256 len, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterResolveCreated(opts *bind.FilterOpts) (*ClaimsContractBindingV050ResolveCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050ResolveCreatedIterator{contract: _ClaimsContractBindingV050.contract, event: "ResolveCreated", logs: logs, sub: sub}, nil
}

// WatchResolveCreated is a free log subscription operation binding the contract event 0xe3c5a61bbb499af1d75605281b4c8c6e4635121595fe0cdc3ef366e28204dbef.
//
// Solidity: event ResolveCreated(uint256 pod, address use, uint256 len, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchResolveCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050ResolveCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050ResolveCreated)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
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

// ParseResolveCreated is a log parse operation binding the contract event 0xe3c5a61bbb499af1d75605281b4c8c6e4635121595fe0cdc3ef366e28204dbef.
//
// Solidity: event ResolveCreated(uint256 pod, address use, uint256 len, uint64 exp)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseResolveCreated(log types.Log) (*ClaimsContractBindingV050ResolveCreated, error) {
	event := new(ClaimsContractBindingV050ResolveCreated)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleAdminChangedIterator struct {
	Event *ClaimsContractBindingV050RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050RoleAdminChanged)
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
		it.Event = new(ClaimsContractBindingV050RoleAdminChanged)
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
func (it *ClaimsContractBindingV050RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050RoleAdminChanged represents a RoleAdminChanged event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ClaimsContractBindingV050RoleAdminChangedIterator, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050RoleAdminChangedIterator{contract: _ClaimsContractBindingV050.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050RoleAdminChanged)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseRoleAdminChanged(log types.Log) (*ClaimsContractBindingV050RoleAdminChanged, error) {
	event := new(ClaimsContractBindingV050RoleAdminChanged)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleGrantedIterator struct {
	Event *ClaimsContractBindingV050RoleGranted // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050RoleGranted)
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
		it.Event = new(ClaimsContractBindingV050RoleGranted)
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
func (it *ClaimsContractBindingV050RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050RoleGranted represents a RoleGranted event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsContractBindingV050RoleGrantedIterator, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050RoleGrantedIterator{contract: _ClaimsContractBindingV050.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050RoleGranted)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseRoleGranted(log types.Log) (*ClaimsContractBindingV050RoleGranted, error) {
	event := new(ClaimsContractBindingV050RoleGranted)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingV050RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleRevokedIterator struct {
	Event *ClaimsContractBindingV050RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingV050RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingV050RoleRevoked)
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
		it.Event = new(ClaimsContractBindingV050RoleRevoked)
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
func (it *ClaimsContractBindingV050RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingV050RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingV050RoleRevoked represents a RoleRevoked event raised by the ClaimsContractBindingV050 contract.
type ClaimsContractBindingV050RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsContractBindingV050RoleRevokedIterator, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingV050RoleRevokedIterator{contract: _ClaimsContractBindingV050.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingV050RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBindingV050.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingV050RoleRevoked)
				if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ClaimsContractBindingV050 *ClaimsContractBindingV050Filterer) ParseRoleRevoked(log types.Log) (*ClaimsContractBindingV050RoleRevoked, error) {
	event := new(ClaimsContractBindingV050RoleRevoked)
	if err := _ClaimsContractBindingV050.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
