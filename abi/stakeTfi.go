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

// StakeTfiMetaData contains all meta data concerning the StakeTfi contract.
var StakeTfiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tfi\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CloseCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OpenCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnStakeCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getUsersByIndex\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tfi\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userTfiAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeTfiABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeTfiMetaData.ABI instead.
var StakeTfiABI = StakeTfiMetaData.ABI

// StakeTfi is an auto generated Go binding around an Ethereum contract.
type StakeTfi struct {
	StakeTfiCaller     // Read-only binding to the contract
	StakeTfiTransactor // Write-only binding to the contract
	StakeTfiFilterer   // Log filterer for contract events
}

// StakeTfiCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeTfiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTfiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTfiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTfiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeTfiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTfiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeTfiSession struct {
	Contract     *StakeTfi         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeTfiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeTfiCallerSession struct {
	Contract *StakeTfiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StakeTfiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTfiTransactorSession struct {
	Contract     *StakeTfiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakeTfiRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeTfiRaw struct {
	Contract *StakeTfi // Generic contract binding to access the raw methods on
}

// StakeTfiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeTfiCallerRaw struct {
	Contract *StakeTfiCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTfiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTfiTransactorRaw struct {
	Contract *StakeTfiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeTfi creates a new instance of StakeTfi, bound to a specific deployed contract.
func NewStakeTfi(address common.Address, backend bind.ContractBackend) (*StakeTfi, error) {
	contract, err := bindStakeTfi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeTfi{StakeTfiCaller: StakeTfiCaller{contract: contract}, StakeTfiTransactor: StakeTfiTransactor{contract: contract}, StakeTfiFilterer: StakeTfiFilterer{contract: contract}}, nil
}

// NewStakeTfiCaller creates a new read-only instance of StakeTfi, bound to a specific deployed contract.
func NewStakeTfiCaller(address common.Address, caller bind.ContractCaller) (*StakeTfiCaller, error) {
	contract, err := bindStakeTfi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTfiCaller{contract: contract}, nil
}

// NewStakeTfiTransactor creates a new write-only instance of StakeTfi, bound to a specific deployed contract.
func NewStakeTfiTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTfiTransactor, error) {
	contract, err := bindStakeTfi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTfiTransactor{contract: contract}, nil
}

// NewStakeTfiFilterer creates a new log filterer instance of StakeTfi, bound to a specific deployed contract.
func NewStakeTfiFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeTfiFilterer, error) {
	contract, err := bindStakeTfi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeTfiFilterer{contract: contract}, nil
}

// bindStakeTfi binds a generic wrapper to an already deployed contract.
func bindStakeTfi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeTfiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeTfi *StakeTfiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeTfi.Contract.StakeTfiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeTfi *StakeTfiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeTfi.Contract.StakeTfiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeTfi *StakeTfiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeTfi.Contract.StakeTfiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeTfi *StakeTfiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeTfi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeTfi *StakeTfiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeTfi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeTfi *StakeTfiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeTfi.Contract.contract.Transact(opts, method, params...)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_StakeTfi *StakeTfiCaller) GetUserLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "getUserLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_StakeTfi *StakeTfiSession) GetUserLength() (*big.Int, error) {
	return _StakeTfi.Contract.GetUserLength(&_StakeTfi.CallOpts)
}

