// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"AddedBlackLists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"name\":\"RemovedBlackLists\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_evilUser\",\"type\":\"address[]\"}],\"name\":\"addBlackLists\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerToSet\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"mixedAddrVal\",\"type\":\"uint256[]\"}],\"name\":\"multiTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_clearedUser\",\"type\":\"address[]\"}],\"name\":\"removeBlackLists\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x60a06040523480156200001157600080fd5b506040516200184c3803806200184c833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060409081526020828101519290910151600080546001600160a01b0319163317905586519294509250620001d39160069187019062000230565b508251620001e990600790602086019062000230565b5060f81b7fff000000000000000000000000000000000000000000000000000000000000001660805260058190553360009081526003602052604090205550620002cc9050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200027357805160ff1916838001178555620002a3565b82800160010185558215620002a3579182015b82811115620002a357825182559160200191906001019062000286565b50620002b1929150620002b5565b5090565b5b80821115620002b15760008155600101620002b6565b60805160f81c611562620002ea6000398061067f52506115626000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c80638da5cb5b116100b8578063a9059cbb1161007c578063a9059cbb14610462578063b33fcc7a1461048e578063bc5920ba146104fe578063d4ee1d9014610506578063dd62ed3e1461050e578063e47d60601461053c57610137565b80638da5cb5b1461037457806395d89b411461039857806398313ca1146103a0578063a457c2d714610410578063a6f9dae11461043c57610137565b806339509351116100ff578063395093511461026757806342966c6814610293578063440878b4146102b257806370a082311461032257806379cc67901461034857610137565b806306fdde031461013c578063095ea7b3146101b957806318160ddd146101f957806323b872dd14610213578063313ce56714610249575b600080fd5b610144610562565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017e578181015183820152602001610166565b50505050905090810190601f1680156101ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360408110156101cf57600080fd5b506001600160a01b0381351690602001356105f8565b604080519115158252519081900360200190f35b61020161060e565b60408051918252519081900360200190f35b6101e56004803603606081101561022957600080fd5b506001600160a01b03813581169160208101359091169060400135610614565b61025161067d565b6040805160ff9092168252519081900360200190f35b6101e56004803603604081101561027d57600080fd5b506001600160a01b0381351690602001356106a1565b6102b0600480360360208110156102a957600080fd5b50356106d7565b005b6102b0600480360360208110156102c857600080fd5b8101906020810181356401000000008111156102e357600080fd5b8201836020820111156102f557600080fd5b8035906020019184602083028401116401000000008311171561031757600080fd5b5090925090506106e4565b6102016004803603602081101561033857600080fd5b50356001600160a01b03166107eb565b6102b06004803603604081101561035e57600080fd5b506001600160a01b038135169060200135610806565b61037c610852565b604080516001600160a01b039092168252519081900360200190f35b610144610861565b6102b0600480360360208110156103b657600080fd5b8101906020810181356401000000008111156103d157600080fd5b8201836020820111156103e357600080fd5b8035906020019184602083028401116401000000008311171561040557600080fd5b5090925090506108c2565b6101e56004803603604081101561042657600080fd5b506001600160a01b0381351690602001356109c0565b6102b06004803603602081101561045257600080fd5b50356001600160a01b0316610a0f565b6101e56004803603604081101561047857600080fd5b506001600160a01b038135169060200135610b59565b6101e5600480360360208110156104a457600080fd5b8101906020810181356401000000008111156104bf57600080fd5b8201836020820111156104d157600080fd5b803590602001918460208302840111640100000000831117156104f357600080fd5b509092509050610b66565b6102b0610bc7565b61037c610c65565b6102016004803603604081101561052457600080fd5b506001600160a01b0381358116916020013516610c74565b6101e56004803603602081101561055257600080fd5b50356001600160a01b0316610c9f565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105ee5780601f106105c3576101008083540402835291602001916105ee565b820191906000526020600020905b8154815290600101906020018083116105d157829003601f168201915b5050505050905090565b6000610605338484610cbd565b50600192915050565b60055490565b6000610621848484610da9565b610673843361066e856040518060600160405280602f8152602001611282602f91396001600160a01b038a1660009081526004602090815260408083203384529091529020549190610f05565b610cbd565b5060019392505050565b7f000000000000000000000000000000000000000000000000000000000000000090565b3360008181526004602090815260408083206001600160a01b0387168452909152812054909161060591859061066e9086610f9c565b6106e13382610ffd565b50565b6000546001600160a01b0316331461072d5760405162461bcd60e51b815260040180806020018281038252602581526020018061125d6025913960400191505060405180910390fd5b60005b818110156107835760016002600085858581811061074a57fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff1916911515919091179055600101610730565b507fe9b487e36d4279af2f695c62e0e801c5d45017d600659f53f7889ebea442dd77828260405180806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f19169092018290039550909350505050a15050565b6001600160a01b031660009081526003602052604090205490565b6000610836826040518060600160405280602b81526020016113b4602b913961082f8633610c74565b9190610f05565b9050610843833383610cbd565b61084d8383610ffd565b505050565b6000546001600160a01b031690565b60078054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156105ee5780601f106105c3576101008083540402835291602001916105ee565b6000546001600160a01b0316331461090b5760405162461bcd60e51b815260040180806020018281038252602581526020018061125d6025913960400191505060405180910390fd5b60005b81811015610958576002600084848481811061092657fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff1916905560010161090e565b507f98feb506554d800c8b843a711fb4a52365bf0230afbf171e4e0c8491ad57eda4828260405180806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f19169092018290039550909350505050a15050565b6000610605338461066e856040518060600160405280602c8152602001611231602c91393360009081526004602090815260408083206001600160a01b038d1684529091529020549190610f05565b6000546001600160a01b03163314610a585760405162461bcd60e51b815260040180806020018281038252602581526020018061125d6025913960400191505060405180910390fd5b6001600160a01b038116610a9d5760405162461bcd60e51b815260040180806020018281038252602381526020018061120e6023913960400191505060405180910390fd5b6000546001600160a01b0382811691161415610aea5760405162461bcd60e51b81526004018080602001828103825260348152602001806114176034913960400191505060405180910390fd5b6001546001600160a01b0382811691161415610b375760405162461bcd60e51b81526004018080602001828103825260388152602001806113df6038913960400191505060405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000610605338484610da9565b6000805b828110156106735760006060858584818110610b8257fe5b90506020020135901c90506000858584818110610b9b57fe5b905060200201356bffffffffffffffffffffffff169050610bbd338383610da9565b5050600101610b6a565b6001546001600160a01b03163314610c105760405162461bcd60e51b81526004018080602001828103825260298152602001806113096029913960400191505060405180910390fd5b600154600080546001600160a01b0319166001600160a01b03909216918217905560408051918252517fa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf369181900360200190a1565b6001546001600160a01b031690565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b6001600160a01b031660009081526002602052604090205460ff1690565b6001600160a01b038316610d025760405162461bcd60e51b815260040180806020018281038252602b815260200180611481602b913960400191505060405180910390fd5b6001600160a01b038216610d475760405162461bcd60e51b81526004018080602001828103825260298152602001806114ac6029913960400191505060405180910390fd5b6001600160a01b03808416600081815260046020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b038316610dee5760405162461bcd60e51b815260040180806020018281038252602c815260200180611332602c913960400191505060405180910390fd5b6001600160a01b038216610e335760405162461bcd60e51b815260040180806020018281038252602a8152602001806112df602a913960400191505060405180910390fd5b610e3d83836110f8565b610e7a816040518060600160405280602d815260200161135e602d91396001600160a01b0386166000908152600360205260409020549190610f05565b6001600160a01b038085166000908152600360205260408082209390935590841681522054610ea99082610f9c565b6001600160a01b0380841660008181526003602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610f945760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610f59578181015183820152602001610f41565b50505050905090810190601f168015610f865780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610ff6576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166110425760405162461bcd60e51b81526004018080602001828103825260288152602001806114d56028913960400191505060405180910390fd5b61104d8260006110f8565b61108a8160405180606001604052806029815260200161138b602991396001600160a01b0385166000908152600360205260409020549190610f05565b6001600160a01b0383166000908152600360205260409020556005546110b090826111cb565b6005556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b61110133610c9f565b1561113d5760405162461bcd60e51b815260040180806020018281038252603681526020018061144b6036913960400191505060405180910390fd5b61114682610c9f565b156111825760405162461bcd60e51b81526004018080602001828103825260308152602001806114fd6030913960400191505060405180910390fd5b61118b81610c9f565b156111c75760405162461bcd60e51b815260040180806020018281038252602e8152602001806112b1602e913960400191505060405180910390fd5b5050565b6000610ff683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610f0556fe4f6e6553776170546f6b656e3a20494e56414c49445f4f574e45525f414444524553534f6e6553776170546f6b656e3a204445435245415345445f414c4c4f57414e43455f42454c4f575f5a45524f4f6e6553776170546f6b656e3a204d53475f53454e4445525f49535f4e4f545f4f574e45524f6e6553776170546f6b656e3a205452414e534645525f414d4f554e545f455843454544535f414c4c4f57414e43454f6e6553776170546f6b656e3a20544f5f49535f424c41434b4c49535445445f42595f544f4b454e5f4f574e45524f6e6553776170546f6b656e3a205452414e534645525f544f5f5448455f5a45524f5f414444524553534f6e6553776170546f6b656e3a204d53475f53454e4445525f49535f4e4f545f4e45575f4f574e45524f6e6553776170546f6b656e3a205452414e534645525f46524f4d5f5448455f5a45524f5f414444524553534f6e6553776170546f6b656e3a205452414e534645525f414d4f554e545f455843454544535f42414c414e43454f6e6553776170546f6b656e3a204255524e5f414d4f554e545f455843454544535f42414c414e43454f6e6553776170546f6b656e3a204255524e5f414d4f554e545f455843454544535f414c4c4f57414e43454f6e6553776170546f6b656e3a204e45575f4f574e45525f49535f5448455f53414d455f41535f43555252454e545f4e45575f4f574e45524f6e6553776170546f6b656e3a204e45575f4f574e45525f49535f5448455f53414d455f41535f43555252454e545f4f574e45524f6e6553776170546f6b656e3a204d53475f53454e4445525f49535f424c41434b4c49535445445f42595f544f4b454e5f4f574e45524f6e6553776170546f6b656e3a20415050524f56455f46524f4d5f5448455f5a45524f5f414444524553534f6e6553776170546f6b656e3a20415050524f56455f544f5f5448455f5a45524f5f414444524553534f6e6553776170546f6b656e3a204255524e5f46524f4d5f5448455f5a45524f5f414444524553534f6e6553776170546f6b656e3a2046524f4d5f49535f424c41434b4c49535445445f42595f544f4b454e5f4f574e4552a2646970667358221220984c4fde1ac71d7c0c5225058173c32c622ca353b0ec382260fab21822ad8ccd64736f6c634300060c0033"

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, supply *big.Int, decimals uint8) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, name, symbol, supply, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address user) view returns(bool)
func (_Token *TokenCaller) IsBlackListed(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isBlackListed", user)
	return *ret0, err
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address user) view returns(bool)
func (_Token *TokenSession) IsBlackListed(user common.Address) (bool, error) {
	return _Token.Contract.IsBlackListed(&_Token.CallOpts, user)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address user) view returns(bool)
