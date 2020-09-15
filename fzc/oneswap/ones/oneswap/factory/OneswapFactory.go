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
const FactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ones\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairLogic\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBPS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"getTokensFromPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ones\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_bps\",\"type\":\"uint32\"}],\"name\":\"setFeeBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implLogic\",\"type\":\"address\"}],\"name\":\"setPairLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"tokensToPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FactoryBin is the compiled bytecode used for deploying new contracts.
var FactoryBin = "0x60c06040526001805463ffffffff60a01b1916601960a11b17905534801561002657600080fd5b50604051610fde380380610fde8339818101604052608081101561004957600080fd5b50805160208201516040830151606093840151600180546001600160a01b03199081166001600160a01b039687161790915583861b6001600160601b03199081166080529583901b90951660a052600280549095169084161790935581169116610f096100d5600039806106f9528061094452508061031252806103795280610a185250610f096000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636c87813e11610097578063a2e74af611610066578063a2e74af61461022b578063c0225aa114610251578063ca215225146102a6578063f46901ed146102cc576100f5565b80636c87813e146101ab57806382dfdce4146101e357806387dcb20e1461021b5780638ef582b014610223576100f5565b80631a1c6e53116100d35780631a1c6e531461012e5780631e3dd18b1461014f5780633be9d9301461016c578063574f2ba314610191576100f5565b8063017e7e58146100fa578063094b74151461011e57806312d43a5114610126575b600080fd5b6101026102f2565b604080516001600160a01b039092168252519081900360200190f35b610102610301565b610102610310565b610136610334565b6040805163ffffffff9092168252519081900360200190f35b6101026004803603602081101561016557600080fd5b5035610347565b61018f6004803603602081101561018257600080fd5b503563ffffffff1661036e565b005b61019961046d565b60408051918252519081900360200190f35b610102600480360360608110156101c157600080fd5b506001600160a01b038135811691602081013590911690604001351515610473565b610102600480360360608110156101f957600080fd5b506001600160a01b0381358116916020810135909116906040013515156104de565b610102610933565b610102610942565b61018f6004803603602081101561024157600080fd5b50356001600160a01b0316610966565b6102776004803603602081101561026757600080fd5b50356001600160a01b03166109e3565b60405180836001600160a01b03168152602001826001600160a01b031681526020019250505060405180910390f35b61018f600480360360208110156102bc57600080fd5b50356001600160a01b0316610a0d565b61018f600480360360208110156102e257600080fd5b50356001600160a01b0316610aac565b6000546001600160a01b031681565b6001546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600154600160a01b900463ffffffff1681565b6005818154811061035457fe5b6000918252602090912001546001600160a01b0316905081565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103eb576040805162461bcd60e51b815260206004820152601f60248201527f4f6e6553776170466163746f72793a205345545445525f4d49534d4154434800604482015290519081900360640190fd5b60328163ffffffff161115610447576040805162461bcd60e51b815260206004820181905260248201527f4f6e6553776170466163746f72793a204250535f4f55545f4f465f52414e4745604482015290519081900360640190fd5b6001805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b60055490565b604080516bffffffffffffffffffffffff19606095861b81166020808401919091529490951b909416603485015290151560f81b604884015280516029818503018152604990930181528251928201929092206000908152600490915220546001600160a01b031690565b6000826001600160a01b0316846001600160a01b031614156105315760405162461bcd60e51b8152600401808060200182810382526023815260200180610e8a6023913960400191505060405180910390fd5b600061053c84610b29565b9050600061054986610b29565b90508060171015801561055a575060015b6105955760405162461bcd60e51b815260040180806020018281038252602c815260200180610e5e602c913960400191505060405180910390fd5b6000600482106105a6575060031981015b6001806000848611156105cf57846013018611156105c6575060016105cf565b848603600a0a92505b858511156105f357856013018511156105ea575060016105f3565b858503600a0a91505b80156106305760405162461bcd60e51b8152600401808060200182810382526027815260200180610ead6027913960400191505060405180910390fd5b6040805160608c811b6bffffffffffffffffffffffff19908116602080850191909152918d901b1660348301528a151560f81b6048830152825180830360290181526049909201835281519181019190912060008181526004909252919020546001600160a01b0316156106eb576040805162461bcd60e51b815260206004820152601b60248201527f4f6e6553776170466163746f72793a20504149525f4558495354530000000000604482015290519081900360640190fd5b6000818c8c8c89600a0a89897f000000000000000000000000000000000000000000000000000000000000000060405161072490610bb1565b6001600160a01b039788168152958716602087015293151560408087019190915267ffffffffffffffff9384166060870152918316608086015290911660a0840152921660c08201529051829181900360e001906000f590508015801561078f573d6000803e3d6000fd5b5090508098506005899080600181540180825580915050600190039060005260206000200160009091909190916101000a8154816001600160a01b0302191690836001600160a01b03160217905550886004600084815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808d6001600160a01b031681526020018c6001600160a01b0316815250600360008b6001600160a01b03166001600160a01b0316815260200190815260200160002060008201518160000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060208201518160010160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550905050886001600160a01b03167fd6b196957cdc774a3f446a0dcd127f49e6352507784aa7fd6137c1d13326b6eb8d8d8d60405180846001600160a01b03168152602001836001600160a01b031681526020018215158152602001935050505060405180910390a250505050505050509392505050565b6002546001600160a01b031681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6001546001600160a01b031633146109c1576040805162461bcd60e51b815260206004820152601960248201527827b732a9bbb0b82330b1ba37b93c9d102327a92124a22222a760391b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b039081166000908152600360205260409020805460019091015490821692911690565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a8a576040805162461bcd60e51b815260206004820152601f60248201527f4f6e6553776170466163746f72793a205345545445525f4d49534d4154434800604482015290519081900360640190fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314610b07576040805162461bcd60e51b815260206004820152601960248201527827b732a9bbb0b82330b1ba37b93c9d102327a92124a22222a760391b604482015290519081900360640190fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b038216610b4157506012610bac565b816001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610b7a57600080fd5b505afa158015610b8e573d6000803e3d6000fd5b505050506040513d6020811015610ba457600080fd5b505160ff1690505b919050565b61029f80610bbf8339019056fe61012060405234801561001157600080fd5b5060405161029f38038061029f833981810160405260e081101561003457600080fd5b50805160208201516040830151606084015160808086015160a08088015160c098890151339094526001600160a01b03808816909252818816909852821660e052949593949293919291906000851561008b575060015b604090811b6001600160401b0395861617811b9385169390931790921b9216919091176101005250506001600a55505060805160a05160c05160e051610100516101aa6100f56000398060a552508060835250806061525080603f525080601d52506101aa6000f3fe60806040523661000b57005b604080516343ee590760e11b815290517f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000916000916001600160a01b038816916387dcb20e91600480830192602092919082900301818787803b1580156100f957600080fd5b505af115801561010d573d6000803e3d6000fd5b505050506040513d602081101561012357600080fd5b505160405190915036806000833781810188815260208101889052604081018790526060810186905260800184905260a0016000808284865af43d9150816000843e808015610170578284f35b8284fdfea2646970667358221220f4f3e45c612e755091cd2baea22bae196b1ad7993acfd369375cddd1831dbdb564736f6c634300060c00334f6e6553776170466163746f72793a2053544f434b5f444543494d414c535f4e4f545f535550504f525445444f6e6553776170466163746f72793a204944454e544943414c5f4144445245535345534f6e6553776170466163746f72793a20444543494d414c535f444946465f544f4f5f4c41524745a264697066735822122086c983d43a8770169e4515e240525c30cf4007ee2d5e4ed954de0d0da61ee08764736f6c634300060c0033"

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _feeToSetter common.Address, _gov common.Address, _ones common.Address, _pairLogic common.Address) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactoryBin), backend, _feeToSetter, _gov, _ones, _pairLogic)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

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

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "allPairs", arg0)
	return *ret0, err
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.AllPairs(&_Factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Factory *FactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Factory.Contract.AllPairs(&_Factory.CallOpts, arg0)
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

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Factory *FactoryCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "gov")
	return *ret0, err
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Factory *FactorySession) Gov() (common.Address, error) {
	return _Factory.Contract.Gov(&_Factory.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Factory *FactoryCallerSession) Gov() (common.Address, error) {
	return _Factory.Contract.Gov(&_Factory.CallOpts)
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Factory *FactoryCaller) Ones(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "ones")
	return *ret0, err
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Factory *FactorySession) Ones() (common.Address, error) {
	return _Factory.Contract.Ones(&_Factory.CallOpts)
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Factory *FactoryCallerSession) Ones() (common.Address, error) {
	return _Factory.Contract.Ones(&_Factory.CallOpts)
}

