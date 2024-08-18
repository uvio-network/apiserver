// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package randomizercontract

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

// RandomizerMetaData contains all meta data concerning the Randomizer contract.
var RandomizerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_markets\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isUniqueVoter\",\"inputs\":[{\"name\":\"userClaimKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isUnique\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keepers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"m\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMarkets\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prepareVote\",\"inputs\":[{\"name\":\"_yeaVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_nayVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"_keeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVotersLimit\",\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerPrepareVote\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userClaimKey\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"votersLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ArrayLengthExceedsLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotKeeper\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotStaker\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotUnique\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// RandomizerABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomizerMetaData.ABI instead.
var RandomizerABI = RandomizerMetaData.ABI

// Randomizer is an auto generated Go binding around an Ethereum contract.
type Randomizer struct {
	RandomizerCaller     // Read-only binding to the contract
	RandomizerTransactor // Write-only binding to the contract
	RandomizerFilterer   // Log filterer for contract events
}

// RandomizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomizerSession struct {
	Contract     *Randomizer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RandomizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomizerCallerSession struct {
	Contract *RandomizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RandomizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomizerTransactorSession struct {
	Contract     *RandomizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RandomizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomizerRaw struct {
	Contract *Randomizer // Generic contract binding to access the raw methods on
}

// RandomizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomizerCallerRaw struct {
	Contract *RandomizerCaller // Generic read-only contract binding to access the raw methods on
}

// RandomizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomizerTransactorRaw struct {
	Contract *RandomizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomizer creates a new instance of Randomizer, bound to a specific deployed contract.
func NewRandomizer(address common.Address, backend bind.ContractBackend) (*Randomizer, error) {
	contract, err := bindRandomizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Randomizer{RandomizerCaller: RandomizerCaller{contract: contract}, RandomizerTransactor: RandomizerTransactor{contract: contract}, RandomizerFilterer: RandomizerFilterer{contract: contract}}, nil
}

// NewRandomizerCaller creates a new read-only instance of Randomizer, bound to a specific deployed contract.
func NewRandomizerCaller(address common.Address, caller bind.ContractCaller) (*RandomizerCaller, error) {
	contract, err := bindRandomizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomizerCaller{contract: contract}, nil
}

// NewRandomizerTransactor creates a new write-only instance of Randomizer, bound to a specific deployed contract.
func NewRandomizerTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomizerTransactor, error) {
	contract, err := bindRandomizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomizerTransactor{contract: contract}, nil
}

// NewRandomizerFilterer creates a new log filterer instance of Randomizer, bound to a specific deployed contract.
func NewRandomizerFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomizerFilterer, error) {
	contract, err := bindRandomizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomizerFilterer{contract: contract}, nil
}

// bindRandomizer binds a generic wrapper to an already deployed contract.
func bindRandomizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomizerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Randomizer *RandomizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Randomizer.Contract.RandomizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Randomizer *RandomizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Randomizer.Contract.RandomizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Randomizer *RandomizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Randomizer.Contract.RandomizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Randomizer *RandomizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Randomizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Randomizer *RandomizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Randomizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Randomizer *RandomizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Randomizer.Contract.contract.Transact(opts, method, params...)
}

// IsUniqueVoter is a free data retrieval call binding the contract method 0xa8cae7eb.
//
// Solidity: function isUniqueVoter(bytes32 userClaimKey) view returns(bool isUnique)
func (_Randomizer *RandomizerCaller) IsUniqueVoter(opts *bind.CallOpts, userClaimKey [32]byte) (bool, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "isUniqueVoter", userClaimKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUniqueVoter is a free data retrieval call binding the contract method 0xa8cae7eb.
//
// Solidity: function isUniqueVoter(bytes32 userClaimKey) view returns(bool isUnique)
func (_Randomizer *RandomizerSession) IsUniqueVoter(userClaimKey [32]byte) (bool, error) {
	return _Randomizer.Contract.IsUniqueVoter(&_Randomizer.CallOpts, userClaimKey)
}

// IsUniqueVoter is a free data retrieval call binding the contract method 0xa8cae7eb.
//
// Solidity: function isUniqueVoter(bytes32 userClaimKey) view returns(bool isUnique)
func (_Randomizer *RandomizerCallerSession) IsUniqueVoter(userClaimKey [32]byte) (bool, error) {
	return _Randomizer.Contract.IsUniqueVoter(&_Randomizer.CallOpts, userClaimKey)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Randomizer *RandomizerCaller) Keepers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "keepers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Randomizer *RandomizerSession) Keepers(arg0 common.Address) (bool, error) {
	return _Randomizer.Contract.Keepers(&_Randomizer.CallOpts, arg0)
}

