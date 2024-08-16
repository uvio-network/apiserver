// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketscontract

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

// IMarketsClaim is an auto generated low-level Go binding around an user-defined struct.
type IMarketsClaim struct {
	MetadataURI     string
	Expiration      *big.Int
	Proposer        common.Address
	Asset           common.Address
	IsNullifyMarket bool
	Stake           IMarketsStake
	Vote            IMarketsVote
	Status          uint8
}

// IMarketsPrice is an auto generated low-level Go binding around an user-defined struct.
type IMarketsPrice struct {
	CurveType uint8
	Steepness *big.Int
}

// IMarketsPropose is an auto generated low-level Go binding around an user-defined struct.
type IMarketsPropose struct {
	MetadataURI       string
	MarketId          *big.Int
	NullifyMarketId   *big.Int
	Amount            *big.Int
	Asset             common.Address
	ClaimExpiration   *big.Int
	StakingExpiration *big.Int
	Yea               bool
	Dispute           bool
	Price             IMarketsPrice
}

// IMarketsStake is an auto generated low-level Go binding around an user-defined struct.
type IMarketsStake struct {
	Yea        *big.Int
	Nay        *big.Int
	Expiration *big.Int
	Start      *big.Int
	YeaStakers []common.Address
	NayStakers []common.Address
	Price      IMarketsPrice
}

// IMarketsVote is an auto generated low-level Go binding around an user-defined struct.
type IMarketsVote struct {
	Yea               *big.Int
	Nay               *big.Int
	Expiration        *big.Int
	DisputeExpiration *big.Int
	YeaVoters         []common.Address
	NayVoters         []common.Address
	Outcome           uint8
}

// MarketsMetaData contains all meta data concerning the Markets contract.
var MarketsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_settings\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_CLAIMS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimKey\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"claimProceeds\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_claimId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claims\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[4]\",\"internalType\":\"structIMarkets.Claim[4]\",\"components\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isNullifyMarket\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Stake\",\"components\":[{\"name\":\"yea\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"start\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"yeaStakers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"nayStakers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Price\",\"components\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"steepness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"vote\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Vote\",\"components\":[{\"name\":\"yea\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nay\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"disputeExpiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"yeaVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"nayVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"outcome\",\"type\":\"uint8\",\"internalType\":\"enumIMarkets.Outcome\"}]},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIMarkets.ClaimStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimsLength\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"committeeResolve\",\"inputs\":[{\"name\":\"_outcome\",\"type\":\"uint8\",\"internalType\":\"enumIMarkets.Outcome\"},{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_reciever\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endVote\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"marketId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marketMinStake\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nullifiedMarkets\",\"inputs\":[{\"name\":\"targetMarketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"nullifyMarketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prepareVote\",\"inputs\":[{\"name\":\"_yeaVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_nayVoters\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"propose\",\"inputs\":[{\"name\":\"_propose\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Propose\",\"components\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nullifyMarketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimExpiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"stakingExpiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"yea\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dispute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Price\",\"components\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"steepness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISettings\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_yea\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"_timeWeightedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userBalance\",\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userClaimed\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userStake\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userStakeStatus\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarkets.VoteStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userVoteStatus\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_userVoteKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIMarkets.VoteStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"_marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_yea\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_stake\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_reciever\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClaimProceeds\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proceeds\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"earned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CommitteeResolve\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reciever\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EndVote\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeToProposer\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrepareVote\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Proposed\",\"inputs\":[{\"name\":\"propose\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIMarkets.Propose\",\"components\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nullifyMarketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimExpiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"stakingExpiration\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"yea\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"dispute\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"tuple\",\"internalType\":\"structIMarkets.Price\",\"components\":[{\"name\":\"curveType\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"steepness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"claimId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resolve\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Staked\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timeWeightedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"yea\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Voted\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"claimId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"yea\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"stake\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reciever\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AssetNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpiration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMarketId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNullifyMarketId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNullifyMarketParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPriceParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxClaimsReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPendingCommitteeResolution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NullifyMarketNotResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UserNotWhitelisted\",\"inputs\":[]}]",
}

// MarketsABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketsMetaData.ABI instead.
var MarketsABI = MarketsMetaData.ABI

// Markets is an auto generated Go binding around an Ethereum contract.
type Markets struct {
	MarketsCaller     // Read-only binding to the contract
	MarketsTransactor // Write-only binding to the contract
	MarketsFilterer   // Log filterer for contract events
}

// MarketsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketsSession struct {
	Contract     *Markets          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketsCallerSession struct {
	Contract *MarketsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MarketsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketsTransactorSession struct {
	Contract     *MarketsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MarketsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketsRaw struct {
	Contract *Markets // Generic contract binding to access the raw methods on
}

// MarketsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketsCallerRaw struct {
	Contract *MarketsCaller // Generic read-only contract binding to access the raw methods on
}

// MarketsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketsTransactorRaw struct {
	Contract *MarketsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarkets creates a new instance of Markets, bound to a specific deployed contract.
func NewMarkets(address common.Address, backend bind.ContractBackend) (*Markets, error) {
	contract, err := bindMarkets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Markets{MarketsCaller: MarketsCaller{contract: contract}, MarketsTransactor: MarketsTransactor{contract: contract}, MarketsFilterer: MarketsFilterer{contract: contract}}, nil
}

// NewMarketsCaller creates a new read-only instance of Markets, bound to a specific deployed contract.
func NewMarketsCaller(address common.Address, caller bind.ContractCaller) (*MarketsCaller, error) {
	contract, err := bindMarkets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketsCaller{contract: contract}, nil
}

// NewMarketsTransactor creates a new write-only instance of Markets, bound to a specific deployed contract.
func NewMarketsTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketsTransactor, error) {
	contract, err := bindMarkets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketsTransactor{contract: contract}, nil
}

// NewMarketsFilterer creates a new log filterer instance of Markets, bound to a specific deployed contract.
func NewMarketsFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketsFilterer, error) {
	contract, err := bindMarkets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketsFilterer{contract: contract}, nil
}