// PairLogic is a free data retrieval call binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() view returns(address)
func (_Factory *FactoryCaller) PairLogic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "pairLogic")
	return *ret0, err
}

// PairLogic is a free data retrieval call binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() view returns(address)
func (_Factory *FactorySession) PairLogic() (common.Address, error) {
	return _Factory.Contract.PairLogic(&_Factory.CallOpts)
}

// PairLogic is a free data retrieval call binding the contract method 0x87dcb20e.
//
// Solidity: function pairLogic() view returns(address)
func (_Factory *FactoryCallerSession) PairLogic() (common.Address, error) {
	return _Factory.Contract.PairLogic(&_Factory.CallOpts)
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

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 _bps) returns()
func (_Factory *FactoryTransactor) SetFeeBPS(opts *bind.TransactOpts, _bps uint32) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeBPS", _bps)
}

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 _bps) returns()
func (_Factory *FactorySession) SetFeeBPS(_bps uint32) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeBPS(&_Factory.TransactOpts, _bps)
}

// SetFeeBPS is a paid mutator transaction binding the contract method 0x3be9d930.
//
// Solidity: function setFeeBPS(uint32 _bps) returns()
func (_Factory *FactoryTransactorSession) SetFeeBPS(_bps uint32) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeBPS(&_Factory.TransactOpts, _bps)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Factory *FactoryTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Factory *FactorySession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeTo(&_Factory.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Factory *FactoryTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeTo(&_Factory.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Factory *FactoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Factory *FactorySession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeToSetter(&_Factory.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Factory *FactoryTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Factory.Contract.SetFeeToSetter(&_Factory.TransactOpts, _feeToSetter)
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
