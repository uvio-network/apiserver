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

// ClaimsMetaData contains all meta data concerning the Claims contract.
var ClaimsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unx\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Mapping\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"DisputeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"DisputeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ProposeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"all\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nah\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tot\",\"type\":\"uint256\"}],\"name\":\"ProposeSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"ResolveCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ADDRESS_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BASIS_TOTAL\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_ADDRESS_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_P\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_R\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_BALANCE_U\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_P\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_R\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_EXPIRY_T\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_A\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_B\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_C\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_D\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MID_UINT256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_STAKE_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_N\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_S\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOTE_TRUTH_Y\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basisFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basisProposer\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basisProtocol\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dis\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"createDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"tok\",\"type\":\"address[]\"}],\"name\":\"createPropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"ind\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"exp\",\"type\":\"uint64\"}],\"name\":\"createResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationBasis\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMax\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"durationMin\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"use\",\"type\":\"address\"}],\"name\":\"searchBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchExpired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchPropose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"ind\",\"type\":\"uint8\"}],\"name\":\"searchResolve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchSamples\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lef\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rig\",\"type\":\"uint256\"}],\"name\":\"searchStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"}],\"name\":\"searchVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cla\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"updateBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"bas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"min\",\"type\":\"uint64\"}],\"name\":\"updateDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"psr\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"ptc\",\"type\":\"uint16\"}],\"name\":\"updateFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cla\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"}],\"name\":\"updatePropose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pro\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"vot\",\"type\":\"bool\"}],\"name\":\"updateResolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ClaimsABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimsMetaData.ABI instead.
var ClaimsABI = ClaimsMetaData.ABI

// Claims is an auto generated Go binding around an Ethereum contract.
type Claims struct {
	ClaimsCaller     // Read-only binding to the contract
	ClaimsTransactor // Write-only binding to the contract
	ClaimsFilterer   // Log filterer for contract events
}

// ClaimsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimsSession struct {
	Contract     *Claims           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimsCallerSession struct {
	Contract *ClaimsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClaimsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimsTransactorSession struct {
	Contract     *ClaimsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimsRaw struct {
	Contract *Claims // Generic contract binding to access the raw methods on
}

// ClaimsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimsCallerRaw struct {
	Contract *ClaimsCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimsTransactorRaw struct {
	Contract *ClaimsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaims creates a new instance of Claims, bound to a specific deployed contract.
func NewClaims(address common.Address, backend bind.ContractBackend) (*Claims, error) {
	contract, err := bindClaims(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claims{ClaimsCaller: ClaimsCaller{contract: contract}, ClaimsTransactor: ClaimsTransactor{contract: contract}, ClaimsFilterer: ClaimsFilterer{contract: contract}}, nil
}

// NewClaimsCaller creates a new read-only instance of Claims, bound to a specific deployed contract.
func NewClaimsCaller(address common.Address, caller bind.ContractCaller) (*ClaimsCaller, error) {
	contract, err := bindClaims(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsCaller{contract: contract}, nil
}

// NewClaimsTransactor creates a new write-only instance of Claims, bound to a specific deployed contract.
func NewClaimsTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimsTransactor, error) {
	contract, err := bindClaims(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimsTransactor{contract: contract}, nil
}

// NewClaimsFilterer creates a new log filterer instance of Claims, bound to a specific deployed contract.
func NewClaimsFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimsFilterer, error) {
	contract, err := bindClaims(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimsFilterer{contract: contract}, nil
}

// bindClaims binds a generic wrapper to an already deployed contract.
func bindClaims(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claims *ClaimsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claims.Contract.ClaimsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claims *ClaimsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claims.Contract.ClaimsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claims *ClaimsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claims.Contract.ClaimsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claims *ClaimsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Claims.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claims *ClaimsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claims.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claims *ClaimsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claims.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCaller) ADDRESSSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "ADDRESS_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_Claims *ClaimsSession) ADDRESSSTAKEN() (uint8, error) {
	return _Claims.Contract.ADDRESSSTAKEN(&_Claims.CallOpts)
}

// ADDRESSSTAKEN is a free data retrieval call binding the contract method 0x232c33bc.
//
// Solidity: function ADDRESS_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) ADDRESSSTAKEN() (uint8, error) {
	return _Claims.Contract.ADDRESSSTAKEN(&_Claims.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCaller) ADDRESSSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "ADDRESS_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsSession) ADDRESSSTAKEY() (uint8, error) {
	return _Claims.Contract.ADDRESSSTAKEY(&_Claims.CallOpts)
}

// ADDRESSSTAKEY is a free data retrieval call binding the contract method 0xed5822b1.
//
// Solidity: function ADDRESS_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) ADDRESSSTAKEY() (uint8, error) {
	return _Claims.Contract.ADDRESSSTAKEY(&_Claims.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_Claims *ClaimsCaller) BASISTOTAL(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "BASIS_TOTAL")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_Claims *ClaimsSession) BASISTOTAL() (uint16, error) {
	return _Claims.Contract.BASISTOTAL(&_Claims.CallOpts)
}

// BASISTOTAL is a free data retrieval call binding the contract method 0x9a61d8c3.
//
// Solidity: function BASIS_TOTAL() view returns(uint16)
func (_Claims *ClaimsCallerSession) BASISTOTAL() (uint16, error) {
	return _Claims.Contract.BASISTOTAL(&_Claims.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_Claims *ClaimsCaller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_Claims *ClaimsSession) BOTROLE() ([32]byte, error) {
	return _Claims.Contract.BOTROLE(&_Claims.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_Claims *ClaimsCallerSession) BOTROLE() ([32]byte, error) {
	return _Claims.Contract.BOTROLE(&_Claims.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMADDRESSN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_ADDRESS_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMADDRESSN() (uint8, error) {
	return _Claims.Contract.CLAIMADDRESSN(&_Claims.CallOpts)
}

// CLAIMADDRESSN is a free data retrieval call binding the contract method 0x5722c818.
//
// Solidity: function CLAIM_ADDRESS_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMADDRESSN() (uint8, error) {
	return _Claims.Contract.CLAIMADDRESSN(&_Claims.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMADDRESSY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_ADDRESS_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMADDRESSY() (uint8, error) {
	return _Claims.Contract.CLAIMADDRESSY(&_Claims.CallOpts)
}

// CLAIMADDRESSY is a free data retrieval call binding the contract method 0x14ea97a1.
//
// Solidity: function CLAIM_ADDRESS_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMADDRESSY() (uint8, error) {
	return _Claims.Contract.CLAIMADDRESSY(&_Claims.CallOpts)
}

// CLAIMBALANCEP is a free data retrieval call binding the contract method 0x409f9ba8.
//
// Solidity: function CLAIM_BALANCE_P() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMBALANCEP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_BALANCE_P")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCEP is a free data retrieval call binding the contract method 0x409f9ba8.
//
// Solidity: function CLAIM_BALANCE_P() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMBALANCEP() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCEP(&_Claims.CallOpts)
}

// CLAIMBALANCEP is a free data retrieval call binding the contract method 0x409f9ba8.
//
// Solidity: function CLAIM_BALANCE_P() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMBALANCEP() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCEP(&_Claims.CallOpts)
}

// CLAIMBALANCER is a free data retrieval call binding the contract method 0x098e92bd.
//
// Solidity: function CLAIM_BALANCE_R() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMBALANCER(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_BALANCE_R")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCER is a free data retrieval call binding the contract method 0x098e92bd.
//
// Solidity: function CLAIM_BALANCE_R() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMBALANCER() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCER(&_Claims.CallOpts)
}

// CLAIMBALANCER is a free data retrieval call binding the contract method 0x098e92bd.
//
// Solidity: function CLAIM_BALANCE_R() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMBALANCER() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCER(&_Claims.CallOpts)
}

// CLAIMBALANCEU is a free data retrieval call binding the contract method 0x4043ae78.
//
// Solidity: function CLAIM_BALANCE_U() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMBALANCEU(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_BALANCE_U")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMBALANCEU is a free data retrieval call binding the contract method 0x4043ae78.
//
// Solidity: function CLAIM_BALANCE_U() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMBALANCEU() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCEU(&_Claims.CallOpts)
}

// CLAIMBALANCEU is a free data retrieval call binding the contract method 0x4043ae78.
//
// Solidity: function CLAIM_BALANCE_U() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMBALANCEU() (uint8, error) {
	return _Claims.Contract.CLAIMBALANCEU(&_Claims.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMEXPIRYP(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_EXPIRY_P")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMEXPIRYP() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYP(&_Claims.CallOpts)
}

// CLAIMEXPIRYP is a free data retrieval call binding the contract method 0x55151f23.
//
// Solidity: function CLAIM_EXPIRY_P() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMEXPIRYP() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYP(&_Claims.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMEXPIRYR(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_EXPIRY_R")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMEXPIRYR() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYR(&_Claims.CallOpts)
}

// CLAIMEXPIRYR is a free data retrieval call binding the contract method 0x4884a67b.
//
// Solidity: function CLAIM_EXPIRY_R() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMEXPIRYR() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYR(&_Claims.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMEXPIRYT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_EXPIRY_T")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMEXPIRYT() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYT(&_Claims.CallOpts)
}

