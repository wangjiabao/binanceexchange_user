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

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userMaxTime\",\"type\":\"uint256\"}],\"name\":\"closeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userMaxTime\",\"type\":\"uint256\"}],\"name\":\"openCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userMaxTime\",\"type\":\"uint256\"}],\"name\":\"stakeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userMaxTime\",\"type\":\"uint256\"}],\"name\":\"unStakeCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"updateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCurrentOpen\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMaxTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userOpenTokenRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userUsdtAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stake *StakeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stake *StakeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stake.Contract.DEFAULTADMINROLE(&_Stake.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stake *StakeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stake.Contract.DEFAULTADMINROLE(&_Stake.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stake *StakeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stake *StakeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stake.Contract.GetRoleAdmin(&_Stake.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stake *StakeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stake.Contract.GetRoleAdmin(&_Stake.CallOpts, role)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Stake *StakeCaller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Stake *StakeSession) GetTokens() ([]common.Address, error) {
	return _Stake.Contract.GetTokens(&_Stake.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Stake *StakeCallerSession) GetTokens() ([]common.Address, error) {
	return _Stake.Contract.GetTokens(&_Stake.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stake *StakeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stake *StakeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stake.Contract.HasRole(&_Stake.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stake *StakeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stake.Contract.HasRole(&_Stake.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stake.Contract.SupportsInterface(&_Stake.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stake.Contract.SupportsInterface(&_Stake.CallOpts, interfaceId)
}

// UpdateTime is a free data retrieval call binding the contract method 0x976f06a6.
//
// Solidity: function updateTime(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UpdateTime(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "updateTime", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateTime is a free data retrieval call binding the contract method 0x976f06a6.
//
// Solidity: function updateTime(address , address ) view returns(uint256)
func (_Stake *StakeSession) UpdateTime(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UpdateTime(&_Stake.CallOpts, arg0, arg1)
}

// UpdateTime is a free data retrieval call binding the contract method 0x976f06a6.
//
// Solidity: function updateTime(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UpdateTime(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UpdateTime(&_Stake.CallOpts, arg0, arg1)
}

// UserCurrentOpen is a free data retrieval call binding the contract method 0xe1098f14.
//
// Solidity: function userCurrentOpen(address ) view returns(address)
func (_Stake *StakeCaller) UserCurrentOpen(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userCurrentOpen", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserCurrentOpen is a free data retrieval call binding the contract method 0xe1098f14.
//
// Solidity: function userCurrentOpen(address ) view returns(address)
func (_Stake *StakeSession) UserCurrentOpen(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.UserCurrentOpen(&_Stake.CallOpts, arg0)
}

// UserCurrentOpen is a free data retrieval call binding the contract method 0xe1098f14.
//
// Solidity: function userCurrentOpen(address ) view returns(address)
func (_Stake *StakeCallerSession) UserCurrentOpen(arg0 common.Address) (common.Address, error) {
	return _Stake.Contract.UserCurrentOpen(&_Stake.CallOpts, arg0)
}

// UserMaxTime is a free data retrieval call binding the contract method 0xf3f157e1.
//
// Solidity: function userMaxTime(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UserMaxTime(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userMaxTime", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMaxTime is a free data retrieval call binding the contract method 0xf3f157e1.
//
// Solidity: function userMaxTime(address , address ) view returns(uint256)
func (_Stake *StakeSession) UserMaxTime(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserMaxTime(&_Stake.CallOpts, arg0, arg1)
}

// UserMaxTime is a free data retrieval call binding the contract method 0xf3f157e1.
//
// Solidity: function userMaxTime(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UserMaxTime(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserMaxTime(&_Stake.CallOpts, arg0, arg1)
}

// UserOpen is a free data retrieval call binding the contract method 0xa98a7743.
//
// Solidity: function userOpen(address , address ) view returns(bool)
func (_Stake *StakeCaller) UserOpen(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userOpen", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserOpen is a free data retrieval call binding the contract method 0xa98a7743.
//
// Solidity: function userOpen(address , address ) view returns(bool)
func (_Stake *StakeSession) UserOpen(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Stake.Contract.UserOpen(&_Stake.CallOpts, arg0, arg1)
}

// UserOpen is a free data retrieval call binding the contract method 0xa98a7743.
//
// Solidity: function userOpen(address , address ) view returns(bool)
func (_Stake *StakeCallerSession) UserOpen(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Stake.Contract.UserOpen(&_Stake.CallOpts, arg0, arg1)
}

// UserOpenTokenRate is a free data retrieval call binding the contract method 0x3c293ca6.
//
// Solidity: function userOpenTokenRate(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UserOpenTokenRate(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userOpenTokenRate", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserOpenTokenRate is a free data retrieval call binding the contract method 0x3c293ca6.
//
// Solidity: function userOpenTokenRate(address , address ) view returns(uint256)
func (_Stake *StakeSession) UserOpenTokenRate(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserOpenTokenRate(&_Stake.CallOpts, arg0, arg1)
}

// UserOpenTokenRate is a free data retrieval call binding the contract method 0x3c293ca6.
//
// Solidity: function userOpenTokenRate(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UserOpenTokenRate(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserOpenTokenRate(&_Stake.CallOpts, arg0, arg1)
}

// UserStakeAmount is a free data retrieval call binding the contract method 0x80d24a92.
//
// Solidity: function userStakeAmount(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UserStakeAmount(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userStakeAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakeAmount is a free data retrieval call binding the contract method 0x80d24a92.
//
// Solidity: function userStakeAmount(address , address ) view returns(uint256)
func (_Stake *StakeSession) UserStakeAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserStakeAmount(&_Stake.CallOpts, arg0, arg1)
}

// UserStakeAmount is a free data retrieval call binding the contract method 0x80d24a92.
//
// Solidity: function userStakeAmount(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UserStakeAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserStakeAmount(&_Stake.CallOpts, arg0, arg1)
}

// UserUsdtAmount is a free data retrieval call binding the contract method 0x382e48f1.
//
// Solidity: function userUsdtAmount(address , address ) view returns(uint256)
func (_Stake *StakeCaller) UserUsdtAmount(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "userUsdtAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUsdtAmount is a free data retrieval call binding the contract method 0x382e48f1.
//
// Solidity: function userUsdtAmount(address , address ) view returns(uint256)
func (_Stake *StakeSession) UserUsdtAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserUsdtAmount(&_Stake.CallOpts, arg0, arg1)
}

// UserUsdtAmount is a free data retrieval call binding the contract method 0x382e48f1.
//
// Solidity: function userUsdtAmount(address , address ) view returns(uint256)
func (_Stake *StakeCallerSession) UserUsdtAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Stake.Contract.UserUsdtAmount(&_Stake.CallOpts, arg0, arg1)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address token) returns()
func (_Stake *StakeTransactor) Close(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "close", token)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address token) returns()
func (_Stake *StakeSession) Close(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Close(&_Stake.TransactOpts, token)
}

// Close is a paid mutator transaction binding the contract method 0xc74073a1.
//
// Solidity: function close(address token) returns()
func (_Stake *StakeTransactorSession) Close(token common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Close(&_Stake.TransactOpts, token)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stake *StakeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stake *StakeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.Contract.GrantRole(&_Stake.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stake *StakeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.Contract.GrantRole(&_Stake.TransactOpts, role, account)
}

// Open is a paid mutator transaction binding the contract method 0x149e7cb9.
//
// Solidity: function open(uint256 usdtAmount, address token, address token0) returns()
func (_Stake *StakeTransactor) Open(opts *bind.TransactOpts, usdtAmount *big.Int, token common.Address, token0 common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "open", usdtAmount, token, token0)
}

// Open is a paid mutator transaction binding the contract method 0x149e7cb9.
//
// Solidity: function open(uint256 usdtAmount, address token, address token0) returns()
func (_Stake *StakeSession) Open(usdtAmount *big.Int, token common.Address, token0 common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Open(&_Stake.TransactOpts, usdtAmount, token, token0)
}

// Open is a paid mutator transaction binding the contract method 0x149e7cb9.
//
// Solidity: function open(uint256 usdtAmount, address token, address token0) returns()
func (_Stake *StakeTransactorSession) Open(usdtAmount *big.Int, token common.Address, token0 common.Address) (*types.Transaction, error) {
	return _Stake.Contract.Open(&_Stake.TransactOpts, usdtAmount, token, token0)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stake *StakeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stake *StakeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RenounceRole(&_Stake.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Stake *StakeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RenounceRole(&_Stake.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stake *StakeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stake *StakeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RevokeRole(&_Stake.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stake *StakeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stake.Contract.RevokeRole(&_Stake.TransactOpts, role, account)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool add) returns()
func (_Stake *StakeTransactor) SetToken(opts *bind.TransactOpts, token common.Address, add bool) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setToken", token, add)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool add) returns()
func (_Stake *StakeSession) SetToken(token common.Address, add bool) (*types.Transaction, error) {
	return _Stake.Contract.SetToken(&_Stake.TransactOpts, token, add)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address token, bool add) returns()
func (_Stake *StakeTransactorSession) SetToken(token common.Address, add bool) (*types.Transaction, error) {
	return _Stake.Contract.SetToken(&_Stake.TransactOpts, token, add)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeTransactor) Stake(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "stake", token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeSession) Stake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address token, uint256 amount) returns()
func (_Stake *StakeTransactorSession) Stake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Stake(&_Stake.TransactOpts, token, amount)
}

// UnStake is a paid mutator transaction binding the contract method 0xd9393814.
//
// Solidity: function unStake(address token, uint256 amount) returns()
func (_Stake *StakeTransactor) UnStake(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "unStake", token, amount)
}

// UnStake is a paid mutator transaction binding the contract method 0xd9393814.
//
// Solidity: function unStake(address token, uint256 amount) returns()
func (_Stake *StakeSession) UnStake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.UnStake(&_Stake.TransactOpts, token, amount)
}

// UnStake is a paid mutator transaction binding the contract method 0xd9393814.
//
// Solidity: function unStake(address token, uint256 amount) returns()
func (_Stake *StakeTransactorSession) UnStake(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.UnStake(&_Stake.TransactOpts, token, amount)
}

// StakeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stake contract.
type StakeRoleAdminChangedIterator struct {
	Event *StakeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRoleAdminChanged)
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
		it.Event = new(StakeRoleAdminChanged)
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
func (it *StakeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRoleAdminChanged represents a RoleAdminChanged event raised by the Stake contract.
type StakeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stake *StakeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Stake.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakeRoleAdminChangedIterator{contract: _Stake.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stake *StakeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Stake.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRoleAdminChanged)
				if err := _Stake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Stake *StakeFilterer) ParseRoleAdminChanged(log types.Log) (*StakeRoleAdminChanged, error) {
	event := new(StakeRoleAdminChanged)
	if err := _Stake.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stake contract.
type StakeRoleGrantedIterator struct {
	Event *StakeRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRoleGranted)
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
		it.Event = new(StakeRoleGranted)
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
func (it *StakeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRoleGranted represents a RoleGranted event raised by the Stake contract.
type StakeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stake *StakeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakeRoleGrantedIterator, error) {

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

	logs, sub, err := _Stake.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakeRoleGrantedIterator{contract: _Stake.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stake *StakeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stake.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRoleGranted)
				if err := _Stake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Stake *StakeFilterer) ParseRoleGranted(log types.Log) (*StakeRoleGranted, error) {
	event := new(StakeRoleGranted)
	if err := _Stake.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stake contract.
type StakeRoleRevokedIterator struct {
	Event *StakeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeRoleRevoked)
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
		it.Event = new(StakeRoleRevoked)
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
func (it *StakeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeRoleRevoked represents a RoleRevoked event raised by the Stake contract.
type StakeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stake *StakeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakeRoleRevokedIterator, error) {

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

	logs, sub, err := _Stake.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakeRoleRevokedIterator{contract: _Stake.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stake *StakeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stake.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeRoleRevoked)
				if err := _Stake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Stake *StakeFilterer) ParseRoleRevoked(log types.Log) (*StakeRoleRevoked, error) {
	event := new(StakeRoleRevoked)
	if err := _Stake.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeCloseCompletedIterator is returned from FilterCloseCompleted and is used to iterate over the raw logs and unpacked data for CloseCompleted events raised by the Stake contract.
type StakeCloseCompletedIterator struct {
	Event *StakeCloseCompleted // Event containing the contract specifics and raw log

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
func (it *StakeCloseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeCloseCompleted)
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
		it.Event = new(StakeCloseCompleted)
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
func (it *StakeCloseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeCloseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeCloseCompleted represents a CloseCompleted event raised by the Stake contract.
type StakeCloseCompleted struct {
	Account     common.Address
	Token       common.Address
	UserMaxTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCloseCompleted is a free log retrieval operation binding the contract event 0xea0d9b8d6dd79197e355e14c6cad606ca1a03a643fcb79d7b0562349d98b40b1.
//
// Solidity: event closeCompleted(address indexed account, address token, uint256 userMaxTime)
func (_Stake *StakeFilterer) FilterCloseCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeCloseCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "closeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeCloseCompletedIterator{contract: _Stake.contract, event: "closeCompleted", logs: logs, sub: sub}, nil
}

// WatchCloseCompleted is a free log subscription operation binding the contract event 0xea0d9b8d6dd79197e355e14c6cad606ca1a03a643fcb79d7b0562349d98b40b1.
//
// Solidity: event closeCompleted(address indexed account, address token, uint256 userMaxTime)
func (_Stake *StakeFilterer) WatchCloseCompleted(opts *bind.WatchOpts, sink chan<- *StakeCloseCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "closeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeCloseCompleted)
				if err := _Stake.contract.UnpackLog(event, "closeCompleted", log); err != nil {
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

// ParseCloseCompleted is a log parse operation binding the contract event 0xea0d9b8d6dd79197e355e14c6cad606ca1a03a643fcb79d7b0562349d98b40b1.
//
// Solidity: event closeCompleted(address indexed account, address token, uint256 userMaxTime)
func (_Stake *StakeFilterer) ParseCloseCompleted(log types.Log) (*StakeCloseCompleted, error) {
	event := new(StakeCloseCompleted)
	if err := _Stake.contract.UnpackLog(event, "closeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeOpenCompletedIterator is returned from FilterOpenCompleted and is used to iterate over the raw logs and unpacked data for OpenCompleted events raised by the Stake contract.
type StakeOpenCompletedIterator struct {
	Event *StakeOpenCompleted // Event containing the contract specifics and raw log

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
func (it *StakeOpenCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOpenCompleted)
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
		it.Event = new(StakeOpenCompleted)
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
func (it *StakeOpenCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOpenCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOpenCompleted represents a OpenCompleted event raised by the Stake contract.
type StakeOpenCompleted struct {
	Account     common.Address
	Token       common.Address
	Token0      common.Address
	UsdtAmount  *big.Int
	UserMaxTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpenCompleted is a free log retrieval operation binding the contract event 0x26f26478175d93f68a0289cff1e2e3d1d4969d3eaf5bd0421a215764c1a7267b.
//
// Solidity: event openCompleted(address indexed account, address token, address token0, uint256 usdtAmount, uint256 userMaxTime)
func (_Stake *StakeFilterer) FilterOpenCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeOpenCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "openCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeOpenCompletedIterator{contract: _Stake.contract, event: "openCompleted", logs: logs, sub: sub}, nil
}

// WatchOpenCompleted is a free log subscription operation binding the contract event 0x26f26478175d93f68a0289cff1e2e3d1d4969d3eaf5bd0421a215764c1a7267b.
//
// Solidity: event openCompleted(address indexed account, address token, address token0, uint256 usdtAmount, uint256 userMaxTime)
func (_Stake *StakeFilterer) WatchOpenCompleted(opts *bind.WatchOpts, sink chan<- *StakeOpenCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "openCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOpenCompleted)
				if err := _Stake.contract.UnpackLog(event, "openCompleted", log); err != nil {
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

// ParseOpenCompleted is a log parse operation binding the contract event 0x26f26478175d93f68a0289cff1e2e3d1d4969d3eaf5bd0421a215764c1a7267b.
//
// Solidity: event openCompleted(address indexed account, address token, address token0, uint256 usdtAmount, uint256 userMaxTime)
func (_Stake *StakeFilterer) ParseOpenCompleted(log types.Log) (*StakeOpenCompleted, error) {
	event := new(StakeOpenCompleted)
	if err := _Stake.contract.UnpackLog(event, "openCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakeCompletedIterator is returned from FilterStakeCompleted and is used to iterate over the raw logs and unpacked data for StakeCompleted events raised by the Stake contract.
type StakeStakeCompletedIterator struct {
	Event *StakeStakeCompleted // Event containing the contract specifics and raw log

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
func (it *StakeStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeCompleted)
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
		it.Event = new(StakeStakeCompleted)
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
func (it *StakeStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeCompleted represents a StakeCompleted event raised by the Stake contract.
type StakeStakeCompleted struct {
	Account     common.Address
	Token       common.Address
	Amount      *big.Int
	UserMaxTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeCompleted is a free log retrieval operation binding the contract event 0x58db6d2d7fa2d53531c6c5b4b677cafec910a8436671b93c2a86ca08c3f279db.
//
// Solidity: event stakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) FilterStakeCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeStakeCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "stakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeCompletedIterator{contract: _Stake.contract, event: "stakeCompleted", logs: logs, sub: sub}, nil
}

// WatchStakeCompleted is a free log subscription operation binding the contract event 0x58db6d2d7fa2d53531c6c5b4b677cafec910a8436671b93c2a86ca08c3f279db.
//
// Solidity: event stakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) WatchStakeCompleted(opts *bind.WatchOpts, sink chan<- *StakeStakeCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "stakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeCompleted)
				if err := _Stake.contract.UnpackLog(event, "stakeCompleted", log); err != nil {
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

// ParseStakeCompleted is a log parse operation binding the contract event 0x58db6d2d7fa2d53531c6c5b4b677cafec910a8436671b93c2a86ca08c3f279db.
//
// Solidity: event stakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) ParseStakeCompleted(log types.Log) (*StakeStakeCompleted, error) {
	event := new(StakeStakeCompleted)
	if err := _Stake.contract.UnpackLog(event, "stakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeUnStakeCompletedIterator is returned from FilterUnStakeCompleted and is used to iterate over the raw logs and unpacked data for UnStakeCompleted events raised by the Stake contract.
type StakeUnStakeCompletedIterator struct {
	Event *StakeUnStakeCompleted // Event containing the contract specifics and raw log

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
func (it *StakeUnStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeUnStakeCompleted)
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
		it.Event = new(StakeUnStakeCompleted)
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
func (it *StakeUnStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeUnStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeUnStakeCompleted represents a UnStakeCompleted event raised by the Stake contract.
type StakeUnStakeCompleted struct {
	Account     common.Address
	Token       common.Address
	Amount      *big.Int
	UserMaxTime *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnStakeCompleted is a free log retrieval operation binding the contract event 0x8134d1e198731ab72ca96a08bf2f01171b8ea827a660dc7eb04c3dad75edbd1b.
//
// Solidity: event unStakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) FilterUnStakeCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeUnStakeCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "unStakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeUnStakeCompletedIterator{contract: _Stake.contract, event: "unStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchUnStakeCompleted is a free log subscription operation binding the contract event 0x8134d1e198731ab72ca96a08bf2f01171b8ea827a660dc7eb04c3dad75edbd1b.
//
// Solidity: event unStakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) WatchUnStakeCompleted(opts *bind.WatchOpts, sink chan<- *StakeUnStakeCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "unStakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeUnStakeCompleted)
				if err := _Stake.contract.UnpackLog(event, "unStakeCompleted", log); err != nil {
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

// ParseUnStakeCompleted is a log parse operation binding the contract event 0x8134d1e198731ab72ca96a08bf2f01171b8ea827a660dc7eb04c3dad75edbd1b.
//
// Solidity: event unStakeCompleted(address indexed account, address token, uint256 amount, uint256 userMaxTime)
func (_Stake *StakeFilterer) ParseUnStakeCompleted(log types.Log) (*StakeUnStakeCompleted, error) {
	event := new(StakeUnStakeCompleted)
	if err := _Stake.contract.UnpackLog(event, "unStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
