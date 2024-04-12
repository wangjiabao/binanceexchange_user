// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// TradingBoxMetaData contains all meta data concerning the TradingBox contract.
var TradingBoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCOME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"beginTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxOpenTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxPerSecondHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxPerSecondRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boxTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"income\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openTimeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openTimeStamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perSecondBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"perSecondRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"safeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setBeginTimeStamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setBoxPerSecondHistory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"incomeAmount\",\"type\":\"int256\"}],\"name\":\"setIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"int256[]\",\"name\":\"incomeAmounts\",\"type\":\"int256[]\"}],\"name\":\"setIncomeAny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate_\",\"type\":\"uint256\"}],\"name\":\"setPerSecondRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeStamp_\",\"type\":\"uint256\"}],\"name\":\"setTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setWithdrawIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TradingBoxABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingBoxMetaData.ABI instead.
var TradingBoxABI = TradingBoxMetaData.ABI

// TradingBox is an auto generated Go binding around an Ethereum contract.
type TradingBox struct {
	TradingBoxCaller     // Read-only binding to the contract
	TradingBoxTransactor // Write-only binding to the contract
	TradingBoxFilterer   // Log filterer for contract events
}

// TradingBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingBoxSession struct {
	Contract     *TradingBox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingBoxCallerSession struct {
	Contract *TradingBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradingBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingBoxTransactorSession struct {
	Contract     *TradingBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradingBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingBoxRaw struct {
	Contract *TradingBox // Generic contract binding to access the raw methods on
}

// TradingBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingBoxCallerRaw struct {
	Contract *TradingBoxCaller // Generic read-only contract binding to access the raw methods on
}

// TradingBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingBoxTransactorRaw struct {
	Contract *TradingBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingBox creates a new instance of TradingBox, bound to a specific deployed contract.
func NewTradingBox(address common.Address, backend bind.ContractBackend) (*TradingBox, error) {
	contract, err := bindTradingBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradingBox{TradingBoxCaller: TradingBoxCaller{contract: contract}, TradingBoxTransactor: TradingBoxTransactor{contract: contract}, TradingBoxFilterer: TradingBoxFilterer{contract: contract}}, nil
}

// NewTradingBoxCaller creates a new read-only instance of TradingBox, bound to a specific deployed contract.
func NewTradingBoxCaller(address common.Address, caller bind.ContractCaller) (*TradingBoxCaller, error) {
	contract, err := bindTradingBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingBoxCaller{contract: contract}, nil
}

// NewTradingBoxTransactor creates a new write-only instance of TradingBox, bound to a specific deployed contract.
func NewTradingBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingBoxTransactor, error) {
	contract, err := bindTradingBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingBoxTransactor{contract: contract}, nil
}

// NewTradingBoxFilterer creates a new log filterer instance of TradingBox, bound to a specific deployed contract.
func NewTradingBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingBoxFilterer, error) {
	contract, err := bindTradingBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingBoxFilterer{contract: contract}, nil
}

// bindTradingBox binds a generic wrapper to an already deployed contract.
func bindTradingBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingBox *TradingBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingBox.Contract.TradingBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingBox *TradingBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBox.Contract.TradingBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingBox *TradingBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingBox.Contract.TradingBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingBox *TradingBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingBox *TradingBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingBox *TradingBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingBox.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingBox.Contract.DEFAULTADMINROLE(&_TradingBox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingBox.Contract.DEFAULTADMINROLE(&_TradingBox.CallOpts)
}

// INCOMEROLE is a free data retrieval call binding the contract method 0x0a4fb640.
//
// Solidity: function INCOME_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCaller) INCOMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "INCOME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INCOMEROLE is a free data retrieval call binding the contract method 0x0a4fb640.
//
// Solidity: function INCOME_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxSession) INCOMEROLE() ([32]byte, error) {
	return _TradingBox.Contract.INCOMEROLE(&_TradingBox.CallOpts)
}

// INCOMEROLE is a free data retrieval call binding the contract method 0x0a4fb640.
//
// Solidity: function INCOME_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCallerSession) INCOMEROLE() ([32]byte, error) {
	return _TradingBox.Contract.INCOMEROLE(&_TradingBox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxSession) PAUSERROLE() ([32]byte, error) {
	return _TradingBox.Contract.PAUSERROLE(&_TradingBox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _TradingBox.Contract.PAUSERROLE(&_TradingBox.CallOpts)
}

// UPDATEROLE is a free data retrieval call binding the contract method 0x00497846.
//
// Solidity: function UPDATE_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCaller) UPDATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "UPDATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATEROLE is a free data retrieval call binding the contract method 0x00497846.
//
// Solidity: function UPDATE_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxSession) UPDATEROLE() ([32]byte, error) {
	return _TradingBox.Contract.UPDATEROLE(&_TradingBox.CallOpts)
}