// CLAIMEXPIRYT is a free data retrieval call binding the contract method 0x970be447.
//
// Solidity: function CLAIM_EXPIRY_T() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMEXPIRYT() (uint8, error) {
	return _Claims.Contract.CLAIMEXPIRYT(&_Claims.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKEA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_A")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKEA() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEA(&_Claims.CallOpts)
}

// CLAIMSTAKEA is a free data retrieval call binding the contract method 0xe4b2d71e.
//
// Solidity: function CLAIM_STAKE_A() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKEA() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEA(&_Claims.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKEB(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_B")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKEB() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEB(&_Claims.CallOpts)
}

// CLAIMSTAKEB is a free data retrieval call binding the contract method 0xecc94e8f.
//
// Solidity: function CLAIM_STAKE_B() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKEB() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEB(&_Claims.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKEC(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_C")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKEC() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEC(&_Claims.CallOpts)
}

// CLAIMSTAKEC is a free data retrieval call binding the contract method 0x8781fa03.
//
// Solidity: function CLAIM_STAKE_C() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKEC() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEC(&_Claims.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_D")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKED() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKED(&_Claims.CallOpts)
}

// CLAIMSTAKED is a free data retrieval call binding the contract method 0x62f487a1.
//
// Solidity: function CLAIM_STAKE_D() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKED() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKED(&_Claims.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKEN() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEN(&_Claims.CallOpts)
}

// CLAIMSTAKEN is a free data retrieval call binding the contract method 0x4e38bd03.
//
// Solidity: function CLAIM_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKEN() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEN(&_Claims.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMSTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMSTAKEY() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEY(&_Claims.CallOpts)
}

// CLAIMSTAKEY is a free data retrieval call binding the contract method 0x4ded0e71.
//
// Solidity: function CLAIM_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMSTAKEY() (uint8, error) {
	return _Claims.Contract.CLAIMSTAKEY(&_Claims.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMTRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMTRUTHN() (uint8, error) {
	return _Claims.Contract.CLAIMTRUTHN(&_Claims.CallOpts)
}

// CLAIMTRUTHN is a free data retrieval call binding the contract method 0x0e532ff2.
//
// Solidity: function CLAIM_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMTRUTHN() (uint8, error) {
	return _Claims.Contract.CLAIMTRUTHN(&_Claims.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsCaller) CLAIMTRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "CLAIM_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsSession) CLAIMTRUTHY() (uint8, error) {
	return _Claims.Contract.CLAIMTRUTHY(&_Claims.CallOpts)
}

