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

// UserCredentialsBscMetaData contains all meta data concerning the UserCredentialsBsc contract.
var UserCredentialsBscMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"apiKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"apiSecret\",\"type\":\"string\"}],\"name\":\"setUserCredentialsBsc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCredentialsBsc\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserCredentialsBscBySystemRole\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLengthBySystemRole\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndexAndSystemRole\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsersBySystemRole\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UserCredentialsBscABI is the input ABI used to generate the binding from.
// Deprecated: Use UserCredentialsBscMetaData.ABI instead.
var UserCredentialsBscABI = UserCredentialsBscMetaData.ABI

// UserCredentialsBsc is an auto generated Go binding around an Ethereum contract.
type UserCredentialsBsc struct {
	UserCredentialsBscCaller     // Read-only binding to the contract
	UserCredentialsBscTransactor // Write-only binding to the contract
	UserCredentialsBscFilterer   // Log filterer for contract events
}

// UserCredentialsBscCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCredentialsBscCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCredentialsBscTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserCredentialsBscTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCredentialsBscFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserCredentialsBscFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserCredentialsBscSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserCredentialsBscSession struct {
	Contract     *UserCredentialsBsc // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UserCredentialsBscCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCredentialsBscCallerSession struct {
	Contract *UserCredentialsBscCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// UserCredentialsBscTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserCredentialsBscTransactorSession struct {
	Contract     *UserCredentialsBscTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UserCredentialsBscRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserCredentialsBscRaw struct {
	Contract *UserCredentialsBsc // Generic contract binding to access the raw methods on
}

// UserCredentialsBscCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCredentialsBscCallerRaw struct {
	Contract *UserCredentialsBscCaller // Generic read-only contract binding to access the raw methods on
}

// UserCredentialsBscTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserCredentialsBscTransactorRaw struct {
	Contract *UserCredentialsBscTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserCredentialsBsc creates a new instance of UserCredentialsBsc, bound to a specific deployed contract.
func NewUserCredentialsBsc(address common.Address, backend bind.ContractBackend) (*UserCredentialsBsc, error) {
	contract, err := bindUserCredentialsBsc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBsc{UserCredentialsBscCaller: UserCredentialsBscCaller{contract: contract}, UserCredentialsBscTransactor: UserCredentialsBscTransactor{contract: contract}, UserCredentialsBscFilterer: UserCredentialsBscFilterer{contract: contract}}, nil
}

// NewUserCredentialsBscCaller creates a new read-only instance of UserCredentialsBsc, bound to a specific deployed contract.
func NewUserCredentialsBscCaller(address common.Address, caller bind.ContractCaller) (*UserCredentialsBscCaller, error) {
	contract, err := bindUserCredentialsBsc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscCaller{contract: contract}, nil
}

// NewUserCredentialsBscTransactor creates a new write-only instance of UserCredentialsBsc, bound to a specific deployed contract.
func NewUserCredentialsBscTransactor(address common.Address, transactor bind.ContractTransactor) (*UserCredentialsBscTransactor, error) {
	contract, err := bindUserCredentialsBsc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscTransactor{contract: contract}, nil
}

// NewUserCredentialsBscFilterer creates a new log filterer instance of UserCredentialsBsc, bound to a specific deployed contract.
func NewUserCredentialsBscFilterer(address common.Address, filterer bind.ContractFilterer) (*UserCredentialsBscFilterer, error) {
	contract, err := bindUserCredentialsBsc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscFilterer{contract: contract}, nil
}

