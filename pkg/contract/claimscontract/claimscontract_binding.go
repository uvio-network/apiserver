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

// ClaimsContractBindingMetaData contains all meta data concerning the ClaimsContractBinding contract.
var ClaimsContractBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unx\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Mapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"BalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"ClaimUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"DisputeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ProposeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"ProposeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ResolveCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_FEE\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROPOSER\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_TOTAL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_S\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_P\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_R\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_T\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_A\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_B\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_C\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_D\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MID_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_S\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"con\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"createDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"con\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"tok\",\"type\":\"address[]\"}],\"name\":\"createPropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ind\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"createResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationBasis\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMax\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMin\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"}],\"name\":\"searchBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchContent\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchExpired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchHistory\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchLatest\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchPropose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"ind\",\"type\":\"uint8\"}],\"name\":\"searchResolve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchResults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchSamples\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"}],\"name\":\"searchVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"}],\"name\":\"updatePropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pod\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"}],\"name\":\"updateResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ClaimsContractBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimsContractBindingMetaData.ABI instead.
var ClaimsContractBindingABI = ClaimsContractBindingMetaData.ABI

// ClaimsContractBinding is an auto generated Go binding around an Ethereum contract.
type ClaimsContractBinding struct {
	ClaimsContractBindingCaller     // Read-only binding to the contract
	ClaimsContractBindingTransactor // Write-only binding to the contract
	ClaimsContractBindingFilterer   // Log filterer for contract events
}

// ClaimsContractBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimsContractBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimsContractBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimsContractBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsContractBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimsContractBindingSession struct {
	Contract     *ClaimsContractBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ClaimsContractBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimsContractBindingCallerSession struct {
	Contract *ClaimsContractBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ClaimsContractBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimsContractBindingTransactorSession struct {
	Contract     *ClaimsContractBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ClaimsContractBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimsContractBindingRaw struct {
	Contract *ClaimsContractBinding // Generic contract binding to access the raw methods on
}

// ClaimsContractBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimsContractBindingCallerRaw struct {
	Contract *ClaimsContractBindingCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimsContractBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimsContractBindingTransactorRaw struct {
	Contract *ClaimsContractBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimsContractBinding creates a new instance of ClaimsContractBinding, bound to a specific deployed contract.
func NewClaimsContractBinding(address common.Address, backend bind.ContractBackend) (*ClaimsContractBinding, error) {
	contract, err := bindClaimsContractBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBinding{ClaimsContractBindingCaller: ClaimsContractBindingCaller{contract: contract}, ClaimsContractBindingTransactor: ClaimsContractBindingTransactor{contract: contract}, ClaimsContractBindingFilterer: ClaimsContractBindingFilterer{contract: contract}}, nil
}

// NewClaimsContractBindingCaller creates a new read-only instance of ClaimsContractBinding, bound to a specific deployed contract.
func NewClaimsContractBindingCaller(address common.Address, caller bind.ContractCaller) (*ClaimsContractBindingCaller, error) {
	contract, err := bindClaimsContractBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingCaller{contract: contract}, nil
}

// NewClaimsContractBindingTransactor creates a new write-only instance of ClaimsContractBinding, bound to a specific deployed contract.
func NewClaimsContractBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimsContractBindingTransactor, error) {
	contract, err := bindClaimsContractBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingTransactor{contract: contract}, nil
}

// NewClaimsContractBindingFilterer creates a new log filterer instance of ClaimsContractBinding, bound to a specific deployed contract.
func NewClaimsContractBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimsContractBindingFilterer, error) {
	contract, err := bindClaimsContractBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingFilterer{contract: contract}, nil
}

// bindClaimsContractBinding binds a generic wrapper to an already deployed contract.
func bindClaimsContractBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimsContractBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimsContractBinding *ClaimsContractBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimsContractBinding.Contract.ClaimsContractBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimsContractBinding *ClaimsContractBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.ClaimsContractBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimsContractBinding *ClaimsContractBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.ClaimsContractBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimsContractBinding *ClaimsContractBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimsContractBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimsContractBinding *ClaimsContractBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimsContractBinding *ClaimsContractBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) ADDRESSSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "ADDRESS_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) ADDRESSSTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.ADDRESSSTAKEN(&_ClaimsContractBinding.CallOpts)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) ADDRESSSTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.ADDRESSSTAKEN(&_ClaimsContractBinding.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) ADDRESSSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "ADDRESS_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) ADDRESSSTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.ADDRESSSTAKEY(&_ClaimsContractBinding.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) ADDRESSSTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.ADDRESSSTAKEY(&_ClaimsContractBinding.CallOpts)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) BASISFEE(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "BASIS_FEE")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingSession) BASISFEE() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISFEE(&_ClaimsContractBinding.CallOpts)
}

