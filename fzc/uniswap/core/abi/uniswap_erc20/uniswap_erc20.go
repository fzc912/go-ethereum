// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap_erc20

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

// UniswapErc20ABI is the input ABI used to generate the binding from.
const UniswapErc20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniswapErc20 is an auto generated Go binding around an Ethereum contract.
type UniswapErc20 struct {
	UniswapErc20Caller     // Read-only binding to the contract
	UniswapErc20Transactor // Write-only binding to the contract
	UniswapErc20Filterer   // Log filterer for contract events
}

// UniswapErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapErc20Session struct {
	Contract     *UniswapErc20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapErc20CallerSession struct {
	Contract *UniswapErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// UniswapErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapErc20TransactorSession struct {
	Contract     *UniswapErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapErc20Raw struct {
	Contract *UniswapErc20 // Generic contract binding to access the raw methods on
}

// UniswapErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapErc20CallerRaw struct {
	Contract *UniswapErc20Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapErc20TransactorRaw struct {
	Contract *UniswapErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapErc20 creates a new instance of UniswapErc20, bound to a specific deployed contract.
func NewUniswapErc20(address common.Address, backend bind.ContractBackend) (*UniswapErc20, error) {
	contract, err := bindUniswapErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20{UniswapErc20Caller: UniswapErc20Caller{contract: contract}, UniswapErc20Transactor: UniswapErc20Transactor{contract: contract}, UniswapErc20Filterer: UniswapErc20Filterer{contract: contract}}, nil
}

// NewUniswapErc20Caller creates a new read-only instance of UniswapErc20, bound to a specific deployed contract.
func NewUniswapErc20Caller(address common.Address, caller bind.ContractCaller) (*UniswapErc20Caller, error) {
	contract, err := bindUniswapErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20Caller{contract: contract}, nil
}

// NewUniswapErc20Transactor creates a new write-only instance of UniswapErc20, bound to a specific deployed contract.
func NewUniswapErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapErc20Transactor, error) {
	contract, err := bindUniswapErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20Transactor{contract: contract}, nil
}

// NewUniswapErc20Filterer creates a new log filterer instance of UniswapErc20, bound to a specific deployed contract.
func NewUniswapErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapErc20Filterer, error) {
	contract, err := bindUniswapErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20Filterer{contract: contract}, nil
}

// bindUniswapErc20 binds a generic wrapper to an already deployed contract.
func bindUniswapErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapErc20 *UniswapErc20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniswapErc20.Contract.UniswapErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapErc20 *UniswapErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapErc20.Contract.UniswapErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapErc20 *UniswapErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapErc20.Contract.UniswapErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapErc20 *UniswapErc20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UniswapErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapErc20 *UniswapErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapErc20 *UniswapErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapErc20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapErc20 *UniswapErc20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapErc20 *UniswapErc20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapErc20.Contract.DOMAINSEPARATOR(&_UniswapErc20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_UniswapErc20 *UniswapErc20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _UniswapErc20.Contract.DOMAINSEPARATOR(&_UniswapErc20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_UniswapErc20 *UniswapErc20Caller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_UniswapErc20 *UniswapErc20Session) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapErc20.Contract.PERMITTYPEHASH(&_UniswapErc20.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_UniswapErc20 *UniswapErc20CallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _UniswapErc20.Contract.PERMITTYPEHASH(&_UniswapErc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.Allowance(&_UniswapErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_UniswapErc20 *UniswapErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.Allowance(&_UniswapErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.BalanceOf(&_UniswapErc20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.BalanceOf(&_UniswapErc20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_UniswapErc20 *UniswapErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_UniswapErc20 *UniswapErc20Session) Decimals() (uint8, error) {
	return _UniswapErc20.Contract.Decimals(&_UniswapErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_UniswapErc20 *UniswapErc20CallerSession) Decimals() (uint8, error) {
	return _UniswapErc20.Contract.Decimals(&_UniswapErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_UniswapErc20 *UniswapErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_UniswapErc20 *UniswapErc20Session) Name() (string, error) {
	return _UniswapErc20.Contract.Name(&_UniswapErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_UniswapErc20 *UniswapErc20CallerSession) Name() (string, error) {
	return _UniswapErc20.Contract.Name(&_UniswapErc20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Caller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "nonces", owner)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20Session) Nonces(owner common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.Nonces(&_UniswapErc20.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_UniswapErc20 *UniswapErc20CallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _UniswapErc20.Contract.Nonces(&_UniswapErc20.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_UniswapErc20 *UniswapErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_UniswapErc20 *UniswapErc20Session) Symbol() (string, error) {
	return _UniswapErc20.Contract.Symbol(&_UniswapErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_UniswapErc20 *UniswapErc20CallerSession) Symbol() (string, error) {
	return _UniswapErc20.Contract.Symbol(&_UniswapErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapErc20 *UniswapErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UniswapErc20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapErc20 *UniswapErc20Session) TotalSupply() (*big.Int, error) {
	return _UniswapErc20.Contract.TotalSupply(&_UniswapErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_UniswapErc20 *UniswapErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _UniswapErc20.Contract.TotalSupply(&_UniswapErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Approve(&_UniswapErc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Approve(&_UniswapErc20.TransactOpts, spender, value)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapErc20 *UniswapErc20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapErc20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapErc20 *UniswapErc20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Permit(&_UniswapErc20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_UniswapErc20 *UniswapErc20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Permit(&_UniswapErc20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Transfer(&_UniswapErc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.Transfer(&_UniswapErc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.TransferFrom(&_UniswapErc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_UniswapErc20 *UniswapErc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UniswapErc20.Contract.TransferFrom(&_UniswapErc20.TransactOpts, from, to, value)
}

// UniswapErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the UniswapErc20 contract.
type UniswapErc20ApprovalIterator struct {
	Event *UniswapErc20Approval // Event containing the contract specifics and raw log

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
func (it *UniswapErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapErc20Approval)
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
		it.Event = new(UniswapErc20Approval)
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
func (it *UniswapErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapErc20Approval represents a Approval event raised by the UniswapErc20 contract.
type UniswapErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*UniswapErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UniswapErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20ApprovalIterator{contract: _UniswapErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UniswapErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _UniswapErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapErc20Approval)
				if err := _UniswapErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) ParseApproval(log types.Log) (*UniswapErc20Approval, error) {
	event := new(UniswapErc20Approval)
	if err := _UniswapErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the UniswapErc20 contract.
type UniswapErc20TransferIterator struct {
	Event *UniswapErc20Transfer // Event containing the contract specifics and raw log

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
func (it *UniswapErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapErc20Transfer)
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
		it.Event = new(UniswapErc20Transfer)
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
func (it *UniswapErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapErc20Transfer represents a Transfer event raised by the UniswapErc20 contract.
type UniswapErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UniswapErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapErc20TransferIterator{contract: _UniswapErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UniswapErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UniswapErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapErc20Transfer)
				if err := _UniswapErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_UniswapErc20 *UniswapErc20Filterer) ParseTransfer(log types.Log) (*UniswapErc20Transfer, error) {
	event := new(UniswapErc20Transfer)
	if err := _UniswapErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