// GetUserLength is a free data retrieval call binding the contract method 0x7456fed6.
//
// Solidity: function getUserLength() view returns(uint256)
func (_StakeTfi *StakeTfiCallerSession) GetUserLength() (*big.Int, error) {
	return _StakeTfi.Contract.GetUserLength(&_StakeTfi.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_StakeTfi *StakeTfiCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "getUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_StakeTfi *StakeTfiSession) GetUsers() ([]common.Address, error) {
	return _StakeTfi.Contract.GetUsers(&_StakeTfi.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() view returns(address[])
func (_StakeTfi *StakeTfiCallerSession) GetUsers() ([]common.Address, error) {
	return _StakeTfi.Contract.GetUsers(&_StakeTfi.CallOpts)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_StakeTfi *StakeTfiCaller) GetUsersByIndex(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "getUsersByIndex", startIndex, endIndex)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_StakeTfi *StakeTfiSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _StakeTfi.Contract.GetUsersByIndex(&_StakeTfi.CallOpts, startIndex, endIndex)
}

// GetUsersByIndex is a free data retrieval call binding the contract method 0xfe36c56c.
//
// Solidity: function getUsersByIndex(uint256 startIndex, uint256 endIndex) view returns(address[])
func (_StakeTfi *StakeTfiCallerSession) GetUsersByIndex(startIndex *big.Int, endIndex *big.Int) ([]common.Address, error) {
	return _StakeTfi.Contract.GetUsersByIndex(&_StakeTfi.CallOpts, startIndex, endIndex)
}

// Tfi is a free data retrieval call binding the contract method 0x5b5cff16.
//
// Solidity: function tfi() view returns(address)
func (_StakeTfi *StakeTfiCaller) Tfi(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "tfi")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tfi is a free data retrieval call binding the contract method 0x5b5cff16.
//
// Solidity: function tfi() view returns(address)
func (_StakeTfi *StakeTfiSession) Tfi() (common.Address, error) {
	return _StakeTfi.Contract.Tfi(&_StakeTfi.CallOpts)
}

// Tfi is a free data retrieval call binding the contract method 0x5b5cff16.
//
// Solidity: function tfi() view returns(address)
func (_StakeTfi *StakeTfiCallerSession) Tfi() (common.Address, error) {
	return _StakeTfi.Contract.Tfi(&_StakeTfi.CallOpts)
}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_StakeTfi *StakeTfiCaller) UserCost(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "userCost", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_StakeTfi *StakeTfiSession) UserCost(arg0 common.Address) (*big.Int, error) {
	return _StakeTfi.Contract.UserCost(&_StakeTfi.CallOpts, arg0)
}

// UserCost is a free data retrieval call binding the contract method 0xc47a03a0.
//
// Solidity: function userCost(address ) view returns(uint256)
func (_StakeTfi *StakeTfiCallerSession) UserCost(arg0 common.Address) (*big.Int, error) {
	return _StakeTfi.Contract.UserCost(&_StakeTfi.CallOpts, arg0)
}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_StakeTfi *StakeTfiCaller) UserOpen(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "userOpen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_StakeTfi *StakeTfiSession) UserOpen(arg0 common.Address) (bool, error) {
	return _StakeTfi.Contract.UserOpen(&_StakeTfi.CallOpts, arg0)
}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_StakeTfi *StakeTfiCallerSession) UserOpen(arg0 common.Address) (bool, error) {
	return _StakeTfi.Contract.UserOpen(&_StakeTfi.CallOpts, arg0)
}

