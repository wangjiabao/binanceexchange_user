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

// DepositOpenMetaData contains all meta data concerning the DepositOpen contract.
var DepositOpenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CloseCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OpenCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DepositOpenABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositOpenMetaData.ABI instead.
var DepositOpenABI = DepositOpenMetaData.ABI

// DepositOpen is an auto generated Go binding around an Ethereum contract.
type DepositOpen struct {
	DepositOpenCaller     // Read-only binding to the contract
	DepositOpenTransactor // Write-only binding to the contract
	DepositOpenFilterer   // Log filterer for contract events
}

// DepositOpenCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositOpenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositOpenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositOpenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositOpenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositOpenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositOpenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositOpenSession struct {
	Contract     *DepositOpen      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositOpenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositOpenCallerSession struct {
	Contract *DepositOpenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DepositOpenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositOpenTransactorSession struct {
	Contract     *DepositOpenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DepositOpenRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositOpenRaw struct {
	Contract *DepositOpen // Generic contract binding to access the raw methods on
}

// DepositOpenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositOpenCallerRaw struct {
	Contract *DepositOpenCaller // Generic read-only contract binding to access the raw methods on
}

// DepositOpenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositOpenTransactorRaw struct {
	Contract *DepositOpenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositOpen creates a new instance of DepositOpen, bound to a specific deployed contract.
func NewDepositOpen(address common.Address, backend bind.ContractBackend) (*DepositOpen, error) {
	contract, err := bindDepositOpen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositOpen{DepositOpenCaller: DepositOpenCaller{contract: contract}, DepositOpenTransactor: DepositOpenTransactor{contract: contract}, DepositOpenFilterer: DepositOpenFilterer{contract: contract}}, nil
}

// NewDepositOpenCaller creates a new read-only instance of DepositOpen, bound to a specific deployed contract.
func NewDepositOpenCaller(address common.Address, caller bind.ContractCaller) (*DepositOpenCaller, error) {
	contract, err := bindDepositOpen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositOpenCaller{contract: contract}, nil
}

// NewDepositOpenTransactor creates a new write-only instance of DepositOpen, bound to a specific deployed contract.
func NewDepositOpenTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositOpenTransactor, error) {
	contract, err := bindDepositOpen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositOpenTransactor{contract: contract}, nil
}

// NewDepositOpenFilterer creates a new log filterer instance of DepositOpen, bound to a specific deployed contract.
func NewDepositOpenFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositOpenFilterer, error) {
	contract, err := bindDepositOpen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositOpenFilterer{contract: contract}, nil
}

