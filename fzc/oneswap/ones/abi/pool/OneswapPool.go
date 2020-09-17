// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stockAndMoneyAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stockAndMoneyAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveStockAndMoney\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"moneyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBooked\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"bookedStock\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"bookedMoney\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"firstBuyID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserveStock\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserveMoney\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"firstSellID\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"internalStatus\",\"outputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"res\",\"type\":\"uint256[3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"money\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// GetBooked is a free data retrieval call binding the contract method 0x1b857340.
//
// Solidity: function getBooked() view returns(uint112 bookedStock, uint112 bookedMoney, uint32 firstBuyID)
func (_Pool *PoolCaller) GetBooked(opts *bind.CallOpts) (struct {
	BookedStock *big.Int
	BookedMoney *big.Int
	FirstBuyID  uint32
}, error) {
	ret := new(struct {
		BookedStock *big.Int
		BookedMoney *big.Int
		FirstBuyID  uint32
	})
	out := ret
	err := _Pool.contract.Call(opts, out, "getBooked")
	return *ret, err
}

// GetBooked is a free data retrieval call binding the contract method 0x1b857340.
//
// Solidity: function getBooked() view returns(uint112 bookedStock, uint112 bookedMoney, uint32 firstBuyID)
func (_Pool *PoolSession) GetBooked() (struct {
	BookedStock *big.Int
	BookedMoney *big.Int
	FirstBuyID  uint32
}, error) {
	return _Pool.Contract.GetBooked(&_Pool.CallOpts)
}

