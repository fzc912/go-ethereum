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
const RouterABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"moneyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"money\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOnlySwap\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountStockDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStockMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountStock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoney\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevKey\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"stockAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"limitOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStockMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoneyMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountStock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountMoney\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// RouterBin is the compiled bytecode used for deploying new contracts.
var RouterBin = "0x60a060405234801561001057600080fd5b506040516117783803806117788339818101604052602081101561003357600080fd5b5051606081901b6001600160601b0319166080526001600160a01b0316611702610076600039806105ee52806106725280610a845280610aab52506117026000f3fe60806040526004361061004a5760003560e01c806301f3cc5a1461004f5780635a47ddc3146100a357806396c92f5e14610121578063a70002601461018d578063c45a015514610278575b600080fd5b6100a1600480360360e081101561006557600080fd5b5080351515906001600160a01b036020820135169060408101359060608101359063ffffffff6080820135169060a08101359060c001356102a9565b005b61010360048036036101208110156100ba57600080fd5b506001600160a01b0381358116916020810135821691604082013515159160608101359160808201359160a08101359160c08201359160e0810135909116906101000135610503565b60408051938452602084019290925282820152519081900360600190f35b34801561012d57600080fd5b50610174600480360360c081101561014457600080fd5b506001600160a01b0381358116916020810135916040820135916060810135916080820135169060a00135610823565b6040805192835260208301919091528051918290030190f35b610228600480360360c08110156101a357600080fd5b6001600160a01b0382351691602081013591604082013591908101906080810160608201356401000000008111156101da57600080fd5b8201836020820111156101ec57600080fd5b8035906020019184602083028401116401000000008311171561020e57600080fd5b91935091506001600160a01b03813516906020013561089d565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561026457818101518382015260200161024c565b505050509050019250505060405180910390f35b34801561028457600080fd5b5061028d610a82565b604080516001600160a01b039092168252519081900360200190f35b80428110156102f8576040805162461bcd60e51b815260206004820152601660248201527513db9954ddd85c149bdd5d195c8e881156141254915160521b604482015290519081900360640190fd5b60008061030489610aa6565b6040805163c2dab57b60e01b815267ffffffffffffffff8916600482015263ffffffff8b166024820152815193955091935060009283926001600160a01b038e169263c2dab57b92604480840193829003018186803b15801561036657600080fd5b505afa15801561037a573d6000803e3d6000fd5b505050506040513d604081101561039057600080fd5b50805160209091015190925090508b15610401576001600160a01b038316156103f05734156103f05760405162461bcd60e51b81526004018080602001828103825260228152602001806115cb6022913960400191505060405180910390fd5b6103fc83338d84610bc2565b610459565b6001600160a01b0384161561044d57341561044d5760405162461bcd60e51b81526004018080602001828103825260228152602001806115cb6022913960400191505060405180910390fd5b61045984338d85610bc2565b50506040805163c018f41360e01b81528b1515600482015233602482015267ffffffffffffffff8716604482015263ffffffff808a1660648301528816608482015268ffffffffffffffffff8a1660a482015290516001600160a01b038b169163c018f4139160c480830192600092919082900301818387803b1580156104df57600080fd5b505af11580156104f3573d6000803e3d6000fd5b5050505050505050505050505050565b60008060008342811015610557576040805162461bcd60e51b815260206004820152601660248201527513db9954ddd85c149bdd5d195c8e881156141254915160521b604482015290519081900360640190fd5b6001600160a01b038d161580159061057757506001600160a01b038c1615155b156105b95734156105b95760405162461bcd60e51b81526004018080602001828103825260228152602001806115cb6022913960400191505060405180910390fd5b60408051633643c09f60e11b81526001600160a01b038f811660048301528e811660248301528d1515604483015291516000927f00000000000000000000000000000000000000000000000000000000000000001691636c87813e916064808301926020929190829003018186803b15801561063457600080fd5b505afa158015610648573d6000803e3d6000fd5b505050506040513d602081101561065e57600080fd5b505190506001600160a01b038116610729577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166382dfdce48f8f8f6040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b0316815260200182151581526020019350505050602060405180830381600087803b1580156106fa57600080fd5b505af115801561070e573d6000803e3d6000fd5b505050506040513d602081101561072457600080fd5b505190505b610736818c8c8c8c610ea4565b90955093506107478e338388610bc2565b6107538d338387610bc2565b806001600160a01b0316636a627842886040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b1580156107a257600080fd5b505af11580156107b6573d6000803e3d6000fd5b505050506040513d60208110156107cc57600080fd5b5051604080518781526020810187905280820183905290519194507ff75993dbe1645872cbbea6395e1feebee76b435baf0e4d62d7eac269c6f57b24919081900360600190a1505099509950999650505050505050565b6000808242811015610875576040805162461bcd60e51b815260206004820152601660248201527513db9954ddd85c149bdd5d195c8e881156141254915160521b604482015290519081900360640190fd5b61087e89610aa6565b505061088d898989898961100e565b909a909950975050505050505050565b606081428110156108ee576040805162461bcd60e51b815260206004820152601660248201527513db9954ddd85c149bdd5d195c8e881156141254915160521b604482015290519081900360640190fd5b6001600160a01b0389161561093a57341561093a5760405162461bcd60e51b81526004018080602001828103825260228152602001806115cb6022913960400191505060405180910390fd5b6001851015610990576040805162461bcd60e51b815260206004820152601b60248201527f4f6e6553776170526f757465723a20494e56414c49445f504154480000000000604482015290519081900360640190fd5b6109b5868660008181106109a057fe5b905060200201356001600160a01b0316610aa6565b50506109df8933888860008181106109c957fe5b905060200201356001600160a01b03168b610bc2565b610a1f89898888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a925061119d915050565b915086828787905081518110610a3157fe5b60200260200101511015610a765760405162461bcd60e51b81526004018080602001828103825260298152602001806116816029913960400191505060405180910390fd5b50979650505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c0225aa1846040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050604080518083038186803b158015610b1557600080fd5b505afa158015610b29573d6000803e3d6000fd5b505050506040513d6040811015610b3f57600080fd5b50805160209091015190925090506001600160a01b038216151580610b6c57506001600160a01b03811615155b610bbd576040805162461bcd60e51b815260206004820152601d60248201527f4f6e6553776170526f757465723a20504149525f4e4f545f4558495354000000604482015290519081900360640190fd5b915091565b6001600160a01b038416610bf557610bda8282611383565b3481811115610bef57610bef33838303611383565b50610e9e565b6000846001600160a01b03166370a08231846040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610c4457600080fd5b505afa158015610c58573d6000803e3d6000fd5b505050506040513d6020811015610c6e57600080fd5b5051604080516001600160a01b0387811660248301528681166044830152606480830187905283518084039091018152608490920183526020820180516001600160e01b03166323b872dd60e01b17815292518251949550600094606094928b16939282918083835b60208310610cf65780518252601f199092019160209182019101610cd7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610d58576040519150601f19603f3d011682016040523d82523d6000602084013e610d5d565b606091505b5091509150818015610d8b575080511580610d8b5750808060200190516020811015610d8857600080fd5b50515b610dc65760405162461bcd60e51b815260040180806020018281038252602381526020018061165e6023913960400191505060405180910390fd5b6000876001600160a01b03166370a08231876040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610e1557600080fd5b505afa158015610e29573d6000803e3d6000fd5b505050506040513d6020811015610e3f57600080fd5b505190508385018114610e99576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6553776170526f757465723a205452414e534645525f4641494c45440000604482015290519081900360640190fd5b505050505b50505050565b600080600080886001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015610ee357600080fd5b505afa158015610ef7573d6000803e3d6000fd5b505050506040513d6060811015610f0d57600080fd5b5080516020909101516001600160701b03918216935016905081158015610f32575080155b15610f4257879350869250611002565b6000610f4f89848461147b565b9050878111610fa25785811015610f975760405162461bcd60e51b81526004018080602001828103825260288152602001806116366028913960400191505060405180910390fd5b889450925082611000565b6000610faf89848661147b565b905089811115610fbb57fe5b87811015610ffa5760405162461bcd60e51b81526004018080602001828103825260288152602001806115ed6028913960400191505060405180910390fd5b94508793505b505b50509550959350505050565b604080516323b872dd60e01b81523360048201526001600160a01b03871660248201819052604482018790529151600092839290916323b872dd9160648082019260209290919082900301818787803b15801561106a57600080fd5b505af115801561107e573d6000803e3d6000fd5b505050506040513d602081101561109457600080fd5b50506040805163226bf2d160e21b81526001600160a01b0385811660048301528251908a16926389afcb4492602480820193918290030181600087803b1580156110dd57600080fd5b505af11580156110f1573d6000803e3d6000fd5b505050506040513d604081101561110757600080fd5b5080516020909101519092509050848210156111545760405162461bcd60e51b81526004018080602001828103825260288152602001806115ed6028913960400191505060405180910390fd5b838110156111935760405162461bcd60e51b81526004018080602001828103825260288152602001806116366028913960400191505060405180910390fd5b9550959350505050565b6060825160010167ffffffffffffffff811180156111ba57600080fd5b506040519080825280602002602001820160405280156111e4578160200160208202803683370190505b50905083816000815181106111f557fe5b60200260200101818152505060005b835181101561137a57600080600186510383106112235784600161123d565b85836001018151811061123257fe5b602002602001015160005b9150915085838151811061124d57fe5b60200260200101516001600160a01b031663a6e81533898487878151811061127157fe5b60200260200101516040518463ffffffff1660e01b815260040180846001600160a01b03168152602001836001600160a01b03168152602001826001600160701b031681526020019350505050602060405180830381600087803b1580156112d857600080fd5b505af11580156112ec573d6000803e3d6000fd5b505050506040513d602081101561130257600080fd5b5051845185906001860190811061131557fe5b602002602001018181525050806113705760008061134588868151811061133857fe5b6020026020010151610aa6565b91509150896001600160a01b0316826001600160a01b03161415611369578061136b565b815b995050505b5050600101611204565b50949350505050565b604080516000808252602082019092526001600160a01b0384169083906040518082805190602001908083835b602083106113cf5780518252601f1990920191602091820191016113b0565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611431576040519150601f19603f3d011682016040523d82523d6000602084013e611436565b606091505b50509050806114765760405162461bcd60e51b81526004018080602001828103825260238152602001806116aa6023913960400191505060405180910390fd5b505050565b60008084116114bb5760405162461bcd60e51b81526004018080602001828103825260228152602001806115846022913960400191505060405180910390fd5b6000831180156114cb5750600082115b6115065760405162461bcd60e51b81526004018080602001828103825260258152602001806115a66025913960400191505060405180910390fd5b826115118584611521565b8161151857fe5b04949350505050565b6000826115305750600061157d565b8282028284828161153d57fe5b041461157a5760405162461bcd60e51b81526004018080602001828103825260218152602001806116156021913960400191505060405180910390fd5b90505b9291505056fe4f6e6553776170526f757465723a20494e53554646494349454e545f414d4f554e544f6e6553776170526f757465723a20494e53554646494349454e545f4c49515549444954594f6e6553776170526f757465723a204e4f545f454e5445525f4554485f56414c55454f6e6553776170526f757465723a20494e53554646494349454e545f53544f434b5f414d4f554e54536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f774f6e6553776170526f757465723a20494e53554646494349454e545f4d4f4e45595f414d4f554e544f6e6553776170526f757465723a205452414e534645525f46524f4d5f4641494c45444f6e6553776170526f757465723a20494e53554646494349454e545f4f55545055545f414d4f554e545472616e7366657248656c7065723a204554485f5452414e534645525f4641494c4544a264697066735822122022dc59f6297de2004ecb69a0ab505d13b31b0595d0ec41fb21184623fed6b30764736f6c634300060c0033"

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RouterBin), backend, _factory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

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
// Solidity: function factory() view returns(address)
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
// Solidity: function factory() view returns(address)
func (_Router *RouterSession) Factory() (common.Address, error) {
	return _Router.Contract.Factory(&_Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
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