// CLAIMTRUTHY is a free data retrieval call binding the contract method 0x33d9a313.
//
// Solidity: function CLAIM_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) CLAIMTRUTHY() (uint8, error) {
	return _Claims.Contract.CLAIMTRUTHY(&_Claims.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Claims *ClaimsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Claims *ClaimsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Claims.Contract.DEFAULTADMINROLE(&_Claims.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Claims *ClaimsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Claims.Contract.DEFAULTADMINROLE(&_Claims.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_Claims *ClaimsCaller) MAXUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "MAX_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_Claims *ClaimsSession) MAXUINT256() (*big.Int, error) {
	return _Claims.Contract.MAXUINT256(&_Claims.CallOpts)
}

// MAXUINT256 is a free data retrieval call binding the contract method 0x33a581d2.
//
// Solidity: function MAX_UINT256() view returns(uint256)
func (_Claims *ClaimsCallerSession) MAXUINT256() (*big.Int, error) {
	return _Claims.Contract.MAXUINT256(&_Claims.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_Claims *ClaimsCaller) MIDUINT256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "MID_UINT256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_Claims *ClaimsSession) MIDUINT256() (*big.Int, error) {
	return _Claims.Contract.MIDUINT256(&_Claims.CallOpts)
}

// MIDUINT256 is a free data retrieval call binding the contract method 0x795433da.
//
// Solidity: function MID_UINT256() view returns(uint256)
func (_Claims *ClaimsCallerSession) MIDUINT256() (*big.Int, error) {
	return _Claims.Contract.MIDUINT256(&_Claims.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Claims *ClaimsCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Claims *ClaimsSession) VERSION() (string, error) {
	return _Claims.Contract.VERSION(&_Claims.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Claims *ClaimsCallerSession) VERSION() (string, error) {
	return _Claims.Contract.VERSION(&_Claims.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCaller) VOTESTAKEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_STAKE_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_Claims *ClaimsSession) VOTESTAKEN() (uint8, error) {
	return _Claims.Contract.VOTESTAKEN(&_Claims.CallOpts)
}

// VOTESTAKEN is a free data retrieval call binding the contract method 0xc61f0ae8.
//
// Solidity: function VOTE_STAKE_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTESTAKEN() (uint8, error) {
	return _Claims.Contract.VOTESTAKEN(&_Claims.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCaller) VOTESTAKEY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_STAKE_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsSession) VOTESTAKEY() (uint8, error) {
	return _Claims.Contract.VOTESTAKEY(&_Claims.CallOpts)
}

// VOTESTAKEY is a free data retrieval call binding the contract method 0x54785755.
//
// Solidity: function VOTE_STAKE_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTESTAKEY() (uint8, error) {
	return _Claims.Contract.VOTESTAKEY(&_Claims.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsCaller) VOTETRUTHN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_TRUTH_N")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsSession) VOTETRUTHN() (uint8, error) {
	return _Claims.Contract.VOTETRUTHN(&_Claims.CallOpts)
}

// VOTETRUTHN is a free data retrieval call binding the contract method 0x78997fc2.
//
// Solidity: function VOTE_TRUTH_N() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTETRUTHN() (uint8, error) {
	return _Claims.Contract.VOTETRUTHN(&_Claims.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_Claims *ClaimsCaller) VOTETRUTHS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_TRUTH_S")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_Claims *ClaimsSession) VOTETRUTHS() (uint8, error) {
	return _Claims.Contract.VOTETRUTHS(&_Claims.CallOpts)
}

// VOTETRUTHS is a free data retrieval call binding the contract method 0x8a10d237.
//
// Solidity: function VOTE_TRUTH_S() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTETRUTHS() (uint8, error) {
	return _Claims.Contract.VOTETRUTHS(&_Claims.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_Claims *ClaimsCaller) VOTETRUTHV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_TRUTH_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_Claims *ClaimsSession) VOTETRUTHV() (uint8, error) {
	return _Claims.Contract.VOTETRUTHV(&_Claims.CallOpts)
}

// VOTETRUTHV is a free data retrieval call binding the contract method 0xcc250655.
//
// Solidity: function VOTE_TRUTH_V() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTETRUTHV() (uint8, error) {
	return _Claims.Contract.VOTETRUTHV(&_Claims.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsCaller) VOTETRUTHY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "VOTE_TRUTH_Y")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsSession) VOTETRUTHY() (uint8, error) {
	return _Claims.Contract.VOTETRUTHY(&_Claims.CallOpts)
}

// VOTETRUTHY is a free data retrieval call binding the contract method 0x00916264.
//
// Solidity: function VOTE_TRUTH_Y() view returns(uint8)
func (_Claims *ClaimsCallerSession) VOTETRUTHY() (uint8, error) {
	return _Claims.Contract.VOTETRUTHY(&_Claims.CallOpts)
}

// BasisFee is a free data retrieval call binding the contract method 0x971e1a66.
//
// Solidity: function basisFee() view returns(uint16)
func (_Claims *ClaimsCaller) BasisFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "basisFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BasisFee is a free data retrieval call binding the contract method 0x971e1a66.
//
// Solidity: function basisFee() view returns(uint16)
func (_Claims *ClaimsSession) BasisFee() (uint16, error) {
	return _Claims.Contract.BasisFee(&_Claims.CallOpts)
}

// BasisFee is a free data retrieval call binding the contract method 0x971e1a66.
//
// Solidity: function basisFee() view returns(uint16)
func (_Claims *ClaimsCallerSession) BasisFee() (uint16, error) {
	return _Claims.Contract.BasisFee(&_Claims.CallOpts)
}

// BasisProposer is a free data retrieval call binding the contract method 0x3b0ec0fd.
//
// Solidity: function basisProposer() view returns(uint16)
func (_Claims *ClaimsCaller) BasisProposer(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "basisProposer")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BasisProposer is a free data retrieval call binding the contract method 0x3b0ec0fd.
//
// Solidity: function basisProposer() view returns(uint16)
func (_Claims *ClaimsSession) BasisProposer() (uint16, error) {
	return _Claims.Contract.BasisProposer(&_Claims.CallOpts)
}

// BasisProposer is a free data retrieval call binding the contract method 0x3b0ec0fd.
//
// Solidity: function basisProposer() view returns(uint16)
func (_Claims *ClaimsCallerSession) BasisProposer() (uint16, error) {
	return _Claims.Contract.BasisProposer(&_Claims.CallOpts)
}