// GetBooked is a free data retrieval call binding the contract method 0x1b857340.
//
// Solidity: function getBooked() view returns(uint112 bookedStock, uint112 bookedMoney, uint32 firstBuyID)
func (_Pool *PoolCallerSession) GetBooked() (struct {
	BookedStock *big.Int
	BookedMoney *big.Int
	FirstBuyID  uint32
}, error) {
	return _Pool.Contract.GetBooked(&_Pool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveStock, uint112 reserveMoney, uint32 firstSellID)
func (_Pool *PoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	ReserveStock *big.Int
	ReserveMoney *big.Int
	FirstSellID  uint32
}, error) {
	ret := new(struct {
		ReserveStock *big.Int
		ReserveMoney *big.Int
		FirstSellID  uint32
	})
	out := ret
	err := _Pool.contract.Call(opts, out, "getReserves")
	return *ret, err
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveStock, uint112 reserveMoney, uint32 firstSellID)
func (_Pool *PoolSession) GetReserves() (struct {
	ReserveStock *big.Int
	ReserveMoney *big.Int
	FirstSellID  uint32
}, error) {
	return _Pool.Contract.GetReserves(&_Pool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserveStock, uint112 reserveMoney, uint32 firstSellID)
func (_Pool *PoolCallerSession) GetReserves() (struct {
	ReserveStock *big.Int
	ReserveMoney *big.Int
	FirstSellID  uint32
}, error) {
	return _Pool.Contract.GetReserves(&_Pool.CallOpts)
}

// InternalStatus is a free data retrieval call binding the contract method 0xe6ba4136.
//
// Solidity: function internalStatus() view returns(uint256[3] res)
func (_Pool *PoolCaller) InternalStatus(opts *bind.CallOpts) ([3]*big.Int, error) {
	var (
		ret0 = new([3]*big.Int)
	)
	out := ret0
	err := _Pool.contract.Call(opts, out, "internalStatus")
	return *ret0, err
}

// InternalStatus is a free data retrieval call binding the contract method 0xe6ba4136.
//
// Solidity: function internalStatus() view returns(uint256[3] res)
func (_Pool *PoolSession) InternalStatus() ([3]*big.Int, error) {
	return _Pool.Contract.InternalStatus(&_Pool.CallOpts)
}

// InternalStatus is a free data retrieval call binding the contract method 0xe6ba4136.
//
// Solidity: function internalStatus() view returns(uint256[3] res)
func (_Pool *PoolCallerSession) InternalStatus() ([3]*big.Int, error) {
	return _Pool.Contract.InternalStatus(&_Pool.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 stockAmount, uint256 moneyAmount)
func (_Pool *PoolTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 stockAmount, uint256 moneyAmount)
func (_Pool *PoolSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 stockAmount, uint256 moneyAmount)
func (_Pool *PoolTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Burn(&_Pool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pool *PoolSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pool *PoolTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, to)
}

// Money is a paid mutator transaction binding the contract method 0x4ddd108a.
//
// Solidity: function money() returns(address)
func (_Pool *PoolTransactor) Money(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "money")
}

// Money is a paid mutator transaction binding the contract method 0x4ddd108a.
//
// Solidity: function money() returns(address)
func (_Pool *PoolSession) Money() (*types.Transaction, error) {
	return _Pool.Contract.Money(&_Pool.TransactOpts)
}

// Money is a paid mutator transaction binding the contract method 0x4ddd108a.
//
// Solidity: function money() returns(address)
func (_Pool *PoolTransactorSession) Money() (*types.Transaction, error) {
	return _Pool.Contract.Money(&_Pool.TransactOpts)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pool *PoolTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pool *PoolSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Skim(&_Pool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pool *PoolTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Skim(&_Pool.TransactOpts, to)
}

// Stock is a paid mutator transaction binding the contract method 0xbdf3c4ae.
//
// Solidity: function stock() returns(address)
func (_Pool *PoolTransactor) Stock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stock")
}

// Stock is a paid mutator transaction binding the contract method 0xbdf3c4ae.
//
// Solidity: function stock() returns(address)
func (_Pool *PoolSession) Stock() (*types.Transaction, error) {
	return _Pool.Contract.Stock(&_Pool.TransactOpts)
}

// Stock is a paid mutator transaction binding the contract method 0xbdf3c4ae.
//
// Solidity: function stock() returns(address)
func (_Pool *PoolTransactorSession) Stock() (*types.Transaction, error) {
	return _Pool.Contract.Stock(&_Pool.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pool *PoolTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pool *PoolSession) Sync() (*types.Transaction, error) {
	return _Pool.Contract.Sync(&_Pool.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pool *PoolTransactorSession) Sync() (*types.Transaction, error) {
	return _Pool.Contract.Sync(&_Pool.TransactOpts)
}

// PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pool contract.
type PoolBurnIterator struct {
	Event *PoolBurn // Event containing the contract specifics and raw log

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
func (it *PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBurn)
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
		it.Event = new(PoolBurn)
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
func (it *PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBurn represents a Burn event raised by the Pool contract.
type PoolBurn struct {
	Sender              common.Address
	StockAndMoneyAmount *big.Int
	To                  common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdbdf9b8e4b75e75b162d151ec8fc7f0561cabab5fcccfa2600be62223e4300c4.
//
// Solidity: event Burn(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolBurnIterator{contract: _Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdbdf9b8e4b75e75b162d151ec8fc7f0561cabab5fcccfa2600be62223e4300c4.
//
// Solidity: event Burn(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBurn)
				if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdbdf9b8e4b75e75b162d151ec8fc7f0561cabab5fcccfa2600be62223e4300c4.
//
// Solidity: event Burn(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) ParseBurn(log types.Log) (*PoolBurn, error) {
	event := new(PoolBurn)
	if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pool contract.
type PoolMintIterator struct {
	Event *PoolMint // Event containing the contract specifics and raw log

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
func (it *PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMint)
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
		it.Event = new(PoolMint)
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
func (it *PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMint represents a Mint event raised by the Pool contract.
type PoolMint struct {
	Sender              common.Address
	StockAndMoneyAmount *big.Int
	To                  common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xbcad3d7d3dfccb90d49c6063bf70f828901fefc88937d90af74e58e6e55bc39d.
//
// Solidity: event Mint(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolMintIterator{contract: _Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xbcad3d7d3dfccb90d49c6063bf70f828901fefc88937d90af74e58e6e55bc39d.
//
// Solidity: event Mint(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMint)
				if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xbcad3d7d3dfccb90d49c6063bf70f828901fefc88937d90af74e58e6e55bc39d.
//
// Solidity: event Mint(address indexed sender, uint256 stockAndMoneyAmount, address indexed to)
func (_Pool *PoolFilterer) ParseMint(log types.Log) (*PoolMint, error) {
	event := new(PoolMint)
	if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pool contract.
type PoolSyncIterator struct {
	Event *PoolSync // Event containing the contract specifics and raw log

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
func (it *PoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSync)
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
		it.Event = new(PoolSync)
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
func (it *PoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSync represents a Sync event raised by the Pool contract.
type PoolSync struct {
	ReserveStockAndMoney *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x8a0df8ef054fae2c3d2d19a7b322e864870cc9fd3cb07fb9526309c596244bf4.
//
// Solidity: event Sync(uint256 reserveStockAndMoney)
func (_Pool *PoolFilterer) FilterSync(opts *bind.FilterOpts) (*PoolSyncIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &PoolSyncIterator{contract: _Pool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x8a0df8ef054fae2c3d2d19a7b322e864870cc9fd3cb07fb9526309c596244bf4.
//
// Solidity: event Sync(uint256 reserveStockAndMoney)
func (_Pool *PoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PoolSync) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSync)
				if err := _Pool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x8a0df8ef054fae2c3d2d19a7b322e864870cc9fd3cb07fb9526309c596244bf4.
//
// Solidity: event Sync(uint256 reserveStockAndMoney)
func (_Pool *PoolFilterer) ParseSync(log types.Log) (*PoolSync, error) {
	event := new(PoolSync)
	if err := _Pool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	return event, nil
}
