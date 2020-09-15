// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package callee

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

// CalleeABI is the input ABI used to generate the binding from.
const CalleeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Callee is an auto generated Go binding around an Ethereum contract.
type Callee struct {
	CalleeCaller     // Read-only binding to the contract
	CalleeTransactor // Write-only binding to the contract
	CalleeFilterer   // Log filterer for contract events
}

// CalleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalleeSession struct {
	Contract     *Callee           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalleeCallerSession struct {
	Contract *CalleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CalleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalleeTransactorSession struct {
	Contract     *CalleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalleeRaw struct {
	Contract *Callee // Generic contract binding to access the raw methods on
}

// CalleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalleeCallerRaw struct {
	Contract *CalleeCaller // Generic read-only contract binding to access the raw methods on
}

// CalleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalleeTransactorRaw struct {
	Contract *CalleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCallee creates a new instance of Callee, bound to a specific deployed contract.
func NewCallee(address common.Address, backend bind.ContractBackend) (*Callee, error) {
	contract, err := bindCallee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Callee{CalleeCaller: CalleeCaller{contract: contract}, CalleeTransactor: CalleeTransactor{contract: contract}, CalleeFilterer: CalleeFilterer{contract: contract}}, nil
}

// NewCalleeCaller creates a new read-only instance of Callee, bound to a specific deployed contract.
func NewCalleeCaller(address common.Address, caller bind.ContractCaller) (*CalleeCaller, error) {
	contract, err := bindCallee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeCaller{contract: contract}, nil
}

// NewCalleeTransactor creates a new write-only instance of Callee, bound to a specific deployed contract.
func NewCalleeTransactor(address common.Address, transactor bind.ContractTransactor) (*CalleeTransactor, error) {
	contract, err := bindCallee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalleeTransactor{contract: contract}, nil
}

// NewCalleeFilterer creates a new log filterer instance of Callee, bound to a specific deployed contract.
func NewCalleeFilterer(address common.Address, filterer bind.ContractFilterer) (*CalleeFilterer, error) {
	contract, err := bindCallee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalleeFilterer{contract: contract}, nil
}

// bindCallee binds a generic wrapper to an already deployed contract.
func bindCallee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalleeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.CalleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.CalleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Callee *CalleeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Callee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Callee *CalleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Callee *CalleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Callee.Contract.contract.Transact(opts, method, params...)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Callee *CalleeTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Callee.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Callee *CalleeSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Callee.Contract.UniswapV2Call(&_Callee.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Callee *CalleeTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Callee.Contract.UniswapV2Call(&_Callee.TransactOpts, sender, amount0, amount1, data)
}
