// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBPS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"getTokensFromPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"bps\",\"type\":\"uint32\"}],\"name\":\"setFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implLogic\",\"type\":\"address\"}],\"name\":\"setPairLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"tokensToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "allPairsLength")
	return *ret0, err
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactorySession) AllPairsLength() (*big.Int, error) {
	return _Factory.Contract.AllPairsLength(&_Factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Factory *FactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Factory.Contract.AllPairsLength(&_Factory.CallOpts)
}

// FeeBPS is a free data retrieval call binding the contract method 0x1a1c6e53.
//
// Solidity: function feeBPS() view returns(uint32)
func (_Factory *FactoryCaller) FeeBPS(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "feeBPS")
	return *ret0, err
}

// FeeBPS is a free data retrieval call binding the contract method 0x1a1c6e53.
//
// Solidity: function feeBPS() view returns(uint32)
func (_Factory *FactorySession) FeeBPS() (uint32, error) {
	return _Factory.Contract.FeeBPS(&_Factory.CallOpts)
}

// FeeBPS is a free data retrieval call binding the contract method 0x1a1c6e53.
//
// Solidity: function feeBPS() view returns(uint32)
func (_Factory *FactoryCallerSession) FeeBPS() (uint32, error) {
	return _Factory.Contract.FeeBPS(&_Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Factory *FactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "feeTo")
	return *ret0, err
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Factory *FactorySession) FeeTo() (common.Address, error) {
	return _Factory.Contract.FeeTo(&_Factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Factory *FactoryCallerSession) FeeTo() (common.Address, error) {
	return _Factory.Contract.FeeTo(&_Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Factory *FactoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "feeToSetter")
	return *ret0, err
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Factory *FactorySession) FeeToSetter() (common.Address, error) {
	return _Factory.Contract.FeeToSetter(&_Factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Factory *FactoryCallerSession) FeeToSetter() (common.Address, error) {
	return _Factory.Contract.FeeToSetter(&_Factory.CallOpts)
}

// GetTokensFromPair is a free data retrieval call binding the contract method 0xc0225aa1.
//
// Solidity: function getTokensFromPair(address pair) view returns(address stock, address money)
func (_Factory *FactoryCaller) GetTokensFromPair(opts *bind.CallOpts, pair common.Address) (struct {
	Stock common.Address
	Money common.Address
}, error) {
	ret := new(struct {
		Stock common.Address
		Money common.Address
	})
	out := ret
	err := _Factory.contract.Call(opts, out, "getTokensFromPair", pair)
	return *ret, err
}

// GetTokensFromPair is a free data retrieval call binding the contract method 0xc0225aa1.
//
// Solidity: function getTokensFromPair(address pair) view returns(address stock, address money)
func (_Factory *FactorySession) GetTokensFromPair(pair common.Address) (struct {
	Stock common.Address
	Money common.Address
}, error) {
	return _Factory.Contract.GetTokensFromPair(&_Factory.CallOpts, pair)
}

// GetTokensFromPair is a free data retrieval call binding the contract method 0xc0225aa1.
//
// Solidity: function getTokensFromPair(address pair) view returns(address stock, address money)
func (_Factory *FactoryCallerSession) GetTokensFromPair(pair common.Address) (struct {
	Stock common.Address
	Money common.Address
}, error) {
	return _Factory.Contract.GetTokensFromPair(&_Factory.CallOpts, pair)
}

// TokensToPair is a free data retrieval call binding the contract method 0x6c87813e.
//
// Solidity: function tokensToPair(address stock, address money, bool isOnlySwap) view returns(address pair)
func (_Factory *FactoryCaller) TokensToPair(opts *bind.CallOpts, stock common.Address, money common.Address, isOnlySwap bool) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "tokensToPair", stock, money, isOnlySwap)
	return *ret0, err
}

// TokensToPair is a free data retrieval call binding the contract method 0x6c87813e.
//
// Solidity: function tokensToPair(address stock, address money, bool isOnlySwap) view returns(address pair)
func (_Factory *FactorySession) TokensToPair(stock common.Address, money common.Address, isOnlySwap bool) (common.Address, error) {
	return _Factory.Contract.TokensToPair(&_Factory.CallOpts, stock, money, isOnlySwap)
}

// TokensToPair is a free data retrieval call binding the contract method 0x6c87813e.
//
// Solidity: function tokensToPair(address stock, address money, bool isOnlySwap) view returns(address pair)
func (_Factory *FactoryCallerSession) TokensToPair(stock common.Address, money common.Address, isOnlySwap bool) (common.Address, error) {
	return _Factory.Contract.TokensToPair(&_Factory.CallOpts, stock, money, isOnlySwap)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address stock, address money, bool isOnlySwap) returns(address pair)
