// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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

// RouterABI is the input ABI used to generate the binding from.
const RouterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moneyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountStockDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStockMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountStock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoney\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"limitOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStockMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountStock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoney\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_Router *RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Router.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() pure returns(address)
func (_Router *RouterCallerSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address stock, address money, bool isOnlySwap, uint256 amountStockDesired, uint256 amountMoneyDesired, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) payable returns(uint256 amountStock, uint256 amountMoney, uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, stock common.Address, money common.Address, isOnlySwap bool, amountStockDesired *big.Int, amountMoneyDesired *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", stock, money, isOnlySwap, amountStockDesired, amountMoneyDesired, amountStockMin, amountMoneyMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address stock, address money, bool isOnlySwap, uint256 amountStockDesired, uint256 amountMoneyDesired, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) payable returns(uint256 amountStock, uint256 amountMoney, uint256 liquidity)
func (_Router *RouterSession) AddLiquidity(stock common.Address, money common.Address, isOnlySwap bool, amountStockDesired *big.Int, amountMoneyDesired *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, stock, money, isOnlySwap, amountStockDesired, amountMoneyDesired, amountStockMin, amountMoneyMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x5a47ddc3.
//
// Solidity: function addLiquidity(address stock, address money, bool isOnlySwap, uint256 amountStockDesired, uint256 amountMoneyDesired, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) payable returns(uint256 amountStock, uint256 amountMoney, uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidity(stock common.Address, money common.Address, isOnlySwap bool, amountStockDesired *big.Int, amountMoneyDesired *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, stock, money, isOnlySwap, amountStockDesired, amountMoneyDesired, amountStockMin, amountMoneyMin, to, deadline)
}

// LimitOrder is a paid mutator transaction binding the contract method 0x01f3cc5a.
//
// Solidity: function limitOrder(bool isBuy, address pair, uint256 prevKey, uint256 price, uint32 id, uint256 stockAmount, uint256 deadline) payable returns()
func (_Router *RouterTransactor) LimitOrder(opts *bind.TransactOpts, isBuy bool, pair common.Address, prevKey *big.Int, price *big.Int, id uint32, stockAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "limitOrder", isBuy, pair, prevKey, price, id, stockAmount, deadline)
}

// LimitOrder is a paid mutator transaction binding the contract method 0x01f3cc5a.
//
// Solidity: function limitOrder(bool isBuy, address pair, uint256 prevKey, uint256 price, uint32 id, uint256 stockAmount, uint256 deadline) payable returns()
func (_Router *RouterSession) LimitOrder(isBuy bool, pair common.Address, prevKey *big.Int, price *big.Int, id uint32, stockAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.LimitOrder(&_Router.TransactOpts, isBuy, pair, prevKey, price, id, stockAmount, deadline)
}