// bindMarkets binds a generic wrapper to an already deployed contract.
func bindMarkets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Markets *MarketsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Markets.Contract.MarketsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Markets *MarketsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Markets.Contract.MarketsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Markets *MarketsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Markets.Contract.MarketsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Markets *MarketsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Markets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Markets *MarketsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Markets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Markets *MarketsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Markets.Contract.contract.Transact(opts, method, params...)
}

// MAXCLAIMS is a free data retrieval call binding the contract method 0xcf62ba92.
//
// Solidity: function MAX_CLAIMS() view returns(uint256)
func (_Markets *MarketsCaller) MAXCLAIMS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "MAX_CLAIMS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCLAIMS is a free data retrieval call binding the contract method 0xcf62ba92.
//
// Solidity: function MAX_CLAIMS() view returns(uint256)
func (_Markets *MarketsSession) MAXCLAIMS() (*big.Int, error) {
	return _Markets.Contract.MAXCLAIMS(&_Markets.CallOpts)
}

// MAXCLAIMS is a free data retrieval call binding the contract method 0xcf62ba92.
//
// Solidity: function MAX_CLAIMS() view returns(uint256)
func (_Markets *MarketsCallerSession) MAXCLAIMS() (*big.Int, error) {
	return _Markets.Contract.MAXCLAIMS(&_Markets.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Markets *MarketsCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Markets *MarketsSession) PRECISION() (*big.Int, error) {
	return _Markets.Contract.PRECISION(&_Markets.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Markets *MarketsCallerSession) PRECISION() (*big.Int, error) {
	return _Markets.Contract.PRECISION(&_Markets.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Markets *MarketsCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Markets *MarketsSession) PRICEPRECISION() (*big.Int, error) {
	return _Markets.Contract.PRICEPRECISION(&_Markets.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Markets *MarketsCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Markets.Contract.PRICEPRECISION(&_Markets.CallOpts)
}

// ClaimKey is a free data retrieval call binding the contract method 0xaa43319a.
//
// Solidity: function claimKey(uint256 _marketId, uint256 _claimId) pure returns(bytes32)
func (_Markets *MarketsCaller) ClaimKey(opts *bind.CallOpts, _marketId *big.Int, _claimId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "claimKey", _marketId, _claimId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimKey is a free data retrieval call binding the contract method 0xaa43319a.
//
// Solidity: function claimKey(uint256 _marketId, uint256 _claimId) pure returns(bytes32)
func (_Markets *MarketsSession) ClaimKey(_marketId *big.Int, _claimId *big.Int) ([32]byte, error) {
	return _Markets.Contract.ClaimKey(&_Markets.CallOpts, _marketId, _claimId)
}

// ClaimKey is a free data retrieval call binding the contract method 0xaa43319a.
//
// Solidity: function claimKey(uint256 _marketId, uint256 _claimId) pure returns(bytes32)
func (_Markets *MarketsCallerSession) ClaimKey(_marketId *big.Int, _claimId *big.Int) ([32]byte, error) {
	return _Markets.Contract.ClaimKey(&_Markets.CallOpts, _marketId, _claimId)
}

// Claims is a free data retrieval call binding the contract method 0xa888c2cd.
//
// Solidity: function claims(uint256 _marketId) view returns((string,uint40,address,address,bool,(uint256,uint256,uint40,uint40,address[],address[],(uint8,uint256)),(uint256,uint256,uint40,uint40,address[],address[],uint8),uint8)[4])
func (_Markets *MarketsCaller) Claims(opts *bind.CallOpts, _marketId *big.Int) ([4]IMarketsClaim, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "claims", _marketId)

	if err != nil {
		return *new([4]IMarketsClaim), err
	}

	out0 := *abi.ConvertType(out[0], new([4]IMarketsClaim)).(*[4]IMarketsClaim)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xa888c2cd.
//
// Solidity: function claims(uint256 _marketId) view returns((string,uint40,address,address,bool,(uint256,uint256,uint40,uint40,address[],address[],(uint8,uint256)),(uint256,uint256,uint40,uint40,address[],address[],uint8),uint8)[4])
func (_Markets *MarketsSession) Claims(_marketId *big.Int) ([4]IMarketsClaim, error) {
	return _Markets.Contract.Claims(&_Markets.CallOpts, _marketId)
}

// Claims is a free data retrieval call binding the contract method 0xa888c2cd.
//
// Solidity: function claims(uint256 _marketId) view returns((string,uint40,address,address,bool,(uint256,uint256,uint40,uint40,address[],address[],(uint8,uint256)),(uint256,uint256,uint40,uint40,address[],address[],uint8),uint8)[4])
func (_Markets *MarketsCallerSession) Claims(_marketId *big.Int) ([4]IMarketsClaim, error) {
	return _Markets.Contract.Claims(&_Markets.CallOpts, _marketId)
}

// ClaimsLength is a free data retrieval call binding the contract method 0x21c0fba7.
//
// Solidity: function claimsLength(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsCaller) ClaimsLength(opts *bind.CallOpts, _marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "claimsLength", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimsLength is a free data retrieval call binding the contract method 0x21c0fba7.
//
// Solidity: function claimsLength(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsSession) ClaimsLength(_marketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.ClaimsLength(&_Markets.CallOpts, _marketId)
}

// ClaimsLength is a free data retrieval call binding the contract method 0x21c0fba7.
//
// Solidity: function claimsLength(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsCallerSession) ClaimsLength(_marketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.ClaimsLength(&_Markets.CallOpts, _marketId)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_Markets *MarketsCaller) MarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "marketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_Markets *MarketsSession) MarketId() (*big.Int, error) {
	return _Markets.Contract.MarketId(&_Markets.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_Markets *MarketsCallerSession) MarketId() (*big.Int, error) {
	return _Markets.Contract.MarketId(&_Markets.CallOpts)
}

// MarketMinStake is a free data retrieval call binding the contract method 0x487cd1b2.
//
// Solidity: function marketMinStake(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsCaller) MarketMinStake(opts *bind.CallOpts, _marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "marketMinStake", _marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketMinStake is a free data retrieval call binding the contract method 0x487cd1b2.
//
// Solidity: function marketMinStake(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsSession) MarketMinStake(_marketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.MarketMinStake(&_Markets.CallOpts, _marketId)
}

// MarketMinStake is a free data retrieval call binding the contract method 0x487cd1b2.
//
// Solidity: function marketMinStake(uint256 _marketId) view returns(uint256)
func (_Markets *MarketsCallerSession) MarketMinStake(_marketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.MarketMinStake(&_Markets.CallOpts, _marketId)
}

// NullifiedMarkets is a free data retrieval call binding the contract method 0x3938b126.
//
// Solidity: function nullifiedMarkets(uint256 targetMarketId) view returns(uint256 nullifyMarketId)
func (_Markets *MarketsCaller) NullifiedMarkets(opts *bind.CallOpts, targetMarketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "nullifiedMarkets", targetMarketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NullifiedMarkets is a free data retrieval call binding the contract method 0x3938b126.
//
// Solidity: function nullifiedMarkets(uint256 targetMarketId) view returns(uint256 nullifyMarketId)
func (_Markets *MarketsSession) NullifiedMarkets(targetMarketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.NullifiedMarkets(&_Markets.CallOpts, targetMarketId)
}

// NullifiedMarkets is a free data retrieval call binding the contract method 0x3938b126.
//
// Solidity: function nullifiedMarkets(uint256 targetMarketId) view returns(uint256 nullifyMarketId)
func (_Markets *MarketsCallerSession) NullifiedMarkets(targetMarketId *big.Int) (*big.Int, error) {
	return _Markets.Contract.NullifiedMarkets(&_Markets.CallOpts, targetMarketId)
}

// S is a free data retrieval call binding the contract method 0x86b714e2.
//
// Solidity: function s() view returns(address)
func (_Markets *MarketsCaller) S(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "s")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// S is a free data retrieval call binding the contract method 0x86b714e2.
//
// Solidity: function s() view returns(address)
func (_Markets *MarketsSession) S() (common.Address, error) {
	return _Markets.Contract.S(&_Markets.CallOpts)
}

// S is a free data retrieval call binding the contract method 0x86b714e2.
//
// Solidity: function s() view returns(address)
func (_Markets *MarketsCallerSession) S() (common.Address, error) {
	return _Markets.Contract.S(&_Markets.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x62a4b0a9.
//
// Solidity: function userBalance(address _asset, address _user) view returns(uint256)
func (_Markets *MarketsCaller) UserBalance(opts *bind.CallOpts, _asset common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "userBalance", _asset, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x62a4b0a9.
//
// Solidity: function userBalance(address _asset, address _user) view returns(uint256)
func (_Markets *MarketsSession) UserBalance(_asset common.Address, _user common.Address) (*big.Int, error) {
	return _Markets.Contract.UserBalance(&_Markets.CallOpts, _asset, _user)
}

// UserBalance is a free data retrieval call binding the contract method 0x62a4b0a9.
//
// Solidity: function userBalance(address _asset, address _user) view returns(uint256)
func (_Markets *MarketsCallerSession) UserBalance(_asset common.Address, _user common.Address) (*big.Int, error) {
	return _Markets.Contract.UserBalance(&_Markets.CallOpts, _asset, _user)
}

// UserClaimed is a free data retrieval call binding the contract method 0x63a554fa.
//
// Solidity: function userClaimed(address _user, bytes32 _claimKey) view returns(bool)
func (_Markets *MarketsCaller) UserClaimed(opts *bind.CallOpts, _user common.Address, _claimKey [32]byte) (bool, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "userClaimed", _user, _claimKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserClaimed is a free data retrieval call binding the contract method 0x63a554fa.
//
// Solidity: function userClaimed(address _user, bytes32 _claimKey) view returns(bool)
func (_Markets *MarketsSession) UserClaimed(_user common.Address, _claimKey [32]byte) (bool, error) {
	return _Markets.Contract.UserClaimed(&_Markets.CallOpts, _user, _claimKey)
}

// UserClaimed is a free data retrieval call binding the contract method 0x63a554fa.
//
// Solidity: function userClaimed(address _user, bytes32 _claimKey) view returns(bool)
func (_Markets *MarketsCallerSession) UserClaimed(_user common.Address, _claimKey [32]byte) (bool, error) {
	return _Markets.Contract.UserClaimed(&_Markets.CallOpts, _user, _claimKey)
}

// UserStake is a free data retrieval call binding the contract method 0x7f727479.
//
// Solidity: function userStake(address _user, bytes32 _claimKey) view returns(uint256)
func (_Markets *MarketsCaller) UserStake(opts *bind.CallOpts, _user common.Address, _claimKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "userStake", _user, _claimKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStake is a free data retrieval call binding the contract method 0x7f727479.
//
// Solidity: function userStake(address _user, bytes32 _claimKey) view returns(uint256)
func (_Markets *MarketsSession) UserStake(_user common.Address, _claimKey [32]byte) (*big.Int, error) {
	return _Markets.Contract.UserStake(&_Markets.CallOpts, _user, _claimKey)
}

// UserStake is a free data retrieval call binding the contract method 0x7f727479.
//
// Solidity: function userStake(address _user, bytes32 _claimKey) view returns(uint256)
func (_Markets *MarketsCallerSession) UserStake(_user common.Address, _claimKey [32]byte) (*big.Int, error) {
	return _Markets.Contract.UserStake(&_Markets.CallOpts, _user, _claimKey)
}

// UserStakeStatus is a free data retrieval call binding the contract method 0x41d643ee.
//
// Solidity: function userStakeStatus(address _user, bytes32 _claimKey) view returns(uint8)
func (_Markets *MarketsCaller) UserStakeStatus(opts *bind.CallOpts, _user common.Address, _claimKey [32]byte) (uint8, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "userStakeStatus", _user, _claimKey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UserStakeStatus is a free data retrieval call binding the contract method 0x41d643ee.
//
// Solidity: function userStakeStatus(address _user, bytes32 _claimKey) view returns(uint8)
func (_Markets *MarketsSession) UserStakeStatus(_user common.Address, _claimKey [32]byte) (uint8, error) {
	return _Markets.Contract.UserStakeStatus(&_Markets.CallOpts, _user, _claimKey)
}

// UserStakeStatus is a free data retrieval call binding the contract method 0x41d643ee.
//
// Solidity: function userStakeStatus(address _user, bytes32 _claimKey) view returns(uint8)
func (_Markets *MarketsCallerSession) UserStakeStatus(_user common.Address, _claimKey [32]byte) (uint8, error) {
	return _Markets.Contract.UserStakeStatus(&_Markets.CallOpts, _user, _claimKey)
}

// UserVoteStatus is a free data retrieval call binding the contract method 0x03980ff2.
//
// Solidity: function userVoteStatus(address _user, bytes32 _userVoteKey) view returns(uint8)
func (_Markets *MarketsCaller) UserVoteStatus(opts *bind.CallOpts, _user common.Address, _userVoteKey [32]byte) (uint8, error) {
	var out []interface{}
	err := _Markets.contract.Call(opts, &out, "userVoteStatus", _user, _userVoteKey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UserVoteStatus is a free data retrieval call binding the contract method 0x03980ff2.
//
// Solidity: function userVoteStatus(address _user, bytes32 _userVoteKey) view returns(uint8)
func (_Markets *MarketsSession) UserVoteStatus(_user common.Address, _userVoteKey [32]byte) (uint8, error) {
	return _Markets.Contract.UserVoteStatus(&_Markets.CallOpts, _user, _userVoteKey)
}

// UserVoteStatus is a free data retrieval call binding the contract method 0x03980ff2.
//
// Solidity: function userVoteStatus(address _user, bytes32 _userVoteKey) view returns(uint8)
func (_Markets *MarketsCallerSession) UserVoteStatus(_user common.Address, _userVoteKey [32]byte) (uint8, error) {
	return _Markets.Contract.UserVoteStatus(&_Markets.CallOpts, _user, _userVoteKey)
}

// ClaimProceeds is a paid mutator transaction binding the contract method 0x21722b0f.
//
// Solidity: function claimProceeds(uint256 _marketId, uint256 _claimId, address _user) returns()
func (_Markets *MarketsTransactor) ClaimProceeds(opts *bind.TransactOpts, _marketId *big.Int, _claimId *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "claimProceeds", _marketId, _claimId, _user)
}

// ClaimProceeds is a paid mutator transaction binding the contract method 0x21722b0f.
//
// Solidity: function claimProceeds(uint256 _marketId, uint256 _claimId, address _user) returns()
func (_Markets *MarketsSession) ClaimProceeds(_marketId *big.Int, _claimId *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Markets.Contract.ClaimProceeds(&_Markets.TransactOpts, _marketId, _claimId, _user)
}

// ClaimProceeds is a paid mutator transaction binding the contract method 0x21722b0f.
//
// Solidity: function claimProceeds(uint256 _marketId, uint256 _claimId, address _user) returns()
func (_Markets *MarketsTransactorSession) ClaimProceeds(_marketId *big.Int, _claimId *big.Int, _user common.Address) (*types.Transaction, error) {
	return _Markets.Contract.ClaimProceeds(&_Markets.TransactOpts, _marketId, _claimId, _user)
}

// CommitteeResolve is a paid mutator transaction binding the contract method 0xdb6de3e0.
//
// Solidity: function committeeResolve(uint8 _outcome, uint256 _marketId) returns()
func (_Markets *MarketsTransactor) CommitteeResolve(opts *bind.TransactOpts, _outcome uint8, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "committeeResolve", _outcome, _marketId)
}

// CommitteeResolve is a paid mutator transaction binding the contract method 0xdb6de3e0.
//
// Solidity: function committeeResolve(uint8 _outcome, uint256 _marketId) returns()
func (_Markets *MarketsSession) CommitteeResolve(_outcome uint8, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.CommitteeResolve(&_Markets.TransactOpts, _outcome, _marketId)
}

// CommitteeResolve is a paid mutator transaction binding the contract method 0xdb6de3e0.
//
// Solidity: function committeeResolve(uint8 _outcome, uint256 _marketId) returns()
func (_Markets *MarketsTransactorSession) CommitteeResolve(_outcome uint8, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.CommitteeResolve(&_Markets.TransactOpts, _outcome, _marketId)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "deposit", _amount, _asset, _reciever)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsSession) Deposit(_amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.Contract.Deposit(&_Markets.TransactOpts, _amount, _asset, _reciever)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsTransactorSession) Deposit(_amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.Contract.Deposit(&_Markets.TransactOpts, _amount, _asset, _reciever)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _marketId) returns()
func (_Markets *MarketsTransactor) EndVote(opts *bind.TransactOpts, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "endVote", _marketId)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _marketId) returns()
func (_Markets *MarketsSession) EndVote(_marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.EndVote(&_Markets.TransactOpts, _marketId)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _marketId) returns()
func (_Markets *MarketsTransactorSession) EndVote(_marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.EndVote(&_Markets.TransactOpts, _marketId)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Markets *MarketsTransactor) PrepareVote(opts *bind.TransactOpts, _yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "prepareVote", _yeaVoters, _nayVoters, _marketId)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Markets *MarketsSession) PrepareVote(_yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.PrepareVote(&_Markets.TransactOpts, _yeaVoters, _nayVoters, _marketId)
}

// PrepareVote is a paid mutator transaction binding the contract method 0xdd548d0e.
//
// Solidity: function prepareVote(address[] _yeaVoters, address[] _nayVoters, uint256 _marketId) returns()
func (_Markets *MarketsTransactorSession) PrepareVote(_yeaVoters []common.Address, _nayVoters []common.Address, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.PrepareVote(&_Markets.TransactOpts, _yeaVoters, _nayVoters, _marketId)
}

// Propose is a paid mutator transaction binding the contract method 0x5e630eef.
//
// Solidity: function propose((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) _propose) returns(uint256, uint256)
func (_Markets *MarketsTransactor) Propose(opts *bind.TransactOpts, _propose IMarketsPropose) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "propose", _propose)
}

// Propose is a paid mutator transaction binding the contract method 0x5e630eef.
//
// Solidity: function propose((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) _propose) returns(uint256, uint256)
func (_Markets *MarketsSession) Propose(_propose IMarketsPropose) (*types.Transaction, error) {
	return _Markets.Contract.Propose(&_Markets.TransactOpts, _propose)
}

// Propose is a paid mutator transaction binding the contract method 0x5e630eef.
//
// Solidity: function propose((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) _propose) returns(uint256, uint256)
func (_Markets *MarketsTransactorSession) Propose(_propose IMarketsPropose) (*types.Transaction, error) {
	return _Markets.Contract.Propose(&_Markets.TransactOpts, _propose)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 _marketId) returns()
func (_Markets *MarketsTransactor) Resolve(opts *bind.TransactOpts, _marketId *big.Int) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "resolve", _marketId)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 _marketId) returns()
func (_Markets *MarketsSession) Resolve(_marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.Resolve(&_Markets.TransactOpts, _marketId)
}

// Resolve is a paid mutator transaction binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 _marketId) returns()
func (_Markets *MarketsTransactorSession) Resolve(_marketId *big.Int) (*types.Transaction, error) {
	return _Markets.Contract.Resolve(&_Markets.TransactOpts, _marketId)
}

// Stake is a paid mutator transaction binding the contract method 0x561b2463.
//
// Solidity: function stake(uint256 _marketId, uint256 _amount, bool _yea) returns(uint256 _timeWeightedAmount)
func (_Markets *MarketsTransactor) Stake(opts *bind.TransactOpts, _marketId *big.Int, _amount *big.Int, _yea bool) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "stake", _marketId, _amount, _yea)
}

// Stake is a paid mutator transaction binding the contract method 0x561b2463.
//
// Solidity: function stake(uint256 _marketId, uint256 _amount, bool _yea) returns(uint256 _timeWeightedAmount)
func (_Markets *MarketsSession) Stake(_marketId *big.Int, _amount *big.Int, _yea bool) (*types.Transaction, error) {
	return _Markets.Contract.Stake(&_Markets.TransactOpts, _marketId, _amount, _yea)
}

// Stake is a paid mutator transaction binding the contract method 0x561b2463.
//
// Solidity: function stake(uint256 _marketId, uint256 _amount, bool _yea) returns(uint256 _timeWeightedAmount)
func (_Markets *MarketsTransactorSession) Stake(_marketId *big.Int, _amount *big.Int, _yea bool) (*types.Transaction, error) {
	return _Markets.Contract.Stake(&_Markets.TransactOpts, _marketId, _amount, _yea)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _marketId, bool _yea, bool _stake) returns()
func (_Markets *MarketsTransactor) Vote(opts *bind.TransactOpts, _marketId *big.Int, _yea bool, _stake bool) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "vote", _marketId, _yea, _stake)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _marketId, bool _yea, bool _stake) returns()
func (_Markets *MarketsSession) Vote(_marketId *big.Int, _yea bool, _stake bool) (*types.Transaction, error) {
	return _Markets.Contract.Vote(&_Markets.TransactOpts, _marketId, _yea, _stake)
}