// bindDepositOpen binds a generic wrapper to an already deployed contract.
func bindDepositOpen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositOpenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositOpen *DepositOpenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositOpen.Contract.DepositOpenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositOpen *DepositOpenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositOpen.Contract.DepositOpenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositOpen *DepositOpenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositOpen.Contract.DepositOpenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositOpen *DepositOpenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositOpen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositOpen *DepositOpenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositOpen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositOpen *DepositOpenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositOpen.Contract.contract.Transact(opts, method, params...)
}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_DepositOpen *DepositOpenCaller) UserOpen(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _DepositOpen.contract.Call(opts, &out, "userOpen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_DepositOpen *DepositOpenSession) UserOpen(arg0 common.Address) (bool, error) {
	return _DepositOpen.Contract.UserOpen(&_DepositOpen.CallOpts, arg0)
}

// UserOpen is a free data retrieval call binding the contract method 0x03c22aaa.
//
// Solidity: function userOpen(address ) view returns(bool)
func (_DepositOpen *DepositOpenCallerSession) UserOpen(arg0 common.Address) (bool, error) {
	return _DepositOpen.Contract.UserOpen(&_DepositOpen.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DepositOpen *DepositOpenTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositOpen.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DepositOpen *DepositOpenSession) Close() (*types.Transaction, error) {
	return _DepositOpen.Contract.Close(&_DepositOpen.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_DepositOpen *DepositOpenTransactorSession) Close() (*types.Transaction, error) {
	return _DepositOpen.Contract.Close(&_DepositOpen.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DepositOpen *DepositOpenTransactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositOpen.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DepositOpen *DepositOpenSession) Open() (*types.Transaction, error) {
	return _DepositOpen.Contract.Open(&_DepositOpen.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_DepositOpen *DepositOpenTransactorSession) Open() (*types.Transaction, error) {
	return _DepositOpen.Contract.Open(&_DepositOpen.TransactOpts)
}

// DepositOpenCloseCompletedIterator is returned from FilterCloseCompleted and is used to iterate over the raw logs and unpacked data for CloseCompleted events raised by the DepositOpen contract.
type DepositOpenCloseCompletedIterator struct {
	Event *DepositOpenCloseCompleted // Event containing the contract specifics and raw log

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
func (it *DepositOpenCloseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositOpenCloseCompleted)
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
		it.Event = new(DepositOpenCloseCompleted)
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
func (it *DepositOpenCloseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositOpenCloseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositOpenCloseCompleted represents a CloseCompleted event raised by the DepositOpen contract.
type DepositOpenCloseCompleted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCloseCompleted is a free log retrieval operation binding the contract event 0x2e6b11d6ee8ee0b4cb12ac7bf6f9e25bb2579bf6225fa04cae2ce8e50d0e8142.
//
// Solidity: event CloseCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) FilterCloseCompleted(opts *bind.FilterOpts, account []common.Address) (*DepositOpenCloseCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositOpen.contract.FilterLogs(opts, "CloseCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &DepositOpenCloseCompletedIterator{contract: _DepositOpen.contract, event: "CloseCompleted", logs: logs, sub: sub}, nil
}

// WatchCloseCompleted is a free log subscription operation binding the contract event 0x2e6b11d6ee8ee0b4cb12ac7bf6f9e25bb2579bf6225fa04cae2ce8e50d0e8142.
//
// Solidity: event CloseCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) WatchCloseCompleted(opts *bind.WatchOpts, sink chan<- *DepositOpenCloseCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositOpen.contract.WatchLogs(opts, "CloseCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositOpenCloseCompleted)
				if err := _DepositOpen.contract.UnpackLog(event, "CloseCompleted", log); err != nil {
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

// ParseCloseCompleted is a log parse operation binding the contract event 0x2e6b11d6ee8ee0b4cb12ac7bf6f9e25bb2579bf6225fa04cae2ce8e50d0e8142.
//
// Solidity: event CloseCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) ParseCloseCompleted(log types.Log) (*DepositOpenCloseCompleted, error) {
	event := new(DepositOpenCloseCompleted)
	if err := _DepositOpen.contract.UnpackLog(event, "CloseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositOpenOpenCompletedIterator is returned from FilterOpenCompleted and is used to iterate over the raw logs and unpacked data for OpenCompleted events raised by the DepositOpen contract.
type DepositOpenOpenCompletedIterator struct {
	Event *DepositOpenOpenCompleted // Event containing the contract specifics and raw log

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
func (it *DepositOpenOpenCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositOpenOpenCompleted)
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
		it.Event = new(DepositOpenOpenCompleted)
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
func (it *DepositOpenOpenCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositOpenOpenCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositOpenOpenCompleted represents a OpenCompleted event raised by the DepositOpen contract.
type DepositOpenOpenCompleted struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOpenCompleted is a free log retrieval operation binding the contract event 0x5370723db71bb9dc233e426366365df8ccd9d476f689590f572fcf83f55c48ea.
//
// Solidity: event OpenCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) FilterOpenCompleted(opts *bind.FilterOpts, account []common.Address) (*DepositOpenOpenCompletedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositOpen.contract.FilterLogs(opts, "OpenCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return &DepositOpenOpenCompletedIterator{contract: _DepositOpen.contract, event: "OpenCompleted", logs: logs, sub: sub}, nil
}

// WatchOpenCompleted is a free log subscription operation binding the contract event 0x5370723db71bb9dc233e426366365df8ccd9d476f689590f572fcf83f55c48ea.
//
// Solidity: event OpenCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) WatchOpenCompleted(opts *bind.WatchOpts, sink chan<- *DepositOpenOpenCompleted, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _DepositOpen.contract.WatchLogs(opts, "OpenCompleted", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositOpenOpenCompleted)
				if err := _DepositOpen.contract.UnpackLog(event, "OpenCompleted", log); err != nil {
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

// ParseOpenCompleted is a log parse operation binding the contract event 0x5370723db71bb9dc233e426366365df8ccd9d476f689590f572fcf83f55c48ea.
//
// Solidity: event OpenCompleted(address indexed account)
func (_DepositOpen *DepositOpenFilterer) ParseOpenCompleted(log types.Log) (*DepositOpenOpenCompleted, error) {
	event := new(DepositOpenOpenCompleted)
	if err := _DepositOpen.contract.UnpackLog(event, "OpenCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