// LimitOrder is a paid mutator transaction binding the contract method 0x01f3cc5a.
//
// Solidity: function limitOrder(bool isBuy, address pair, uint256 prevKey, uint256 price, uint32 id, uint256 stockAmount, uint256 deadline) payable returns()
func (_Router *RouterTransactorSession) LimitOrder(isBuy bool, pair common.Address, prevKey *big.Int, price *big.Int, id uint32, stockAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.LimitOrder(&_Router.TransactOpts, isBuy, pair, prevKey, price, id, stockAmount, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x96c92f5e.
//
// Solidity: function removeLiquidity(address pair, uint256 liquidity, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) returns(uint256 amountStock, uint256 amountMoney)
func (_Router *RouterTransactor) RemoveLiquidity(opts *bind.TransactOpts, pair common.Address, liquidity *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeLiquidity", pair, liquidity, amountStockMin, amountMoneyMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x96c92f5e.
//
// Solidity: function removeLiquidity(address pair, uint256 liquidity, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) returns(uint256 amountStock, uint256 amountMoney)
func (_Router *RouterSession) RemoveLiquidity(pair common.Address, liquidity *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, pair, liquidity, amountStockMin, amountMoneyMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x96c92f5e.
//
// Solidity: function removeLiquidity(address pair, uint256 liquidity, uint256 amountStockMin, uint256 amountMoneyMin, address to, uint256 deadline) returns(uint256 amountStock, uint256 amountMoney)
func (_Router *RouterTransactorSession) RemoveLiquidity(pair common.Address, liquidity *big.Int, amountStockMin *big.Int, amountMoneyMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.RemoveLiquidity(&_Router.TransactOpts, pair, liquidity, amountStockMin, amountMoneyMin, to, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0xa7000260.
//
// Solidity: function swapToken(address token, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterTransactor) SwapToken(opts *bind.TransactOpts, token common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapToken", token, amountIn, amountOutMin, path, to, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0xa7000260.
//
// Solidity: function swapToken(address token, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterSession) SwapToken(token common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapToken(&_Router.TransactOpts, token, amountIn, amountOutMin, path, to, deadline)
}

// SwapToken is a paid mutator transaction binding the contract method 0xa7000260.
//
// Solidity: function swapToken(address token, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Router *RouterTransactorSession) SwapToken(token common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SwapToken(&_Router.TransactOpts, token, amountIn, amountOutMin, path, to, deadline)
}

// RouterAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Router contract.
type RouterAddLiquidityIterator struct {
	Event *RouterAddLiquidity // Event containing the contract specifics and raw log

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
func (it *RouterAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterAddLiquidity)
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
		it.Event = new(RouterAddLiquidity)
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
func (it *RouterAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterAddLiquidity represents a AddLiquidity event raised by the Router contract.
type RouterAddLiquidity struct {
	StockAmount *big.Int
	MoneyAmount *big.Int
	Liquidity   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xf75993dbe1645872cbbea6395e1feebee76b435baf0e4d62d7eac269c6f57b24.
//
// Solidity: event AddLiquidity(uint256 stockAmount, uint256 moneyAmount, uint256 liquidity)
func (_Router *RouterFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*RouterAddLiquidityIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &RouterAddLiquidityIterator{contract: _Router.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xf75993dbe1645872cbbea6395e1feebee76b435baf0e4d62d7eac269c6f57b24.
//
// Solidity: event AddLiquidity(uint256 stockAmount, uint256 moneyAmount, uint256 liquidity)
func (_Router *RouterFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *RouterAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterAddLiquidity)
				if err := _Router.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xf75993dbe1645872cbbea6395e1feebee76b435baf0e4d62d7eac269c6f57b24.
//
// Solidity: event AddLiquidity(uint256 stockAmount, uint256 moneyAmount, uint256 liquidity)
func (_Router *RouterFilterer) ParseAddLiquidity(log types.Log) (*RouterAddLiquidity, error) {
	event := new(RouterAddLiquidity)
	if err := _Router.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RouterPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Router contract.
type RouterPairCreatedIterator struct {
	Event *RouterPairCreated // Event containing the contract specifics and raw log

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
func (it *RouterPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterPairCreated)
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
		it.Event = new(RouterPairCreated)
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
func (it *RouterPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterPairCreated represents a PairCreated event raised by the Router contract.
type RouterPairCreated struct {
	Pair       common.Address
	Stock      common.Address
	Money      common.Address
	IsOnlySwap bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb.
//
// Solidity: event PairCreated(address indexed pair, address stock, address money, bool isOnlySwap)
func (_Router *RouterFilterer) FilterPairCreated(opts *bind.FilterOpts, pair []common.Address) (*RouterPairCreatedIterator, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "PairCreated", pairRule)
	if err != nil {
		return nil, err
	}
	return &RouterPairCreatedIterator{contract: _Router.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb.
//
// Solidity: event PairCreated(address indexed pair, address stock, address money, bool isOnlySwap)
func (_Router *RouterFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *RouterPairCreated, pair []common.Address) (event.Subscription, error) {

	var pairRule []interface{}
	for _, pairItem := range pair {
		pairRule = append(pairRule, pairItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "PairCreated", pairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterPairCreated)
				if err := _Router.contract.UnpackLog(event, "PairCreated", log); err != nil {
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
func (_Router *RouterFilterer) ParsePairCreated(log types.Log) (*RouterPairCreated, error) {
	event := new(RouterPairCreated)
	if err := _Router.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