// Keepers is a free data retrieval call binding the contract method 0x3bbd64bc.
//
// Solidity: function keepers(address ) view returns(bool)
func (_Randomizer *RandomizerCallerSession) Keepers(arg0 common.Address) (bool, error) {
	return _Randomizer.Contract.Keepers(&_Randomizer.CallOpts, arg0)
}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(address)
func (_Randomizer *RandomizerCaller) M(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "m")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(address)
func (_Randomizer *RandomizerSession) M() (common.Address, error) {
	return _Randomizer.Contract.M(&_Randomizer.CallOpts)
}

// M is a free data retrieval call binding the contract method 0x5a2ee019.
//
// Solidity: function m() view returns(address)
func (_Randomizer *RandomizerCallerSession) M() (common.Address, error) {
	return _Randomizer.Contract.M(&_Randomizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Randomizer *RandomizerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Randomizer *RandomizerSession) Owner() (common.Address, error) {
	return _Randomizer.Contract.Owner(&_Randomizer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Randomizer *RandomizerCallerSession) Owner() (common.Address, error) {
	return _Randomizer.Contract.Owner(&_Randomizer.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Randomizer *RandomizerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Randomizer *RandomizerSession) PendingOwner() (common.Address, error) {
	return _Randomizer.Contract.PendingOwner(&_Randomizer.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Randomizer *RandomizerCallerSession) PendingOwner() (common.Address, error) {
	return _Randomizer.Contract.PendingOwner(&_Randomizer.CallOpts)
}

// TriggerPrepareVote is a free data retrieval call binding the contract method 0xca398749.
//
// Solidity: function triggerPrepareVote(uint256 _marketId) view returns(bool)
func (_Randomizer *RandomizerCaller) TriggerPrepareVote(opts *bind.CallOpts, _marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "triggerPrepareVote", _marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TriggerPrepareVote is a free data retrieval call binding the contract method 0xca398749.
//
// Solidity: function triggerPrepareVote(uint256 _marketId) view returns(bool)
func (_Randomizer *RandomizerSession) TriggerPrepareVote(_marketId *big.Int) (bool, error) {
	return _Randomizer.Contract.TriggerPrepareVote(&_Randomizer.CallOpts, _marketId)
}

// TriggerPrepareVote is a free data retrieval call binding the contract method 0xca398749.
//
// Solidity: function triggerPrepareVote(uint256 _marketId) view returns(bool)
func (_Randomizer *RandomizerCallerSession) TriggerPrepareVote(_marketId *big.Int) (bool, error) {
	return _Randomizer.Contract.TriggerPrepareVote(&_Randomizer.CallOpts, _marketId)
}

// UserClaimKey is a free data retrieval call binding the contract method 0x7a70305c.
//
// Solidity: function userClaimKey(uint256 _marketId, uint256 _claimId, address _user) pure returns(bytes32)
func (_Randomizer *RandomizerCaller) UserClaimKey(opts *bind.CallOpts, _marketId *big.Int, _claimId *big.Int, _user common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "userClaimKey", _marketId, _claimId, _user)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserClaimKey is a free data retrieval call binding the contract method 0x7a70305c.
//
// Solidity: function userClaimKey(uint256 _marketId, uint256 _claimId, address _user) pure returns(bytes32)
func (_Randomizer *RandomizerSession) UserClaimKey(_marketId *big.Int, _claimId *big.Int, _user common.Address) ([32]byte, error) {
	return _Randomizer.Contract.UserClaimKey(&_Randomizer.CallOpts, _marketId, _claimId, _user)
}

// UserClaimKey is a free data retrieval call binding the contract method 0x7a70305c.
//
// Solidity: function userClaimKey(uint256 _marketId, uint256 _claimId, address _user) pure returns(bytes32)
func (_Randomizer *RandomizerCallerSession) UserClaimKey(_marketId *big.Int, _claimId *big.Int, _user common.Address) ([32]byte, error) {
	return _Randomizer.Contract.UserClaimKey(&_Randomizer.CallOpts, _marketId, _claimId, _user)
}

// VotersLimit is a free data retrieval call binding the contract method 0x544d6f6d.
//
// Solidity: function votersLimit() view returns(uint256)
func (_Randomizer *RandomizerCaller) VotersLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Randomizer.contract.Call(opts, &out, "votersLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotersLimit is a free data retrieval call binding the contract method 0x544d6f6d.
//
// Solidity: function votersLimit() view returns(uint256)
func (_Randomizer *RandomizerSession) VotersLimit() (*big.Int, error) {
	return _Randomizer.Contract.VotersLimit(&_Randomizer.CallOpts)
}

// VotersLimit is a free data retrieval call binding the contract method 0x544d6f6d.
//
// Solidity: function votersLimit() view returns(uint256)
func (_Randomizer *RandomizerCallerSession) VotersLimit() (*big.Int, error) {
	return _Randomizer.Contract.VotersLimit(&_Randomizer.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Randomizer *RandomizerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Randomizer *RandomizerSession) AcceptOwnership() (*types.Transaction, error) {
	return _Randomizer.Contract.AcceptOwnership(&_Randomizer.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Randomizer *RandomizerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Randomizer.Contract.AcceptOwnership(&_Randomizer.TransactOpts)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Randomizer *RandomizerTransactor) PrepareVote(opts *bind.TransactOpts, _yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "prepareVote", _yeaVoters, _nayVoters, _marketId)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Randomizer *RandomizerSession) PrepareVote(_yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Randomizer.Contract.PrepareVote(&_Randomizer.TransactOpts, _yeaVoters, _nayVoters, _marketId)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Randomizer *RandomizerTransactorSession) PrepareVote(_yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Randomizer.Contract.PrepareVote(&_Randomizer.TransactOpts, _yeaVoters, _nayVoters, _marketId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Randomizer *RandomizerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Randomizer *RandomizerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Randomizer.Contract.RenounceOwnership(&_Randomizer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Randomizer *RandomizerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Randomizer.Contract.RenounceOwnership(&_Randomizer.TransactOpts)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _status) returns()
func (_Randomizer *RandomizerTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address, _status bool) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "setKeeper", _keeper, _status)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _status) returns()
func (_Randomizer *RandomizerSession) SetKeeper(_keeper common.Address, _status bool) (*types.Transaction, error) {
	return _Randomizer.Contract.SetKeeper(&_Randomizer.TransactOpts, _keeper, _status)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _status) returns()
func (_Randomizer *RandomizerTransactorSession) SetKeeper(_keeper common.Address, _status bool) (*types.Transaction, error) {
	return _Randomizer.Contract.SetKeeper(&_Randomizer.TransactOpts, _keeper, _status)
}

// SetVotersLimit is a paid mutator transaction binding the contract method 0xdc115972.
//
// Solidity: function setVotersLimit(uint256 _limit) returns()
func (_Randomizer *RandomizerTransactor) SetVotersLimit(opts *bind.TransactOpts, _limit *big.Int) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "setVotersLimit", _limit)
}

// SetVotersLimit is a paid mutator transaction binding the contract method 0xdc115972.
//
// Solidity: function setVotersLimit(uint256 _limit) returns()
func (_Randomizer *RandomizerSession) SetVotersLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Randomizer.Contract.SetVotersLimit(&_Randomizer.TransactOpts, _limit)
}

// SetVotersLimit is a paid mutator transaction binding the contract method 0xdc115972.
//
// Solidity: function setVotersLimit(uint256 _limit) returns()
func (_Randomizer *RandomizerTransactorSession) SetVotersLimit(_limit *big.Int) (*types.Transaction, error) {
	return _Randomizer.Contract.SetVotersLimit(&_Randomizer.TransactOpts, _limit)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Randomizer *RandomizerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Randomizer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Randomizer *RandomizerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Randomizer.Contract.TransferOwnership(&_Randomizer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Randomizer *RandomizerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Randomizer.Contract.TransferOwnership(&_Randomizer.TransactOpts, newOwner)
}

// RandomizerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Randomizer contract.
type RandomizerOwnershipTransferStartedIterator struct {
	Event *RandomizerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RandomizerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomizerOwnershipTransferStarted)
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
		it.Event = new(RandomizerOwnershipTransferStarted)
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
func (it *RandomizerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomizerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomizerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Randomizer contract.
type RandomizerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RandomizerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Randomizer.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RandomizerOwnershipTransferStartedIterator{contract: _Randomizer.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RandomizerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Randomizer.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomizerOwnershipTransferStarted)
				if err := _Randomizer.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) ParseOwnershipTransferStarted(log types.Log) (*RandomizerOwnershipTransferStarted, error) {
	event := new(RandomizerOwnershipTransferStarted)
	if err := _Randomizer.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomizerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Randomizer contract.
type RandomizerOwnershipTransferredIterator struct {
	Event *RandomizerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RandomizerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomizerOwnershipTransferred)
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
		it.Event = new(RandomizerOwnershipTransferred)
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
func (it *RandomizerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomizerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomizerOwnershipTransferred represents a OwnershipTransferred event raised by the Randomizer contract.
type RandomizerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RandomizerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Randomizer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RandomizerOwnershipTransferredIterator{contract: _Randomizer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RandomizerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Randomizer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomizerOwnershipTransferred)
				if err := _Randomizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Randomizer *RandomizerFilterer) ParseOwnershipTransferred(log types.Log) (*RandomizerOwnershipTransferred, error) {
	event := new(RandomizerOwnershipTransferred)
	if err := _Randomizer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