// UPDATEROLE is a free data retrieval call binding the contract method 0x00497846.
//
// Solidity: function UPDATE_ROLE() view returns(bytes32)
func (_TradingBox *TradingBoxCallerSession) UPDATEROLE() ([32]byte, error) {
	return _TradingBox.Contract.UPDATEROLE(&_TradingBox.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TradingBox *TradingBoxSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TradingBox.Contract.BalanceOf(&_TradingBox.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TradingBox.Contract.BalanceOf(&_TradingBox.CallOpts, owner)
}

// BeginTimeStamp is a free data retrieval call binding the contract method 0x48b59885.
//
// Solidity: function beginTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BeginTimeStamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "beginTimeStamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BeginTimeStamp is a free data retrieval call binding the contract method 0x48b59885.
//
// Solidity: function beginTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) BeginTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BeginTimeStamp(&_TradingBox.CallOpts, arg0)
}

// BeginTimeStamp is a free data retrieval call binding the contract method 0x48b59885.
//
// Solidity: function beginTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BeginTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BeginTimeStamp(&_TradingBox.CallOpts, arg0)
}

// BoxOpenTimeLimit is a free data retrieval call binding the contract method 0xd015acb0.
//
// Solidity: function boxOpenTimeLimit(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BoxOpenTimeLimit(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "boxOpenTimeLimit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoxOpenTimeLimit is a free data retrieval call binding the contract method 0xd015acb0.
//
// Solidity: function boxOpenTimeLimit(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) BoxOpenTimeLimit(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxOpenTimeLimit(&_TradingBox.CallOpts, arg0)
}

// BoxOpenTimeLimit is a free data retrieval call binding the contract method 0xd015acb0.
//
// Solidity: function boxOpenTimeLimit(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BoxOpenTimeLimit(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxOpenTimeLimit(&_TradingBox.CallOpts, arg0)
}

// BoxPerSecondHistory is a free data retrieval call binding the contract method 0x5273a7eb.
//
// Solidity: function boxPerSecondHistory(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BoxPerSecondHistory(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "boxPerSecondHistory", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoxPerSecondHistory is a free data retrieval call binding the contract method 0x5273a7eb.
//
// Solidity: function boxPerSecondHistory(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) BoxPerSecondHistory(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxPerSecondHistory(&_TradingBox.CallOpts, arg0)
}

// BoxPerSecondHistory is a free data retrieval call binding the contract method 0x5273a7eb.
//
// Solidity: function boxPerSecondHistory(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BoxPerSecondHistory(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxPerSecondHistory(&_TradingBox.CallOpts, arg0)
}

// BoxPerSecondRate is a free data retrieval call binding the contract method 0xaada9041.
//
// Solidity: function boxPerSecondRate(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BoxPerSecondRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "boxPerSecondRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoxPerSecondRate is a free data retrieval call binding the contract method 0xaada9041.
//
// Solidity: function boxPerSecondRate(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) BoxPerSecondRate(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxPerSecondRate(&_TradingBox.CallOpts, arg0)
}

// BoxPerSecondRate is a free data retrieval call binding the contract method 0xaada9041.
//
// Solidity: function boxPerSecondRate(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BoxPerSecondRate(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxPerSecondRate(&_TradingBox.CallOpts, arg0)
}

// BoxTokenAmount is a free data retrieval call binding the contract method 0xe5eaee7b.
//
// Solidity: function boxTokenAmount(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) BoxTokenAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "boxTokenAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoxTokenAmount is a free data retrieval call binding the contract method 0xe5eaee7b.
//
// Solidity: function boxTokenAmount(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) BoxTokenAmount(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxTokenAmount(&_TradingBox.CallOpts, arg0)
}

// BoxTokenAmount is a free data retrieval call binding the contract method 0xe5eaee7b.
//
// Solidity: function boxTokenAmount(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) BoxTokenAmount(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.BoxTokenAmount(&_TradingBox.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TradingBox.Contract.GetApproved(&_TradingBox.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TradingBox.Contract.GetApproved(&_TradingBox.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingBox *TradingBoxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingBox *TradingBoxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingBox.Contract.GetRoleAdmin(&_TradingBox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingBox *TradingBoxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingBox.Contract.GetRoleAdmin(&_TradingBox.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingBox *TradingBoxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingBox *TradingBoxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingBox.Contract.HasRole(&_TradingBox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingBox *TradingBoxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingBox.Contract.HasRole(&_TradingBox.CallOpts, role, account)
}

