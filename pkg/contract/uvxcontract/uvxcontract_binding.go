// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uvxcontract

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

// UvxContractBindingMetaData contains all meta data concerning the UvxContractBinding contract.
var UvxContractBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOAN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rec\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tk1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tk2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bl1\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outstanding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrict\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UvxContractBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use UvxContractBindingMetaData.ABI instead.
var UvxContractBindingABI = UvxContractBindingMetaData.ABI

// UvxContractBinding is an auto generated Go binding around an Ethereum contract.
type UvxContractBinding struct {
	UvxContractBindingCaller     // Read-only binding to the contract
	UvxContractBindingTransactor // Write-only binding to the contract
	UvxContractBindingFilterer   // Log filterer for contract events
}

// UvxContractBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type UvxContractBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UvxContractBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UvxContractBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UvxContractBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UvxContractBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UvxContractBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UvxContractBindingSession struct {
	Contract     *UvxContractBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UvxContractBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UvxContractBindingCallerSession struct {
	Contract *UvxContractBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UvxContractBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UvxContractBindingTransactorSession struct {
	Contract     *UvxContractBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UvxContractBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type UvxContractBindingRaw struct {
	Contract *UvxContractBinding // Generic contract binding to access the raw methods on
}

// UvxContractBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UvxContractBindingCallerRaw struct {
	Contract *UvxContractBindingCaller // Generic read-only contract binding to access the raw methods on
}

// UvxContractBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UvxContractBindingTransactorRaw struct {
	Contract *UvxContractBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUvxContractBinding creates a new instance of UvxContractBinding, bound to a specific deployed contract.
func NewUvxContractBinding(address common.Address, backend bind.ContractBackend) (*UvxContractBinding, error) {
	contract, err := bindUvxContractBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UvxContractBinding{UvxContractBindingCaller: UvxContractBindingCaller{contract: contract}, UvxContractBindingTransactor: UvxContractBindingTransactor{contract: contract}, UvxContractBindingFilterer: UvxContractBindingFilterer{contract: contract}}, nil
}

// NewUvxContractBindingCaller creates a new read-only instance of UvxContractBinding, bound to a specific deployed contract.
func NewUvxContractBindingCaller(address common.Address, caller bind.ContractCaller) (*UvxContractBindingCaller, error) {
	contract, err := bindUvxContractBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingCaller{contract: contract}, nil
}

// NewUvxContractBindingTransactor creates a new write-only instance of UvxContractBinding, bound to a specific deployed contract.
func NewUvxContractBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*UvxContractBindingTransactor, error) {
	contract, err := bindUvxContractBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingTransactor{contract: contract}, nil
}

// NewUvxContractBindingFilterer creates a new log filterer instance of UvxContractBinding, bound to a specific deployed contract.
func NewUvxContractBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*UvxContractBindingFilterer, error) {
	contract, err := bindUvxContractBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingFilterer{contract: contract}, nil
}