func (_Token *TokenCallerSession) IsBlackListed(user common.Address) (bool, error) {
	return _Token.Contract.IsBlackListed(&_Token.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "newOwner")
	return *ret0, err
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_Token *TokenCallerSession) NewOwner() (common.Address, error) {
	return _Token.Contract.NewOwner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// AddBlackLists is a paid mutator transaction binding the contract method 0x440878b4.
//
// Solidity: function addBlackLists(address[] _evilUser) returns()
func (_Token *TokenTransactor) AddBlackLists(opts *bind.TransactOpts, _evilUser []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addBlackLists", _evilUser)
}

// AddBlackLists is a paid mutator transaction binding the contract method 0x440878b4.
//
// Solidity: function addBlackLists(address[] _evilUser) returns()
func (_Token *TokenSession) AddBlackLists(_evilUser []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddBlackLists(&_Token.TransactOpts, _evilUser)
}

// AddBlackLists is a paid mutator transaction binding the contract method 0x440878b4.
//
// Solidity: function addBlackLists(address[] _evilUser) returns()
func (_Token *TokenTransactorSession) AddBlackLists(_evilUser []common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddBlackLists(&_Token.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Token *TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Token *TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.BurnFrom(&_Token.TransactOpts, account, amount)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address ownerToSet) returns()
func (_Token *TokenTransactor) ChangeOwner(opts *bind.TransactOpts, ownerToSet common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "changeOwner", ownerToSet)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address ownerToSet) returns()
func (_Token *TokenSession) ChangeOwner(ownerToSet common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, ownerToSet)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address ownerToSet) returns()
func (_Token *TokenTransactorSession) ChangeOwner(ownerToSet common.Address) (*types.Transaction, error) {
	return _Token.Contract.ChangeOwner(&_Token.TransactOpts, ownerToSet)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xb33fcc7a.
//
// Solidity: function multiTransfer(uint256[] mixedAddrVal) returns(bool)
func (_Token *TokenTransactor) MultiTransfer(opts *bind.TransactOpts, mixedAddrVal []*big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "multiTransfer", mixedAddrVal)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xb33fcc7a.
//
// Solidity: function multiTransfer(uint256[] mixedAddrVal) returns(bool)
func (_Token *TokenSession) MultiTransfer(mixedAddrVal []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiTransfer(&_Token.TransactOpts, mixedAddrVal)
}

// MultiTransfer is a paid mutator transaction binding the contract method 0xb33fcc7a.
//
// Solidity: function multiTransfer(uint256[] mixedAddrVal) returns(bool)
func (_Token *TokenTransactorSession) MultiTransfer(mixedAddrVal []*big.Int) (*types.Transaction, error) {
	return _Token.Contract.MultiTransfer(&_Token.TransactOpts, mixedAddrVal)
}

// RemoveBlackLists is a paid mutator transaction binding the contract method 0x98313ca1.
//
// Solidity: function removeBlackLists(address[] _clearedUser) returns()
func (_Token *TokenTransactor) RemoveBlackLists(opts *bind.TransactOpts, _clearedUser []common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "removeBlackLists", _clearedUser)
}