// Income is a free data retrieval call binding the contract method 0x84ed8c56.
//
// Solidity: function income(uint256 ) view returns(int256)
func (_TradingBox *TradingBoxCaller) Income(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "income", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Income is a free data retrieval call binding the contract method 0x84ed8c56.
//
// Solidity: function income(uint256 ) view returns(int256)
func (_TradingBox *TradingBoxSession) Income(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.Income(&_TradingBox.CallOpts, arg0)
}

// Income is a free data retrieval call binding the contract method 0x84ed8c56.
//
// Solidity: function income(uint256 ) view returns(int256)
func (_TradingBox *TradingBoxCallerSession) Income(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.Income(&_TradingBox.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TradingBox *TradingBoxCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TradingBox *TradingBoxSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TradingBox.Contract.IsApprovedForAll(&_TradingBox.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TradingBox *TradingBoxCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TradingBox.Contract.IsApprovedForAll(&_TradingBox.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TradingBox *TradingBoxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TradingBox *TradingBoxSession) Name() (string, error) {
	return _TradingBox.Contract.Name(&_TradingBox.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TradingBox *TradingBoxCallerSession) Name() (string, error) {
	return _TradingBox.Contract.Name(&_TradingBox.CallOpts)
}

// OpenTimeLimit is a free data retrieval call binding the contract method 0x00c91d28.
//
// Solidity: function openTimeLimit() view returns(uint256)
func (_TradingBox *TradingBoxCaller) OpenTimeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "openTimeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenTimeLimit is a free data retrieval call binding the contract method 0x00c91d28.
//
// Solidity: function openTimeLimit() view returns(uint256)
func (_TradingBox *TradingBoxSession) OpenTimeLimit() (*big.Int, error) {
	return _TradingBox.Contract.OpenTimeLimit(&_TradingBox.CallOpts)
}

// OpenTimeLimit is a free data retrieval call binding the contract method 0x00c91d28.
//
// Solidity: function openTimeLimit() view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) OpenTimeLimit() (*big.Int, error) {
	return _TradingBox.Contract.OpenTimeLimit(&_TradingBox.CallOpts)
}

// OpenTimeStamp is a free data retrieval call binding the contract method 0x66267084.
//
// Solidity: function openTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) OpenTimeStamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "openTimeStamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenTimeStamp is a free data retrieval call binding the contract method 0x66267084.
//
// Solidity: function openTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) OpenTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.OpenTimeStamp(&_TradingBox.CallOpts, arg0)
}

// OpenTimeStamp is a free data retrieval call binding the contract method 0x66267084.
//
// Solidity: function openTimeStamp(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) OpenTimeStamp(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.OpenTimeStamp(&_TradingBox.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TradingBox.Contract.OwnerOf(&_TradingBox.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TradingBox *TradingBoxCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TradingBox.Contract.OwnerOf(&_TradingBox.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingBox *TradingBoxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingBox *TradingBoxSession) Paused() (bool, error) {
	return _TradingBox.Contract.Paused(&_TradingBox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingBox *TradingBoxCallerSession) Paused() (bool, error) {
	return _TradingBox.Contract.Paused(&_TradingBox.CallOpts)
}

// PerSecondBase is a free data retrieval call binding the contract method 0x0caa3d46.
//
// Solidity: function perSecondBase() view returns(uint256)
func (_TradingBox *TradingBoxCaller) PerSecondBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "perSecondBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerSecondBase is a free data retrieval call binding the contract method 0x0caa3d46.
//
// Solidity: function perSecondBase() view returns(uint256)
func (_TradingBox *TradingBoxSession) PerSecondBase() (*big.Int, error) {
	return _TradingBox.Contract.PerSecondBase(&_TradingBox.CallOpts)
}

// PerSecondBase is a free data retrieval call binding the contract method 0x0caa3d46.
//
// Solidity: function perSecondBase() view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) PerSecondBase() (*big.Int, error) {
	return _TradingBox.Contract.PerSecondBase(&_TradingBox.CallOpts)
}

// PerSecondRate is a free data retrieval call binding the contract method 0x705f2ec8.
//
// Solidity: function perSecondRate() view returns(uint256)
func (_TradingBox *TradingBoxCaller) PerSecondRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "perSecondRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerSecondRate is a free data retrieval call binding the contract method 0x705f2ec8.
//
// Solidity: function perSecondRate() view returns(uint256)
func (_TradingBox *TradingBoxSession) PerSecondRate() (*big.Int, error) {
	return _TradingBox.Contract.PerSecondRate(&_TradingBox.CallOpts)
}

// PerSecondRate is a free data retrieval call binding the contract method 0x705f2ec8.
//
// Solidity: function perSecondRate() view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) PerSecondRate() (*big.Int, error) {
	return _TradingBox.Contract.PerSecondRate(&_TradingBox.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(bool)
func (_TradingBox *TradingBoxCaller) Status(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "status", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(bool)
func (_TradingBox *TradingBoxSession) Status(arg0 *big.Int) (bool, error) {
	return _TradingBox.Contract.Status(&_TradingBox.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x42d21ef7.
//
// Solidity: function status(uint256 ) view returns(bool)
func (_TradingBox *TradingBoxCallerSession) Status(arg0 *big.Int) (bool, error) {
	return _TradingBox.Contract.Status(&_TradingBox.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingBox *TradingBoxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingBox *TradingBoxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingBox.Contract.SupportsInterface(&_TradingBox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingBox *TradingBoxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingBox.Contract.SupportsInterface(&_TradingBox.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TradingBox *TradingBoxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TradingBox *TradingBoxSession) Symbol() (string, error) {
	return _TradingBox.Contract.Symbol(&_TradingBox.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TradingBox *TradingBoxCallerSession) Symbol() (string, error) {
	return _TradingBox.Contract.Symbol(&_TradingBox.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.TokenByIndex(&_TradingBox.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.TokenByIndex(&_TradingBox.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.TokenOfOwnerByIndex(&_TradingBox.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.TokenOfOwnerByIndex(&_TradingBox.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TradingBox *TradingBoxCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TradingBox *TradingBoxSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TradingBox.Contract.TokenURI(&_TradingBox.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TradingBox *TradingBoxCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TradingBox.Contract.TokenURI(&_TradingBox.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TradingBox *TradingBoxCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TradingBox *TradingBoxSession) TotalSupply() (*big.Int, error) {
	return _TradingBox.Contract.TotalSupply(&_TradingBox.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) TotalSupply() (*big.Int, error) {
	return _TradingBox.Contract.TotalSupply(&_TradingBox.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingBox *TradingBoxCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingBox *TradingBoxSession) Usdt() (common.Address, error) {
	return _TradingBox.Contract.Usdt(&_TradingBox.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingBox *TradingBoxCallerSession) Usdt() (common.Address, error) {
	return _TradingBox.Contract.Usdt(&_TradingBox.CallOpts)
}

// WithdrawIncome is a free data retrieval call binding the contract method 0x9273bbb6.
//
// Solidity: function withdrawIncome(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCaller) WithdrawIncome(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingBox.contract.Call(opts, &out, "withdrawIncome", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawIncome is a free data retrieval call binding the contract method 0x9273bbb6.
//
// Solidity: function withdrawIncome(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxSession) WithdrawIncome(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.WithdrawIncome(&_TradingBox.CallOpts, arg0)
}

// WithdrawIncome is a free data retrieval call binding the contract method 0x9273bbb6.
//
// Solidity: function withdrawIncome(uint256 ) view returns(uint256)
func (_TradingBox *TradingBoxCallerSession) WithdrawIncome(arg0 *big.Int) (*big.Int, error) {
	return _TradingBox.Contract.WithdrawIncome(&_TradingBox.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Approve(&_TradingBox.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Approve(&_TradingBox.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Burn(&_TradingBox.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Burn(&_TradingBox.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.GrantRole(&_TradingBox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.GrantRole(&_TradingBox.TransactOpts, role, account)
}

// Open is a paid mutator transaction binding the contract method 0x0a0e5c9d.
//
// Solidity: function open(address account, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) Open(opts *bind.TransactOpts, account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "open", account, tokenId)
}

// Open is a paid mutator transaction binding the contract method 0x0a0e5c9d.
//
// Solidity: function open(address account, uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) Open(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Open(&_TradingBox.TransactOpts, account, tokenId)
}

// Open is a paid mutator transaction binding the contract method 0x0a0e5c9d.
//
// Solidity: function open(address account, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) Open(account common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.Open(&_TradingBox.TransactOpts, account, tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TradingBox *TradingBoxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TradingBox *TradingBoxSession) Pause() (*types.Transaction, error) {
	return _TradingBox.Contract.Pause(&_TradingBox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TradingBox *TradingBoxTransactorSession) Pause() (*types.Transaction, error) {
	return _TradingBox.Contract.Pause(&_TradingBox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingBox *TradingBoxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingBox *TradingBoxSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.RenounceRole(&_TradingBox.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingBox *TradingBoxTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.RenounceRole(&_TradingBox.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.RevokeRole(&_TradingBox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingBox *TradingBoxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingBox.Contract.RevokeRole(&_TradingBox.TransactOpts, role, account)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc772745a.
//
// Solidity: function safeMint(address to, string uri, uint256 amount) returns(uint256)
func (_TradingBox *TradingBoxTransactor) SafeMint(opts *bind.TransactOpts, to common.Address, uri string, amount *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "safeMint", to, uri, amount)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc772745a.
//
// Solidity: function safeMint(address to, string uri, uint256 amount) returns(uint256)
func (_TradingBox *TradingBoxSession) SafeMint(to common.Address, uri string, amount *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeMint(&_TradingBox.TransactOpts, to, uri, amount)
}

// SafeMint is a paid mutator transaction binding the contract method 0xc772745a.
//
// Solidity: function safeMint(address to, string uri, uint256 amount) returns(uint256)
func (_TradingBox *TradingBoxTransactorSession) SafeMint(to common.Address, uri string, amount *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeMint(&_TradingBox.TransactOpts, to, uri, amount)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeTransferFrom(&_TradingBox.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeTransferFrom(&_TradingBox.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TradingBox *TradingBoxTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TradingBox *TradingBoxSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeTransferFrom0(&_TradingBox.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TradingBox *TradingBoxTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TradingBox.Contract.SafeTransferFrom0(&_TradingBox.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TradingBox *TradingBoxTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TradingBox *TradingBoxSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TradingBox.Contract.SetApprovalForAll(&_TradingBox.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TradingBox *TradingBoxTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TradingBox.Contract.SetApprovalForAll(&_TradingBox.TransactOpts, operator, approved)
}

// SetBeginTimeStamp is a paid mutator transaction binding the contract method 0xf67910ce.
//
// Solidity: function setBeginTimeStamp(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) SetBeginTimeStamp(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setBeginTimeStamp", tokenId)
}

// SetBeginTimeStamp is a paid mutator transaction binding the contract method 0xf67910ce.
//
// Solidity: function setBeginTimeStamp(uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) SetBeginTimeStamp(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetBeginTimeStamp(&_TradingBox.TransactOpts, tokenId)
}

// SetBeginTimeStamp is a paid mutator transaction binding the contract method 0xf67910ce.
//
// Solidity: function setBeginTimeStamp(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) SetBeginTimeStamp(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetBeginTimeStamp(&_TradingBox.TransactOpts, tokenId)
}

// SetBoxPerSecondHistory is a paid mutator transaction binding the contract method 0x66b570ec.
//
// Solidity: function setBoxPerSecondHistory(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) SetBoxPerSecondHistory(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setBoxPerSecondHistory", tokenId)
}

// SetBoxPerSecondHistory is a paid mutator transaction binding the contract method 0x66b570ec.
//
// Solidity: function setBoxPerSecondHistory(uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) SetBoxPerSecondHistory(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetBoxPerSecondHistory(&_TradingBox.TransactOpts, tokenId)
}

// SetBoxPerSecondHistory is a paid mutator transaction binding the contract method 0x66b570ec.
//
// Solidity: function setBoxPerSecondHistory(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) SetBoxPerSecondHistory(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetBoxPerSecondHistory(&_TradingBox.TransactOpts, tokenId)
}

// SetIncome is a paid mutator transaction binding the contract method 0x79398fb0.
//
// Solidity: function setIncome(uint256 tokenId, int256 incomeAmount) returns()
func (_TradingBox *TradingBoxTransactor) SetIncome(opts *bind.TransactOpts, tokenId *big.Int, incomeAmount *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setIncome", tokenId, incomeAmount)
}

// SetIncome is a paid mutator transaction binding the contract method 0x79398fb0.
//
// Solidity: function setIncome(uint256 tokenId, int256 incomeAmount) returns()
func (_TradingBox *TradingBoxSession) SetIncome(tokenId *big.Int, incomeAmount *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetIncome(&_TradingBox.TransactOpts, tokenId, incomeAmount)
}

// SetIncome is a paid mutator transaction binding the contract method 0x79398fb0.
//
// Solidity: function setIncome(uint256 tokenId, int256 incomeAmount) returns()
func (_TradingBox *TradingBoxTransactorSession) SetIncome(tokenId *big.Int, incomeAmount *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetIncome(&_TradingBox.TransactOpts, tokenId, incomeAmount)
}

// SetIncomeAny is a paid mutator transaction binding the contract method 0x3d21e1a9.
//
// Solidity: function setIncomeAny(uint256[] tokenIds, int256[] incomeAmounts) returns()
func (_TradingBox *TradingBoxTransactor) SetIncomeAny(opts *bind.TransactOpts, tokenIds []*big.Int, incomeAmounts []*big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setIncomeAny", tokenIds, incomeAmounts)
}

// SetIncomeAny is a paid mutator transaction binding the contract method 0x3d21e1a9.
//
// Solidity: function setIncomeAny(uint256[] tokenIds, int256[] incomeAmounts) returns()
func (_TradingBox *TradingBoxSession) SetIncomeAny(tokenIds []*big.Int, incomeAmounts []*big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetIncomeAny(&_TradingBox.TransactOpts, tokenIds, incomeAmounts)
}

// SetIncomeAny is a paid mutator transaction binding the contract method 0x3d21e1a9.
//
// Solidity: function setIncomeAny(uint256[] tokenIds, int256[] incomeAmounts) returns()
func (_TradingBox *TradingBoxTransactorSession) SetIncomeAny(tokenIds []*big.Int, incomeAmounts []*big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetIncomeAny(&_TradingBox.TransactOpts, tokenIds, incomeAmounts)
}

// SetPerSecondRate is a paid mutator transaction binding the contract method 0x6eed3576.
//
// Solidity: function setPerSecondRate(uint256 rate_) returns()
func (_TradingBox *TradingBoxTransactor) SetPerSecondRate(opts *bind.TransactOpts, rate_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setPerSecondRate", rate_)
}

// SetPerSecondRate is a paid mutator transaction binding the contract method 0x6eed3576.
//
// Solidity: function setPerSecondRate(uint256 rate_) returns()
func (_TradingBox *TradingBoxSession) SetPerSecondRate(rate_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetPerSecondRate(&_TradingBox.TransactOpts, rate_)
}

// SetPerSecondRate is a paid mutator transaction binding the contract method 0x6eed3576.
//
// Solidity: function setPerSecondRate(uint256 rate_) returns()
func (_TradingBox *TradingBoxTransactorSession) SetPerSecondRate(rate_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetPerSecondRate(&_TradingBox.TransactOpts, rate_)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeStamp_) returns()
func (_TradingBox *TradingBoxTransactor) SetTimeLimit(opts *bind.TransactOpts, timeStamp_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setTimeLimit", timeStamp_)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeStamp_) returns()
func (_TradingBox *TradingBoxSession) SetTimeLimit(timeStamp_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetTimeLimit(&_TradingBox.TransactOpts, timeStamp_)
}

// SetTimeLimit is a paid mutator transaction binding the contract method 0xe2889c82.
//
// Solidity: function setTimeLimit(uint256 timeStamp_) returns()
func (_TradingBox *TradingBoxTransactorSession) SetTimeLimit(timeStamp_ *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetTimeLimit(&_TradingBox.TransactOpts, timeStamp_)
}

// SetWithdrawIncome is a paid mutator transaction binding the contract method 0x2cd84673.
//
// Solidity: function setWithdrawIncome(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) SetWithdrawIncome(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "setWithdrawIncome", tokenId)
}

// SetWithdrawIncome is a paid mutator transaction binding the contract method 0x2cd84673.
//
// Solidity: function setWithdrawIncome(uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) SetWithdrawIncome(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetWithdrawIncome(&_TradingBox.TransactOpts, tokenId)
}

// SetWithdrawIncome is a paid mutator transaction binding the contract method 0x2cd84673.
//
// Solidity: function setWithdrawIncome(uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) SetWithdrawIncome(tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.SetWithdrawIncome(&_TradingBox.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.TransferFrom(&_TradingBox.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TradingBox *TradingBoxTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TradingBox.Contract.TransferFrom(&_TradingBox.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TradingBox *TradingBoxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingBox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TradingBox *TradingBoxSession) Unpause() (*types.Transaction, error) {
	return _TradingBox.Contract.Unpause(&_TradingBox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TradingBox *TradingBoxTransactorSession) Unpause() (*types.Transaction, error) {
	return _TradingBox.Contract.Unpause(&_TradingBox.TransactOpts)
}

// TradingBoxApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TradingBox contract.
type TradingBoxApprovalIterator struct {
	Event *TradingBoxApproval // Event containing the contract specifics and raw log

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
func (it *TradingBoxApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxApproval)
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
		it.Event = new(TradingBoxApproval)
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
func (it *TradingBoxApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxApproval represents a Approval event raised by the TradingBox contract.
type TradingBoxApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TradingBoxApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxApprovalIterator{contract: _TradingBox.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TradingBoxApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxApproval)
				if err := _TradingBox.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) ParseApproval(log types.Log) (*TradingBoxApproval, error) {
	event := new(TradingBoxApproval)
	if err := _TradingBox.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TradingBox contract.
type TradingBoxApprovalForAllIterator struct {
	Event *TradingBoxApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TradingBoxApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxApprovalForAll)
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
		it.Event = new(TradingBoxApprovalForAll)
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
func (it *TradingBoxApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxApprovalForAll represents a ApprovalForAll event raised by the TradingBox contract.
type TradingBoxApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TradingBox *TradingBoxFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TradingBoxApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxApprovalForAllIterator{contract: _TradingBox.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TradingBox *TradingBoxFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TradingBoxApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxApprovalForAll)
				if err := _TradingBox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TradingBox *TradingBoxFilterer) ParseApprovalForAll(log types.Log) (*TradingBoxApprovalForAll, error) {
	event := new(TradingBoxApprovalForAll)
	if err := _TradingBox.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the TradingBox contract.
type TradingBoxBatchMetadataUpdateIterator struct {
	Event *TradingBoxBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TradingBoxBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxBatchMetadataUpdate)
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
		it.Event = new(TradingBoxBatchMetadataUpdate)
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
func (it *TradingBoxBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the TradingBox contract.
type TradingBoxBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TradingBox *TradingBoxFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TradingBoxBatchMetadataUpdateIterator, error) {

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TradingBoxBatchMetadataUpdateIterator{contract: _TradingBox.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TradingBox *TradingBoxFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TradingBoxBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxBatchMetadataUpdate)
				if err := _TradingBox.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TradingBox *TradingBoxFilterer) ParseBatchMetadataUpdate(log types.Log) (*TradingBoxBatchMetadataUpdate, error) {
	event := new(TradingBoxBatchMetadataUpdate)
	if err := _TradingBox.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the TradingBox contract.
type TradingBoxMetadataUpdateIterator struct {
	Event *TradingBoxMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TradingBoxMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxMetadataUpdate)
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
		it.Event = new(TradingBoxMetadataUpdate)
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
func (it *TradingBoxMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxMetadataUpdate represents a MetadataUpdate event raised by the TradingBox contract.
type TradingBoxMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TradingBox *TradingBoxFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*TradingBoxMetadataUpdateIterator, error) {

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TradingBoxMetadataUpdateIterator{contract: _TradingBox.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TradingBox *TradingBoxFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TradingBoxMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxMetadataUpdate)
				if err := _TradingBox.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TradingBox *TradingBoxFilterer) ParseMetadataUpdate(log types.Log) (*TradingBoxMetadataUpdate, error) {
	event := new(TradingBoxMetadataUpdate)
	if err := _TradingBox.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TradingBox contract.
type TradingBoxPausedIterator struct {
	Event *TradingBoxPaused // Event containing the contract specifics and raw log

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
func (it *TradingBoxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxPaused)
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
		it.Event = new(TradingBoxPaused)
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
func (it *TradingBoxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxPaused represents a Paused event raised by the TradingBox contract.
type TradingBoxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TradingBox *TradingBoxFilterer) FilterPaused(opts *bind.FilterOpts) (*TradingBoxPausedIterator, error) {

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TradingBoxPausedIterator{contract: _TradingBox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TradingBox *TradingBoxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TradingBoxPaused) (event.Subscription, error) {

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxPaused)
				if err := _TradingBox.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TradingBox *TradingBoxFilterer) ParsePaused(log types.Log) (*TradingBoxPaused, error) {
	event := new(TradingBoxPaused)
	if err := _TradingBox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TradingBox contract.
type TradingBoxRoleAdminChangedIterator struct {
	Event *TradingBoxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TradingBoxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxRoleAdminChanged)
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
		it.Event = new(TradingBoxRoleAdminChanged)
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
func (it *TradingBoxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxRoleAdminChanged represents a RoleAdminChanged event raised by the TradingBox contract.
type TradingBoxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingBox *TradingBoxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TradingBoxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxRoleAdminChangedIterator{contract: _TradingBox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingBox *TradingBoxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TradingBoxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxRoleAdminChanged)
				if err := _TradingBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TradingBox *TradingBoxFilterer) ParseRoleAdminChanged(log types.Log) (*TradingBoxRoleAdminChanged, error) {
	event := new(TradingBoxRoleAdminChanged)
	if err := _TradingBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TradingBox contract.
type TradingBoxRoleGrantedIterator struct {
	Event *TradingBoxRoleGranted // Event containing the contract specifics and raw log

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
func (it *TradingBoxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxRoleGranted)
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
		it.Event = new(TradingBoxRoleGranted)
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
func (it *TradingBoxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxRoleGranted represents a RoleGranted event raised by the TradingBox contract.
type TradingBoxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingBox *TradingBoxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingBoxRoleGrantedIterator, error) {

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

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxRoleGrantedIterator{contract: _TradingBox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingBox *TradingBoxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TradingBoxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxRoleGranted)
				if err := _TradingBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TradingBox *TradingBoxFilterer) ParseRoleGranted(log types.Log) (*TradingBoxRoleGranted, error) {
	event := new(TradingBoxRoleGranted)
	if err := _TradingBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TradingBox contract.
type TradingBoxRoleRevokedIterator struct {
	Event *TradingBoxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TradingBoxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxRoleRevoked)
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
		it.Event = new(TradingBoxRoleRevoked)
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
func (it *TradingBoxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxRoleRevoked represents a RoleRevoked event raised by the TradingBox contract.
type TradingBoxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingBox *TradingBoxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingBoxRoleRevokedIterator, error) {

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

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxRoleRevokedIterator{contract: _TradingBox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingBox *TradingBoxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TradingBoxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxRoleRevoked)
				if err := _TradingBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TradingBox *TradingBoxFilterer) ParseRoleRevoked(log types.Log) (*TradingBoxRoleRevoked, error) {
	event := new(TradingBoxRoleRevoked)
	if err := _TradingBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TradingBox contract.
type TradingBoxTransferIterator struct {
	Event *TradingBoxTransfer // Event containing the contract specifics and raw log

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
func (it *TradingBoxTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxTransfer)
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
		it.Event = new(TradingBoxTransfer)
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
func (it *TradingBoxTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxTransfer represents a Transfer event raised by the TradingBox contract.
type TradingBoxTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TradingBoxTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxTransferIterator{contract: _TradingBox.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TradingBoxTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxTransfer)
				if err := _TradingBox.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TradingBox *TradingBoxFilterer) ParseTransfer(log types.Log) (*TradingBoxTransfer, error) {
	event := new(TradingBoxTransfer)
	if err := _TradingBox.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TradingBox contract.
type TradingBoxUnpausedIterator struct {
	Event *TradingBoxUnpaused // Event containing the contract specifics and raw log

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
func (it *TradingBoxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxUnpaused)
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
		it.Event = new(TradingBoxUnpaused)
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
func (it *TradingBoxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxUnpaused represents a Unpaused event raised by the TradingBox contract.
type TradingBoxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TradingBox *TradingBoxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TradingBoxUnpausedIterator, error) {

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TradingBoxUnpausedIterator{contract: _TradingBox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TradingBox *TradingBoxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TradingBoxUnpaused) (event.Subscription, error) {

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxUnpaused)
				if err := _TradingBox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TradingBox *TradingBoxFilterer) ParseUnpaused(log types.Log) (*TradingBoxUnpaused, error) {
	event := new(TradingBoxUnpaused)
	if err := _TradingBox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingBoxOpenBoxIterator is returned from FilterOpenBox and is used to iterate over the raw logs and unpacked data for OpenBox events raised by the TradingBox contract.
type TradingBoxOpenBoxIterator struct {
	Event *TradingBoxOpenBox // Event containing the contract specifics and raw log

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
func (it *TradingBoxOpenBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingBoxOpenBox)
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
		it.Event = new(TradingBoxOpenBox)
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
func (it *TradingBoxOpenBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingBoxOpenBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingBoxOpenBox represents a OpenBox event raised by the TradingBox contract.
type TradingBoxOpenBox struct {
	TokenId   *big.Int
	TimeStamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOpenBox is a free log retrieval operation binding the contract event 0xf35dea5a32f1d9d37d89b725cbf125aae7e2faa6211c19a87bb844c096671628.
//
// Solidity: event openBox(uint256 indexed tokenId, uint256 indexed timeStamp)
func (_TradingBox *TradingBoxFilterer) FilterOpenBox(opts *bind.FilterOpts, tokenId []*big.Int, timeStamp []*big.Int) (*TradingBoxOpenBoxIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeStampRule []interface{}
	for _, timeStampItem := range timeStamp {
		timeStampRule = append(timeStampRule, timeStampItem)
	}

	logs, sub, err := _TradingBox.contract.FilterLogs(opts, "openBox", tokenIdRule, timeStampRule)
	if err != nil {
		return nil, err
	}
	return &TradingBoxOpenBoxIterator{contract: _TradingBox.contract, event: "openBox", logs: logs, sub: sub}, nil
}

// WatchOpenBox is a free log subscription operation binding the contract event 0xf35dea5a32f1d9d37d89b725cbf125aae7e2faa6211c19a87bb844c096671628.
//
// Solidity: event openBox(uint256 indexed tokenId, uint256 indexed timeStamp)
func (_TradingBox *TradingBoxFilterer) WatchOpenBox(opts *bind.WatchOpts, sink chan<- *TradingBoxOpenBox, tokenId []*big.Int, timeStamp []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var timeStampRule []interface{}
	for _, timeStampItem := range timeStamp {
		timeStampRule = append(timeStampRule, timeStampItem)
	}

	logs, sub, err := _TradingBox.contract.WatchLogs(opts, "openBox", tokenIdRule, timeStampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingBoxOpenBox)
				if err := _TradingBox.contract.UnpackLog(event, "openBox", log); err != nil {
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

// ParseOpenBox is a log parse operation binding the contract event 0xf35dea5a32f1d9d37d89b725cbf125aae7e2faa6211c19a87bb844c096671628.
//
// Solidity: event openBox(uint256 indexed tokenId, uint256 indexed timeStamp)
func (_TradingBox *TradingBoxFilterer) ParseOpenBox(log types.Log) (*TradingBoxOpenBox, error) {
	event := new(TradingBoxOpenBox)
	if err := _TradingBox.contract.UnpackLog(event, "openBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