// bindUvxContractBinding binds a generic wrapper to an already deployed contract.
func bindUvxContractBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UvxContractBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UvxContractBinding *UvxContractBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UvxContractBinding.Contract.UvxContractBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UvxContractBinding *UvxContractBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UvxContractBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UvxContractBinding *UvxContractBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UvxContractBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UvxContractBinding *UvxContractBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UvxContractBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UvxContractBinding *UvxContractBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UvxContractBinding *UvxContractBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.contract.Transact(opts, method, params...)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) BOTROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.BOTROLE(&_UvxContractBinding.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) BOTROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.BOTROLE(&_UvxContractBinding.CallOpts)
}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) CONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) CONTRACTROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.CONTRACTROLE(&_UvxContractBinding.CallOpts)
}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) CONTRACTROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.CONTRACTROLE(&_UvxContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.DEFAULTADMINROLE(&_UvxContractBinding.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.DEFAULTADMINROLE(&_UvxContractBinding.CallOpts)
}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) LOANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "LOAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) LOANROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.LOANROLE(&_UvxContractBinding.CallOpts)
}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) LOANROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.LOANROLE(&_UvxContractBinding.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) TOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) TOKENROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.TOKENROLE(&_UvxContractBinding.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) TOKENROLE() ([32]byte, error) {
	return _UvxContractBinding.Contract.TOKENROLE(&_UvxContractBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UvxContractBinding *UvxContractBindingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UvxContractBinding *UvxContractBindingSession) VERSION() (string, error) {
	return _UvxContractBinding.Contract.VERSION(&_UvxContractBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UvxContractBinding *UvxContractBindingCallerSession) VERSION() (string, error) {
	return _UvxContractBinding.Contract.VERSION(&_UvxContractBinding.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UvxContractBinding.Contract.Allowance(&_UvxContractBinding.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UvxContractBinding.Contract.Allowance(&_UvxContractBinding.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UvxContractBinding.Contract.BalanceOf(&_UvxContractBinding.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UvxContractBinding.Contract.BalanceOf(&_UvxContractBinding.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UvxContractBinding *UvxContractBindingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UvxContractBinding *UvxContractBindingSession) Decimals() (uint8, error) {
	return _UvxContractBinding.Contract.Decimals(&_UvxContractBinding.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UvxContractBinding *UvxContractBindingCallerSession) Decimals() (uint8, error) {
	return _UvxContractBinding.Contract.Decimals(&_UvxContractBinding.CallOpts)
}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UvxContractBinding *UvxContractBindingCaller) Freeze(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "freeze")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) Freeze() (bool, error) {
	return _UvxContractBinding.Contract.Freeze(&_UvxContractBinding.CallOpts)
}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UvxContractBinding *UvxContractBindingCallerSession) Freeze() (bool, error) {
	return _UvxContractBinding.Contract.Freeze(&_UvxContractBinding.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UvxContractBinding.Contract.GetRoleAdmin(&_UvxContractBinding.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UvxContractBinding *UvxContractBindingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UvxContractBinding.Contract.GetRoleAdmin(&_UvxContractBinding.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UvxContractBinding *UvxContractBindingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UvxContractBinding *UvxContractBindingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UvxContractBinding.Contract.GetRoleMember(&_UvxContractBinding.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UvxContractBinding *UvxContractBindingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UvxContractBinding.Contract.GetRoleMember(&_UvxContractBinding.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UvxContractBinding.Contract.GetRoleMemberCount(&_UvxContractBinding.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UvxContractBinding.Contract.GetRoleMemberCount(&_UvxContractBinding.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UvxContractBinding *UvxContractBindingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UvxContractBinding.Contract.HasRole(&_UvxContractBinding.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UvxContractBinding *UvxContractBindingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UvxContractBinding.Contract.HasRole(&_UvxContractBinding.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UvxContractBinding *UvxContractBindingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UvxContractBinding *UvxContractBindingSession) Name() (string, error) {
	return _UvxContractBinding.Contract.Name(&_UvxContractBinding.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UvxContractBinding *UvxContractBindingCallerSession) Name() (string, error) {
	return _UvxContractBinding.Contract.Name(&_UvxContractBinding.CallOpts)
}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCaller) Outstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "outstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingSession) Outstanding() (*big.Int, error) {
	return _UvxContractBinding.Contract.Outstanding(&_UvxContractBinding.CallOpts)
}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCallerSession) Outstanding() (*big.Int, error) {
	return _UvxContractBinding.Contract.Outstanding(&_UvxContractBinding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UvxContractBinding *UvxContractBindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UvxContractBinding *UvxContractBindingSession) Owner() (common.Address, error) {
	return _UvxContractBinding.Contract.Owner(&_UvxContractBinding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UvxContractBinding *UvxContractBindingCallerSession) Owner() (common.Address, error) {
	return _UvxContractBinding.Contract.Owner(&_UvxContractBinding.CallOpts)
}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UvxContractBinding *UvxContractBindingCaller) Restrict(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "restrict")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) Restrict() (bool, error) {
	return _UvxContractBinding.Contract.Restrict(&_UvxContractBinding.CallOpts)
}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UvxContractBinding *UvxContractBindingCallerSession) Restrict() (bool, error) {
	return _UvxContractBinding.Contract.Restrict(&_UvxContractBinding.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UvxContractBinding *UvxContractBindingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UvxContractBinding.Contract.SupportsInterface(&_UvxContractBinding.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UvxContractBinding *UvxContractBindingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UvxContractBinding.Contract.SupportsInterface(&_UvxContractBinding.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UvxContractBinding *UvxContractBindingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UvxContractBinding *UvxContractBindingSession) Symbol() (string, error) {
	return _UvxContractBinding.Contract.Symbol(&_UvxContractBinding.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UvxContractBinding *UvxContractBindingCallerSession) Symbol() (string, error) {
	return _UvxContractBinding.Contract.Symbol(&_UvxContractBinding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UvxContractBinding.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingSession) TotalSupply() (*big.Int, error) {
	return _UvxContractBinding.Contract.TotalSupply(&_UvxContractBinding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UvxContractBinding *UvxContractBindingCallerSession) TotalSupply() (*big.Int, error) {
	return _UvxContractBinding.Contract.TotalSupply(&_UvxContractBinding.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Approve(&_UvxContractBinding.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Approve(&_UvxContractBinding.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Burn(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "burn", tok, bal)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingSession) Burn(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Burn(&_UvxContractBinding.TransactOpts, tok, bal)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Burn(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Burn(&_UvxContractBinding.TransactOpts, tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Fund(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "fund", tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingSession) Fund(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Fund(&_UvxContractBinding.TransactOpts, tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Fund(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Fund(&_UvxContractBinding.TransactOpts, tok, bal)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.GrantRole(&_UvxContractBinding.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.GrantRole(&_UvxContractBinding.TransactOpts, role, account)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Lend(opts *bind.TransactOpts, rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "lend", rec, tk1, tk2, bl1)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UvxContractBinding *UvxContractBindingSession) Lend(rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Lend(&_UvxContractBinding.TransactOpts, rec, tk1, tk2, bl1)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Lend(rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Lend(&_UvxContractBinding.TransactOpts, rec, tk1, tk2, bl1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Mint(opts *bind.TransactOpts, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "mint", dst, bal)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingSession) Mint(dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Mint(&_UvxContractBinding.TransactOpts, dst, bal)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Mint(dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Mint(&_UvxContractBinding.TransactOpts, dst, bal)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UvxContractBinding *UvxContractBindingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.RenounceRole(&_UvxContractBinding.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.RenounceRole(&_UvxContractBinding.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.RevokeRole(&_UvxContractBinding.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.RevokeRole(&_UvxContractBinding.TransactOpts, role, account)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Sell(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "sell", tok, bal)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingSession) Sell(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Sell(&_UvxContractBinding.TransactOpts, tok, bal)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Sell(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Sell(&_UvxContractBinding.TransactOpts, tok, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "transfer", to, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) Transfer(to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Transfer(&_UvxContractBinding.TransactOpts, to, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactorSession) Transfer(to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Transfer(&_UvxContractBinding.TransactOpts, to, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "transferFrom", src, dst, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingSession) TransferFrom(src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.TransferFrom(&_UvxContractBinding.TransactOpts, src, dst, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UvxContractBinding *UvxContractBindingTransactorSession) TransferFrom(src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.TransferFrom(&_UvxContractBinding.TransactOpts, src, dst, bal)
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UvxContractBinding *UvxContractBindingTransactor) UpdateFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "updateFreeze")
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UvxContractBinding *UvxContractBindingSession) UpdateFreeze() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateFreeze(&_UvxContractBinding.TransactOpts)
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) UpdateFreeze() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateFreeze(&_UvxContractBinding.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UvxContractBinding *UvxContractBindingTransactor) UpdateOwner(opts *bind.TransactOpts, own common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "updateOwner", own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UvxContractBinding *UvxContractBindingSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateOwner(&_UvxContractBinding.TransactOpts, own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateOwner(&_UvxContractBinding.TransactOpts, own)
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UvxContractBinding *UvxContractBindingTransactor) UpdateRestrict(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UvxContractBinding.contract.Transact(opts, "updateRestrict")
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UvxContractBinding *UvxContractBindingSession) UpdateRestrict() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateRestrict(&_UvxContractBinding.TransactOpts)
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) UpdateRestrict() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.UpdateRestrict(&_UvxContractBinding.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UvxContractBinding *UvxContractBindingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UvxContractBinding.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UvxContractBinding *UvxContractBindingSession) Receive() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Receive(&_UvxContractBinding.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UvxContractBinding *UvxContractBindingTransactorSession) Receive() (*types.Transaction, error) {
	return _UvxContractBinding.Contract.Receive(&_UvxContractBinding.TransactOpts)
}

// UvxContractBindingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UvxContractBinding contract.
type UvxContractBindingApprovalIterator struct {
	Event *UvxContractBindingApproval // Event containing the contract specifics and raw log

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
func (it *UvxContractBindingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UvxContractBindingApproval)
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
		it.Event = new(UvxContractBindingApproval)
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
func (it *UvxContractBindingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UvxContractBindingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UvxContractBindingApproval represents a Approval event raised by the UvxContractBinding contract.
type UvxContractBindingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UvxContractBindingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UvxContractBinding.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingApprovalIterator{contract: _UvxContractBinding.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UvxContractBindingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UvxContractBinding.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UvxContractBindingApproval)
				if err := _UvxContractBinding.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) ParseApproval(log types.Log) (*UvxContractBindingApproval, error) {
	event := new(UvxContractBindingApproval)
	if err := _UvxContractBinding.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UvxContractBindingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UvxContractBinding contract.
type UvxContractBindingRoleAdminChangedIterator struct {
	Event *UvxContractBindingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UvxContractBindingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UvxContractBindingRoleAdminChanged)
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
		it.Event = new(UvxContractBindingRoleAdminChanged)
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
func (it *UvxContractBindingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UvxContractBindingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UvxContractBindingRoleAdminChanged represents a RoleAdminChanged event raised by the UvxContractBinding contract.
type UvxContractBindingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UvxContractBinding *UvxContractBindingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UvxContractBindingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _UvxContractBinding.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingRoleAdminChangedIterator{contract: _UvxContractBinding.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UvxContractBinding *UvxContractBindingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UvxContractBindingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _UvxContractBinding.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UvxContractBindingRoleAdminChanged)
				if err := _UvxContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_UvxContractBinding *UvxContractBindingFilterer) ParseRoleAdminChanged(log types.Log) (*UvxContractBindingRoleAdminChanged, error) {
	event := new(UvxContractBindingRoleAdminChanged)
	if err := _UvxContractBinding.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UvxContractBindingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UvxContractBinding contract.
type UvxContractBindingRoleGrantedIterator struct {
	Event *UvxContractBindingRoleGranted // Event containing the contract specifics and raw log

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
func (it *UvxContractBindingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UvxContractBindingRoleGranted)
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
		it.Event = new(UvxContractBindingRoleGranted)
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
func (it *UvxContractBindingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UvxContractBindingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UvxContractBindingRoleGranted represents a RoleGranted event raised by the UvxContractBinding contract.
type UvxContractBindingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UvxContractBinding *UvxContractBindingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UvxContractBindingRoleGrantedIterator, error) {

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

	logs, sub, err := _UvxContractBinding.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingRoleGrantedIterator{contract: _UvxContractBinding.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UvxContractBinding *UvxContractBindingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UvxContractBindingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UvxContractBinding.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UvxContractBindingRoleGranted)
				if err := _UvxContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UvxContractBinding *UvxContractBindingFilterer) ParseRoleGranted(log types.Log) (*UvxContractBindingRoleGranted, error) {
	event := new(UvxContractBindingRoleGranted)
	if err := _UvxContractBinding.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UvxContractBindingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UvxContractBinding contract.
type UvxContractBindingRoleRevokedIterator struct {
	Event *UvxContractBindingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UvxContractBindingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UvxContractBindingRoleRevoked)
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
		it.Event = new(UvxContractBindingRoleRevoked)
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
func (it *UvxContractBindingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UvxContractBindingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UvxContractBindingRoleRevoked represents a RoleRevoked event raised by the UvxContractBinding contract.
type UvxContractBindingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UvxContractBinding *UvxContractBindingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UvxContractBindingRoleRevokedIterator, error) {

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

	logs, sub, err := _UvxContractBinding.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingRoleRevokedIterator{contract: _UvxContractBinding.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UvxContractBinding *UvxContractBindingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UvxContractBindingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UvxContractBinding.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UvxContractBindingRoleRevoked)
				if err := _UvxContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UvxContractBinding *UvxContractBindingFilterer) ParseRoleRevoked(log types.Log) (*UvxContractBindingRoleRevoked, error) {
	event := new(UvxContractBindingRoleRevoked)
	if err := _UvxContractBinding.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UvxContractBindingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UvxContractBinding contract.
type UvxContractBindingTransferIterator struct {
	Event *UvxContractBindingTransfer // Event containing the contract specifics and raw log

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
func (it *UvxContractBindingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UvxContractBindingTransfer)
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
		it.Event = new(UvxContractBindingTransfer)
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
func (it *UvxContractBindingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UvxContractBindingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UvxContractBindingTransfer represents a Transfer event raised by the UvxContractBinding contract.
type UvxContractBindingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UvxContractBindingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UvxContractBinding.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UvxContractBindingTransferIterator{contract: _UvxContractBinding.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UvxContractBindingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UvxContractBinding.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UvxContractBindingTransfer)
				if err := _UvxContractBinding.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UvxContractBinding *UvxContractBindingFilterer) ParseTransfer(log types.Log) (*UvxContractBindingTransfer, error) {
	event := new(UvxContractBindingTransfer)
	if err := _UvxContractBinding.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