// RemoveBlackLists is a paid mutator transaction binding the contract method 0x98313ca1.
//
// Solidity: function removeBlackLists(address[] _clearedUser) returns()
func (_Token *TokenSession) RemoveBlackLists(_clearedUser []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveBlackLists(&_Token.TransactOpts, _clearedUser)
}

// RemoveBlackLists is a paid mutator transaction binding the contract method 0x98313ca1.
//
// Solidity: function removeBlackLists(address[] _clearedUser) returns()
func (_Token *TokenTransactorSession) RemoveBlackLists(_clearedUser []common.Address) (*types.Transaction, error) {
	return _Token.Contract.RemoveBlackLists(&_Token.TransactOpts, _clearedUser)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Token *TokenTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Token *TokenSession) UpdateOwner() (*types.Transaction, error) {
	return _Token.Contract.UpdateOwner(&_Token.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_Token *TokenTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _Token.Contract.UpdateOwner(&_Token.TransactOpts)
}

// TokenAddedBlackListsIterator is returned from FilterAddedBlackLists and is used to iterate over the raw logs and unpacked data for AddedBlackLists events raised by the Token contract.
type TokenAddedBlackListsIterator struct {
	Event *TokenAddedBlackLists // Event containing the contract specifics and raw log

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
func (it *TokenAddedBlackListsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenAddedBlackLists)
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
		it.Event = new(TokenAddedBlackLists)
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
func (it *TokenAddedBlackListsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenAddedBlackListsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenAddedBlackLists represents a AddedBlackLists event raised by the Token contract.
type TokenAddedBlackLists struct {
	Arg0 []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackLists is a free log retrieval operation binding the contract event 0xe9b487e36d4279af2f695c62e0e801c5d45017d600659f53f7889ebea442dd77.
//
// Solidity: event AddedBlackLists(address[] arg0)
func (_Token *TokenFilterer) FilterAddedBlackLists(opts *bind.FilterOpts) (*TokenAddedBlackListsIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "AddedBlackLists")
	if err != nil {
		return nil, err
	}
	return &TokenAddedBlackListsIterator{contract: _Token.contract, event: "AddedBlackLists", logs: logs, sub: sub}, nil
}

// WatchAddedBlackLists is a free log subscription operation binding the contract event 0xe9b487e36d4279af2f695c62e0e801c5d45017d600659f53f7889ebea442dd77.
//
// Solidity: event AddedBlackLists(address[] arg0)
func (_Token *TokenFilterer) WatchAddedBlackLists(opts *bind.WatchOpts, sink chan<- *TokenAddedBlackLists) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "AddedBlackLists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenAddedBlackLists)
				if err := _Token.contract.UnpackLog(event, "AddedBlackLists", log); err != nil {
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

// ParseAddedBlackLists is a log parse operation binding the contract event 0xe9b487e36d4279af2f695c62e0e801c5d45017d600659f53f7889ebea442dd77.
//
// Solidity: event AddedBlackLists(address[] arg0)
func (_Token *TokenFilterer) ParseAddedBlackLists(log types.Log) (*TokenAddedBlackLists, error) {
	event := new(TokenAddedBlackLists)
	if err := _Token.contract.UnpackLog(event, "AddedBlackLists", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Token contract.
type TokenOwnerChangedIterator struct {
	Event *TokenOwnerChanged // Event containing the contract specifics and raw log

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
func (it *TokenOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenOwnerChanged)
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
		it.Event = new(TokenOwnerChanged)
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
func (it *TokenOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenOwnerChanged represents a OwnerChanged event raised by the Token contract.
type TokenOwnerChanged struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address arg0)
func (_Token *TokenFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*TokenOwnerChangedIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &TokenOwnerChangedIterator{contract: _Token.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address arg0)
func (_Token *TokenFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *TokenOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenOwnerChanged)
				if err := _Token.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xa2ea9883a321a3e97b8266c2b078bfeec6d50c711ed71f874a90d500ae2eaf36.
//
// Solidity: event OwnerChanged(address arg0)
func (_Token *TokenFilterer) ParseOwnerChanged(log types.Log) (*TokenOwnerChanged, error) {
	event := new(TokenOwnerChanged)
	if err := _Token.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenRemovedBlackListsIterator is returned from FilterRemovedBlackLists and is used to iterate over the raw logs and unpacked data for RemovedBlackLists events raised by the Token contract.
type TokenRemovedBlackListsIterator struct {
	Event *TokenRemovedBlackLists // Event containing the contract specifics and raw log

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
func (it *TokenRemovedBlackListsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRemovedBlackLists)
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
		it.Event = new(TokenRemovedBlackLists)
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
func (it *TokenRemovedBlackListsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRemovedBlackListsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRemovedBlackLists represents a RemovedBlackLists event raised by the Token contract.
type TokenRemovedBlackLists struct {
	Arg0 []common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackLists is a free log retrieval operation binding the contract event 0x98feb506554d800c8b843a711fb4a52365bf0230afbf171e4e0c8491ad57eda4.
//
// Solidity: event RemovedBlackLists(address[] arg0)
func (_Token *TokenFilterer) FilterRemovedBlackLists(opts *bind.FilterOpts) (*TokenRemovedBlackListsIterator, error) {

	logs, sub, err := _Token.contract.FilterLogs(opts, "RemovedBlackLists")
	if err != nil {
		return nil, err
	}
	return &TokenRemovedBlackListsIterator{contract: _Token.contract, event: "RemovedBlackLists", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackLists is a free log subscription operation binding the contract event 0x98feb506554d800c8b843a711fb4a52365bf0230afbf171e4e0c8491ad57eda4.
//
// Solidity: event RemovedBlackLists(address[] arg0)
func (_Token *TokenFilterer) WatchRemovedBlackLists(opts *bind.WatchOpts, sink chan<- *TokenRemovedBlackLists) (event.Subscription, error) {

	logs, sub, err := _Token.contract.WatchLogs(opts, "RemovedBlackLists")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRemovedBlackLists)
				if err := _Token.contract.UnpackLog(event, "RemovedBlackLists", log); err != nil {
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

// ParseRemovedBlackLists is a log parse operation binding the contract event 0x98feb506554d800c8b843a711fb4a52365bf0230afbf171e4e0c8491ad57eda4.
//
// Solidity: event RemovedBlackLists(address[] arg0)
func (_Token *TokenFilterer) ParseRemovedBlackLists(log types.Log) (*TokenRemovedBlackLists, error) {
	event := new(TokenRemovedBlackLists)
	if err := _Token.contract.UnpackLog(event, "RemovedBlackLists", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