// BASISFEE is a free data retrieval call binding the contract method 0x44c6d89b.
//
// Solidity: function BASIS_FEE() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) BASISFEE() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISFEE(&_ClaimsContractBinding.CallOpts)
}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) BASISPROPOSER(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "BASIS_PROPOSER")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingSession) BASISPROPOSER() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISPROPOSER(&_ClaimsContractBinding.CallOpts)
}

// BASISPROPOSER is a free data retrieval call binding the contract method 0xa2b15ab5.
//
// Solidity: function BASIS_PROPOSER() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) BASISPROPOSER() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISPROPOSER(&_ClaimsContractBinding.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) BASISPROTOCOL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "BASIS_PROTOCOL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingSession) BASISPROTOCOL() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISPROTOCOL(&_ClaimsContractBinding.CallOpts)
}

// BASISPROTOCOL is a free data retrieval call binding the contract method 0xc926e70f.
//
// Solidity: function BASIS_PROTOCOL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) BASISPROTOCOL() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISPROTOCOL(&_ClaimsContractBinding.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) BASISTOTAL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "BASIS_TOTAL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingSession) BASISTOTAL() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISTOTAL(&_ClaimsContractBinding.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) BASISTOTAL() (uint16, error) {
	return _ClaimsContractBinding.Contract.BASISTOTAL(&_ClaimsContractBinding.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingSession) BOTROLE() ([32]byte, error) {
	return _ClaimsContractBinding.Contract.BOTROLE(&_ClaimsContractBinding.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) BOTROLE() ([32]byte, error) {
	return _ClaimsContractBinding.Contract.BOTROLE(&_ClaimsContractBinding.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMADDRESSN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_ADDRESS_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMADDRESSN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMADDRESSN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMADDRESSN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMADDRESSN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMADDRESSY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_ADDRESS_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMADDRESSY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMADDRESSY(&_ClaimsContractBinding.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMADDRESSY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMADDRESSY(&_ClaimsContractBinding.CallOpts)
}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMBALANCES(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_BALANCE_S")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMBALANCES() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMBALANCES(&_ClaimsContractBinding.CallOpts)
}

// CLAIMBALANCES is a free data retrieval call binding the contract method 0x379ef117.
//
// Solidity: function CLAIM_BALANCE_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMBALANCES() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMBALANCES(&_ClaimsContractBinding.CallOpts)
}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMBALANCEV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_BALANCE_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMBALANCEV() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMBALANCEV(&_ClaimsContractBinding.CallOpts)
}

// CLAIMBALANCEV is a free data retrieval call binding the contract method 0xb2b8fd75.
//
// Solidity: function CLAIM_BALANCE_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMBALANCEV() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMBALANCEV(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMEXPIRYP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_EXPIRY_P")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMEXPIRYP() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYP(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMEXPIRYP() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYP(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMEXPIRYR(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_EXPIRY_R")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMEXPIRYR() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYR(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMEXPIRYR() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYR(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMEXPIRYT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_EXPIRY_T")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMEXPIRYT() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYT(&_ClaimsContractBinding.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMEXPIRYT() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMEXPIRYT(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKEA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_A")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKEA() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEA(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKEA() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEA(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKEB(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_B")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKEB() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEB(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKEB() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEB(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKEC(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_C")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKEC() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEC(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKEC() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEC(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_D")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKED() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKED(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKED() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKED(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMSTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEY(&_ClaimsContractBinding.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMSTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMSTAKEY(&_ClaimsContractBinding.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMTRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMTRUTHN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMTRUTHN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMTRUTHN() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMTRUTHN(&_ClaimsContractBinding.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) CLAIMTRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "CLAIM_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) CLAIMTRUTHY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMTRUTHY(&_ClaimsContractBinding.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) CLAIMTRUTHY() (uint8, error) {
	return _ClaimsContractBinding.Contract.CLAIMTRUTHY(&_ClaimsContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ClaimsContractBinding.Contract.DEFAULTADMINROLE(&_ClaimsContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ClaimsContractBinding.Contract.DEFAULTADMINROLE(&_ClaimsContractBinding.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) MAXUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "MAX_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingSession) MAXUINT256() (*big.Int, error) {
	return _ClaimsContractBinding.Contract.MAXUINT256(&_ClaimsContractBinding.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) MAXUINT256() (*big.Int, error) {
	return _ClaimsContractBinding.Contract.MAXUINT256(&_ClaimsContractBinding.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) MIDUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "MID_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingSession) MIDUINT256() (*big.Int, error) {
	return _ClaimsContractBinding.Contract.MIDUINT256(&_ClaimsContractBinding.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) MIDUINT256() (*big.Int, error) {
	return _ClaimsContractBinding.Contract.MIDUINT256(&_ClaimsContractBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VERSION() (string, error) {
	return _ClaimsContractBinding.Contract.VERSION(&_ClaimsContractBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VERSION() (string, error) {
	return _ClaimsContractBinding.Contract.VERSION(&_ClaimsContractBinding.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTESTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTESTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTESTAKEN(&_ClaimsContractBinding.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTESTAKEN() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTESTAKEN(&_ClaimsContractBinding.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTESTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTESTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTESTAKEY(&_ClaimsContractBinding.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTESTAKEY() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTESTAKEY(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTETRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTETRUTHN() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHN(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTETRUTHN() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHN(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTETRUTHS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_TRUTH_S")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTETRUTHS() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHS(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTETRUTHS() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHS(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTETRUTHV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_TRUTH_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTETRUTHV() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHV(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTETRUTHV() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHV(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) VOTETRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "VOTE_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingSession) VOTETRUTHY() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHY(&_ClaimsContractBinding.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) VOTETRUTHY() (uint8, error) {
	return _ClaimsContractBinding.Contract.VOTETRUTHY(&_ClaimsContractBinding.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) DurationBasis(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "durationBasis")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingSession) DurationBasis() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationBasis(&_ClaimsContractBinding.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) DurationBasis() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationBasis(&_ClaimsContractBinding.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) DurationMax(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "durationMax")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingSession) DurationMax() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationMax(&_ClaimsContractBinding.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) DurationMax() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationMax(&_ClaimsContractBinding.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) DurationMin(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "durationMin")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingSession) DurationMin() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationMin(&_ClaimsContractBinding.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) DurationMin() (uint64, error) {
	return _ClaimsContractBinding.Contract.DurationMin(&_ClaimsContractBinding.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ClaimsContractBinding.Contract.GetRoleAdmin(&_ClaimsContractBinding.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ClaimsContractBinding.Contract.GetRoleAdmin(&_ClaimsContractBinding.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ClaimsContractBinding.Contract.GetRoleMember(&_ClaimsContractBinding.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ClaimsContractBinding.Contract.GetRoleMember(&_ClaimsContractBinding.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ClaimsContractBinding.Contract.GetRoleMemberCount(&_ClaimsContractBinding.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ClaimsContractBinding.Contract.GetRoleMemberCount(&_ClaimsContractBinding.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ClaimsContractBinding.Contract.HasRole(&_ClaimsContractBinding.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ClaimsContractBinding.Contract.HasRole(&_ClaimsContractBinding.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingSession) Owner() (common.Address, error) {
	return _ClaimsContractBinding.Contract.Owner(&_ClaimsContractBinding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) Owner() (common.Address, error) {
	return _ClaimsContractBinding.Contract.Owner(&_ClaimsContractBinding.CallOpts)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchBalance(opts *bind.CallOpts, use common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchBalance", use)

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
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchBalance(&_ClaimsContractBinding.CallOpts, use)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchBalance(&_ClaimsContractBinding.CallOpts, use)
}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchContent(opts *bind.CallOpts, pod *big.Int) (string, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchContent", pod)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchContent(pod *big.Int) (string, error) {
	return _ClaimsContractBinding.Contract.SearchContent(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchContent is a free data retrieval call binding the contract method 0xa7236020.
//
// Solidity: function searchContent(uint256 pod) view returns(string)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchContent(pod *big.Int) (string, error) {
	return _ClaimsContractBinding.Contract.SearchContent(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchExpired(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchExpired", pod)

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
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchExpired(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchExpired(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchExpired(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchExpired(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchHistory(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchHistory", pod, lef, rig)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchHistory(pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchHistory(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchHistory is a free data retrieval call binding the contract method 0xe69f040e.
//
// Solidity: function searchHistory(uint256 pod, uint256 lef, uint256 rig) view returns(uint256[])
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchHistory(pod *big.Int, lef *big.Int, rig *big.Int) ([]*big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchHistory(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchIndices(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchIndices", pod)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchIndices(pod *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchIndices(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pod) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchIndices(pod *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchIndices(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchLatest is a free data retrieval call binding the contract method 0x5d0d720f.
//
// Solidity: function searchLatest(uint256 pod) view returns(uint8, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchLatest(opts *bind.CallOpts, pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchLatest", pod)

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
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchLatest(pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchLatest(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchLatest is a free data retrieval call binding the contract method 0x5d0d720f.
//
// Solidity: function searchLatest(uint256 pod) view returns(uint8, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchLatest(pod *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchLatest(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pod) view returns(uint256, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchPropose(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchPropose", pod)

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
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchPropose(pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchPropose(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pod) view returns(uint256, uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchPropose(pod *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchPropose(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchResolve(opts *bind.CallOpts, pod *big.Int, ind uint8) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchResolve", pod, ind)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchResolve(pod *big.Int, ind uint8) (bool, error) {
	return _ClaimsContractBinding.Contract.SearchResolve(&_ClaimsContractBinding.CallOpts, pod, ind)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pod, uint8 ind) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchResolve(pod *big.Int, ind uint8) (bool, error) {
	return _ClaimsContractBinding.Contract.SearchResolve(&_ClaimsContractBinding.CallOpts, pod, ind)
}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchResults(opts *bind.CallOpts, pod *big.Int) (bool, bool, bool, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchResults", pod)

	if err != nil {
		return *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchResults(pod *big.Int) (bool, bool, bool, error) {
	return _ClaimsContractBinding.Contract.SearchResults(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchResults is a free data retrieval call binding the contract method 0xaf9db5a7.
//
// Solidity: function searchResults(uint256 pod) view returns(bool, bool, bool)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchResults(pod *big.Int) (bool, bool, bool, error) {
	return _ClaimsContractBinding.Contract.SearchResults(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchSamples(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchSamples", pod, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchSamples(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchSamples(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchSamples(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchSamples(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchStakers(opts *bind.CallOpts, pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchStakers", pod, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchStakers(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchStakers(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pod, uint256 lef, uint256 rig) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchStakers(pod *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchStakers(&_ClaimsContractBinding.CallOpts, pod, lef, rig)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchToken(opts *bind.CallOpts, pro *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchToken", pro)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchToken(&_ClaimsContractBinding.CallOpts, pro)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _ClaimsContractBinding.Contract.SearchToken(&_ClaimsContractBinding.CallOpts, pro)
}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SearchVotes(opts *bind.CallOpts, pod *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "searchVotes", pod)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SearchVotes(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchVotes(&_ClaimsContractBinding.CallOpts, pod)
}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pod) view returns(uint256, uint256)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SearchVotes(pod *big.Int) (*big.Int, *big.Int, error) {
	return _ClaimsContractBinding.Contract.SearchVotes(&_ClaimsContractBinding.CallOpts, pod)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ClaimsContractBinding.Contract.SupportsInterface(&_ClaimsContractBinding.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ClaimsContractBinding.Contract.SupportsInterface(&_ClaimsContractBinding.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimsContractBinding.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingSession) Token() (common.Address, error) {
	return _ClaimsContractBinding.Contract.Token(&_ClaimsContractBinding.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ClaimsContractBinding *ClaimsContractBindingCallerSession) Token() (common.Address, error) {
	return _ClaimsContractBinding.Contract.Token(&_ClaimsContractBinding.CallOpts)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) CreateDispute(opts *bind.TransactOpts, dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "createDispute", dis, bal, vot, exp, con, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreateDispute(&_ClaimsContractBinding.TransactOpts, dis, bal, vot, exp, con, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0x60528d35.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, string con, uint256 pro) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, con string, pro *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreateDispute(&_ClaimsContractBinding.TransactOpts, dis, bal, vot, exp, con, pro)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) CreatePropose(opts *bind.TransactOpts, pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "createPropose", pro, bal, vot, exp, con, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreatePropose(&_ClaimsContractBinding.TransactOpts, pro, bal, vot, exp, con, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0x6e43e824.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, string con, address[] tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, con string, tok []common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreatePropose(&_ClaimsContractBinding.TransactOpts, pro, bal, vot, exp, con, tok)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) CreateResolve(opts *bind.TransactOpts, pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "createResolve", pod, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) CreateResolve(pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreateResolve(&_ClaimsContractBinding.TransactOpts, pod, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pod, uint256[] ind, uint64 exp) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) CreateResolve(pod *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.CreateResolve(&_ClaimsContractBinding.TransactOpts, pod, ind, exp)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.GrantRole(&_ClaimsContractBinding.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.GrantRole(&_ClaimsContractBinding.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.RenounceRole(&_ClaimsContractBinding.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.RenounceRole(&_ClaimsContractBinding.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.RevokeRole(&_ClaimsContractBinding.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.RevokeRole(&_ClaimsContractBinding.TransactOpts, role, account)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) UpdateBalance(opts *bind.TransactOpts, pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "updateBalance", pod, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) UpdateBalance(pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateBalance(&_ClaimsContractBinding.TransactOpts, pod, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 pod, uint256 max) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) UpdateBalance(pod *big.Int, max *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateBalance(&_ClaimsContractBinding.TransactOpts, pod, max)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) UpdateDuration(opts *bind.TransactOpts, bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "updateDuration", bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateDuration(&_ClaimsContractBinding.TransactOpts, bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateDuration(&_ClaimsContractBinding.TransactOpts, bas, max, min)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) UpdateOwner(opts *bind.TransactOpts, own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "updateOwner", own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateOwner(&_ClaimsContractBinding.TransactOpts, own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateOwner(&_ClaimsContractBinding.TransactOpts, own)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) UpdatePropose(opts *bind.TransactOpts, pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "updatePropose", pod, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) UpdatePropose(pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdatePropose(&_ClaimsContractBinding.TransactOpts, pod, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 pod, uint256 bal, bool vot, uint256 tok) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) UpdatePropose(pod *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdatePropose(&_ClaimsContractBinding.TransactOpts, pod, bal, vot, tok)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) UpdateResolve(opts *bind.TransactOpts, pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "updateResolve", pod, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) UpdateResolve(pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateResolve(&_ClaimsContractBinding.TransactOpts, pod, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pod, bool vot) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) UpdateResolve(pod *big.Int, vot bool) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.UpdateResolve(&_ClaimsContractBinding.TransactOpts, pod, vot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) Withdraw(opts *bind.TransactOpts, bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.Transact(opts, "withdraw", bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.Withdraw(&_ClaimsContractBinding.TransactOpts, bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.Withdraw(&_ClaimsContractBinding.TransactOpts, bal)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimsContractBinding.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBinding *ClaimsContractBindingSession) Receive() (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.Receive(&_ClaimsContractBinding.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ClaimsContractBinding *ClaimsContractBindingTransactorSession) Receive() (*types.Transaction, error) {
	return _ClaimsContractBinding.Contract.Receive(&_ClaimsContractBinding.TransactOpts)
}

// ClaimsContractBindingBalanceUpdatedIterator is returned from FilterBalanceUpdated and is used to iterate over the raw logs and unpacked data for BalanceUpdated events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingBalanceUpdatedIterator struct {
	Event *ClaimsContractBindingBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingBalanceUpdated)
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
		it.Event = new(ClaimsContractBindingBalanceUpdated)
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
func (it *ClaimsContractBindingBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingBalanceUpdated represents a BalanceUpdated event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingBalanceUpdated struct {
	Pod *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBalanceUpdated is a free log retrieval operation binding the contract event 0x04be25a4c0fe4c39136ca055daaac08ac6b3041b2e1f2a1b698bac12f1cbdc4a.
//
// Solidity: event BalanceUpdated(uint256 pod)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterBalanceUpdated(opts *bind.FilterOpts) (*ClaimsContractBindingBalanceUpdatedIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "BalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingBalanceUpdatedIterator{contract: _ClaimsContractBinding.contract, event: "BalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceUpdated is a free log subscription operation binding the contract event 0x04be25a4c0fe4c39136ca055daaac08ac6b3041b2e1f2a1b698bac12f1cbdc4a.
//
// Solidity: event BalanceUpdated(uint256 pod)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "BalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingBalanceUpdated)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseBalanceUpdated(log types.Log) (*ClaimsContractBindingBalanceUpdated, error) {
	event := new(ClaimsContractBindingBalanceUpdated)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingClaimUpdatedIterator is returned from FilterClaimUpdated and is used to iterate over the raw logs and unpacked data for ClaimUpdated events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingClaimUpdatedIterator struct {
	Event *ClaimsContractBindingClaimUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingClaimUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingClaimUpdated)
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
		it.Event = new(ClaimsContractBindingClaimUpdated)
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
func (it *ClaimsContractBindingClaimUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingClaimUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingClaimUpdated represents a ClaimUpdated event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingClaimUpdated struct {
	Pod *big.Int
	Use common.Address
	Bal *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimUpdated is a free log retrieval operation binding the contract event 0x456e7c3ba307b1c39b4aa965dad80444be2ddffd40b0d6e930507c5858faf33e.
//
// Solidity: event ClaimUpdated(uint256 pod, address use, uint256 bal)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterClaimUpdated(opts *bind.FilterOpts) (*ClaimsContractBindingClaimUpdatedIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "ClaimUpdated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingClaimUpdatedIterator{contract: _ClaimsContractBinding.contract, event: "ClaimUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimUpdated is a free log subscription operation binding the contract event 0x456e7c3ba307b1c39b4aa965dad80444be2ddffd40b0d6e930507c5858faf33e.
//
// Solidity: event ClaimUpdated(uint256 pod, address use, uint256 bal)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchClaimUpdated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingClaimUpdated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "ClaimUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingClaimUpdated)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseClaimUpdated(log types.Log) (*ClaimsContractBindingClaimUpdated, error) {
	event := new(ClaimsContractBindingClaimUpdated)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "ClaimUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingDisputeCreatedIterator is returned from FilterDisputeCreated and is used to iterate over the raw logs and unpacked data for DisputeCreated events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingDisputeCreatedIterator struct {
	Event *ClaimsContractBindingDisputeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingDisputeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingDisputeCreated)
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
		it.Event = new(ClaimsContractBindingDisputeCreated)
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
func (it *ClaimsContractBindingDisputeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingDisputeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingDisputeCreated represents a DisputeCreated event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingDisputeCreated struct {
	Dis *big.Int
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisputeCreated is a free log retrieval operation binding the contract event 0x848ff8058c972bb6cd367b414b3cd1799ef42a941e7800185f0e580bda9c3cc5.
//
// Solidity: event DisputeCreated(uint256 dis, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterDisputeCreated(opts *bind.FilterOpts) (*ClaimsContractBindingDisputeCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingDisputeCreatedIterator{contract: _ClaimsContractBinding.contract, event: "DisputeCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCreated is a free log subscription operation binding the contract event 0x848ff8058c972bb6cd367b414b3cd1799ef42a941e7800185f0e580bda9c3cc5.
//
// Solidity: event DisputeCreated(uint256 dis, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchDisputeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingDisputeCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingDisputeCreated)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseDisputeCreated(log types.Log) (*ClaimsContractBindingDisputeCreated, error) {
	event := new(ClaimsContractBindingDisputeCreated)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingDisputeSettledIterator is returned from FilterDisputeSettled and is used to iterate over the raw logs and unpacked data for DisputeSettled events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingDisputeSettledIterator struct {
	Event *ClaimsContractBindingDisputeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingDisputeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingDisputeSettled)
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
		it.Event = new(ClaimsContractBindingDisputeSettled)
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
func (it *ClaimsContractBindingDisputeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingDisputeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingDisputeSettled represents a DisputeSettled event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingDisputeSettled struct {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterDisputeSettled(opts *bind.FilterOpts) (*ClaimsContractBindingDisputeSettledIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingDisputeSettledIterator{contract: _ClaimsContractBinding.contract, event: "DisputeSettled", logs: logs, sub: sub}, nil
}

// WatchDisputeSettled is a free log subscription operation binding the contract event 0x885d609f8d9d1fc5b04cdedac72087d79d587398e79a263f148e64f9275fc929.
//
// Solidity: event DisputeSettled(uint256 dis, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchDisputeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingDisputeSettled) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingDisputeSettled)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseDisputeSettled(log types.Log) (*ClaimsContractBindingDisputeSettled, error) {
	event := new(ClaimsContractBindingDisputeSettled)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingProposeCreatedIterator is returned from FilterProposeCreated and is used to iterate over the raw logs and unpacked data for ProposeCreated events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingProposeCreatedIterator struct {
	Event *ClaimsContractBindingProposeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingProposeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingProposeCreated)
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
		it.Event = new(ClaimsContractBindingProposeCreated)
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
func (it *ClaimsContractBindingProposeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingProposeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingProposeCreated represents a ProposeCreated event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingProposeCreated struct {
	Pro *big.Int
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposeCreated is a free log retrieval operation binding the contract event 0x560897c1e1894395c161cb6ee32a257c436c4f5e19d550b5b80a6043fd166c78.
//
// Solidity: event ProposeCreated(uint256 pro, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterProposeCreated(opts *bind.FilterOpts) (*ClaimsContractBindingProposeCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingProposeCreatedIterator{contract: _ClaimsContractBinding.contract, event: "ProposeCreated", logs: logs, sub: sub}, nil
}

// WatchProposeCreated is a free log subscription operation binding the contract event 0x560897c1e1894395c161cb6ee32a257c436c4f5e19d550b5b80a6043fd166c78.
//
// Solidity: event ProposeCreated(uint256 pro, address use, uint256 bal, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchProposeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingProposeCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingProposeCreated)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseProposeCreated(log types.Log) (*ClaimsContractBindingProposeCreated, error) {
	event := new(ClaimsContractBindingProposeCreated)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingProposeSettledIterator is returned from FilterProposeSettled and is used to iterate over the raw logs and unpacked data for ProposeSettled events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingProposeSettledIterator struct {
	Event *ClaimsContractBindingProposeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingProposeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingProposeSettled)
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
		it.Event = new(ClaimsContractBindingProposeSettled)
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
func (it *ClaimsContractBindingProposeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingProposeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingProposeSettled represents a ProposeSettled event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingProposeSettled struct {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterProposeSettled(opts *bind.FilterOpts) (*ClaimsContractBindingProposeSettledIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingProposeSettledIterator{contract: _ClaimsContractBinding.contract, event: "ProposeSettled", logs: logs, sub: sub}, nil
}

// WatchProposeSettled is a free log subscription operation binding the contract event 0x02b6b8b1e0e988010b10993e9424cd6b9a8f599a653942c3fec3756d53f14e68.
//
// Solidity: event ProposeSettled(uint256 pro, uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchProposeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingProposeSettled) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingProposeSettled)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseProposeSettled(log types.Log) (*ClaimsContractBindingProposeSettled, error) {
	event := new(ClaimsContractBindingProposeSettled)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingResolveCreatedIterator is returned from FilterResolveCreated and is used to iterate over the raw logs and unpacked data for ResolveCreated events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingResolveCreatedIterator struct {
	Event *ClaimsContractBindingResolveCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingResolveCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingResolveCreated)
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
		it.Event = new(ClaimsContractBindingResolveCreated)
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
func (it *ClaimsContractBindingResolveCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingResolveCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingResolveCreated represents a ResolveCreated event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingResolveCreated struct {
	Pod *big.Int
	Use common.Address
	Len *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResolveCreated is a free log retrieval operation binding the contract event 0xe3c5a61bbb499af1d75605281b4c8c6e4635121595fe0cdc3ef366e28204dbef.
//
// Solidity: event ResolveCreated(uint256 pod, address use, uint256 len, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterResolveCreated(opts *bind.FilterOpts) (*ClaimsContractBindingResolveCreatedIterator, error) {

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingResolveCreatedIterator{contract: _ClaimsContractBinding.contract, event: "ResolveCreated", logs: logs, sub: sub}, nil
}

// WatchResolveCreated is a free log subscription operation binding the contract event 0xe3c5a61bbb499af1d75605281b4c8c6e4635121595fe0cdc3ef366e28204dbef.
//
// Solidity: event ResolveCreated(uint256 pod, address use, uint256 len, uint64 exp)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchResolveCreated(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingResolveCreated) (event.Subscription, error) {

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingResolveCreated)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseResolveCreated(log types.Log) (*ClaimsContractBindingResolveCreated, error) {
	event := new(ClaimsContractBindingResolveCreated)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleAdminChangedIterator struct {
	Event *ClaimsContractBindingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingRoleAdminChanged)
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
		it.Event = new(ClaimsContractBindingRoleAdminChanged)
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
func (it *ClaimsContractBindingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingRoleAdminChanged represents a RoleAdminChanged event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ClaimsContractBindingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingRoleAdminChangedIterator{contract: _ClaimsContractBinding.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingRoleAdminChanged)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseRoleAdminChanged(log types.Log) (*ClaimsContractBindingRoleAdminChanged, error) {
	event := new(ClaimsContractBindingRoleAdminChanged)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleGrantedIterator struct {
	Event *ClaimsContractBindingRoleGranted // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingRoleGranted)
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
		it.Event = new(ClaimsContractBindingRoleGranted)
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
func (it *ClaimsContractBindingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingRoleGranted represents a RoleGranted event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsContractBindingRoleGrantedIterator, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingRoleGrantedIterator{contract: _ClaimsContractBinding.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingRoleGranted)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseRoleGranted(log types.Log) (*ClaimsContractBindingRoleGranted, error) {
	event := new(ClaimsContractBindingRoleGranted)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsContractBindingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleRevokedIterator struct {
	Event *ClaimsContractBindingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ClaimsContractBindingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsContractBindingRoleRevoked)
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
		it.Event = new(ClaimsContractBindingRoleRevoked)
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
func (it *ClaimsContractBindingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsContractBindingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsContractBindingRoleRevoked represents a RoleRevoked event raised by the ClaimsContractBinding contract.
type ClaimsContractBindingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsContractBindingRoleRevokedIterator, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsContractBindingRoleRevokedIterator{contract: _ClaimsContractBinding.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ClaimsContractBindingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ClaimsContractBinding.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsContractBindingRoleRevoked)
				if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ClaimsContractBinding *ClaimsContractBindingFilterer) ParseRoleRevoked(log types.Log) (*ClaimsContractBindingRoleRevoked, error) {
	event := new(ClaimsContractBindingRoleRevoked)
	if err := _ClaimsContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