// BasisProtocol is a free data retrieval call binding the contract method 0x1cc13509.
//
// Solidity: function basisProtocol() view returns(uint16)
func (_Claims *ClaimsCaller) BasisProtocol(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "basisProtocol")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BasisProtocol is a free data retrieval call binding the contract method 0x1cc13509.
//
// Solidity: function basisProtocol() view returns(uint16)
func (_Claims *ClaimsSession) BasisProtocol() (uint16, error) {
	return _Claims.Contract.BasisProtocol(&_Claims.CallOpts)
}

// BasisProtocol is a free data retrieval call binding the contract method 0x1cc13509.
//
// Solidity: function basisProtocol() view returns(uint16)
func (_Claims *ClaimsCallerSession) BasisProtocol() (uint16, error) {
	return _Claims.Contract.BasisProtocol(&_Claims.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_Claims *ClaimsCaller) DurationBasis(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "durationBasis")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_Claims *ClaimsSession) DurationBasis() (uint64, error) {
	return _Claims.Contract.DurationBasis(&_Claims.CallOpts)
}

// DurationBasis is a free data retrieval call binding the contract method 0x9062a805.
//
// Solidity: function durationBasis() view returns(uint64)
func (_Claims *ClaimsCallerSession) DurationBasis() (uint64, error) {
	return _Claims.Contract.DurationBasis(&_Claims.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_Claims *ClaimsCaller) DurationMax(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "durationMax")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_Claims *ClaimsSession) DurationMax() (uint64, error) {
	return _Claims.Contract.DurationMax(&_Claims.CallOpts)
}

// DurationMax is a free data retrieval call binding the contract method 0xc522a9d4.
//
// Solidity: function durationMax() view returns(uint64)
func (_Claims *ClaimsCallerSession) DurationMax() (uint64, error) {
	return _Claims.Contract.DurationMax(&_Claims.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_Claims *ClaimsCaller) DurationMin(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "durationMin")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_Claims *ClaimsSession) DurationMin() (uint64, error) {
	return _Claims.Contract.DurationMin(&_Claims.CallOpts)
}

