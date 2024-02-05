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

// DepositMetaData contains all meta data concerning the Deposit contract.
var DepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ttdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tupair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceUsdt\",\"type\":\"uint256\"}],\"name\":\"depositCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tupair\",\"type\":\"address\"}],\"name\":\"setTupair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ttdc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tupair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userdepositForUsdtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DepositABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositMetaData.ABI instead.
var DepositABI = DepositMetaData.ABI

// Deposit is an auto generated Go binding around an Ethereum contract.
type Deposit struct {
	DepositCaller     // Read-only binding to the contract
	DepositTransactor // Write-only binding to the contract
	DepositFilterer   // Log filterer for contract events
}

// DepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositSession struct {
	Contract     *Deposit          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositCallerSession struct {
	Contract *DepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositTransactorSession struct {
	Contract     *DepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositRaw struct {
	Contract *Deposit // Generic contract binding to access the raw methods on
}

// DepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositCallerRaw struct {
	Contract *DepositCaller // Generic read-only contract binding to access the raw methods on
}

// DepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositTransactorRaw struct {
	Contract *DepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeposit creates a new instance of Deposit, bound to a specific deployed contract.
func NewDeposit(address common.Address, backend bind.ContractBackend) (*Deposit, error) {
	contract, err := bindDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deposit{DepositCaller: DepositCaller{contract: contract}, DepositTransactor: DepositTransactor{contract: contract}, DepositFilterer: DepositFilterer{contract: contract}}, nil
}

// NewDepositCaller creates a new read-only instance of Deposit, bound to a specific deployed contract.
func NewDepositCaller(address common.Address, caller bind.ContractCaller) (*DepositCaller, error) {
	contract, err := bindDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositCaller{contract: contract}, nil
}

// NewDepositTransactor creates a new write-only instance of Deposit, bound to a specific deployed contract.
func NewDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositTransactor, error) {
	contract, err := bindDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositTransactor{contract: contract}, nil
}

// NewDepositFilterer creates a new log filterer instance of Deposit, bound to a specific deployed contract.
func NewDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositFilterer, error) {
	contract, err := bindDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositFilterer{contract: contract}, nil
}