// bindUserCredentialsBsc binds a generic wrapper to an already deployed contract.
func bindUserCredentialsBsc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserCredentialsBscABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserCredentialsBsc *UserCredentialsBscRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserCredentialsBsc.Contract.UserCredentialsBscCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserCredentialsBsc *UserCredentialsBscRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.UserCredentialsBscTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserCredentialsBsc *UserCredentialsBscRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.UserCredentialsBscTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserCredentialsBsc *UserCredentialsBscCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserCredentialsBsc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserCredentialsBsc *UserCredentialsBscTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserCredentialsBsc *UserCredentialsBscTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UserCredentialsBsc.Contract.DEFAULTADMINROLE(&_UserCredentialsBsc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UserCredentialsBsc.Contract.DEFAULTADMINROLE(&_UserCredentialsBsc.CallOpts)
}

// SYSTEMROLE is a free data retrieval call binding the contract method 0x75071d2a.
//
// Solidity: function SYSTEM_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCaller) SYSTEMROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "SYSTEM_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SYSTEMROLE is a free data retrieval call binding the contract method 0x75071d2a.
//
// Solidity: function SYSTEM_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscSession) SYSTEMROLE() ([32]byte, error) {
	return _UserCredentialsBsc.Contract.SYSTEMROLE(&_UserCredentialsBsc.CallOpts)
}

// SYSTEMROLE is a free data retrieval call binding the contract method 0x75071d2a.
//
// Solidity: function SYSTEM_ROLE() view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) SYSTEMROLE() ([32]byte, error) {
	return _UserCredentialsBsc.Contract.SYSTEMROLE(&_UserCredentialsBsc.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UserCredentialsBsc.Contract.GetRoleAdmin(&_UserCredentialsBsc.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UserCredentialsBsc.Contract.GetRoleAdmin(&_UserCredentialsBsc.CallOpts, role)
}

// GetUserCredentialsBsc is a free data retrieval call binding the contract method 0x16d9160f.
//
// Solidity: function getUserCredentialsBsc() view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetUserCredentialsBsc(opts *bind.CallOpts) (string, string, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getUserCredentialsBsc")

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetUserCredentialsBsc is a free data retrieval call binding the contract method 0x16d9160f.
//
// Solidity: function getUserCredentialsBsc() view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscSession) GetUserCredentialsBsc() (string, string, error) {
	return _UserCredentialsBsc.Contract.GetUserCredentialsBsc(&_UserCredentialsBsc.CallOpts)
}

// GetUserCredentialsBsc is a free data retrieval call binding the contract method 0x16d9160f.
//
// Solidity: function getUserCredentialsBsc() view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetUserCredentialsBsc() (string, string, error) {
	return _UserCredentialsBsc.Contract.GetUserCredentialsBsc(&_UserCredentialsBsc.CallOpts)
}

// GetUserCredentialsBscBySystemRole is a free data retrieval call binding the contract method 0xfc01685c.
//
// Solidity: function getUserCredentialsBscBySystemRole(address account) view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetUserCredentialsBscBySystemRole(opts *bind.CallOpts, account common.Address) (string, string, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getUserCredentialsBscBySystemRole", account)

	if err != nil {
		return *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetUserCredentialsBscBySystemRole is a free data retrieval call binding the contract method 0xfc01685c.
//
// Solidity: function getUserCredentialsBscBySystemRole(address account) view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscSession) GetUserCredentialsBscBySystemRole(account common.Address) (string, string, error) {
	return _UserCredentialsBsc.Contract.GetUserCredentialsBscBySystemRole(&_UserCredentialsBsc.CallOpts, account)
}

// GetUserCredentialsBscBySystemRole is a free data retrieval call binding the contract method 0xfc01685c.
//
// Solidity: function getUserCredentialsBscBySystemRole(address account) view returns(string, string)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetUserCredentialsBscBySystemRole(account common.Address) (string, string, error) {
	return _UserCredentialsBsc.Contract.GetUserCredentialsBscBySystemRole(&_UserCredentialsBsc.CallOpts, account)
}