func (_Factory *FactoryTransactor) CreatePair(opts *bind.TransactOpts, stock common.Address, money common.Address, isOnlySwap bool) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createPair", stock, money, isOnlySwap)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address stock, address money, bool isOnlySwap) returns(address pair)
func (_Factory *FactorySession) CreatePair(stock common.Address, money common.Address, isOnlySwap bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePair(&_Factory.TransactOpts, stock, money, isOnlySwap)
}

// CreatePair is a paid mutator transaction binding the contract method 0x82dfdce4.
//
// Solidity: function createPair(address stock, address money, bool isOnlySwap) returns(address pair)
func (_Factory *FactoryTransactorSession) CreatePair(stock common.Address, money common.Address, isOnlySwap bool) (*types.Transaction, error) {
	return _Factory.Contract.CreatePair(&_Factory.TransactOpts, stock, money, isOnlySwap)
}

// PairLogic is a paid mutator transaction binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() returns(address)
func (_Factory *FactoryTransactor) PairLogic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "pairLogic")
}

// PairLogic is a paid mutator transaction binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() returns(address)
func (_Factory *FactorySession) PairLogic() (*types.Transaction, error) {
	return _Factory.Contract.PairLogic(&_Factory.TransactOpts)
}

// PairLogic is a paid mutator transaction binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() returns(address)
func (_Factory *FactoryTransactorSession) PairLogic() (*types.Transaction, error) {
	return _Factory.Contract.PairLogic(&_Factory.TransactOpts)
}

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 bps) returns()
func (_Factory *FactoryTransactor) SetFeeBPS(opts *bind.TransactOpts, bps uint32) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeBPS", bps)
}

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 bps) returns()
func (_Factory *FactorySession) SetFeeBPS(bps uint32) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeBPS(&_Factory.TransactOpts, bps)
}

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 bps) returns()
func (_Factory *FactoryTransactorSession) SetFeeBPS(bps uint32) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeBPS(&_Factory.TransactOpts, bps)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Factory *FactoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Factory *FactorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeTo(&_Factory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Factory *FactoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeTo(&_Factory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Factory *FactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Factory *FactorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeToSetter(&_Factory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Factory *FactoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeToSetter(&_Factory.TransactOpts, arg0)
}

// SetPairLogic is a paid mutator transaction binding the contract method 0xca215225.
//
// Solidity: function setPairLogic(address implLogic) returns()
func (_Factory *FactoryTransactor) SetPairLogic(opts *bind.TransactOpts, implLogic common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setPairLogic", implLogic)
}

// SetPairLogic is a paid mutator transaction binding the contract method 0xca215225.
//
// Solidity: function setPairLogic(address implLogic) returns()
func (_Factory *FactorySession) SetPairLogic(implLogic common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetPairLogic(&_Factory.TransactOpts, implLogic)
}

// SetPairLogic is a paid mutator transaction binding the contract method 0xca215225.
//
// Solidity: function setPairLogic(address implLogic) returns()
func (_Factory *FactoryTransactorSession) SetPairLogic(implLogic common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetPairLogic(&_Factory.TransactOpts, implLogic)
}

// FactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Factory contract.
type FactoryPairCreatedIterator struct {
	Event *FactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *FactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryPairCreated)
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
		it.Event = new(FactoryPairCreated)
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
func (it *FactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryPairCreated represents a PairCreated event raised by the Factory contract.
type FactoryPairCreated struct {
	Pair       common.Address
	Stock      common.Address
	Money      common.Address
	IsOnlySwap bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb.
//
// Solidity: event PairCreated(address indexed pair, address stock, address money, bool isOnlySwap)
func (_Factory *FactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address) (*FactoryPairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _Factory.contract.FilterLogs(opts, "PairCreated", pairRule)
	if err != nil {
		return nil, err
	}
	return &FactoryPairCreatedIterator{contract: _Factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb.
//
// Solidity: event PairCreated(address indexed pair, address stock, address money, bool isOnlySwap)
func (_Factory *FactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *FactoryPairCreated, pair []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _Factory.contract.WatchLogs(opts, "PairCreated", pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryPairCreated)
				if err := _Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb.
//
// Solidity: event PairCreated(address indexed pair, address stock, address money, bool isOnlySwap)
func (_Factory *FactoryFilterer) ParsePairCreated(log types.Log) (*FactoryPairCreated, error) {
	event := new(FactoryPairCreated)
	if err := _Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