// bindDeposit binds a generic wrapper to an already deployed contract.
func bindDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.DepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.DepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deposit *DepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deposit *DepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deposit *DepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deposit.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deposit *DepositCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deposit *DepositSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deposit.Contract.DEFAULTADMINROLE(&_Deposit.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Deposit *DepositCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Deposit.Contract.DEFAULTADMINROLE(&_Deposit.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deposit *DepositCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deposit *DepositSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deposit.Contract.GetRoleAdmin(&_Deposit.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Deposit *DepositCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Deposit.Contract.GetRoleAdmin(&_Deposit.CallOpts, role)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Deposit *DepositCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Deposit *DepositSession) GetUserLength() (*big.Int, error) {
	return _Deposit.Contract.GetUserLength(&_Deposit.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_Deposit *DepositCallerSession) GetUserLength() (*big.Int, error) {
	return _Deposit.Contract.GetUserLength(&_Deposit.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Deposit *DepositCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Deposit *DepositSession) GetUsers() ([]common.Address, error) {
	return _Deposit.Contract.GetUsers(&_Deposit.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_Deposit *DepositCallerSession) GetUsers() ([]common.Address, error) {
	return _Deposit.Contract.GetUsers(&_Deposit.CallOpts)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Deposit *DepositCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Deposit *DepositSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _Deposit.Contract.GetUsersByIndex(&_Deposit.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_Deposit *DepositCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _Deposit.Contract.GetUsersByIndex(&_Deposit.CallOpts, startIndex, endIndex)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deposit *DepositCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deposit *DepositSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deposit.Contract.HasRole(&_Deposit.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Deposit *DepositCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Deposit.Contract.HasRole(&_Deposit.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deposit *DepositCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deposit *DepositSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deposit.Contract.SupportsInterface(&_Deposit.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Deposit *DepositCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Deposit.Contract.SupportsInterface(&_Deposit.CallOpts, interfaceId)
}

// Tdc is a free data retrieval call binding the contract method 0xcca720f0.
//
// Solidity: function tdc() view returns(address)
func (_Deposit *DepositCaller) Tdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "tdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tdc is a free data retrieval call binding the contract method 0xcca720f0.
//
// Solidity: function tdc() view returns(address)
func (_Deposit *DepositSession) Tdc() (common.Address, error) {
	return _Deposit.Contract.Tdc(&_Deposit.CallOpts)
}

// Tdc is a free data retrieval call binding the contract method 0xcca720f0.
//
// Solidity: function tdc() view returns(address)
func (_Deposit *DepositCallerSession) Tdc() (common.Address, error) {
	return _Deposit.Contract.Tdc(&_Deposit.CallOpts)
}

// Ttdc is a free data retrieval call binding the contract method 0xf347553b.
//
// Solidity: function ttdc() view returns(address)
func (_Deposit *DepositCaller) Ttdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "ttdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ttdc is a free data retrieval call binding the contract method 0xf347553b.
//
// Solidity: function ttdc() view returns(address)
func (_Deposit *DepositSession) Ttdc() (common.Address, error) {
	return _Deposit.Contract.Ttdc(&_Deposit.CallOpts)
}

// Ttdc is a free data retrieval call binding the contract method 0xf347553b.
//
// Solidity: function ttdc() view returns(address)
func (_Deposit *DepositCallerSession) Ttdc() (common.Address, error) {
	return _Deposit.Contract.Ttdc(&_Deposit.CallOpts)
}

// Tupair is a free data retrieval call binding the contract method 0xe06210c7.
//
// Solidity: function tupair() view returns(address)
func (_Deposit *DepositCaller) Tupair(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "tupair")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tupair is a free data retrieval call binding the contract method 0xe06210c7.
//
// Solidity: function tupair() view returns(address)
func (_Deposit *DepositSession) Tupair() (common.Address, error) {
	return _Deposit.Contract.Tupair(&_Deposit.CallOpts)
}

// Tupair is a free data retrieval call binding the contract method 0xe06210c7.
//
// Solidity: function tupair() view returns(address)
func (_Deposit *DepositCallerSession) Tupair() (common.Address, error) {
	return _Deposit.Contract.Tupair(&_Deposit.CallOpts)
}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_Deposit *DepositCaller) UserCost(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "userCost", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_Deposit *DepositSession) UserCost(arg0 common.Address) (*big.Int, error) {
	return _Deposit.Contract.UserCost(&_Deposit.CallOpts, arg0)
}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_Deposit *DepositCallerSession) UserCost(arg0 common.Address) (*big.Int, error) {
	return _Deposit.Contract.UserCost(&_Deposit.CallOpts, arg0)
}

// UserdepositForUsdtAmount is a free data retrieval call binding the contract method 0xfda610ef.
//
// Solidity: function userdepositForUsdtAmount(address ) view returns(uint256)
func (_Deposit *DepositCaller) UserdepositForUsdtAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Deposit.contract.Call(opts, &out, "userdepositForUsdtAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserdepositForUsdtAmount is a free data retrieval call binding the contract method 0xfda610ef.
//
// Solidity: function userdepositForUsdtAmount(address ) view returns(uint256)
func (_Deposit *DepositSession) UserdepositForUsdtAmount(arg0 common.Address) (*big.Int, error) {
	return _Deposit.Contract.UserdepositForUsdtAmount(&_Deposit.CallOpts, arg0)
}

// UserdepositForUsdtAmount is a free data retrieval call binding the contract method 0xfda610ef.
//
// Solidity: function userdepositForUsdtAmount(address ) view returns(uint256)
func (_Deposit *DepositCallerSession) UserdepositForUsdtAmount(arg0 common.Address) (*big.Int, error) {
	return _Deposit.Contract.UserdepositForUsdtAmount(&_Deposit.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 cost) returns()
func (_Deposit *DepositTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int, cost *big.Int) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "deposit", token, amount, cost)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 cost) returns()
func (_Deposit *DepositSession) Deposit(token common.Address, amount *big.Int, cost *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Deposit(&_Deposit.TransactOpts, token, amount, cost)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address token, uint256 amount, uint256 cost) returns()
func (_Deposit *DepositTransactorSession) Deposit(token common.Address, amount *big.Int, cost *big.Int) (*types.Transaction, error) {
	return _Deposit.Contract.Deposit(&_Deposit.TransactOpts, token, amount, cost)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deposit *DepositTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deposit *DepositSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.GrantRole(&_Deposit.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Deposit *DepositTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.GrantRole(&_Deposit.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deposit *DepositTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deposit *DepositSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RenounceRole(&_Deposit.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Deposit *DepositTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RenounceRole(&_Deposit.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deposit *DepositTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deposit *DepositSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RevokeRole(&_Deposit.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Deposit *DepositTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.RevokeRole(&_Deposit.TransactOpts, role, account)
}

// SetTupair is a paid mutator transaction binding the contract method 0x4f47a905.
//
// Solidity: function setTupair(address _tupair) returns()
func (_Deposit *DepositTransactor) SetTupair(opts *bind.TransactOpts, _tupair common.Address) (*types.Transaction, error) {
	return _Deposit.contract.Transact(opts, "setTupair", _tupair)
}

// SetTupair is a paid mutator transaction binding the contract method 0x4f47a905.
//
// Solidity: function setTupair(address _tupair) returns()
func (_Deposit *DepositSession) SetTupair(_tupair common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.SetTupair(&_Deposit.TransactOpts, _tupair)
}

// SetTupair is a paid mutator transaction binding the contract method 0x4f47a905.
//
// Solidity: function setTupair(address _tupair) returns()
func (_Deposit *DepositTransactorSession) SetTupair(_tupair common.Address) (*types.Transaction, error) {
	return _Deposit.Contract.SetTupair(&_Deposit.TransactOpts, _tupair)
}

// DepositRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Deposit contract.
type DepositRoleAdminChangedIterator struct {
	Event *DepositRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DepositRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositRoleAdminChanged)
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
		it.Event = new(DepositRoleAdminChanged)
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
func (it *DepositRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositRoleAdminChanged represents a RoleAdminChanged event raised by the Deposit contract.
type DepositRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deposit *DepositFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DepositRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DepositRoleAdminChangedIterator{contract: _Deposit.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Deposit *DepositFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DepositRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositRoleAdminChanged)
				if err := _Deposit.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Deposit *DepositFilterer) ParseRoleAdminChanged(log types.Log) (*DepositRoleAdminChanged, error) {
	event := new(DepositRoleAdminChanged)
	if err := _Deposit.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Deposit contract.
type DepositRoleGrantedIterator struct {
	Event *DepositRoleGranted // Event containing the contract specifics and raw log

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
func (it *DepositRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositRoleGranted)
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
		it.Event = new(DepositRoleGranted)
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
func (it *DepositRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositRoleGranted represents a RoleGranted event raised by the Deposit contract.
type DepositRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deposit *DepositFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DepositRoleGrantedIterator, error) {

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

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DepositRoleGrantedIterator{contract: _Deposit.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deposit *DepositFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DepositRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositRoleGranted)
				if err := _Deposit.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Deposit *DepositFilterer) ParseRoleGranted(log types.Log) (*DepositRoleGranted, error) {
	event := new(DepositRoleGranted)
	if err := _Deposit.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Deposit contract.
type DepositRoleRevokedIterator struct {
	Event *DepositRoleRevoked // Event containing the contract specifics and raw log

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
func (it *DepositRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositRoleRevoked)
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
		it.Event = new(DepositRoleRevoked)
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
func (it *DepositRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositRoleRevoked represents a RoleRevoked event raised by the Deposit contract.
type DepositRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deposit *DepositFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DepositRoleRevokedIterator, error) {

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

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DepositRoleRevokedIterator{contract: _Deposit.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Deposit *DepositFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DepositRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositRoleRevoked)
				if err := _Deposit.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Deposit *DepositFilterer) ParseRoleRevoked(log types.Log) (*DepositRoleRevoked, error) {
	event := new(DepositRoleRevoked)
	if err := _Deposit.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositDepositCompletedIterator is returned from FilterDepositCompleted and is used to iterate over the raw logs and unpacked data for DepositCompleted events raised by the Deposit contract.
type DepositDepositCompletedIterator struct {
	Event *DepositDepositCompleted // Event containing the contract specifics and raw log

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
func (it *DepositDepositCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositDepositCompleted)
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
		it.Event = new(DepositDepositCompleted)
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
func (it *DepositDepositCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositDepositCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositDepositCompleted represents a DepositCompleted event raised by the Deposit contract.
type DepositDepositCompleted struct {
	Account     common.Address
	Token       common.Address
	Amount      *big.Int
	UsdtAmount  *big.Int
	BalanceUsdt *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositCompleted is a free log retrieval operation binding the contract event 0xfe73a39d5b722c45d191d94fb2191bbfab7b3d454082f0546a4284868a158c7b.
//
// Solidity: event depositCompleted(address indexed account, address token, uint256 amount, uint256 usdtAmount, uint256 balanceUsdt)
func (_Deposit *DepositFilterer) FilterDepositCompleted(opts *bind.FilterOpts, account []common.Address) (*DepositDepositCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Deposit.contract.FilterLogs(opts, "depositCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &DepositDepositCompletedIterator{contract: _Deposit.contract, event: "depositCompleted", logs: logs, sub: sub}, nil
}

// WatchDepositCompleted is a free log subscription operation binding the contract event 0xfe73a39d5b722c45d191d94fb2191bbfab7b3d454082f0546a4284868a158c7b.
//
// Solidity: event depositCompleted(address indexed account, address token, uint256 amount, uint256 usdtAmount, uint256 balanceUsdt)
func (_Deposit *DepositFilterer) WatchDepositCompleted(opts *bind.WatchOpts, sink chan<- *DepositDepositCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Deposit.contract.WatchLogs(opts, "depositCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositDepositCompleted)
				if err := _Deposit.contract.UnpackLog(event, "depositCompleted", log); err != nil {
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

// ParseDepositCompleted is a log parse operation binding the contract event 0xfe73a39d5b722c45d191d94fb2191bbfab7b3d454082f0546a4284868a158c7b.
//
// Solidity: event depositCompleted(address indexed account, address token, uint256 amount, uint256 usdtAmount, uint256 balanceUsdt)
func (_Deposit *DepositFilterer) ParseDepositCompleted(log types.Log) (*DepositDepositCompleted, error) {
	event := new(DepositDepositCompleted)
	if err := _Deposit.contract.UnpackLog(event, "depositCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