// GetUserLengthBySystemRole is a free data retrieval call binding the contract method 0xbc0aa065.
//
// Solidity: function getUserLengthBySystemRole() view returns(uint256)
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetUserLengthBySystemRole(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getUserLengthBySystemRole")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLengthBySystemRole is a free data retrieval call binding the contract method 0xbc0aa065.
//
// Solidity: function getUserLengthBySystemRole() view returns(uint256)
func (_UserCredentialsBsc *UserCredentialsBscSession) GetUserLengthBySystemRole() (*big.Int, error) {
	return _UserCredentialsBsc.Contract.GetUserLengthBySystemRole(&_UserCredentialsBsc.CallOpts)
}

// GetUserLengthBySystemRole is a free data retrieval call binding the contract method 0xbc0aa065.
//
// Solidity: function getUserLengthBySystemRole() view returns(uint256)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetUserLengthBySystemRole() (*big.Int, error) {
	return _UserCredentialsBsc.Contract.GetUserLengthBySystemRole(&_UserCredentialsBsc.CallOpts)
}

// GetUsersByIndexAndSystemRole is a free data retrieval call binding the contract method 0x8544ad44.
//
// Solidity: function getUsersByIndexAndSystemRole(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetUsersByIndexAndSystemRole(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getUsersByIndexAndSystemRole", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndexAndSystemRole is a free data retrieval call binding the contract method 0x8544ad44.
//
// Solidity: function getUsersByIndexAndSystemRole(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscSession) GetUsersByIndexAndSystemRole(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _UserCredentialsBsc.Contract.GetUsersByIndexAndSystemRole(&_UserCredentialsBsc.CallOpts, startIndex, endIndex)
}

// GetUsersByIndexAndSystemRole is a free data retrieval call binding the contract method 0x8544ad44.
//
// Solidity: function getUsersByIndexAndSystemRole(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetUsersByIndexAndSystemRole(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _UserCredentialsBsc.Contract.GetUsersByIndexAndSystemRole(&_UserCredentialsBsc.CallOpts, startIndex, endIndex)
}

// GetUsersBySystemRole is a free data retrieval call binding the contract method 0xb99f0f54.
//
// Solidity: function getUsersBySystemRole() view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscCaller) GetUsersBySystemRole(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "getUsersBySystemRole")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersBySystemRole is a free data retrieval call binding the contract method 0xb99f0f54.
//
// Solidity: function getUsersBySystemRole() view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscSession) GetUsersBySystemRole() ([]common.Address, error) {
	return _UserCredentialsBsc.Contract.GetUsersBySystemRole(&_UserCredentialsBsc.CallOpts)
}

