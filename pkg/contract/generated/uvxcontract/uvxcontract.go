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

// UVXMetaData contains all meta data concerning the UVX contract.
var UVXMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Address\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"why\",\"type\":\"string\"}],\"name\":\"Process\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOAN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeze\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rec\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tk1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tk2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bl1\",\"type\":\"uint256\"}],\"name\":\"lend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outstanding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrict\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tok\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bal\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateFreeze\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UVXABI is the input ABI used to generate the binding from.
// Deprecated: Use UVXMetaData.ABI instead.
var UVXABI = UVXMetaData.ABI

// UVX is an auto generated Go binding around an Ethereum contract.
type UVX struct {
	UVXCaller     // Read-only binding to the contract
	UVXTransactor // Write-only binding to the contract
	UVXFilterer   // Log filterer for contract events
}

// UVXCaller is an auto generated read-only Go binding around an Ethereum contract.
type UVXCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UVXTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UVXTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UVXFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UVXFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UVXSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UVXSession struct {
	Contract     *UVX              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UVXCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UVXCallerSession struct {
	Contract *UVXCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UVXTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UVXTransactorSession struct {
	Contract     *UVXTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UVXRaw is an auto generated low-level Go binding around an Ethereum contract.
type UVXRaw struct {
	Contract *UVX // Generic contract binding to access the raw methods on
}

// UVXCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UVXCallerRaw struct {
	Contract *UVXCaller // Generic read-only contract binding to access the raw methods on
}

// UVXTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UVXTransactorRaw struct {
	Contract *UVXTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUVX creates a new instance of UVX, bound to a specific deployed contract.
func NewUVX(address common.Address, backend bind.ContractBackend) (*UVX, error) {
	contract, err := bindUVX(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UVX{UVXCaller: UVXCaller{contract: contract}, UVXTransactor: UVXTransactor{contract: contract}, UVXFilterer: UVXFilterer{contract: contract}}, nil
}

// NewUVXCaller creates a new read-only instance of UVX, bound to a specific deployed contract.
func NewUVXCaller(address common.Address, caller bind.ContractCaller) (*UVXCaller, error) {
	contract, err := bindUVX(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UVXCaller{contract: contract}, nil
}

// NewUVXTransactor creates a new write-only instance of UVX, bound to a specific deployed contract.
func NewUVXTransactor(address common.Address, transactor bind.ContractTransactor) (*UVXTransactor, error) {
	contract, err := bindUVX(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UVXTransactor{contract: contract}, nil
}

// NewUVXFilterer creates a new log filterer instance of UVX, bound to a specific deployed contract.
func NewUVXFilterer(address common.Address, filterer bind.ContractFilterer) (*UVXFilterer, error) {
	contract, err := bindUVX(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UVXFilterer{contract: contract}, nil
}

// bindUVX binds a generic wrapper to an already deployed contract.
func bindUVX(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UVXMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UVX *UVXRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UVX.Contract.UVXCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UVX *UVXRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UVX.Contract.UVXTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UVX *UVXRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UVX.Contract.UVXTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UVX *UVXCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UVX.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UVX *UVXTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UVX.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UVX *UVXTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UVX.Contract.contract.Transact(opts, method, params...)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UVX *UVXCaller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UVX *UVXSession) BOTROLE() ([32]byte, error) {
	return _UVX.Contract.BOTROLE(&_UVX.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_UVX *UVXCallerSession) BOTROLE() ([32]byte, error) {
	return _UVX.Contract.BOTROLE(&_UVX.CallOpts)
}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UVX *UVXCaller) CONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UVX *UVXSession) CONTRACTROLE() ([32]byte, error) {
	return _UVX.Contract.CONTRACTROLE(&_UVX.CallOpts)
}

// CONTRACTROLE is a free data retrieval call binding the contract method 0x03fe46ab.
//
// Solidity: function CONTRACT_ROLE() view returns(bytes32)
func (_UVX *UVXCallerSession) CONTRACTROLE() ([32]byte, error) {
	return _UVX.Contract.CONTRACTROLE(&_UVX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UVX *UVXCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UVX *UVXSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UVX.Contract.DEFAULTADMINROLE(&_UVX.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UVX *UVXCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UVX.Contract.DEFAULTADMINROLE(&_UVX.CallOpts)
}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UVX *UVXCaller) LOANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "LOAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UVX *UVXSession) LOANROLE() ([32]byte, error) {
	return _UVX.Contract.LOANROLE(&_UVX.CallOpts)
}

// LOANROLE is a free data retrieval call binding the contract method 0xec4bbb4d.
//
// Solidity: function LOAN_ROLE() view returns(bytes32)
func (_UVX *UVXCallerSession) LOANROLE() ([32]byte, error) {
	return _UVX.Contract.LOANROLE(&_UVX.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UVX *UVXCaller) TOKENROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "TOKEN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UVX *UVXSession) TOKENROLE() ([32]byte, error) {
	return _UVX.Contract.TOKENROLE(&_UVX.CallOpts)
}

// TOKENROLE is a free data retrieval call binding the contract method 0x8301057b.
//
// Solidity: function TOKEN_ROLE() view returns(bytes32)
func (_UVX *UVXCallerSession) TOKENROLE() ([32]byte, error) {
	return _UVX.Contract.TOKENROLE(&_UVX.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UVX *UVXCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UVX *UVXSession) VERSION() (string, error) {
	return _UVX.Contract.VERSION(&_UVX.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UVX *UVXCallerSession) VERSION() (string, error) {
	return _UVX.Contract.VERSION(&_UVX.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UVX *UVXCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UVX *UVXSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UVX.Contract.Allowance(&_UVX.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UVX *UVXCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UVX.Contract.Allowance(&_UVX.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UVX *UVXCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UVX *UVXSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UVX.Contract.BalanceOf(&_UVX.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_UVX *UVXCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _UVX.Contract.BalanceOf(&_UVX.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UVX *UVXCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UVX *UVXSession) Decimals() (uint8, error) {
	return _UVX.Contract.Decimals(&_UVX.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_UVX *UVXCallerSession) Decimals() (uint8, error) {
	return _UVX.Contract.Decimals(&_UVX.CallOpts)
}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UVX *UVXCaller) Freeze(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "freeze")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UVX *UVXSession) Freeze() (bool, error) {
	return _UVX.Contract.Freeze(&_UVX.CallOpts)
}

// Freeze is a free data retrieval call binding the contract method 0x62a5af3b.
//
// Solidity: function freeze() view returns(bool)
func (_UVX *UVXCallerSession) Freeze() (bool, error) {
	return _UVX.Contract.Freeze(&_UVX.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UVX *UVXCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UVX *UVXSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UVX.Contract.GetRoleAdmin(&_UVX.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UVX *UVXCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UVX.Contract.GetRoleAdmin(&_UVX.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UVX *UVXCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UVX *UVXSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UVX.Contract.GetRoleMember(&_UVX.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_UVX *UVXCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _UVX.Contract.GetRoleMember(&_UVX.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UVX *UVXCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UVX *UVXSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UVX.Contract.GetRoleMemberCount(&_UVX.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_UVX *UVXCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _UVX.Contract.GetRoleMemberCount(&_UVX.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UVX *UVXCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UVX *UVXSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UVX.Contract.HasRole(&_UVX.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UVX *UVXCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UVX.Contract.HasRole(&_UVX.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UVX *UVXCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UVX *UVXSession) Name() (string, error) {
	return _UVX.Contract.Name(&_UVX.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_UVX *UVXCallerSession) Name() (string, error) {
	return _UVX.Contract.Name(&_UVX.CallOpts)
}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UVX *UVXCaller) Outstanding(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "outstanding")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UVX *UVXSession) Outstanding() (*big.Int, error) {
	return _UVX.Contract.Outstanding(&_UVX.CallOpts)
}

// Outstanding is a free data retrieval call binding the contract method 0xe50d33e3.
//
// Solidity: function outstanding() view returns(uint256)
func (_UVX *UVXCallerSession) Outstanding() (*big.Int, error) {
	return _UVX.Contract.Outstanding(&_UVX.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UVX *UVXCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UVX *UVXSession) Owner() (common.Address, error) {
	return _UVX.Contract.Owner(&_UVX.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UVX *UVXCallerSession) Owner() (common.Address, error) {
	return _UVX.Contract.Owner(&_UVX.CallOpts)
}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UVX *UVXCaller) Restrict(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "restrict")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UVX *UVXSession) Restrict() (bool, error) {
	return _UVX.Contract.Restrict(&_UVX.CallOpts)
}

// Restrict is a free data retrieval call binding the contract method 0x9649d98b.
//
// Solidity: function restrict() view returns(bool)
func (_UVX *UVXCallerSession) Restrict() (bool, error) {
	return _UVX.Contract.Restrict(&_UVX.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UVX *UVXCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UVX *UVXSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UVX.Contract.SupportsInterface(&_UVX.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UVX *UVXCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UVX.Contract.SupportsInterface(&_UVX.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UVX *UVXCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UVX *UVXSession) Symbol() (string, error) {
	return _UVX.Contract.Symbol(&_UVX.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_UVX *UVXCallerSession) Symbol() (string, error) {
	return _UVX.Contract.Symbol(&_UVX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UVX *UVXCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UVX.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UVX *UVXSession) TotalSupply() (*big.Int, error) {
	return _UVX.Contract.TotalSupply(&_UVX.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UVX *UVXCallerSession) TotalSupply() (*big.Int, error) {
	return _UVX.Contract.TotalSupply(&_UVX.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UVX *UVXTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UVX *UVXSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Approve(&_UVX.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UVX *UVXTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Approve(&_UVX.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UVX *UVXTransactor) Burn(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "burn", tok, bal)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UVX *UVXSession) Burn(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Burn(&_UVX.TransactOpts, tok, bal)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address tok, uint256 bal) returns()
func (_UVX *UVXTransactorSession) Burn(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Burn(&_UVX.TransactOpts, tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UVX *UVXTransactor) Fund(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "fund", tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UVX *UVXSession) Fund(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Fund(&_UVX.TransactOpts, tok, bal)
}

// Fund is a paid mutator transaction binding the contract method 0x7b1837de.
//
// Solidity: function fund(address tok, uint256 bal) returns()
func (_UVX *UVXTransactorSession) Fund(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Fund(&_UVX.TransactOpts, tok, bal)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UVX *UVXTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UVX *UVXSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.Contract.GrantRole(&_UVX.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UVX *UVXTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.Contract.GrantRole(&_UVX.TransactOpts, role, account)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UVX *UVXTransactor) Lend(opts *bind.TransactOpts, rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "lend", rec, tk1, tk2, bl1)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UVX *UVXSession) Lend(rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Lend(&_UVX.TransactOpts, rec, tk1, tk2, bl1)
}

// Lend is a paid mutator transaction binding the contract method 0xecee1bf7.
//
// Solidity: function lend(address rec, address tk1, address tk2, uint256 bl1) returns()
func (_UVX *UVXTransactorSession) Lend(rec common.Address, tk1 common.Address, tk2 common.Address, bl1 *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Lend(&_UVX.TransactOpts, rec, tk1, tk2, bl1)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UVX *UVXTransactor) Mint(opts *bind.TransactOpts, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "mint", dst, bal)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UVX *UVXSession) Mint(dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Mint(&_UVX.TransactOpts, dst, bal)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address dst, uint256 bal) returns()
func (_UVX *UVXTransactorSession) Mint(dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Mint(&_UVX.TransactOpts, dst, bal)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UVX *UVXTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UVX *UVXSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UVX.Contract.RenounceRole(&_UVX.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UVX *UVXTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UVX.Contract.RenounceRole(&_UVX.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UVX *UVXTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UVX *UVXSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.Contract.RevokeRole(&_UVX.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UVX *UVXTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UVX.Contract.RevokeRole(&_UVX.TransactOpts, role, account)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UVX *UVXTransactor) Sell(opts *bind.TransactOpts, tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "sell", tok, bal)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UVX *UVXSession) Sell(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Sell(&_UVX.TransactOpts, tok, bal)
}

// Sell is a paid mutator transaction binding the contract method 0x6c197ff5.
//
// Solidity: function sell(address tok, uint256 bal) returns()
func (_UVX *UVXTransactorSession) Sell(tok common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Sell(&_UVX.TransactOpts, tok, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UVX *UVXTransactor) Transfer(opts *bind.TransactOpts, to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "transfer", to, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UVX *UVXSession) Transfer(to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Transfer(&_UVX.TransactOpts, to, bal)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 bal) returns(bool)
func (_UVX *UVXTransactorSession) Transfer(to common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.Transfer(&_UVX.TransactOpts, to, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UVX *UVXTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "transferFrom", src, dst, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UVX *UVXSession) TransferFrom(src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.TransferFrom(&_UVX.TransactOpts, src, dst, bal)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 bal) returns(bool)
func (_UVX *UVXTransactorSession) TransferFrom(src common.Address, dst common.Address, bal *big.Int) (*types.Transaction, error) {
	return _UVX.Contract.TransferFrom(&_UVX.TransactOpts, src, dst, bal)
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UVX *UVXTransactor) UpdateFreeze(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "updateFreeze")
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UVX *UVXSession) UpdateFreeze() (*types.Transaction, error) {
	return _UVX.Contract.UpdateFreeze(&_UVX.TransactOpts)
}

// UpdateFreeze is a paid mutator transaction binding the contract method 0xb49d1067.
//
// Solidity: function updateFreeze() returns()
func (_UVX *UVXTransactorSession) UpdateFreeze() (*types.Transaction, error) {
	return _UVX.Contract.UpdateFreeze(&_UVX.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UVX *UVXTransactor) UpdateOwner(opts *bind.TransactOpts, own common.Address) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "updateOwner", own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UVX *UVXSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _UVX.Contract.UpdateOwner(&_UVX.TransactOpts, own)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address own) returns()
func (_UVX *UVXTransactorSession) UpdateOwner(own common.Address) (*types.Transaction, error) {
	return _UVX.Contract.UpdateOwner(&_UVX.TransactOpts, own)
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UVX *UVXTransactor) UpdateRestrict(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UVX.contract.Transact(opts, "updateRestrict")
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UVX *UVXSession) UpdateRestrict() (*types.Transaction, error) {
	return _UVX.Contract.UpdateRestrict(&_UVX.TransactOpts)
}

// UpdateRestrict is a paid mutator transaction binding the contract method 0x3ccf50e3.
//
// Solidity: function updateRestrict() returns()
func (_UVX *UVXTransactorSession) UpdateRestrict() (*types.Transaction, error) {
	return _UVX.Contract.UpdateRestrict(&_UVX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UVX *UVXTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UVX.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UVX *UVXSession) Receive() (*types.Transaction, error) {
	return _UVX.Contract.Receive(&_UVX.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UVX *UVXTransactorSession) Receive() (*types.Transaction, error) {
	return _UVX.Contract.Receive(&_UVX.TransactOpts)
}

// UVXApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UVX contract.
type UVXApprovalIterator struct {
	Event *UVXApproval // Event containing the contract specifics and raw log

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
func (it *UVXApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UVXApproval)
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
		it.Event = new(UVXApproval)
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
func (it *UVXApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UVXApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UVXApproval represents a Approval event raised by the UVX contract.
type UVXApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UVX *UVXFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UVXApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UVX.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UVXApprovalIterator{contract: _UVX.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UVX *UVXFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UVXApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UVX.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UVXApproval)
				if err := _UVX.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_UVX *UVXFilterer) ParseApproval(log types.Log) (*UVXApproval, error) {
	event := new(UVXApproval)
	if err := _UVX.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UVXRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UVX contract.
type UVXRoleAdminChangedIterator struct {
	Event *UVXRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UVXRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UVXRoleAdminChanged)
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
		it.Event = new(UVXRoleAdminChanged)
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
func (it *UVXRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UVXRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UVXRoleAdminChanged represents a RoleAdminChanged event raised by the UVX contract.
type UVXRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UVX *UVXFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UVXRoleAdminChangedIterator, error) {

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

	logs, sub, err := _UVX.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UVXRoleAdminChangedIterator{contract: _UVX.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UVX *UVXFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UVXRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _UVX.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UVXRoleAdminChanged)
				if err := _UVX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_UVX *UVXFilterer) ParseRoleAdminChanged(log types.Log) (*UVXRoleAdminChanged, error) {
	event := new(UVXRoleAdminChanged)
	if err := _UVX.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UVXRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UVX contract.
type UVXRoleGrantedIterator struct {
	Event *UVXRoleGranted // Event containing the contract specifics and raw log

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
func (it *UVXRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UVXRoleGranted)
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
		it.Event = new(UVXRoleGranted)
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
func (it *UVXRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UVXRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UVXRoleGranted represents a RoleGranted event raised by the UVX contract.
type UVXRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UVX *UVXFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UVXRoleGrantedIterator, error) {

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

	logs, sub, err := _UVX.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UVXRoleGrantedIterator{contract: _UVX.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UVX *UVXFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UVXRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UVX.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UVXRoleGranted)
				if err := _UVX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UVX *UVXFilterer) ParseRoleGranted(log types.Log) (*UVXRoleGranted, error) {
	event := new(UVXRoleGranted)
	if err := _UVX.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UVXRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UVX contract.
type UVXRoleRevokedIterator struct {
	Event *UVXRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UVXRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UVXRoleRevoked)
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
		it.Event = new(UVXRoleRevoked)
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
func (it *UVXRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UVXRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UVXRoleRevoked represents a RoleRevoked event raised by the UVX contract.
type UVXRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UVX *UVXFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UVXRoleRevokedIterator, error) {

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

	logs, sub, err := _UVX.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UVXRoleRevokedIterator{contract: _UVX.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UVX *UVXFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UVXRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UVX.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UVXRoleRevoked)
				if err := _UVX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UVX *UVXFilterer) ParseRoleRevoked(log types.Log) (*UVXRoleRevoked, error) {
	event := new(UVXRoleRevoked)
	if err := _UVX.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UVXTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UVX contract.
type UVXTransferIterator struct {
	Event *UVXTransfer // Event containing the contract specifics and raw log

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
func (it *UVXTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UVXTransfer)
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
		it.Event = new(UVXTransfer)
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
func (it *UVXTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UVXTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UVXTransfer represents a Transfer event raised by the UVX contract.
type UVXTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UVX *UVXFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UVXTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UVX.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UVXTransferIterator{contract: _UVX.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UVX *UVXFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UVXTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UVX.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UVXTransfer)
				if err := _UVX.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_UVX *UVXFilterer) ParseTransfer(log types.Log) (*UVXTransfer, error) {
	event := new(UVXTransfer)
	if err := _UVX.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