// Vote is a paid mutator transaction binding the contract method 0xdf133bca.
//
// Solidity: function vote(uint256 _marketId, bool _yea, bool _stake) returns()
func (_Markets *MarketsTransactorSession) Vote(_marketId *big.Int, _yea bool, _stake bool) (*types.Transaction, error) {
	return _Markets.Contract.Vote(&_Markets.TransactOpts, _marketId, _yea, _stake)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.contract.Transact(opts, "withdraw", _amount, _asset, _reciever)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsSession) Withdraw(_amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.Contract.Withdraw(&_Markets.TransactOpts, _amount, _asset, _reciever)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _amount, address _asset, address _reciever) returns()
func (_Markets *MarketsTransactorSession) Withdraw(_amount *big.Int, _asset common.Address, _reciever common.Address) (*types.Transaction, error) {
	return _Markets.Contract.Withdraw(&_Markets.TransactOpts, _amount, _asset, _reciever)
}

// MarketsClaimProceedsIterator is returned from FilterClaimProceeds and is used to iterate over the raw logs and unpacked data for ClaimProceeds events raised by the Markets contract.
type MarketsClaimProceedsIterator struct {
	Event *MarketsClaimProceeds // Event containing the contract specifics and raw log

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
func (it *MarketsClaimProceedsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsClaimProceeds)
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
		it.Event = new(MarketsClaimProceeds)
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
func (it *MarketsClaimProceedsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsClaimProceedsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsClaimProceeds represents a ClaimProceeds event raised by the Markets contract.
type MarketsClaimProceeds struct {
	User     common.Address
	Proceeds *big.Int
	Earned   *big.Int
	MarketId *big.Int
	ClaimId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimProceeds is a free log retrieval operation binding the contract event 0x4dd1b86d42598980d78cb9407d605fecab0dc71e3a80e7e70edae08076f906ad.
//
// Solidity: event ClaimProceeds(address indexed user, uint256 proceeds, uint256 earned, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) FilterClaimProceeds(opts *bind.FilterOpts, user []common.Address) (*MarketsClaimProceedsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Markets.contract.FilterLogs(opts, "ClaimProceeds", userRule)
	if err != nil {
		return nil, err
	}
	return &MarketsClaimProceedsIterator{contract: _Markets.contract, event: "ClaimProceeds", logs: logs, sub: sub}, nil
}

// WatchClaimProceeds is a free log subscription operation binding the contract event 0x4dd1b86d42598980d78cb9407d605fecab0dc71e3a80e7e70edae08076f906ad.
//
// Solidity: event ClaimProceeds(address indexed user, uint256 proceeds, uint256 earned, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) WatchClaimProceeds(opts *bind.WatchOpts, sink chan<- *MarketsClaimProceeds, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Markets.contract.WatchLogs(opts, "ClaimProceeds", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsClaimProceeds)
				if err := _Markets.contract.UnpackLog(event, "ClaimProceeds", log); err != nil {
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

// ParseClaimProceeds is a log parse operation binding the contract event 0x4dd1b86d42598980d78cb9407d605fecab0dc71e3a80e7e70edae08076f906ad.
//
// Solidity: event ClaimProceeds(address indexed user, uint256 proceeds, uint256 earned, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) ParseClaimProceeds(log types.Log) (*MarketsClaimProceeds, error) {
	event := new(MarketsClaimProceeds)
	if err := _Markets.contract.UnpackLog(event, "ClaimProceeds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsCommitteeResolveIterator is returned from FilterCommitteeResolve and is used to iterate over the raw logs and unpacked data for CommitteeResolve events raised by the Markets contract.
type MarketsCommitteeResolveIterator struct {
	Event *MarketsCommitteeResolve // Event containing the contract specifics and raw log

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
func (it *MarketsCommitteeResolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsCommitteeResolve)
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
		it.Event = new(MarketsCommitteeResolve)
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
func (it *MarketsCommitteeResolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsCommitteeResolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsCommitteeResolve represents a CommitteeResolve event raised by the Markets contract.
type MarketsCommitteeResolve struct {
	MarketId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitteeResolve is a free log retrieval operation binding the contract event 0xe5bd8a76ac5f75b9f37c3f9eda2b98c9a8076f24030aa35d32c2e7b1e9f00223.
//
// Solidity: event CommitteeResolve(uint256 marketId)
func (_Markets *MarketsFilterer) FilterCommitteeResolve(opts *bind.FilterOpts) (*MarketsCommitteeResolveIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "CommitteeResolve")
	if err != nil {
		return nil, err
	}
	return &MarketsCommitteeResolveIterator{contract: _Markets.contract, event: "CommitteeResolve", logs: logs, sub: sub}, nil
}

// WatchCommitteeResolve is a free log subscription operation binding the contract event 0xe5bd8a76ac5f75b9f37c3f9eda2b98c9a8076f24030aa35d32c2e7b1e9f00223.
//
// Solidity: event CommitteeResolve(uint256 marketId)
func (_Markets *MarketsFilterer) WatchCommitteeResolve(opts *bind.WatchOpts, sink chan<- *MarketsCommitteeResolve) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "CommitteeResolve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsCommitteeResolve)
				if err := _Markets.contract.UnpackLog(event, "CommitteeResolve", log); err != nil {
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

// ParseCommitteeResolve is a log parse operation binding the contract event 0xe5bd8a76ac5f75b9f37c3f9eda2b98c9a8076f24030aa35d32c2e7b1e9f00223.
//
// Solidity: event CommitteeResolve(uint256 marketId)
func (_Markets *MarketsFilterer) ParseCommitteeResolve(log types.Log) (*MarketsCommitteeResolve, error) {
	event := new(MarketsCommitteeResolve)
	if err := _Markets.contract.UnpackLog(event, "CommitteeResolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Markets contract.
type MarketsDepositIterator struct {
	Event *MarketsDeposit // Event containing the contract specifics and raw log

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
func (it *MarketsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsDeposit)
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
		it.Event = new(MarketsDeposit)
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
func (it *MarketsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsDeposit represents a Deposit event raised by the Markets contract.
type MarketsDeposit struct {
	Sender   common.Address
	Reciever common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, reciever []common.Address) (*MarketsDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Deposit", senderRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &MarketsDepositIterator{contract: _Markets.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MarketsDeposit, sender []common.Address, reciever []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Deposit", senderRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsDeposit)
				if err := _Markets.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) ParseDeposit(log types.Log) (*MarketsDeposit, error) {
	event := new(MarketsDeposit)
	if err := _Markets.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsEndVoteIterator is returned from FilterEndVote and is used to iterate over the raw logs and unpacked data for EndVote events raised by the Markets contract.
type MarketsEndVoteIterator struct {
	Event *MarketsEndVote // Event containing the contract specifics and raw log

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
func (it *MarketsEndVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsEndVote)
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
		it.Event = new(MarketsEndVote)
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
func (it *MarketsEndVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsEndVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsEndVote represents a EndVote event raised by the Markets contract.
type MarketsEndVote struct {
	MarketId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEndVote is a free log retrieval operation binding the contract event 0xfce20a510f6f83f7014a422dfa9541af4170f706faeed678bf815b2bad2c3b9d.
//
// Solidity: event EndVote(uint256 marketId)
func (_Markets *MarketsFilterer) FilterEndVote(opts *bind.FilterOpts) (*MarketsEndVoteIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "EndVote")
	if err != nil {
		return nil, err
	}
	return &MarketsEndVoteIterator{contract: _Markets.contract, event: "EndVote", logs: logs, sub: sub}, nil
}

// WatchEndVote is a free log subscription operation binding the contract event 0xfce20a510f6f83f7014a422dfa9541af4170f706faeed678bf815b2bad2c3b9d.
//
// Solidity: event EndVote(uint256 marketId)
func (_Markets *MarketsFilterer) WatchEndVote(opts *bind.WatchOpts, sink chan<- *MarketsEndVote) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "EndVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsEndVote)
				if err := _Markets.contract.UnpackLog(event, "EndVote", log); err != nil {
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

// ParseEndVote is a log parse operation binding the contract event 0xfce20a510f6f83f7014a422dfa9541af4170f706faeed678bf815b2bad2c3b9d.
//
// Solidity: event EndVote(uint256 marketId)
func (_Markets *MarketsFilterer) ParseEndVote(log types.Log) (*MarketsEndVote, error) {
	event := new(MarketsEndVote)
	if err := _Markets.contract.UnpackLog(event, "EndVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsFeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the Markets contract.
type MarketsFeeIterator struct {
	Event *MarketsFee // Event containing the contract specifics and raw log

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
func (it *MarketsFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsFee)
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
		it.Event = new(MarketsFee)
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
func (it *MarketsFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsFee represents a Fee event raised by the Markets contract.
type MarketsFee struct {
	Fee           *big.Int
	FeeToProposer *big.Int
	MarketId      *big.Int
	ClaimId       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0x48b09ff26de49c6dfb086a28888f39d8954ae734fa818411eff3051bc315783a.
//
// Solidity: event Fee(uint256 fee, uint256 feeToProposer, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) FilterFee(opts *bind.FilterOpts) (*MarketsFeeIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return &MarketsFeeIterator{contract: _Markets.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0x48b09ff26de49c6dfb086a28888f39d8954ae734fa818411eff3051bc315783a.
//
// Solidity: event Fee(uint256 fee, uint256 feeToProposer, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) WatchFee(opts *bind.WatchOpts, sink chan<- *MarketsFee) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Fee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsFee)
				if err := _Markets.contract.UnpackLog(event, "Fee", log); err != nil {
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

// ParseFee is a log parse operation binding the contract event 0x48b09ff26de49c6dfb086a28888f39d8954ae734fa818411eff3051bc315783a.
//
// Solidity: event Fee(uint256 fee, uint256 feeToProposer, uint256 marketId, uint256 claimId)
func (_Markets *MarketsFilterer) ParseFee(log types.Log) (*MarketsFee, error) {
	event := new(MarketsFee)
	if err := _Markets.contract.UnpackLog(event, "Fee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsPrepareVoteIterator is returned from FilterPrepareVote and is used to iterate over the raw logs and unpacked data for PrepareVote events raised by the Markets contract.
type MarketsPrepareVoteIterator struct {
	Event *MarketsPrepareVote // Event containing the contract specifics and raw log

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
func (it *MarketsPrepareVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsPrepareVote)
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
		it.Event = new(MarketsPrepareVote)
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
func (it *MarketsPrepareVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsPrepareVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsPrepareVote represents a PrepareVote event raised by the Markets contract.
type MarketsPrepareVote struct {
	MarketId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPrepareVote is a free log retrieval operation binding the contract event 0xdd2aa3fdaaf522bd7c073a50a46e1738df7bb414861861fdbf4b0f18be4f5532.
//
// Solidity: event PrepareVote(uint256 marketId)
func (_Markets *MarketsFilterer) FilterPrepareVote(opts *bind.FilterOpts) (*MarketsPrepareVoteIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "PrepareVote")
	if err != nil {
		return nil, err
	}
	return &MarketsPrepareVoteIterator{contract: _Markets.contract, event: "PrepareVote", logs: logs, sub: sub}, nil
}

// WatchPrepareVote is a free log subscription operation binding the contract event 0xdd2aa3fdaaf522bd7c073a50a46e1738df7bb414861861fdbf4b0f18be4f5532.
//
// Solidity: event PrepareVote(uint256 marketId)
func (_Markets *MarketsFilterer) WatchPrepareVote(opts *bind.WatchOpts, sink chan<- *MarketsPrepareVote) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "PrepareVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsPrepareVote)
				if err := _Markets.contract.UnpackLog(event, "PrepareVote", log); err != nil {
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

// ParsePrepareVote is a log parse operation binding the contract event 0xdd2aa3fdaaf522bd7c073a50a46e1738df7bb414861861fdbf4b0f18be4f5532.
//
// Solidity: event PrepareVote(uint256 marketId)
func (_Markets *MarketsFilterer) ParsePrepareVote(log types.Log) (*MarketsPrepareVote, error) {
	event := new(MarketsPrepareVote)
	if err := _Markets.contract.UnpackLog(event, "PrepareVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsProposedIterator is returned from FilterProposed and is used to iterate over the raw logs and unpacked data for Proposed events raised by the Markets contract.
type MarketsProposedIterator struct {
	Event *MarketsProposed // Event containing the contract specifics and raw log

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
func (it *MarketsProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsProposed)
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
		it.Event = new(MarketsProposed)
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
func (it *MarketsProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsProposed represents a Proposed event raised by the Markets contract.
type MarketsProposed struct {
	Propose IMarketsPropose
	ClaimId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProposed is a free log retrieval operation binding the contract event 0x6e8cb6748c0313699a49a56600054e67a16feb6e3c2f76302d5c5294c79f9b03.
//
// Solidity: event Proposed((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) propose, uint256 claimId)
func (_Markets *MarketsFilterer) FilterProposed(opts *bind.FilterOpts) (*MarketsProposedIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return &MarketsProposedIterator{contract: _Markets.contract, event: "Proposed", logs: logs, sub: sub}, nil
}

// WatchProposed is a free log subscription operation binding the contract event 0x6e8cb6748c0313699a49a56600054e67a16feb6e3c2f76302d5c5294c79f9b03.
//
// Solidity: event Proposed((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) propose, uint256 claimId)
func (_Markets *MarketsFilterer) WatchProposed(opts *bind.WatchOpts, sink chan<- *MarketsProposed) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Proposed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsProposed)
				if err := _Markets.contract.UnpackLog(event, "Proposed", log); err != nil {
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

// ParseProposed is a log parse operation binding the contract event 0x6e8cb6748c0313699a49a56600054e67a16feb6e3c2f76302d5c5294c79f9b03.
//
// Solidity: event Proposed((string,uint256,uint256,uint256,address,uint40,uint40,bool,bool,(uint8,uint256)) propose, uint256 claimId)
func (_Markets *MarketsFilterer) ParseProposed(log types.Log) (*MarketsProposed, error) {
	event := new(MarketsProposed)
	if err := _Markets.contract.UnpackLog(event, "Proposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsResolveIterator is returned from FilterResolve and is used to iterate over the raw logs and unpacked data for Resolve events raised by the Markets contract.
type MarketsResolveIterator struct {
	Event *MarketsResolve // Event containing the contract specifics and raw log

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
func (it *MarketsResolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsResolve)
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
		it.Event = new(MarketsResolve)
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
func (it *MarketsResolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsResolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsResolve represents a Resolve event raised by the Markets contract.
type MarketsResolve struct {
	MarketId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResolve is a free log retrieval operation binding the contract event 0xea992c5d8822e09d110191b5e9ea82ee790c417e85079a54f12dc3169ddb2c71.
//
// Solidity: event Resolve(uint256 marketId)
func (_Markets *MarketsFilterer) FilterResolve(opts *bind.FilterOpts) (*MarketsResolveIterator, error) {

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Resolve")
	if err != nil {
		return nil, err
	}
	return &MarketsResolveIterator{contract: _Markets.contract, event: "Resolve", logs: logs, sub: sub}, nil
}

// WatchResolve is a free log subscription operation binding the contract event 0xea992c5d8822e09d110191b5e9ea82ee790c417e85079a54f12dc3169ddb2c71.
//
// Solidity: event Resolve(uint256 marketId)
func (_Markets *MarketsFilterer) WatchResolve(opts *bind.WatchOpts, sink chan<- *MarketsResolve) (event.Subscription, error) {

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Resolve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsResolve)
				if err := _Markets.contract.UnpackLog(event, "Resolve", log); err != nil {
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

// ParseResolve is a log parse operation binding the contract event 0xea992c5d8822e09d110191b5e9ea82ee790c417e85079a54f12dc3169ddb2c71.
//
// Solidity: event Resolve(uint256 marketId)
func (_Markets *MarketsFilterer) ParseResolve(log types.Log) (*MarketsResolve, error) {
	event := new(MarketsResolve)
	if err := _Markets.contract.UnpackLog(event, "Resolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Markets contract.
type MarketsStakedIterator struct {
	Event *MarketsStaked // Event containing the contract specifics and raw log

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
func (it *MarketsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsStaked)
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
		it.Event = new(MarketsStaked)
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
func (it *MarketsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsStaked represents a Staked event raised by the Markets contract.
type MarketsStaked struct {
	Staker             common.Address
	MarketId           *big.Int
	ClaimId            *big.Int
	Amount             *big.Int
	TimeWeightedAmount *big.Int
	Yea                bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xea1d652d20662dd9e57f36998e41bc3a8f5afa8dc39f71c1b5d23c812f6e92ef.
//
// Solidity: event Staked(address indexed staker, uint256 marketId, uint256 claimId, uint256 amount, uint256 timeWeightedAmount, bool yea)
func (_Markets *MarketsFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*MarketsStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MarketsStakedIterator{contract: _Markets.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xea1d652d20662dd9e57f36998e41bc3a8f5afa8dc39f71c1b5d23c812f6e92ef.
//
// Solidity: event Staked(address indexed staker, uint256 marketId, uint256 claimId, uint256 amount, uint256 timeWeightedAmount, bool yea)
func (_Markets *MarketsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *MarketsStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsStaked)
				if err := _Markets.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xea1d652d20662dd9e57f36998e41bc3a8f5afa8dc39f71c1b5d23c812f6e92ef.
//
// Solidity: event Staked(address indexed staker, uint256 marketId, uint256 claimId, uint256 amount, uint256 timeWeightedAmount, bool yea)
func (_Markets *MarketsFilterer) ParseStaked(log types.Log) (*MarketsStaked, error) {
	event := new(MarketsStaked)
	if err := _Markets.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Markets contract.
type MarketsVotedIterator struct {
	Event *MarketsVoted // Event containing the contract specifics and raw log

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
func (it *MarketsVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsVoted)
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
		it.Event = new(MarketsVoted)
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
func (it *MarketsVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsVoted represents a Voted event raised by the Markets contract.
type MarketsVoted struct {
	Voter    common.Address
	MarketId *big.Int
	ClaimId  *big.Int
	Yea      bool
	Stake    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x1c44c0a768d5df8d7c8c1faeef809807b2aa41eb2164631e4b5d021c67446a5d.
//
// Solidity: event Voted(address indexed voter, uint256 marketId, uint256 claimId, bool yea, bool stake)
func (_Markets *MarketsFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address) (*MarketsVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return &MarketsVotedIterator{contract: _Markets.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x1c44c0a768d5df8d7c8c1faeef809807b2aa41eb2164631e4b5d021c67446a5d.
//
// Solidity: event Voted(address indexed voter, uint256 marketId, uint256 claimId, bool yea, bool stake)
func (_Markets *MarketsFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *MarketsVoted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsVoted)
				if err := _Markets.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x1c44c0a768d5df8d7c8c1faeef809807b2aa41eb2164631e4b5d021c67446a5d.
//
// Solidity: event Voted(address indexed voter, uint256 marketId, uint256 claimId, bool yea, bool stake)
func (_Markets *MarketsFilterer) ParseVoted(log types.Log) (*MarketsVoted, error) {
	event := new(MarketsVoted)
	if err := _Markets.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Markets contract.
type MarketsWithdrawIterator struct {
	Event *MarketsWithdraw // Event containing the contract specifics and raw log

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
func (it *MarketsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketsWithdraw)
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
		it.Event = new(MarketsWithdraw)
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
func (it *MarketsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketsWithdraw represents a Withdraw event raised by the Markets contract.
type MarketsWithdraw struct {
	Sender   common.Address
	Reciever common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, reciever []common.Address) (*MarketsWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Markets.contract.FilterLogs(opts, "Withdraw", senderRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return &MarketsWithdrawIterator{contract: _Markets.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MarketsWithdraw, sender []common.Address, reciever []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recieverRule []interface{}
	for _, recieverItem := range reciever {
		recieverRule = append(recieverRule, recieverItem)
	}

	logs, sub, err := _Markets.contract.WatchLogs(opts, "Withdraw", senderRule, recieverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketsWithdraw)
				if err := _Markets.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed sender, address indexed reciever, uint256 amount)
func (_Markets *MarketsFilterer) ParseWithdraw(log types.Log) (*MarketsWithdraw, error) {
	event := new(MarketsWithdraw)
	if err := _Markets.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