// GetUsersBySystemRole is a free data retrieval call binding the contract method 0xb99f0f54.
//
// Solidity: function getUsersBySystemRole() view returns(address[])
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) GetUsersBySystemRole() ([]common.Address, error) {
	return _UserCredentialsBsc.Contract.GetUsersBySystemRole(&_UserCredentialsBsc.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UserCredentialsBsc.Contract.HasRole(&_UserCredentialsBsc.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UserCredentialsBsc.Contract.HasRole(&_UserCredentialsBsc.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UserCredentialsBsc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UserCredentialsBsc.Contract.SupportsInterface(&_UserCredentialsBsc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UserCredentialsBsc *UserCredentialsBscCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UserCredentialsBsc.Contract.SupportsInterface(&_UserCredentialsBsc.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.GrantRole(&_UserCredentialsBsc.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.GrantRole(&_UserCredentialsBsc.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UserCredentialsBsc *UserCredentialsBscSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.RenounceRole(&_UserCredentialsBsc.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.RenounceRole(&_UserCredentialsBsc.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.RevokeRole(&_UserCredentialsBsc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.RevokeRole(&_UserCredentialsBsc.TransactOpts, role, account)
}

// SetUserCredentialsBsc is a paid mutator transaction binding the contract method 0x228b65b5.
//
// Solidity: function setUserCredentialsBsc(string apiKey, string apiSecret) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactor) SetUserCredentialsBsc(opts *bind.TransactOpts, apiKey string, apiSecret string) (*types.Transaction, error) {
	return _UserCredentialsBsc.contract.Transact(opts, "setUserCredentialsBsc", apiKey, apiSecret)
}

// SetUserCredentialsBsc is a paid mutator transaction binding the contract method 0x228b65b5.
//
// Solidity: function setUserCredentialsBsc(string apiKey, string apiSecret) returns()
func (_UserCredentialsBsc *UserCredentialsBscSession) SetUserCredentialsBsc(apiKey string, apiSecret string) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.SetUserCredentialsBsc(&_UserCredentialsBsc.TransactOpts, apiKey, apiSecret)
}

// SetUserCredentialsBsc is a paid mutator transaction binding the contract method 0x228b65b5.
//
// Solidity: function setUserCredentialsBsc(string apiKey, string apiSecret) returns()
func (_UserCredentialsBsc *UserCredentialsBscTransactorSession) SetUserCredentialsBsc(apiKey string, apiSecret string) (*types.Transaction, error) {
	return _UserCredentialsBsc.Contract.SetUserCredentialsBsc(&_UserCredentialsBsc.TransactOpts, apiKey, apiSecret)
}

// UserCredentialsBscRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleAdminChangedIterator struct {
	Event *UserCredentialsBscRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UserCredentialsBscRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserCredentialsBscRoleAdminChanged)
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
		it.Event = new(UserCredentialsBscRoleAdminChanged)
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
func (it *UserCredentialsBscRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserCredentialsBscRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserCredentialsBscRoleAdminChanged represents a RoleAdminChanged event raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UserCredentialsBscRoleAdminChangedIterator, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscRoleAdminChangedIterator{contract: _UserCredentialsBsc.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UserCredentialsBscRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserCredentialsBscRoleAdminChanged)
				if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_UserCredentialsBsc *UserCredentialsBscFilterer) ParseRoleAdminChanged(log types.Log) (*UserCredentialsBscRoleAdminChanged, error) {
	event := new(UserCredentialsBscRoleAdminChanged)
	if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserCredentialsBscRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleGrantedIterator struct {
	Event *UserCredentialsBscRoleGranted // Event containing the contract specifics and raw log

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
func (it *UserCredentialsBscRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserCredentialsBscRoleGranted)
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
		it.Event = new(UserCredentialsBscRoleGranted)
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
func (it *UserCredentialsBscRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserCredentialsBscRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserCredentialsBscRoleGranted represents a RoleGranted event raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UserCredentialsBscRoleGrantedIterator, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscRoleGrantedIterator{contract: _UserCredentialsBsc.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UserCredentialsBscRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserCredentialsBscRoleGranted)
				if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_UserCredentialsBsc *UserCredentialsBscFilterer) ParseRoleGranted(log types.Log) (*UserCredentialsBscRoleGranted, error) {
	event := new(UserCredentialsBscRoleGranted)
	if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserCredentialsBscRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleRevokedIterator struct {
	Event *UserCredentialsBscRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UserCredentialsBscRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserCredentialsBscRoleRevoked)
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
		it.Event = new(UserCredentialsBscRoleRevoked)
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
func (it *UserCredentialsBscRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserCredentialsBscRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserCredentialsBscRoleRevoked represents a RoleRevoked event raised by the UserCredentialsBsc contract.
type UserCredentialsBscRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UserCredentialsBscRoleRevokedIterator, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UserCredentialsBscRoleRevokedIterator{contract: _UserCredentialsBsc.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UserCredentialsBsc *UserCredentialsBscFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UserCredentialsBscRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _UserCredentialsBsc.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserCredentialsBscRoleRevoked)
				if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_UserCredentialsBsc *UserCredentialsBscFilterer) ParseRoleRevoked(log types.Log) (*UserCredentialsBscRoleRevoked, error) {
	event := new(UserCredentialsBscRoleRevoked)
	if err := _UserCredentialsBsc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
