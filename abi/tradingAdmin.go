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

// TradingAdminMetaData contains all meta data concerning the TradingAdmin contract.
var TradingAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradingBox_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"term_\",\"type\":\"uint256\"}],\"name\":\"getTermAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradingTerm\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"termAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradingBox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tradingTerm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pushTerm\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"termAmount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TradingAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingAdminMetaData.ABI instead.
var TradingAdminABI = TradingAdminMetaData.ABI

// TradingAdmin is an auto generated Go binding around an Ethereum contract.
type TradingAdmin struct {
	TradingAdminCaller     // Read-only binding to the contract
	TradingAdminTransactor // Write-only binding to the contract
	TradingAdminFilterer   // Log filterer for contract events
}

// TradingAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingAdminSession struct {
	Contract     *TradingAdmin     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingAdminCallerSession struct {
	Contract *TradingAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TradingAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingAdminTransactorSession struct {
	Contract     *TradingAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TradingAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingAdminRaw struct {
	Contract *TradingAdmin // Generic contract binding to access the raw methods on
}

// TradingAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingAdminCallerRaw struct {
	Contract *TradingAdminCaller // Generic read-only contract binding to access the raw methods on
}

// TradingAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingAdminTransactorRaw struct {
	Contract *TradingAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingAdmin creates a new instance of TradingAdmin, bound to a specific deployed contract.
func NewTradingAdmin(address common.Address, backend bind.ContractBackend) (*TradingAdmin, error) {
	contract, err := bindTradingAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradingAdmin{TradingAdminCaller: TradingAdminCaller{contract: contract}, TradingAdminTransactor: TradingAdminTransactor{contract: contract}, TradingAdminFilterer: TradingAdminFilterer{contract: contract}}, nil
}

// NewTradingAdminCaller creates a new read-only instance of TradingAdmin, bound to a specific deployed contract.
func NewTradingAdminCaller(address common.Address, caller bind.ContractCaller) (*TradingAdminCaller, error) {
	contract, err := bindTradingAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingAdminCaller{contract: contract}, nil
}

// NewTradingAdminTransactor creates a new write-only instance of TradingAdmin, bound to a specific deployed contract.
func NewTradingAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingAdminTransactor, error) {
	contract, err := bindTradingAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingAdminTransactor{contract: contract}, nil
}

// NewTradingAdminFilterer creates a new log filterer instance of TradingAdmin, bound to a specific deployed contract.
func NewTradingAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingAdminFilterer, error) {
	contract, err := bindTradingAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingAdminFilterer{contract: contract}, nil
}