// UserTfiAmount is a free data retrieval call binding the contract method 0x34284f7e.
//
// Solidity: function userTfiAmount(address ) view returns(uint256)
func (_StakeTfi *StakeTfiCaller) UserTfiAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeTfi.contract.Call(opts, &out, "userTfiAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTfiAmount is a free data retrieval call binding the contract method 0x34284f7e.
//
// Solidity: function userTfiAmount(address ) view returns(uint256)
func (_StakeTfi *StakeTfiSession) UserTfiAmount(arg0 common.Address) (*big.Int, error) {
	return _StakeTfi.Contract.UserTfiAmount(&_StakeTfi.CallOpts, arg0)
}

// UserTfiAmount is a free data retrieval call binding the contract method 0x34284f7e.
//
// Solidity: function userTfiAmount(address ) view returns(uint256)
func (_StakeTfi *StakeTfiCallerSession) UserTfiAmount(arg0 common.Address) (*big.Int, error) {
	return _StakeTfi.Contract.UserTfiAmount(&_StakeTfi.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_StakeTfi *StakeTfiTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeTfi.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_StakeTfi *StakeTfiSession) Close() (*types.Transaction, error) {
	return _StakeTfi.Contract.Close(&_StakeTfi.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_StakeTfi *StakeTfiTransactorSession) Close() (*types.Transaction, error) {
	return _StakeTfi.Contract.Close(&_StakeTfi.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 cost) returns()
func (_StakeTfi *StakeTfiTransactor) Open(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _StakeTfi.contract.Transact(opts, "open", cost)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 cost) returns()
func (_StakeTfi *StakeTfiSession) Open(cost *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.Open(&_StakeTfi.TransactOpts, cost)
}

// Open is a paid mutator transaction binding the contract method 0x690e7c09.
//
// Solidity: function open(uint256 cost) returns()
func (_StakeTfi *StakeTfiTransactorSession) Open(cost *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.Open(&_StakeTfi.TransactOpts, cost)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakeTfi *StakeTfiTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakeTfi *StakeTfiSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.Stake(&_StakeTfi.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakeTfi *StakeTfiTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.Stake(&_StakeTfi.TransactOpts, amount)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 amount) returns()
func (_StakeTfi *StakeTfiTransactor) UnStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.contract.Transact(opts, "unStake", amount)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 amount) returns()
func (_StakeTfi *StakeTfiSession) UnStake(amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.UnStake(&_StakeTfi.TransactOpts, amount)
}

// UnStake is a paid mutator transaction binding the contract method 0x5d3eea91.
//
// Solidity: function unStake(uint256 amount) returns()
func (_StakeTfi *StakeTfiTransactorSession) UnStake(amount *big.Int) (*types.Transaction, error) {
	return _StakeTfi.Contract.UnStake(&_StakeTfi.TransactOpts, amount)
}

// StakeTfiCloseCompletedIterator is returned from FilterCloseCompleted and is used to iterate over the raw logs and unpacked data for CloseCompleted events raised by the StakeTfi contract.
type StakeTfiCloseCompletedIterator struct {
	Event *StakeTfiCloseCompleted // Event containing the contract specifics and raw log

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
func (it *StakeTfiCloseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTfiCloseCompleted)
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
		it.Event = new(StakeTfiCloseCompleted)
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
func (it *StakeTfiCloseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTfiCloseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTfiCloseCompleted represents a CloseCompleted event raised by the StakeTfi contract.
type StakeTfiCloseCompleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCloseCompleted is a free log retrieval operation binding the contract event 0xf77b93d1ff45ebc73be645e6fcc1c601f57ee7d1545de06a2dfd194acdb3bed4.
//
// Solidity: event CloseCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) FilterCloseCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeTfiCloseCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.FilterLogs(opts, "CloseCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeTfiCloseCompletedIterator{contract: _StakeTfi.contract, event: "CloseCompleted", logs: logs, sub: sub}, nil
}

// WatchCloseCompleted is a free log subscription operation binding the contract event 0xf77b93d1ff45ebc73be645e6fcc1c601f57ee7d1545de06a2dfd194acdb3bed4.
//
// Solidity: event CloseCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) WatchCloseCompleted(opts *bind.WatchOpts, sink chan<- *StakeTfiCloseCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.WatchLogs(opts, "CloseCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTfiCloseCompleted)
				if err := _StakeTfi.contract.UnpackLog(event, "CloseCompleted", log); err != nil {
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

// ParseCloseCompleted is a log parse operation binding the contract event 0xf77b93d1ff45ebc73be645e6fcc1c601f57ee7d1545de06a2dfd194acdb3bed4.
//
// Solidity: event CloseCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) ParseCloseCompleted(log types.Log) (*StakeTfiCloseCompleted, error) {
	event := new(StakeTfiCloseCompleted)
	if err := _StakeTfi.contract.UnpackLog(event, "CloseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTfiOpenCompletedIterator is returned from FilterOpenCompleted and is used to iterate over the raw logs and unpacked data for OpenCompleted events raised by the StakeTfi contract.
type StakeTfiOpenCompletedIterator struct {
	Event *StakeTfiOpenCompleted // Event containing the contract specifics and raw log

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
func (it *StakeTfiOpenCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTfiOpenCompleted)
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
		it.Event = new(StakeTfiOpenCompleted)
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
func (it *StakeTfiOpenCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTfiOpenCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTfiOpenCompleted represents a OpenCompleted event raised by the StakeTfi contract.
type StakeTfiOpenCompleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenCompleted is a free log retrieval operation binding the contract event 0x344d9bb47d1e7c1dd851fabba2f9d946154db488ff95dc466e36353112c79892.
//
// Solidity: event OpenCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) FilterOpenCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeTfiOpenCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.FilterLogs(opts, "OpenCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeTfiOpenCompletedIterator{contract: _StakeTfi.contract, event: "OpenCompleted", logs: logs, sub: sub}, nil
}

// WatchOpenCompleted is a free log subscription operation binding the contract event 0x344d9bb47d1e7c1dd851fabba2f9d946154db488ff95dc466e36353112c79892.
//
// Solidity: event OpenCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) WatchOpenCompleted(opts *bind.WatchOpts, sink chan<- *StakeTfiOpenCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.WatchLogs(opts, "OpenCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTfiOpenCompleted)
				if err := _StakeTfi.contract.UnpackLog(event, "OpenCompleted", log); err != nil {
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

// ParseOpenCompleted is a log parse operation binding the contract event 0x344d9bb47d1e7c1dd851fabba2f9d946154db488ff95dc466e36353112c79892.
//
// Solidity: event OpenCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) ParseOpenCompleted(log types.Log) (*StakeTfiOpenCompleted, error) {
	event := new(StakeTfiOpenCompleted)
	if err := _StakeTfi.contract.UnpackLog(event, "OpenCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTfiStakeCompletedIterator is returned from FilterStakeCompleted and is used to iterate over the raw logs and unpacked data for StakeCompleted events raised by the StakeTfi contract.
type StakeTfiStakeCompletedIterator struct {
	Event *StakeTfiStakeCompleted // Event containing the contract specifics and raw log

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
func (it *StakeTfiStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTfiStakeCompleted)
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
		it.Event = new(StakeTfiStakeCompleted)
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
func (it *StakeTfiStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTfiStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTfiStakeCompleted represents a StakeCompleted event raised by the StakeTfi contract.
type StakeTfiStakeCompleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeCompleted is a free log retrieval operation binding the contract event 0x1db000f103b8324c7cd49fa51320c82dca922c81eed3d0e1b0010f94fbdd9fc3.
//
// Solidity: event StakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) FilterStakeCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeTfiStakeCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.FilterLogs(opts, "StakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeTfiStakeCompletedIterator{contract: _StakeTfi.contract, event: "StakeCompleted", logs: logs, sub: sub}, nil
}

// WatchStakeCompleted is a free log subscription operation binding the contract event 0x1db000f103b8324c7cd49fa51320c82dca922c81eed3d0e1b0010f94fbdd9fc3.
//
// Solidity: event StakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) WatchStakeCompleted(opts *bind.WatchOpts, sink chan<- *StakeTfiStakeCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.WatchLogs(opts, "StakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTfiStakeCompleted)
				if err := _StakeTfi.contract.UnpackLog(event, "StakeCompleted", log); err != nil {
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

// ParseStakeCompleted is a log parse operation binding the contract event 0x1db000f103b8324c7cd49fa51320c82dca922c81eed3d0e1b0010f94fbdd9fc3.
//
// Solidity: event StakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) ParseStakeCompleted(log types.Log) (*StakeTfiStakeCompleted, error) {
	event := new(StakeTfiStakeCompleted)
	if err := _StakeTfi.contract.UnpackLog(event, "StakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTfiUnStakeCompletedIterator is returned from FilterUnStakeCompleted and is used to iterate over the raw logs and unpacked data for UnStakeCompleted events raised by the StakeTfi contract.
type StakeTfiUnStakeCompletedIterator struct {
	Event *StakeTfiUnStakeCompleted // Event containing the contract specifics and raw log

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
func (it *StakeTfiUnStakeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTfiUnStakeCompleted)
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
		it.Event = new(StakeTfiUnStakeCompleted)
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
func (it *StakeTfiUnStakeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTfiUnStakeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTfiUnStakeCompleted represents a UnStakeCompleted event raised by the StakeTfi contract.
type StakeTfiUnStakeCompleted struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnStakeCompleted is a free log retrieval operation binding the contract event 0x3e97d1f089eda0ccec4c249ce3a73dc16722f788449b4da1ffe4687465c3df1d.
//
// Solidity: event UnStakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) FilterUnStakeCompleted(opts *bind.FilterOpts, account []common.Address) (*StakeTfiUnStakeCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.FilterLogs(opts, "UnStakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeTfiUnStakeCompletedIterator{contract: _StakeTfi.contract, event: "UnStakeCompleted", logs: logs, sub: sub}, nil
}

// WatchUnStakeCompleted is a free log subscription operation binding the contract event 0x3e97d1f089eda0ccec4c249ce3a73dc16722f788449b4da1ffe4687465c3df1d.
//
// Solidity: event UnStakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) WatchUnStakeCompleted(opts *bind.WatchOpts, sink chan<- *StakeTfiUnStakeCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeTfi.contract.WatchLogs(opts, "UnStakeCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTfiUnStakeCompleted)
				if err := _StakeTfi.contract.UnpackLog(event, "UnStakeCompleted", log); err != nil {
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

// ParseUnStakeCompleted is a log parse operation binding the contract event 0x3e97d1f089eda0ccec4c249ce3a73dc16722f788449b4da1ffe4687465c3df1d.
//
// Solidity: event UnStakeCompleted(address indexed account, uint256 amount)
func (_StakeTfi *StakeTfiFilterer) ParseUnStakeCompleted(log types.Log) (*StakeTfiUnStakeCompleted, error) {
	event := new(StakeTfiUnStakeCompleted)
	if err := _StakeTfi.contract.UnpackLog(event, "UnStakeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