// DurationMin is a free data retrieval call binding the contract method 0xb5392984.
//
// Solidity: function durationMin() view returns(uint64)
func (_Claims *ClaimsCallerSession) DurationMin() (uint64, error) {
	return _Claims.Contract.DurationMin(&_Claims.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Claims *ClaimsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Claims *ClaimsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Claims.Contract.GetRoleAdmin(&_Claims.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Claims *ClaimsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Claims.Contract.GetRoleAdmin(&_Claims.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Claims *ClaimsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Claims *ClaimsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Claims.Contract.GetRoleMember(&_Claims.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Claims *ClaimsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Claims.Contract.GetRoleMember(&_Claims.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Claims *ClaimsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Claims *ClaimsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Claims.Contract.GetRoleMemberCount(&_Claims.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Claims *ClaimsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Claims.Contract.GetRoleMemberCount(&_Claims.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Claims *ClaimsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Claims *ClaimsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Claims.Contract.HasRole(&_Claims.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Claims *ClaimsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Claims.Contract.HasRole(&_Claims.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Claims *ClaimsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Claims *ClaimsSession) Owner() (common.Address, error) {
	return _Claims.Contract.Owner(&_Claims.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Claims *ClaimsCallerSession) Owner() (common.Address, error) {
	return _Claims.Contract.Owner(&_Claims.CallOpts)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_Claims *ClaimsCaller) SearchBalance(opts *bind.CallOpts, use common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchBalance", use)

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
func (_Claims *ClaimsSession) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchBalance(&_Claims.CallOpts, use)
}

// SearchBalance is a free data retrieval call binding the contract method 0x778c06f2.
//
// Solidity: function searchBalance(address use) view returns(uint256, uint256)
func (_Claims *ClaimsCallerSession) SearchBalance(use common.Address) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchBalance(&_Claims.CallOpts, use)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsCaller) SearchExpired(opts *bind.CallOpts, pro *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchExpired", pro)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsSession) SearchExpired(pro *big.Int) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchExpired(&_Claims.CallOpts, pro)
}

// SearchExpired is a free data retrieval call binding the contract method 0x33afa97c.
//
// Solidity: function searchExpired(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsCallerSession) SearchExpired(pro *big.Int) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchExpired(&_Claims.CallOpts, pro)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pro) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Claims *ClaimsCaller) SearchIndices(opts *bind.CallOpts, pro *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchIndices", pro)

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
// Solidity: function searchIndices(uint256 pro) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Claims *ClaimsSession) SearchIndices(pro *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Claims.Contract.SearchIndices(&_Claims.CallOpts, pro)
}

// SearchIndices is a free data retrieval call binding the contract method 0x9188d532.
//
// Solidity: function searchIndices(uint256 pro) view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, uint256)
func (_Claims *ClaimsCallerSession) SearchIndices(pro *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Claims.Contract.SearchIndices(&_Claims.CallOpts, pro)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pro) view returns(uint256, uint256, uint256)
func (_Claims *ClaimsCaller) SearchPropose(opts *bind.CallOpts, pro *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchPropose", pro)

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
// Solidity: function searchPropose(uint256 pro) view returns(uint256, uint256, uint256)
func (_Claims *ClaimsSession) SearchPropose(pro *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Claims.Contract.SearchPropose(&_Claims.CallOpts, pro)
}

// SearchPropose is a free data retrieval call binding the contract method 0x209f8eaf.
//
// Solidity: function searchPropose(uint256 pro) view returns(uint256, uint256, uint256)
func (_Claims *ClaimsCallerSession) SearchPropose(pro *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Claims.Contract.SearchPropose(&_Claims.CallOpts, pro)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pro, uint8 ind) view returns(bool)
func (_Claims *ClaimsCaller) SearchResolve(opts *bind.CallOpts, pro *big.Int, ind uint8) (bool, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchResolve", pro, ind)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pro, uint8 ind) view returns(bool)
func (_Claims *ClaimsSession) SearchResolve(pro *big.Int, ind uint8) (bool, error) {
	return _Claims.Contract.SearchResolve(&_Claims.CallOpts, pro, ind)
}

// SearchResolve is a free data retrieval call binding the contract method 0x88dc0bd2.
//
// Solidity: function searchResolve(uint256 pro, uint8 ind) view returns(bool)
func (_Claims *ClaimsCallerSession) SearchResolve(pro *big.Int, ind uint8) (bool, error) {
	return _Claims.Contract.SearchResolve(&_Claims.CallOpts, pro, ind)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsCaller) SearchSamples(opts *bind.CallOpts, pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchSamples", pro, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsSession) SearchSamples(pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchSamples(&_Claims.CallOpts, pro, lef, rig)
}

// SearchSamples is a free data retrieval call binding the contract method 0x286d2207.
//
// Solidity: function searchSamples(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsCallerSession) SearchSamples(pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchSamples(&_Claims.CallOpts, pro, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsCaller) SearchStakers(opts *bind.CallOpts, pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchStakers", pro, lef, rig)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsSession) SearchStakers(pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchStakers(&_Claims.CallOpts, pro, lef, rig)
}

// SearchStakers is a free data retrieval call binding the contract method 0x6b250a9f.
//
// Solidity: function searchStakers(uint256 pro, uint256 lef, uint256 rig) view returns(address[])
func (_Claims *ClaimsCallerSession) SearchStakers(pro *big.Int, lef *big.Int, rig *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchStakers(&_Claims.CallOpts, pro, lef, rig)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_Claims *ClaimsCaller) SearchToken(opts *bind.CallOpts, pro *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchToken", pro)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_Claims *ClaimsSession) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchToken(&_Claims.CallOpts, pro)
}

// SearchToken is a free data retrieval call binding the contract method 0x1a5c9a68.
//
// Solidity: function searchToken(uint256 pro) view returns(address[])
func (_Claims *ClaimsCallerSession) SearchToken(pro *big.Int) ([]common.Address, error) {
	return _Claims.Contract.SearchToken(&_Claims.CallOpts, pro)
}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsCaller) SearchVotes(opts *bind.CallOpts, pro *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "searchVotes", pro)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsSession) SearchVotes(pro *big.Int) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchVotes(&_Claims.CallOpts, pro)
}

// SearchVotes is a free data retrieval call binding the contract method 0xadfef97a.
//
// Solidity: function searchVotes(uint256 pro) view returns(uint256, uint256)
func (_Claims *ClaimsCallerSession) SearchVotes(pro *big.Int) (*big.Int, *big.Int, error) {
	return _Claims.Contract.SearchVotes(&_Claims.CallOpts, pro)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Claims *ClaimsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Claims *ClaimsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Claims.Contract.SupportsInterface(&_Claims.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Claims *ClaimsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Claims.Contract.SupportsInterface(&_Claims.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Claims *ClaimsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Claims.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Claims *ClaimsSession) Token() (common.Address, error) {
	return _Claims.Contract.Token(&_Claims.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Claims *ClaimsCallerSession) Token() (common.Address, error) {
	return _Claims.Contract.Token(&_Claims.CallOpts)
}

// CreateDispute is a paid mutator transaction binding the contract method 0xb30d0a97.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, uint256 pro) returns()
func (_Claims *ClaimsTransactor) CreateDispute(opts *bind.TransactOpts, dis *big.Int, bal *big.Int, vot bool, exp uint64, pro *big.Int) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "createDispute", dis, bal, vot, exp, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0xb30d0a97.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, uint256 pro) returns()
func (_Claims *ClaimsSession) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, pro *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.CreateDispute(&_Claims.TransactOpts, dis, bal, vot, exp, pro)
}

// CreateDispute is a paid mutator transaction binding the contract method 0xb30d0a97.
//
// Solidity: function createDispute(uint256 dis, uint256 bal, bool vot, uint64 exp, uint256 pro) returns()
func (_Claims *ClaimsTransactorSession) CreateDispute(dis *big.Int, bal *big.Int, vot bool, exp uint64, pro *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.CreateDispute(&_Claims.TransactOpts, dis, bal, vot, exp, pro)
}

// CreatePropose is a paid mutator transaction binding the contract method 0xffb02ec8.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, address[] tok) returns()
func (_Claims *ClaimsTransactor) CreatePropose(opts *bind.TransactOpts, pro *big.Int, bal *big.Int, vot bool, exp uint64, tok []common.Address) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "createPropose", pro, bal, vot, exp, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0xffb02ec8.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, address[] tok) returns()
func (_Claims *ClaimsSession) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, tok []common.Address) (*types.Transaction, error) {
	return _Claims.Contract.CreatePropose(&_Claims.TransactOpts, pro, bal, vot, exp, tok)
}

// CreatePropose is a paid mutator transaction binding the contract method 0xffb02ec8.
//
// Solidity: function createPropose(uint256 pro, uint256 bal, bool vot, uint64 exp, address[] tok) returns()
func (_Claims *ClaimsTransactorSession) CreatePropose(pro *big.Int, bal *big.Int, vot bool, exp uint64, tok []common.Address) (*types.Transaction, error) {
	return _Claims.Contract.CreatePropose(&_Claims.TransactOpts, pro, bal, vot, exp, tok)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pro, uint256[] ind, uint64 exp) returns()
func (_Claims *ClaimsTransactor) CreateResolve(opts *bind.TransactOpts, pro *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "createResolve", pro, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pro, uint256[] ind, uint64 exp) returns()
func (_Claims *ClaimsSession) CreateResolve(pro *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _Claims.Contract.CreateResolve(&_Claims.TransactOpts, pro, ind, exp)
}

// CreateResolve is a paid mutator transaction binding the contract method 0xd4e1dbc0.
//
// Solidity: function createResolve(uint256 pro, uint256[] ind, uint64 exp) returns()
func (_Claims *ClaimsTransactorSession) CreateResolve(pro *big.Int, ind []*big.Int, exp uint64) (*types.Transaction, error) {
	return _Claims.Contract.CreateResolve(&_Claims.TransactOpts, pro, ind, exp)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Claims *ClaimsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Claims *ClaimsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.Contract.GrantRole(&_Claims.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Claims *ClaimsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.Contract.GrantRole(&_Claims.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Claims *ClaimsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Claims *ClaimsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Claims.Contract.RenounceRole(&_Claims.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Claims *ClaimsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Claims.Contract.RenounceRole(&_Claims.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Claims *ClaimsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Claims *ClaimsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.Contract.RevokeRole(&_Claims.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Claims *ClaimsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Claims.Contract.RevokeRole(&_Claims.TransactOpts, role, account)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 cla, uint256 max) returns()
func (_Claims *ClaimsTransactor) UpdateBalance(opts *bind.TransactOpts, cla *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updateBalance", cla, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 cla, uint256 max) returns()
func (_Claims *ClaimsSession) UpdateBalance(cla *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.UpdateBalance(&_Claims.TransactOpts, cla, max)
}

// UpdateBalance is a paid mutator transaction binding the contract method 0x58453fef.
//
// Solidity: function updateBalance(uint256 cla, uint256 max) returns()
func (_Claims *ClaimsTransactorSession) UpdateBalance(cla *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.UpdateBalance(&_Claims.TransactOpts, cla, max)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_Claims *ClaimsTransactor) UpdateDuration(opts *bind.TransactOpts, bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updateDuration", bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_Claims *ClaimsSession) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _Claims.Contract.UpdateDuration(&_Claims.TransactOpts, bas, max, min)
}

// UpdateDuration is a paid mutator transaction binding the contract method 0xcf5e2292.
//
// Solidity: function updateDuration(uint64 bas, uint64 max, uint64 min) returns()
func (_Claims *ClaimsTransactorSession) UpdateDuration(bas uint64, max uint64, min uint64) (*types.Transaction, error) {
	return _Claims.Contract.UpdateDuration(&_Claims.TransactOpts, bas, max, min)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x670babe0.
//
// Solidity: function updateFees(uint16 fee, uint16 psr, uint16 ptc) returns()
func (_Claims *ClaimsTransactor) UpdateFees(opts *bind.TransactOpts, fee uint16, psr uint16, ptc uint16) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updateFees", fee, psr, ptc)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x670babe0.
//
// Solidity: function updateFees(uint16 fee, uint16 psr, uint16 ptc) returns()
func (_Claims *ClaimsSession) UpdateFees(fee uint16, psr uint16, ptc uint16) (*types.Transaction, error) {
	return _Claims.Contract.UpdateFees(&_Claims.TransactOpts, fee, psr, ptc)
}

// UpdateFees is a paid mutator transaction binding the contract method 0x670babe0.
//
// Solidity: function updateFees(uint16 fee, uint16 psr, uint16 ptc) returns()
func (_Claims *ClaimsTransactorSession) UpdateFees(fee uint16, psr uint16, ptc uint16) (*types.Transaction, error) {
	return _Claims.Contract.UpdateFees(&_Claims.TransactOpts, fee, psr, ptc)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_Claims *ClaimsTransactor) UpdateOwner(opts *bind.TransactOpts, own common.Address) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updateOwner", own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_Claims *ClaimsSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _Claims.Contract.UpdateOwner(&_Claims.TransactOpts, own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_Claims *ClaimsTransactorSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _Claims.Contract.UpdateOwner(&_Claims.TransactOpts, own)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 cla, uint256 bal, bool vot, uint256 tok) returns()
func (_Claims *ClaimsTransactor) UpdatePropose(opts *bind.TransactOpts, cla *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updatePropose", cla, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 cla, uint256 bal, bool vot, uint256 tok) returns()
func (_Claims *ClaimsSession) UpdatePropose(cla *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.UpdatePropose(&_Claims.TransactOpts, cla, bal, vot, tok)
}

// UpdatePropose is a paid mutator transaction binding the contract method 0x6e1ea6e3.
//
// Solidity: function updatePropose(uint256 cla, uint256 bal, bool vot, uint256 tok) returns()
func (_Claims *ClaimsTransactorSession) UpdatePropose(cla *big.Int, bal *big.Int, vot bool, tok *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.UpdatePropose(&_Claims.TransactOpts, cla, bal, vot, tok)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pro, bool vot) returns()
func (_Claims *ClaimsTransactor) UpdateResolve(opts *bind.TransactOpts, pro *big.Int, vot bool) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "updateResolve", pro, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pro, bool vot) returns()
func (_Claims *ClaimsSession) UpdateResolve(pro *big.Int, vot bool) (*types.Transaction, error) {
	return _Claims.Contract.UpdateResolve(&_Claims.TransactOpts, pro, vot)
}

// UpdateResolve is a paid mutator transaction binding the contract method 0x1b3569ed.
//
// Solidity: function updateResolve(uint256 pro, bool vot) returns()
func (_Claims *ClaimsTransactorSession) UpdateResolve(pro *big.Int, vot bool) (*types.Transaction, error) {
	return _Claims.Contract.UpdateResolve(&_Claims.TransactOpts, pro, vot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_Claims *ClaimsTransactor) Withdraw(opts *bind.TransactOpts, bal *big.Int) (*types.Transaction, error) {
	return _Claims.contract.Transact(opts, "withdraw", bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_Claims *ClaimsSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.Withdraw(&_Claims.TransactOpts, bal)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 bal) returns()
func (_Claims *ClaimsTransactorSession) Withdraw(bal *big.Int) (*types.Transaction, error) {
	return _Claims.Contract.Withdraw(&_Claims.TransactOpts, bal)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Claims *ClaimsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claims.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Claims *ClaimsSession) Receive() (*types.Transaction, error) {
	return _Claims.Contract.Receive(&_Claims.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Claims *ClaimsTransactorSession) Receive() (*types.Transaction, error) {
	return _Claims.Contract.Receive(&_Claims.TransactOpts)
}

// ClaimsDisputeCreatedIterator is returned from FilterDisputeCreated and is used to iterate over the raw logs and unpacked data for DisputeCreated events raised by the Claims contract.
type ClaimsDisputeCreatedIterator struct {
	Event *ClaimsDisputeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsDisputeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsDisputeCreated)
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
		it.Event = new(ClaimsDisputeCreated)
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
func (it *ClaimsDisputeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsDisputeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsDisputeCreated represents a DisputeCreated event raised by the Claims contract.
type ClaimsDisputeCreated struct {
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisputeCreated is a free log retrieval operation binding the contract event 0x2d63b30cb3e22f7f5b576c8e66fa45b23efbd90969ee2fecb0d1a814990db16c.
//
// Solidity: event DisputeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) FilterDisputeCreated(opts *bind.FilterOpts) (*ClaimsDisputeCreatedIterator, error) {

	logs, sub, err := _Claims.contract.FilterLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsDisputeCreatedIterator{contract: _Claims.contract, event: "DisputeCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeCreated is a free log subscription operation binding the contract event 0x2d63b30cb3e22f7f5b576c8e66fa45b23efbd90969ee2fecb0d1a814990db16c.
//
// Solidity: event DisputeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) WatchDisputeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsDisputeCreated) (event.Subscription, error) {

	logs, sub, err := _Claims.contract.WatchLogs(opts, "DisputeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsDisputeCreated)
				if err := _Claims.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
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

// ParseDisputeCreated is a log parse operation binding the contract event 0x2d63b30cb3e22f7f5b576c8e66fa45b23efbd90969ee2fecb0d1a814990db16c.
//
// Solidity: event DisputeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) ParseDisputeCreated(log types.Log) (*ClaimsDisputeCreated, error) {
	event := new(ClaimsDisputeCreated)
	if err := _Claims.contract.UnpackLog(event, "DisputeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsDisputeSettledIterator is returned from FilterDisputeSettled and is used to iterate over the raw logs and unpacked data for DisputeSettled events raised by the Claims contract.
type ClaimsDisputeSettledIterator struct {
	Event *ClaimsDisputeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsDisputeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsDisputeSettled)
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
		it.Event = new(ClaimsDisputeSettled)
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
func (it *ClaimsDisputeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsDisputeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsDisputeSettled represents a DisputeSettled event raised by the Claims contract.
type ClaimsDisputeSettled struct {
	All *big.Int
	Yay *big.Int
	Nah *big.Int
	Tot *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDisputeSettled is a free log retrieval operation binding the contract event 0xe21a3cfb0be3eec935c7dc8d2d7180967cfc67ef41f4b503da1f3d9888a0e1de.
//
// Solidity: event DisputeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) FilterDisputeSettled(opts *bind.FilterOpts) (*ClaimsDisputeSettledIterator, error) {

	logs, sub, err := _Claims.contract.FilterLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsDisputeSettledIterator{contract: _Claims.contract, event: "DisputeSettled", logs: logs, sub: sub}, nil
}

// WatchDisputeSettled is a free log subscription operation binding the contract event 0xe21a3cfb0be3eec935c7dc8d2d7180967cfc67ef41f4b503da1f3d9888a0e1de.
//
// Solidity: event DisputeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) WatchDisputeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsDisputeSettled) (event.Subscription, error) {

	logs, sub, err := _Claims.contract.WatchLogs(opts, "DisputeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsDisputeSettled)
				if err := _Claims.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
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

// ParseDisputeSettled is a log parse operation binding the contract event 0xe21a3cfb0be3eec935c7dc8d2d7180967cfc67ef41f4b503da1f3d9888a0e1de.
//
// Solidity: event DisputeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) ParseDisputeSettled(log types.Log) (*ClaimsDisputeSettled, error) {
	event := new(ClaimsDisputeSettled)
	if err := _Claims.contract.UnpackLog(event, "DisputeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsProposeCreatedIterator is returned from FilterProposeCreated and is used to iterate over the raw logs and unpacked data for ProposeCreated events raised by the Claims contract.
type ClaimsProposeCreatedIterator struct {
	Event *ClaimsProposeCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsProposeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsProposeCreated)
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
		it.Event = new(ClaimsProposeCreated)
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
func (it *ClaimsProposeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsProposeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsProposeCreated represents a ProposeCreated event raised by the Claims contract.
type ClaimsProposeCreated struct {
	Use common.Address
	Bal *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposeCreated is a free log retrieval operation binding the contract event 0x9958345afb967495d1c8349a9f0b1eb6fa8da1d797a8e97d8a27b3ef35538e18.
//
// Solidity: event ProposeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) FilterProposeCreated(opts *bind.FilterOpts) (*ClaimsProposeCreatedIterator, error) {

	logs, sub, err := _Claims.contract.FilterLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsProposeCreatedIterator{contract: _Claims.contract, event: "ProposeCreated", logs: logs, sub: sub}, nil
}

// WatchProposeCreated is a free log subscription operation binding the contract event 0x9958345afb967495d1c8349a9f0b1eb6fa8da1d797a8e97d8a27b3ef35538e18.
//
// Solidity: event ProposeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) WatchProposeCreated(opts *bind.WatchOpts, sink chan<- *ClaimsProposeCreated) (event.Subscription, error) {

	logs, sub, err := _Claims.contract.WatchLogs(opts, "ProposeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsProposeCreated)
				if err := _Claims.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
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

// ParseProposeCreated is a log parse operation binding the contract event 0x9958345afb967495d1c8349a9f0b1eb6fa8da1d797a8e97d8a27b3ef35538e18.
//
// Solidity: event ProposeCreated(address use, uint256 bal, uint64 exp)
func (_Claims *ClaimsFilterer) ParseProposeCreated(log types.Log) (*ClaimsProposeCreated, error) {
	event := new(ClaimsProposeCreated)
	if err := _Claims.contract.UnpackLog(event, "ProposeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsProposeSettledIterator is returned from FilterProposeSettled and is used to iterate over the raw logs and unpacked data for ProposeSettled events raised by the Claims contract.
type ClaimsProposeSettledIterator struct {
	Event *ClaimsProposeSettled // Event containing the contract specifics and raw log

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
func (it *ClaimsProposeSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsProposeSettled)
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
		it.Event = new(ClaimsProposeSettled)
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
func (it *ClaimsProposeSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsProposeSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsProposeSettled represents a ProposeSettled event raised by the Claims contract.
type ClaimsProposeSettled struct {
	All *big.Int
	Yay *big.Int
	Nah *big.Int
	Tot *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposeSettled is a free log retrieval operation binding the contract event 0xc660c9505d3bd7c0767b64bbebae4279d08366b6f01b10ba5126cffdc3198aee.
//
// Solidity: event ProposeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) FilterProposeSettled(opts *bind.FilterOpts) (*ClaimsProposeSettledIterator, error) {

	logs, sub, err := _Claims.contract.FilterLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return &ClaimsProposeSettledIterator{contract: _Claims.contract, event: "ProposeSettled", logs: logs, sub: sub}, nil
}

// WatchProposeSettled is a free log subscription operation binding the contract event 0xc660c9505d3bd7c0767b64bbebae4279d08366b6f01b10ba5126cffdc3198aee.
//
// Solidity: event ProposeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) WatchProposeSettled(opts *bind.WatchOpts, sink chan<- *ClaimsProposeSettled) (event.Subscription, error) {

	logs, sub, err := _Claims.contract.WatchLogs(opts, "ProposeSettled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsProposeSettled)
				if err := _Claims.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
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

// ParseProposeSettled is a log parse operation binding the contract event 0xc660c9505d3bd7c0767b64bbebae4279d08366b6f01b10ba5126cffdc3198aee.
//
// Solidity: event ProposeSettled(uint256 all, uint256 yay, uint256 nah, uint256 tot)
func (_Claims *ClaimsFilterer) ParseProposeSettled(log types.Log) (*ClaimsProposeSettled, error) {
	event := new(ClaimsProposeSettled)
	if err := _Claims.contract.UnpackLog(event, "ProposeSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsResolveCreatedIterator is returned from FilterResolveCreated and is used to iterate over the raw logs and unpacked data for ResolveCreated events raised by the Claims contract.
type ClaimsResolveCreatedIterator struct {
	Event *ClaimsResolveCreated // Event containing the contract specifics and raw log

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
func (it *ClaimsResolveCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsResolveCreated)
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
		it.Event = new(ClaimsResolveCreated)
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
func (it *ClaimsResolveCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsResolveCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsResolveCreated represents a ResolveCreated event raised by the Claims contract.
type ClaimsResolveCreated struct {
	Use common.Address
	Len *big.Int
	Exp uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResolveCreated is a free log retrieval operation binding the contract event 0xf7430e9eb3d2013265aa9b15a468fe43f023d694832c76117e47e30963f82cf9.
//
// Solidity: event ResolveCreated(address use, uint256 len, uint64 exp)
func (_Claims *ClaimsFilterer) FilterResolveCreated(opts *bind.FilterOpts) (*ClaimsResolveCreatedIterator, error) {

	logs, sub, err := _Claims.contract.FilterLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return &ClaimsResolveCreatedIterator{contract: _Claims.contract, event: "ResolveCreated", logs: logs, sub: sub}, nil
}

// WatchResolveCreated is a free log subscription operation binding the contract event 0xf7430e9eb3d2013265aa9b15a468fe43f023d694832c76117e47e30963f82cf9.
//
// Solidity: event ResolveCreated(address use, uint256 len, uint64 exp)
func (_Claims *ClaimsFilterer) WatchResolveCreated(opts *bind.WatchOpts, sink chan<- *ClaimsResolveCreated) (event.Subscription, error) {

	logs, sub, err := _Claims.contract.WatchLogs(opts, "ResolveCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsResolveCreated)
				if err := _Claims.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
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

// ParseResolveCreated is a log parse operation binding the contract event 0xf7430e9eb3d2013265aa9b15a468fe43f023d694832c76117e47e30963f82cf9.
//
// Solidity: event ResolveCreated(address use, uint256 len, uint64 exp)
func (_Claims *ClaimsFilterer) ParseResolveCreated(log types.Log) (*ClaimsResolveCreated, error) {
	event := new(ClaimsResolveCreated)
	if err := _Claims.contract.UnpackLog(event, "ResolveCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Claims contract.
type ClaimsRoleAdminChangedIterator struct {
	Event *ClaimsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ClaimsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsRoleAdminChanged)
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
		it.Event = new(ClaimsRoleAdminChanged)
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
func (it *ClaimsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsRoleAdminChanged represents a RoleAdminChanged event raised by the Claims contract.
type ClaimsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Claims *ClaimsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ClaimsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Claims.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsRoleAdminChangedIterator{contract: _Claims.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Claims *ClaimsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ClaimsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Claims.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsRoleAdminChanged)
				if err := _Claims.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Claims *ClaimsFilterer) ParseRoleAdminChanged(log types.Log) (*ClaimsRoleAdminChanged, error) {
	event := new(ClaimsRoleAdminChanged)
	if err := _Claims.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Claims contract.
type ClaimsRoleGrantedIterator struct {
	Event *ClaimsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ClaimsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsRoleGranted)
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
		it.Event = new(ClaimsRoleGranted)
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
func (it *ClaimsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsRoleGranted represents a RoleGranted event raised by the Claims contract.
type ClaimsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Claims *ClaimsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsRoleGrantedIterator, error) {

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

	logs, sub, err := _Claims.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsRoleGrantedIterator{contract: _Claims.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Claims *ClaimsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ClaimsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Claims.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsRoleGranted)
				if err := _Claims.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Claims *ClaimsFilterer) ParseRoleGranted(log types.Log) (*ClaimsRoleGranted, error) {
	event := new(ClaimsRoleGranted)
	if err := _Claims.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Claims contract.
type ClaimsRoleRevokedIterator struct {
	Event *ClaimsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ClaimsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimsRoleRevoked)
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
		it.Event = new(ClaimsRoleRevoked)
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
func (it *ClaimsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimsRoleRevoked represents a RoleRevoked event raised by the Claims contract.
type ClaimsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Claims *ClaimsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ClaimsRoleRevokedIterator, error) {

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

	logs, sub, err := _Claims.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ClaimsRoleRevokedIterator{contract: _Claims.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Claims *ClaimsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ClaimsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Claims.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimsRoleRevoked)
				if err := _Claims.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Claims *ClaimsFilterer) ParseRoleRevoked(log types.Log) (*ClaimsRoleRevoked, error) {
	event := new(ClaimsRoleRevoked)
	if err := _Claims.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