// bindTradingAdmin binds a generic wrapper to an already deployed contract.
func bindTradingAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradingAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingAdmin *TradingAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingAdmin.Contract.TradingAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingAdmin *TradingAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAdmin.Contract.TradingAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingAdmin *TradingAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingAdmin.Contract.TradingAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingAdmin *TradingAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingAdmin *TradingAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingAdmin *TradingAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingAdmin.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingAdmin *TradingAdminCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingAdmin *TradingAdminSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingAdmin.Contract.DEFAULTADMINROLE(&_TradingAdmin.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingAdmin *TradingAdminCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingAdmin.Contract.DEFAULTADMINROLE(&_TradingAdmin.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingAdmin *TradingAdminCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingAdmin *TradingAdminSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingAdmin.Contract.GetRoleAdmin(&_TradingAdmin.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingAdmin *TradingAdminCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingAdmin.Contract.GetRoleAdmin(&_TradingAdmin.CallOpts, role)
}

// GetTermAmount is a free data retrieval call binding the contract method 0xc794edcd.
//
// Solidity: function getTermAmount(uint256 term_) view returns(uint256)
func (_TradingAdmin *TradingAdminCaller) GetTermAmount(opts *bind.CallOpts, term_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "getTermAmount", term_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTermAmount is a free data retrieval call binding the contract method 0xc794edcd.
//
// Solidity: function getTermAmount(uint256 term_) view returns(uint256)
func (_TradingAdmin *TradingAdminSession) GetTermAmount(term_ *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.GetTermAmount(&_TradingAdmin.CallOpts, term_)
}

// GetTermAmount is a free data retrieval call binding the contract method 0xc794edcd.
//
// Solidity: function getTermAmount(uint256 term_) view returns(uint256)
func (_TradingAdmin *TradingAdminCallerSession) GetTermAmount(term_ *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.GetTermAmount(&_TradingAdmin.CallOpts, term_)
}

// GetTradingTerm is a free data retrieval call binding the contract method 0x7594453d.
//
// Solidity: function getTradingTerm() view returns(uint256[])
func (_TradingAdmin *TradingAdminCaller) GetTradingTerm(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "getTradingTerm")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTradingTerm is a free data retrieval call binding the contract method 0x7594453d.
//
// Solidity: function getTradingTerm() view returns(uint256[])
func (_TradingAdmin *TradingAdminSession) GetTradingTerm() ([]*big.Int, error) {
	return _TradingAdmin.Contract.GetTradingTerm(&_TradingAdmin.CallOpts)
}

// GetTradingTerm is a free data retrieval call binding the contract method 0x7594453d.
//
// Solidity: function getTradingTerm() view returns(uint256[])
func (_TradingAdmin *TradingAdminCallerSession) GetTradingTerm() ([]*big.Int, error) {
	return _TradingAdmin.Contract.GetTradingTerm(&_TradingAdmin.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingAdmin *TradingAdminCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingAdmin *TradingAdminSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingAdmin.Contract.HasRole(&_TradingAdmin.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingAdmin *TradingAdminCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingAdmin.Contract.HasRole(&_TradingAdmin.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingAdmin *TradingAdminCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingAdmin *TradingAdminSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingAdmin.Contract.SupportsInterface(&_TradingAdmin.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingAdmin *TradingAdminCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingAdmin.Contract.SupportsInterface(&_TradingAdmin.CallOpts, interfaceId)
}

// TermAmount is a free data retrieval call binding the contract method 0x1a788c0e.
//
// Solidity: function termAmount(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminCaller) TermAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "termAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TermAmount is a free data retrieval call binding the contract method 0x1a788c0e.
//
// Solidity: function termAmount(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminSession) TermAmount(arg0 *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.TermAmount(&_TradingAdmin.CallOpts, arg0)
}

// TermAmount is a free data retrieval call binding the contract method 0x1a788c0e.
//
// Solidity: function termAmount(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminCallerSession) TermAmount(arg0 *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.TermAmount(&_TradingAdmin.CallOpts, arg0)
}

// TradingBox is a free data retrieval call binding the contract method 0x5040d78d.
//
// Solidity: function tradingBox() view returns(address)
func (_TradingAdmin *TradingAdminCaller) TradingBox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "tradingBox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TradingBox is a free data retrieval call binding the contract method 0x5040d78d.
//
// Solidity: function tradingBox() view returns(address)
func (_TradingAdmin *TradingAdminSession) TradingBox() (common.Address, error) {
	return _TradingAdmin.Contract.TradingBox(&_TradingAdmin.CallOpts)
}

// TradingBox is a free data retrieval call binding the contract method 0x5040d78d.
//
// Solidity: function tradingBox() view returns(address)
func (_TradingAdmin *TradingAdminCallerSession) TradingBox() (common.Address, error) {
	return _TradingAdmin.Contract.TradingBox(&_TradingAdmin.CallOpts)
}

// TradingTerm is a free data retrieval call binding the contract method 0xc5a56969.
//
// Solidity: function tradingTerm(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminCaller) TradingTerm(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "tradingTerm", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TradingTerm is a free data retrieval call binding the contract method 0xc5a56969.
//
// Solidity: function tradingTerm(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminSession) TradingTerm(arg0 *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.TradingTerm(&_TradingAdmin.CallOpts, arg0)
}

// TradingTerm is a free data retrieval call binding the contract method 0xc5a56969.
//
// Solidity: function tradingTerm(uint256 ) view returns(uint256)
func (_TradingAdmin *TradingAdminCallerSession) TradingTerm(arg0 *big.Int) (*big.Int, error) {
	return _TradingAdmin.Contract.TradingTerm(&_TradingAdmin.CallOpts, arg0)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingAdmin *TradingAdminCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradingAdmin.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingAdmin *TradingAdminSession) Usdt() (common.Address, error) {
	return _TradingAdmin.Contract.Usdt(&_TradingAdmin.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_TradingAdmin *TradingAdminCallerSession) Usdt() (common.Address, error) {
	return _TradingAdmin.Contract.Usdt(&_TradingAdmin.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.GrantRole(&_TradingAdmin.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.GrantRole(&_TradingAdmin.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingAdmin *TradingAdminTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingAdmin.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingAdmin *TradingAdminSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.RenounceRole(&_TradingAdmin.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingAdmin *TradingAdminTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.RenounceRole(&_TradingAdmin.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.RevokeRole(&_TradingAdmin.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingAdmin *TradingAdminTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingAdmin.Contract.RevokeRole(&_TradingAdmin.TransactOpts, role, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc393bbce.
//
// Solidity: function withdraw(address account, uint256 amount, bool pushTerm, uint256 termAmount_) returns()
func (_TradingAdmin *TradingAdminTransactor) Withdraw(opts *bind.TransactOpts, account common.Address, amount *big.Int, pushTerm bool, termAmount_ *big.Int) (*types.Transaction, error) {
	return _TradingAdmin.contract.Transact(opts, "withdraw", account, amount, pushTerm, termAmount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc393bbce.
//
// Solidity: function withdraw(address account, uint256 amount, bool pushTerm, uint256 termAmount_) returns()
func (_TradingAdmin *TradingAdminSession) Withdraw(account common.Address, amount *big.Int, pushTerm bool, termAmount_ *big.Int) (*types.Transaction, error) {
	return _TradingAdmin.Contract.Withdraw(&_TradingAdmin.TransactOpts, account, amount, pushTerm, termAmount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc393bbce.
//
// Solidity: function withdraw(address account, uint256 amount, bool pushTerm, uint256 termAmount_) returns()
func (_TradingAdmin *TradingAdminTransactorSession) Withdraw(account common.Address, amount *big.Int, pushTerm bool, termAmount_ *big.Int) (*types.Transaction, error) {
	return _TradingAdmin.Contract.Withdraw(&_TradingAdmin.TransactOpts, account, amount, pushTerm, termAmount_)
}

// TradingAdminRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TradingAdmin contract.
type TradingAdminRoleAdminChangedIterator struct {
	Event *TradingAdminRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TradingAdminRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAdminRoleAdminChanged)
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
		it.Event = new(TradingAdminRoleAdminChanged)
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
func (it *TradingAdminRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAdminRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAdminRoleAdminChanged represents a RoleAdminChanged event raised by the TradingAdmin contract.
type TradingAdminRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingAdmin *TradingAdminFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TradingAdminRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TradingAdmin.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TradingAdminRoleAdminChangedIterator{contract: _TradingAdmin.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingAdmin *TradingAdminFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TradingAdminRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TradingAdmin.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAdminRoleAdminChanged)
				if err := _TradingAdmin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TradingAdmin *TradingAdminFilterer) ParseRoleAdminChanged(log types.Log) (*TradingAdminRoleAdminChanged, error) {
	event := new(TradingAdminRoleAdminChanged)
	if err := _TradingAdmin.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAdminRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TradingAdmin contract.
type TradingAdminRoleGrantedIterator struct {
	Event *TradingAdminRoleGranted // Event containing the contract specifics and raw log

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
func (it *TradingAdminRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAdminRoleGranted)
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
		it.Event = new(TradingAdminRoleGranted)
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
func (it *TradingAdminRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAdminRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAdminRoleGranted represents a RoleGranted event raised by the TradingAdmin contract.
type TradingAdminRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingAdmin *TradingAdminFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingAdminRoleGrantedIterator, error) {

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

	logs, sub, err := _TradingAdmin.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingAdminRoleGrantedIterator{contract: _TradingAdmin.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingAdmin *TradingAdminFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TradingAdminRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingAdmin.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAdminRoleGranted)
				if err := _TradingAdmin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TradingAdmin *TradingAdminFilterer) ParseRoleGranted(log types.Log) (*TradingAdminRoleGranted, error) {
	event := new(TradingAdminRoleGranted)
	if err := _TradingAdmin.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingAdminRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TradingAdmin contract.
type TradingAdminRoleRevokedIterator struct {
	Event *TradingAdminRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TradingAdminRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingAdminRoleRevoked)
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
		it.Event = new(TradingAdminRoleRevoked)
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
func (it *TradingAdminRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingAdminRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingAdminRoleRevoked represents a RoleRevoked event raised by the TradingAdmin contract.
type TradingAdminRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingAdmin *TradingAdminFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingAdminRoleRevokedIterator, error) {

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

	logs, sub, err := _TradingAdmin.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingAdminRoleRevokedIterator{contract: _TradingAdmin.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingAdmin *TradingAdminFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TradingAdminRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingAdmin.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingAdminRoleRevoked)
				if err := _TradingAdmin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TradingAdmin *TradingAdminFilterer) ParseRoleRevoked(log types.Log) (*TradingAdminRoleRevoked, error) {
	event := new(TradingAdminRoleRevoked)
	if err := _TradingAdmin.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
