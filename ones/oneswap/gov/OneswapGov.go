// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

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

// GovABI is the input ABI used to generate the binding from.
const GovABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ones\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"opinion\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"AddVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"deadline\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NewFundsProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"deadline\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"feeBPS\",\"type\":\"uint32\"}],\"name\":\"NewParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"deadline\",\"type\":\"uint32\"}],\"name\":\"NewTextProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"deadline\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pairLogic\",\"type\":\"address\"}],\"name\":\"NewUpgradeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"opinion\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"NewVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"opinion\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"Revote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalID\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pass\",\"type\":\"bool\"}],\"name\":\"TallyResult\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ones\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalInfo\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"deadline\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"totalYes\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"totalNo\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"totalDeposit\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fundsAmt\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"submitFundsProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"feeBPS\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"submitParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"submitTextProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pairLogic\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"submitUpgradeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tally\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"opinion\",\"type\":\"uint8\"},{\"internalType\":\"uint112\",\"name\":\"voteAmt\",\"type\":\"uint112\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"voterInfo\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"votedProposalID\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"votedOpinion\",\"type\":\"uint8\"},{\"internalType\":\"uint112\",\"name\":\"votedAmt\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"depositedAmt\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint112\",\"name\":\"amt\",\"type\":\"uint112\"}],\"name\":\"withdrawOnes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GovBin is the compiled bytecode used for deploying new contracts.
var GovBin = "0x60e060405234801561001057600080fd5b506040516122333803806122338339818101604052602081101561003357600080fd5b50516001600160601b0319606082901b166080526040805163313ce56760e01b815290516000916001600160a01b0384169163313ce56791600480820192602092909190829003018186803b15801561008b57600080fd5b505afa15801561009f573d6000803e3d6000fd5b505050506040513d60208110156100b557600080fd5b5051624c4b4060ff909116600a0a90810260a0526103e80260c052505060805160601c60a05160c05161210f61012460003980611d3c525080610b97525080610c0e5280610e415280611297528061143b528061156852806119345280611baa5280611eb4525061210f6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638ef582b0116100665780638ef582b0146104b1578063a9527d7a146104d5578063cdf58d35146105ec578063f53747ac14610649578063fce2c4541461066f5761009e565b806309a91df0146100a35780632b865a2a14610118578063410673e51461024b5780635aea4f4b1461025357806372b1661914610385575b600080fd5b6100ab61069e565b6040805162ffffff909a168a526001600160a01b0398891660208b015260ff9097168988015263ffffffff909516606089015292909516608087015260a08601526001600160701b0393841660c0860152831660e085015290911661010083015251908190036101200190f35b610249600480360360c081101561012e57600080fd5b810190602081018135600160201b81111561014857600080fd5b82018360208201111561015a57600080fd5b803590602001918460018302840111600160201b8311171561017b57600080fd5b919390929091602081019035600160201b81111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111600160201b831117156101cb57600080fd5b919390929091602081019035600160201b8111156101e857600080fd5b8201836020820111156101fa57600080fd5b803590602001918460018302840111600160201b8311171561021b57600080fd5b919350915080356001600160a01b0390811691602081013590911690604001356001600160701b0316610701565b005b61024961089e565b610249600480360360c081101561026957600080fd5b810190602081018135600160201b81111561028357600080fd5b82018360208201111561029557600080fd5b803590602001918460018302840111600160201b831117156102b657600080fd5b919390929091602081019035600160201b8111156102d357600080fd5b8201836020820111156102e557600080fd5b803590602001918460018302840111600160201b8311171561030657600080fd5b919390929091602081019035600160201b81111561032357600080fd5b82018360208201111561033557600080fd5b803590602001918460018302840111600160201b8311171561035657600080fd5b919350915080356001600160a01b031690602081013563ffffffff1690604001356001600160701b0316610a10565b610249600480360360c081101561039b57600080fd5b810190602081018135600160201b8111156103b557600080fd5b8201836020820111156103c757600080fd5b803590602001918460018302840111600160201b831117156103e857600080fd5b919390929091602081019035600160201b81111561040557600080fd5b82018360208201111561041757600080fd5b803590602001918460018302840111600160201b8311171561043857600080fd5b919390929091602081019035600160201b81111561045557600080fd5b82018360208201111561046757600080fd5b803590602001918460018302840111600160201b8311171561048857600080fd5b919350915080356001600160a01b031690602081013590604001356001600160701b0316610b8f565b6104b9610e3f565b604080516001600160a01b039092168252519081900360200190f35b610249600480360360808110156104eb57600080fd5b810190602081018135600160201b81111561050557600080fd5b82018360208201111561051757600080fd5b803590602001918460018302840111600160201b8311171561053857600080fd5b919390929091602081019035600160201b81111561055557600080fd5b82018360208201111561056757600080fd5b803590602001918460018302840111600160201b8311171561058857600080fd5b919390929091602081019035600160201b8111156105a557600080fd5b8201836020820111156105b757600080fd5b803590602001918460018302840111600160201b831117156105d857600080fd5b9193509150356001600160701b0316610e63565b6106126004803603602081101561060257600080fd5b50356001600160a01b0316610fd6565b6040805162ffffff909516855260ff90931660208501526001600160701b0391821684840152166060830152519081900360800190f35b6102496004803603602081101561065f57600080fd5b50356001600160701b031661105a565b6102496004803603604081101561068557600080fd5b50803560ff1690602001356001600160701b031661130f565b60005460025460015460035460045462ffffff8516956001600160a01b039485169560ff63010000008204169563ffffffff600160201b83041695600160401b909204169390926001600160701b0380831693600160701b909304811692911690565b6001600160a01b03821661075c576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6553776170476f763a20494e56414c49445f504149525f4c4f4749430000604482015290519081900360640190fd5b610772600384846001600160a01b0316846113d9565b7f620d77127ee7d32615eb70c371bd279ee1c4f664cb785673bdec31a49898f03a60008054906101000a900462ffffff168a8a8a8a8a8a600060049054906101000a900463ffffffff168b8b604051808b62ffffff1681526020018060200180602001806020018763ffffffff168152602001866001600160a01b03168152602001856001600160a01b0316815260200184810384528d8d82818152602001925080828437600083820152601f01601f191690910185810384528b815260200190508b8b80828437600083820152601f01601f191690910185810383528981526020019050898980828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a16108936001826116fc565b505050505050505050565b6000546301000000900460ff166108f6576040805162461bcd60e51b815260206004820152601760248201527613db9954ddd85c11dbdd8e881393d7d41493d413d4d053604a1b604482015290519081900360640190fd5b60005442600160201b90910463ffffffff16111561095b576040805162461bcd60e51b815260206004820152601860248201527f4f6e6553776170476f763a205354494c4c5f564f54494e470000000000000000604482015290519081900360640190fd5b6003546000546001546002546001600160701b03600160701b850481169416939093119260ff6301000000840416926001600160a01b03600160401b909104811691166109a6611b58565b84156109bc576109b7848385611b96565b6109c5565b6109c581611d3a565b6000546040805162ffffff9092168252861515602083015280517f709c5e70cd3dcf1f39400585b4da49bf555425a933cbafa65e5746f81bf3ca4c9281900390910190a15050505050565b603263ffffffff83161115610a6c576040805162461bcd60e51b815260206004820152601b60248201527f4f6e6553776170476f763a20494e56414c49445f4645455f4250530000000000604482015290519081900360640190fd5b610a7f6002848463ffffffff16846113d9565b6000546040805162ffffff831680825263ffffffff600160201b9094048416608083018190526001600160a01b03881660a084015293861660c083015260e0602083018181529083018d90527f4623cd1863f26cf17b0b0252ebc774ecf43e15c013a840478f00c241f76e5e199491938e938e938e938e938e938e93928e928e92919082016060830161010084018d8d80828437600083820152601f01601f191690910185810384528b815260200190508b8b80828437600083820152601f01601f191690910185810383528981526020019050898980828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a16108936001826116fc565b8115610d1a577f0000000000000000000000000000000000000000000000000000000000000000821115610c0a576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6553776170476f763a2041534b5f544f4f5f4d414e595f46554e44530000604482015290519081900360640190fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610c7957600080fd5b505afa158015610c8d573d6000803e3d6000fd5b505050506040513d6020811015610ca357600080fd5b50516004549091506001600160701b03168082039082118015610cc65750838110155b610d17576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6553776170476f763a20494e53554646494349454e545f46554e44530000604482015290519081900360640190fd5b50505b610d2760018484846113d9565b7fcefa6a6335c5a8e5b325bf02dae880483f6a9e3632e21083f8994b95191bc00c60008054906101000a900462ffffff168a8a8a8a8a8a600060049054906101000a900463ffffffff168b8b604051808b62ffffff1681526020018060200180602001806020018763ffffffff168152602001866001600160a01b0316815260200185815260200184810384528d8d82818152602001925080828437600083820152601f01601f191690910185810384528b815260200190508b8b80828437600083820152601f01601f191690910185810383528981526020019050898980828437600083820152604051601f909101601f19169092018290039f50909d5050505050505050505050505050a16108936001826116fc565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000544263ffffffff600160201b90920491909116620151800110610ecf576040805162461bcd60e51b815260206004820152601860248201527f4f6e6553776170476f763a20434f4f4c494e475f444f574e0000000000000000604482015290519081900360640190fd5b610edd6004600080846113d9565b6000546040805162ffffff8316808252600160201b90930463ffffffff166080820181905260a0602083018181529083018b90527fedb5790b505fab801089c42a4d9c59dedfab325afc38c2f880f43fab318942b094938c938c938c938c938c938c9392909182016060830160c084018b8b80828437600083820152601f01601f191690910185810384528981526020019050898980828437600083820152601f01601f191690910185810383528781526020019050878780828437600083820152604051601f909101601f19169092018290039d50909b505050505050505050505050a1610fcd6001826116fc565b50505050505050565b600080600080610fe4612048565b5050506001600160a01b039092166000908152600560209081526040918290208251608081018452905462ffffff811680835260ff63010000008304169383018490526001600160701b03600160201b83048116958401869052600160901b909204909116606090920182905295919450919250565b611062612048565b503360009081526005602090815260408083208151608081018352905462ffffff81168252630100000080820460ff90811695840195909552600160201b82046001600160701b0390811694840194909452600160901b90910490921660608201529254041615806110df5750600054815162ffffff9182169116105b611128576040805162461bcd60e51b81526020600482015260156024820152744f6e6553776170476f763a20494e5f564f54494e4760581b604482015290519081900360640190fd5b6000826001600160701b0316118015611157575080606001516001600160701b0316826001600160701b031611155b6111925760405162461bcd60e51b81526004018080602001828103825260238152602001806120916023913960400191505060405180910390fd5b600480546001600160701b031981166001600160701b03918216859003821617909155606082018051849003909116908190526111de5733600090815260056020526040812055611268565b33600090815260056020908152604091829020835181549285015193850151606086015162ffffff1990941662ffffff9092169190911763ff0000001916630100000060ff9095169490940293909317640100000000600160901b031916600160201b6001600160701b0394851602176001600160901b0316600160901b93909216929092021790555b6040805163a9059cbb60e01b81523360048201526001600160701b038416602482015290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163a9059cbb9160448083019260209291908290030181600087803b1580156112df57600080fd5b505af11580156112f3573d6000803e3d6000fd5b505050506040513d602081101561130957600080fd5b50505050565b6000546301000000900460ff16611367576040805162461bcd60e51b815260206004820152601760248201527613db9954ddd85c11dbdd8e881393d7d41493d413d4d053604a1b604482015290519081900360640190fd5b60005442600160201b90910463ffffffff16116113cb576040805162461bcd60e51b815260206004820152601c60248201527f4f6e6553776170476f763a20444541444c494e455f5245414348454400000000604482015290519081900360640190fd5b6113d582826116fc565b5050565b600160ff8516108015906113f15750600460ff851611155b61142c5760405162461bcd60e51b81526004018080602001828103825260218152602001806120706021913960400191505060405180910390fd5b60ff8416600414806114ca57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561149257600080fd5b505afa1580156114a6573d6000803e3d6000fd5b505050506040513d60208110156114bc57600080fd5b50516001600160a01b031633145b61151b576040805162461bcd60e51b815260206004820152601a60248201527f4f6e6553776170476f763a204e4f545f4f4e45535f4f574e4552000000000000604482015290519081900360640190fd5b6000546301000000900460ff16156115645760405162461bcd60e51b81526004018080602001828103825260268152602001806120b46026913960400191505060405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115bf57600080fd5b505afa1580156115d3573d6000803e3d6000fd5b505050506040513d60208110156115e957600080fd5b50519050606481046001600160701b03831681111561164f576040805162461bcd60e51b815260206004820181905260248201527f4f6e6553776170476f763a20564f54455f414d4f554e545f544f4f5f4c455353604482015290519081900360640190fd5b50506000805460028054336001600160a01b0319909116179055600193845562ffffff19811662ffffff918216909401169290921763ff0000001916630100000060ff95909516949094029390931767ffffffff000000001916600160201b6203f480420163ffffffff16021768010000000000000000600160e01b031916600160401b6001600160a01b03939093169290920291909117905550600380546001600160e01b0319169055565b60ff82166001118015906117145750600260ff831611155b611765576040805162461bcd60e51b815260206004820152601b60248201527f4f6e6553776170476f763a20494e56414c49445f4f50494e494f4e0000000000604482015290519081900360640190fd5b6000816001600160701b0316116117c3576040805162461bcd60e51b815260206004820152601c60248201527f4f6e6553776170476f763a205a45524f5f564f54455f414d4f554e5400000000604482015290519081900360640190fd5b60008060008060006117d3611ef9565b9450945094509450945060008562ffffff168562ffffff161480156117fe57508760ff168460ff1614155b156118a55760ff84166001141561184e576003546001600160701b038085169116101561182757fe5b600380546001600160701b03808216869003166001600160701b031990911617905561189d565b6003546001600160701b03808516600160701b90920416101561186d57fe5b600380546001600160701b03600160701b808304821687900390911602600160701b600160e01b03199091161790555b506000915060015b826001600160701b0316826001600160701b031610156118c157fe5b8282036001600160701b0316876001600160701b031611156119b357600480546001600160701b031981168585038a036001600160701b0392831681018316919091178355604080516323b872dd60e01b81523394810194909452306024850152918116604484015290519381019390917f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316916323b872dd916064808201926020929091908290030181600087803b15801561198557600080fd5b505af1158015611999573d6000803e3d6000fd5b505050506040513d60208110156119af57600080fd5b5050505b60ff8816600114156119e557600380546001600160701b038082168a01166001600160701b0319909116179055611a14565b600380546001600160701b03600160701b80830482168b0190911602600160701b600160e01b03199091161790555b91860191611a2486898585611f98565b8015611a84576040805162ffffff8816815233602082015260ff8a16818301526001600160701b038916606082015290517f2054d4cdf4c6c9a38470bc7846d035a1dc962fd95a136c48fca4b36b01fb388b9181900360800190a1611b4e565b866001600160701b0316836001600160701b03161115611af8576040805162ffffff8816815233602082015260ff8a16818301526001600160701b038916606082015290517f79cf911a7c6d5d1b5d658aa8a512828c7539341faa788e21ee4b22997cee5ab09181900360800190a1611b4e565b6040805162ffffff8816815233602082015260ff8a16818301526001600160701b038916606082015290517fac44cb0e78b2dd1c33e8aa0f21a40ac00cc0b08f64cd62163ec42809a58e750d9181900360800190a15b5050505050505050565b60008054600182905567ffffffff01000000600160e01b0319169055600280546001600160a01b0319169055600380546001600160e01b0319169055565b60ff831660011415611c51578015611c4c577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663a9059cbb83836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050602060405180830381600087803b158015611c1f57600080fd5b505af1158015611c33573d6000803e3d6000fd5b505050506040513d6020811015611c4957600080fd5b50505b611d35565b60ff831660021415611cc657816001600160a01b0316633be9d930826040518263ffffffff1660e01b8152600401808263ffffffff168152602001915050600060405180830381600087803b158015611ca957600080fd5b505af1158015611cbd573d6000803e3d6000fd5b50505050611d35565b60ff831660031415611d3557816001600160a01b031663ca215225826040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015611d2157600080fd5b505af1158015610fcd573d6000803e3d6000fd5b505050565b7f0000000000000000000000000000000000000000000000000000000000000000611d63612048565b506001600160a01b0382166000908152600560209081526040918290208251608081018452905462ffffff8116825260ff6301000000820416928201929092526001600160701b03600160201b8304811693820193909352600160901b909104821660608201819052909183161115611dde57806060015191505b600480546001600160701b0380821685900381166001600160701b0319909216919091178255606083018051859003821681526001600160a01b03808716600090815260056020908152604080832088518154938a0151838b015197518916600160901b026001600160901b03988a16600160201b02640100000000600160901b031960ff90931663010000000263ff0000001962ffffff90951662ffffff19909816979097179390931695909517161795909516919091179093558251630852cd8d60e31b81529387169484019490945290517f0000000000000000000000000000000000000000000000000000000000000000909116926342966c68926024808201939182900301818387803b158015611d2157600080fd5b6000805462ffffff1690808080611f0e612048565b5050336000908152600560209081526040918290208251608081018452905462ffffff80821680845260ff6301000000840416948401949094526001600160701b03600160201b8304811695840195909552600160901b909104909316606082018190529290919087161415611f905785945080602001519350806040015192505b509091929394565b6040805160808101825262ffffff958616815260ff94851660208083019182526001600160701b03958616838501908152948616606084019081523360009081526005909252939020915182549151945193518616600160901b026001600160901b0394909616600160201b02640100000000600160901b03199590971663010000000263ff000000199190981662ffffff19909216919091171695909517919091169290921791909116179055565b6040805160808101825260008082526020820181905291810182905260608101919091529056fe4f6e6553776170476f763a20494e56414c49445f50524f504f53414c5f545950454f6e6553776170476f763a20494e56414c49445f57495448445241575f414d4f554e544f6e6553776170476f763a204c4153545f50524f504f53414c5f4e4f545f46494e4953484544a26469706673582212201bbbf8f8e2b47d27b7e62396550155f4694205fb8743d8bcf8da4f4fbc7d465c64736f6c634300060c0033"

// DeployGov deploys a new Ethereum contract, binding an instance of Gov to it.
func DeployGov(auth *bind.TransactOpts, backend bind.ContractBackend, _ones common.Address) (common.Address, *types.Transaction, *Gov, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovBin), backend, _ones)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// Gov is an auto generated Go binding around an Ethereum contract.
type Gov struct {
	GovCaller     // Read-only binding to the contract
	GovTransactor // Write-only binding to the contract
	GovFilterer   // Log filterer for contract events
}

// GovCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovSession struct {
	Contract     *Gov              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovCallerSession struct {
	Contract *GovCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovTransactorSession struct {
	Contract     *GovTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovRaw struct {
	Contract *Gov // Generic contract binding to access the raw methods on
}

// GovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovCallerRaw struct {
	Contract *GovCaller // Generic read-only contract binding to access the raw methods on
}

// GovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovTransactorRaw struct {
	Contract *GovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGov creates a new instance of Gov, bound to a specific deployed contract.
func NewGov(address common.Address, backend bind.ContractBackend) (*Gov, error) {
	contract, err := bindGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gov{GovCaller: GovCaller{contract: contract}, GovTransactor: GovTransactor{contract: contract}, GovFilterer: GovFilterer{contract: contract}}, nil
}

// NewGovCaller creates a new read-only instance of Gov, bound to a specific deployed contract.
func NewGovCaller(address common.Address, caller bind.ContractCaller) (*GovCaller, error) {
	contract, err := bindGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovCaller{contract: contract}, nil
}

// NewGovTransactor creates a new write-only instance of Gov, bound to a specific deployed contract.
func NewGovTransactor(address common.Address, transactor bind.ContractTransactor) (*GovTransactor, error) {
	contract, err := bindGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovTransactor{contract: contract}, nil
}

// NewGovFilterer creates a new log filterer instance of Gov, bound to a specific deployed contract.
func NewGovFilterer(address common.Address, filterer bind.ContractFilterer) (*GovFilterer, error) {
	contract, err := bindGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovFilterer{contract: contract}, nil
}

// bindGov binds a generic wrapper to an already deployed contract.
func bindGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.GovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.GovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gov *GovCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gov *GovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gov *GovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gov.Contract.contract.Transact(opts, method, params...)
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Gov *GovCaller) Ones(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gov.contract.Call(opts, out, "ones")
	return *ret0, err
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Gov *GovSession) Ones() (common.Address, error) {
	return _Gov.Contract.Ones(&_Gov.CallOpts)
}

// Ones is a free data retrieval call binding the contract method 0x8ef582b0.
//
// Solidity: function ones() view returns(address)
func (_Gov *GovCallerSession) Ones() (common.Address, error) {
	return _Gov.Contract.Ones(&_Gov.CallOpts)
}

// ProposalInfo is a free data retrieval call binding the contract method 0x09a91df0.
//
// Solidity: function proposalInfo() view returns(uint24 id, address proposer, uint8 _type, uint32 deadline, address addr, uint256 value, uint112 totalYes, uint112 totalNo, uint112 totalDeposit)
func (_Gov *GovCaller) ProposalInfo(opts *bind.CallOpts) (struct {
	Id           *big.Int
	Proposer     common.Address
	Type         uint8
	Deadline     uint32
	Addr         common.Address
	Value        *big.Int
	TotalYes     *big.Int
	TotalNo      *big.Int
	TotalDeposit *big.Int
}, error) {
	ret := new(struct {
		Id           *big.Int
		Proposer     common.Address
		Type         uint8
		Deadline     uint32
		Addr         common.Address
		Value        *big.Int
		TotalYes     *big.Int
		TotalNo      *big.Int
		TotalDeposit *big.Int
	})
	out := ret
	err := _Gov.contract.Call(opts, out, "proposalInfo")
	return *ret, err
}

// ProposalInfo is a free data retrieval call binding the contract method 0x09a91df0.
//
// Solidity: function proposalInfo() view returns(uint24 id, address proposer, uint8 _type, uint32 deadline, address addr, uint256 value, uint112 totalYes, uint112 totalNo, uint112 totalDeposit)
func (_Gov *GovSession) ProposalInfo() (struct {
	Id           *big.Int
	Proposer     common.Address
	Type         uint8
	Deadline     uint32
	Addr         common.Address
	Value        *big.Int
	TotalYes     *big.Int
	TotalNo      *big.Int
	TotalDeposit *big.Int
}, error) {
	return _Gov.Contract.ProposalInfo(&_Gov.CallOpts)
}

// ProposalInfo is a free data retrieval call binding the contract method 0x09a91df0.
//
// Solidity: function proposalInfo() view returns(uint24 id, address proposer, uint8 _type, uint32 deadline, address addr, uint256 value, uint112 totalYes, uint112 totalNo, uint112 totalDeposit)
func (_Gov *GovCallerSession) ProposalInfo() (struct {
	Id           *big.Int
	Proposer     common.Address
	Type         uint8
	Deadline     uint32
	Addr         common.Address
	Value        *big.Int
	TotalYes     *big.Int
	TotalNo      *big.Int
	TotalDeposit *big.Int
}, error) {
	return _Gov.Contract.ProposalInfo(&_Gov.CallOpts)
}

// VoterInfo is a free data retrieval call binding the contract method 0xcdf58d35.
//
// Solidity: function voterInfo(address voter) view returns(uint24 votedProposalID, uint8 votedOpinion, uint112 votedAmt, uint112 depositedAmt)
func (_Gov *GovCaller) VoterInfo(opts *bind.CallOpts, voter common.Address) (struct {
	VotedProposalID *big.Int
	VotedOpinion    uint8
	VotedAmt        *big.Int
	DepositedAmt    *big.Int
}, error) {
	ret := new(struct {
		VotedProposalID *big.Int
		VotedOpinion    uint8
		VotedAmt        *big.Int
		DepositedAmt    *big.Int
	})
	out := ret
	err := _Gov.contract.Call(opts, out, "voterInfo", voter)
	return *ret, err
}

// VoterInfo is a free data retrieval call binding the contract method 0xcdf58d35.
//
// Solidity: function voterInfo(address voter) view returns(uint24 votedProposalID, uint8 votedOpinion, uint112 votedAmt, uint112 depositedAmt)
func (_Gov *GovSession) VoterInfo(voter common.Address) (struct {
	VotedProposalID *big.Int
	VotedOpinion    uint8
	VotedAmt        *big.Int
	DepositedAmt    *big.Int
}, error) {
	return _Gov.Contract.VoterInfo(&_Gov.CallOpts, voter)
}

// VoterInfo is a free data retrieval call binding the contract method 0xcdf58d35.
//
// Solidity: function voterInfo(address voter) view returns(uint24 votedProposalID, uint8 votedOpinion, uint112 votedAmt, uint112 depositedAmt)
func (_Gov *GovCallerSession) VoterInfo(voter common.Address) (struct {
	VotedProposalID *big.Int
	VotedOpinion    uint8
	VotedAmt        *big.Int
	DepositedAmt    *big.Int
}, error) {
	return _Gov.Contract.VoterInfo(&_Gov.CallOpts, voter)
}

// SubmitFundsProposal is a paid mutator transaction binding the contract method 0x72b16619.
//
// Solidity: function submitFundsProposal(string title, string desc, string url, address beneficiary, uint256 fundsAmt, uint112 voteAmt) returns()
func (_Gov *GovTransactor) SubmitFundsProposal(opts *bind.TransactOpts, title string, desc string, url string, beneficiary common.Address, fundsAmt *big.Int, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "submitFundsProposal", title, desc, url, beneficiary, fundsAmt, voteAmt)
}

// SubmitFundsProposal is a paid mutator transaction binding the contract method 0x72b16619.
//
// Solidity: function submitFundsProposal(string title, string desc, string url, address beneficiary, uint256 fundsAmt, uint112 voteAmt) returns()
func (_Gov *GovSession) SubmitFundsProposal(title string, desc string, url string, beneficiary common.Address, fundsAmt *big.Int, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitFundsProposal(&_Gov.TransactOpts, title, desc, url, beneficiary, fundsAmt, voteAmt)
}

// SubmitFundsProposal is a paid mutator transaction binding the contract method 0x72b16619.
//
// Solidity: function submitFundsProposal(string title, string desc, string url, address beneficiary, uint256 fundsAmt, uint112 voteAmt) returns()
func (_Gov *GovTransactorSession) SubmitFundsProposal(title string, desc string, url string, beneficiary common.Address, fundsAmt *big.Int, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitFundsProposal(&_Gov.TransactOpts, title, desc, url, beneficiary, fundsAmt, voteAmt)
}

// SubmitParamProposal is a paid mutator transaction binding the contract method 0x5aea4f4b.
//
// Solidity: function submitParamProposal(string title, string desc, string url, address factory, uint32 feeBPS, uint112 voteAmt) returns()
func (_Gov *GovTransactor) SubmitParamProposal(opts *bind.TransactOpts, title string, desc string, url string, factory common.Address, feeBPS uint32, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "submitParamProposal", title, desc, url, factory, feeBPS, voteAmt)
}

// SubmitParamProposal is a paid mutator transaction binding the contract method 0x5aea4f4b.
//
// Solidity: function submitParamProposal(string title, string desc, string url, address factory, uint32 feeBPS, uint112 voteAmt) returns()
func (_Gov *GovSession) SubmitParamProposal(title string, desc string, url string, factory common.Address, feeBPS uint32, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitParamProposal(&_Gov.TransactOpts, title, desc, url, factory, feeBPS, voteAmt)
}

// SubmitParamProposal is a paid mutator transaction binding the contract method 0x5aea4f4b.
//
// Solidity: function submitParamProposal(string title, string desc, string url, address factory, uint32 feeBPS, uint112 voteAmt) returns()
func (_Gov *GovTransactorSession) SubmitParamProposal(title string, desc string, url string, factory common.Address, feeBPS uint32, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitParamProposal(&_Gov.TransactOpts, title, desc, url, factory, feeBPS, voteAmt)
}

// SubmitTextProposal is a paid mutator transaction binding the contract method 0xa9527d7a.
//
// Solidity: function submitTextProposal(string title, string desc, string url, uint112 voteAmt) returns()
func (_Gov *GovTransactor) SubmitTextProposal(opts *bind.TransactOpts, title string, desc string, url string, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "submitTextProposal", title, desc, url, voteAmt)
}

// SubmitTextProposal is a paid mutator transaction binding the contract method 0xa9527d7a.
//
// Solidity: function submitTextProposal(string title, string desc, string url, uint112 voteAmt) returns()
func (_Gov *GovSession) SubmitTextProposal(title string, desc string, url string, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitTextProposal(&_Gov.TransactOpts, title, desc, url, voteAmt)
}

// SubmitTextProposal is a paid mutator transaction binding the contract method 0xa9527d7a.
//
// Solidity: function submitTextProposal(string title, string desc, string url, uint112 voteAmt) returns()
func (_Gov *GovTransactorSession) SubmitTextProposal(title string, desc string, url string, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitTextProposal(&_Gov.TransactOpts, title, desc, url, voteAmt)
}

// SubmitUpgradeProposal is a paid mutator transaction binding the contract method 0x2b865a2a.
//
// Solidity: function submitUpgradeProposal(string title, string desc, string url, address factory, address pairLogic, uint112 voteAmt) returns()
func (_Gov *GovTransactor) SubmitUpgradeProposal(opts *bind.TransactOpts, title string, desc string, url string, factory common.Address, pairLogic common.Address, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "submitUpgradeProposal", title, desc, url, factory, pairLogic, voteAmt)
}

// SubmitUpgradeProposal is a paid mutator transaction binding the contract method 0x2b865a2a.
//
// Solidity: function submitUpgradeProposal(string title, string desc, string url, address factory, address pairLogic, uint112 voteAmt) returns()
func (_Gov *GovSession) SubmitUpgradeProposal(title string, desc string, url string, factory common.Address, pairLogic common.Address, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitUpgradeProposal(&_Gov.TransactOpts, title, desc, url, factory, pairLogic, voteAmt)
}

// SubmitUpgradeProposal is a paid mutator transaction binding the contract method 0x2b865a2a.
//
// Solidity: function submitUpgradeProposal(string title, string desc, string url, address factory, address pairLogic, uint112 voteAmt) returns()
func (_Gov *GovTransactorSession) SubmitUpgradeProposal(title string, desc string, url string, factory common.Address, pairLogic common.Address, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.SubmitUpgradeProposal(&_Gov.TransactOpts, title, desc, url, factory, pairLogic, voteAmt)
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns()
func (_Gov *GovTransactor) Tally(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "tally")
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns()
func (_Gov *GovSession) Tally() (*types.Transaction, error) {
	return _Gov.Contract.Tally(&_Gov.TransactOpts)
}

// Tally is a paid mutator transaction binding the contract method 0x410673e5.
//
// Solidity: function tally() returns()
func (_Gov *GovTransactorSession) Tally() (*types.Transaction, error) {
	return _Gov.Contract.Tally(&_Gov.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xfce2c454.
//
// Solidity: function vote(uint8 opinion, uint112 voteAmt) returns()
func (_Gov *GovTransactor) Vote(opts *bind.TransactOpts, opinion uint8, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "vote", opinion, voteAmt)
}

// Vote is a paid mutator transaction binding the contract method 0xfce2c454.
//
// Solidity: function vote(uint8 opinion, uint112 voteAmt) returns()
func (_Gov *GovSession) Vote(opinion uint8, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, opinion, voteAmt)
}

// Vote is a paid mutator transaction binding the contract method 0xfce2c454.
//
// Solidity: function vote(uint8 opinion, uint112 voteAmt) returns()
func (_Gov *GovTransactorSession) Vote(opinion uint8, voteAmt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.Vote(&_Gov.TransactOpts, opinion, voteAmt)
}

// WithdrawOnes is a paid mutator transaction binding the contract method 0xf53747ac.
//
// Solidity: function withdrawOnes(uint112 amt) returns()
func (_Gov *GovTransactor) WithdrawOnes(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _Gov.contract.Transact(opts, "withdrawOnes", amt)
}

// WithdrawOnes is a paid mutator transaction binding the contract method 0xf53747ac.
//
// Solidity: function withdrawOnes(uint112 amt) returns()
func (_Gov *GovSession) WithdrawOnes(amt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.WithdrawOnes(&_Gov.TransactOpts, amt)
}

// WithdrawOnes is a paid mutator transaction binding the contract method 0xf53747ac.
//
// Solidity: function withdrawOnes(uint112 amt) returns()
func (_Gov *GovTransactorSession) WithdrawOnes(amt *big.Int) (*types.Transaction, error) {
	return _Gov.Contract.WithdrawOnes(&_Gov.TransactOpts, amt)
}

// GovAddVoteIterator is returned from FilterAddVote and is used to iterate over the raw logs and unpacked data for AddVote events raised by the Gov contract.
type GovAddVoteIterator struct {
	Event *GovAddVote // Event containing the contract specifics and raw log

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
func (it *GovAddVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovAddVote)
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
		it.Event = new(GovAddVote)
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
func (it *GovAddVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovAddVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovAddVote represents a AddVote event raised by the Gov contract.
type GovAddVote struct {
	ProposalID uint64
	Voter      common.Address
	Opinion    uint8
	VoteAmt    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddVote is a free log retrieval operation binding the contract event 0x79cf911a7c6d5d1b5d658aa8a512828c7539341faa788e21ee4b22997cee5ab0.
//
// Solidity: event AddVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) FilterAddVote(opts *bind.FilterOpts) (*GovAddVoteIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "AddVote")
	if err != nil {
		return nil, err
	}
	return &GovAddVoteIterator{contract: _Gov.contract, event: "AddVote", logs: logs, sub: sub}, nil
}

// WatchAddVote is a free log subscription operation binding the contract event 0x79cf911a7c6d5d1b5d658aa8a512828c7539341faa788e21ee4b22997cee5ab0.
//
// Solidity: event AddVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) WatchAddVote(opts *bind.WatchOpts, sink chan<- *GovAddVote) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "AddVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovAddVote)
				if err := _Gov.contract.UnpackLog(event, "AddVote", log); err != nil {
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

// ParseAddVote is a log parse operation binding the contract event 0x79cf911a7c6d5d1b5d658aa8a512828c7539341faa788e21ee4b22997cee5ab0.
//
// Solidity: event AddVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) ParseAddVote(log types.Log) (*GovAddVote, error) {
	event := new(GovAddVote)
	if err := _Gov.contract.UnpackLog(event, "AddVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovNewFundsProposalIterator is returned from FilterNewFundsProposal and is used to iterate over the raw logs and unpacked data for NewFundsProposal events raised by the Gov contract.
type GovNewFundsProposalIterator struct {
	Event *GovNewFundsProposal // Event containing the contract specifics and raw log

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
func (it *GovNewFundsProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovNewFundsProposal)
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
		it.Event = new(GovNewFundsProposal)
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
func (it *GovNewFundsProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovNewFundsProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovNewFundsProposal represents a NewFundsProposal event raised by the Gov contract.
type GovNewFundsProposal struct {
	ProposalID  uint64
	Title       string
	Desc        string
	Url         string
	Deadline    uint32
	Beneficiary common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewFundsProposal is a free log retrieval operation binding the contract event 0xcefa6a6335c5a8e5b325bf02dae880483f6a9e3632e21083f8994b95191bc00c.
//
// Solidity: event NewFundsProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address beneficiary, uint256 amount)
func (_Gov *GovFilterer) FilterNewFundsProposal(opts *bind.FilterOpts) (*GovNewFundsProposalIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "NewFundsProposal")
	if err != nil {
		return nil, err
	}
	return &GovNewFundsProposalIterator{contract: _Gov.contract, event: "NewFundsProposal", logs: logs, sub: sub}, nil
}

// WatchNewFundsProposal is a free log subscription operation binding the contract event 0xcefa6a6335c5a8e5b325bf02dae880483f6a9e3632e21083f8994b95191bc00c.
//
// Solidity: event NewFundsProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address beneficiary, uint256 amount)
func (_Gov *GovFilterer) WatchNewFundsProposal(opts *bind.WatchOpts, sink chan<- *GovNewFundsProposal) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "NewFundsProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovNewFundsProposal)
				if err := _Gov.contract.UnpackLog(event, "NewFundsProposal", log); err != nil {
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

// ParseNewFundsProposal is a log parse operation binding the contract event 0xcefa6a6335c5a8e5b325bf02dae880483f6a9e3632e21083f8994b95191bc00c.
//
// Solidity: event NewFundsProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address beneficiary, uint256 amount)
func (_Gov *GovFilterer) ParseNewFundsProposal(log types.Log) (*GovNewFundsProposal, error) {
	event := new(GovNewFundsProposal)
	if err := _Gov.contract.UnpackLog(event, "NewFundsProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovNewParamProposalIterator is returned from FilterNewParamProposal and is used to iterate over the raw logs and unpacked data for NewParamProposal events raised by the Gov contract.
type GovNewParamProposalIterator struct {
	Event *GovNewParamProposal // Event containing the contract specifics and raw log

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
func (it *GovNewParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovNewParamProposal)
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
		it.Event = new(GovNewParamProposal)
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
func (it *GovNewParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovNewParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovNewParamProposal represents a NewParamProposal event raised by the Gov contract.
type GovNewParamProposal struct {
	ProposalID uint64
	Title      string
	Desc       string
	Url        string
	Deadline   uint32
	Factory    common.Address
	FeeBPS     uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewParamProposal is a free log retrieval operation binding the contract event 0x4623cd1863f26cf17b0b0252ebc774ecf43e15c013a840478f00c241f76e5e19.
//
// Solidity: event NewParamProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, uint32 feeBPS)
func (_Gov *GovFilterer) FilterNewParamProposal(opts *bind.FilterOpts) (*GovNewParamProposalIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "NewParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovNewParamProposalIterator{contract: _Gov.contract, event: "NewParamProposal", logs: logs, sub: sub}, nil
}

// WatchNewParamProposal is a free log subscription operation binding the contract event 0x4623cd1863f26cf17b0b0252ebc774ecf43e15c013a840478f00c241f76e5e19.
//
// Solidity: event NewParamProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, uint32 feeBPS)
func (_Gov *GovFilterer) WatchNewParamProposal(opts *bind.WatchOpts, sink chan<- *GovNewParamProposal) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "NewParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovNewParamProposal)
				if err := _Gov.contract.UnpackLog(event, "NewParamProposal", log); err != nil {
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

// ParseNewParamProposal is a log parse operation binding the contract event 0x4623cd1863f26cf17b0b0252ebc774ecf43e15c013a840478f00c241f76e5e19.
//
// Solidity: event NewParamProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, uint32 feeBPS)
func (_Gov *GovFilterer) ParseNewParamProposal(log types.Log) (*GovNewParamProposal, error) {
	event := new(GovNewParamProposal)
	if err := _Gov.contract.UnpackLog(event, "NewParamProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovNewTextProposalIterator is returned from FilterNewTextProposal and is used to iterate over the raw logs and unpacked data for NewTextProposal events raised by the Gov contract.
type GovNewTextProposalIterator struct {
	Event *GovNewTextProposal // Event containing the contract specifics and raw log

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
func (it *GovNewTextProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovNewTextProposal)
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
		it.Event = new(GovNewTextProposal)
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
func (it *GovNewTextProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovNewTextProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovNewTextProposal represents a NewTextProposal event raised by the Gov contract.
type GovNewTextProposal struct {
	ProposalID uint64
	Title      string
	Desc       string
	Url        string
	Deadline   uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTextProposal is a free log retrieval operation binding the contract event 0xedb5790b505fab801089c42a4d9c59dedfab325afc38c2f880f43fab318942b0.
//
// Solidity: event NewTextProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline)
func (_Gov *GovFilterer) FilterNewTextProposal(opts *bind.FilterOpts) (*GovNewTextProposalIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "NewTextProposal")
	if err != nil {
		return nil, err
	}
	return &GovNewTextProposalIterator{contract: _Gov.contract, event: "NewTextProposal", logs: logs, sub: sub}, nil
}

// WatchNewTextProposal is a free log subscription operation binding the contract event 0xedb5790b505fab801089c42a4d9c59dedfab325afc38c2f880f43fab318942b0.
//
// Solidity: event NewTextProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline)
func (_Gov *GovFilterer) WatchNewTextProposal(opts *bind.WatchOpts, sink chan<- *GovNewTextProposal) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "NewTextProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovNewTextProposal)
				if err := _Gov.contract.UnpackLog(event, "NewTextProposal", log); err != nil {
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

// ParseNewTextProposal is a log parse operation binding the contract event 0xedb5790b505fab801089c42a4d9c59dedfab325afc38c2f880f43fab318942b0.
//
// Solidity: event NewTextProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline)
func (_Gov *GovFilterer) ParseNewTextProposal(log types.Log) (*GovNewTextProposal, error) {
	event := new(GovNewTextProposal)
	if err := _Gov.contract.UnpackLog(event, "NewTextProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovNewUpgradeProposalIterator is returned from FilterNewUpgradeProposal and is used to iterate over the raw logs and unpacked data for NewUpgradeProposal events raised by the Gov contract.
type GovNewUpgradeProposalIterator struct {
	Event *GovNewUpgradeProposal // Event containing the contract specifics and raw log

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
func (it *GovNewUpgradeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovNewUpgradeProposal)
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
		it.Event = new(GovNewUpgradeProposal)
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
func (it *GovNewUpgradeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovNewUpgradeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovNewUpgradeProposal represents a NewUpgradeProposal event raised by the Gov contract.
type GovNewUpgradeProposal struct {
	ProposalID uint64
	Title      string
	Desc       string
	Url        string
	Deadline   uint32
	Factory    common.Address
	PairLogic  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewUpgradeProposal is a free log retrieval operation binding the contract event 0x620d77127ee7d32615eb70c371bd279ee1c4f664cb785673bdec31a49898f03a.
//
// Solidity: event NewUpgradeProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, address pairLogic)
func (_Gov *GovFilterer) FilterNewUpgradeProposal(opts *bind.FilterOpts) (*GovNewUpgradeProposalIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "NewUpgradeProposal")
	if err != nil {
		return nil, err
	}
	return &GovNewUpgradeProposalIterator{contract: _Gov.contract, event: "NewUpgradeProposal", logs: logs, sub: sub}, nil
}

// WatchNewUpgradeProposal is a free log subscription operation binding the contract event 0x620d77127ee7d32615eb70c371bd279ee1c4f664cb785673bdec31a49898f03a.
//
// Solidity: event NewUpgradeProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, address pairLogic)
func (_Gov *GovFilterer) WatchNewUpgradeProposal(opts *bind.WatchOpts, sink chan<- *GovNewUpgradeProposal) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "NewUpgradeProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovNewUpgradeProposal)
				if err := _Gov.contract.UnpackLog(event, "NewUpgradeProposal", log); err != nil {
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

// ParseNewUpgradeProposal is a log parse operation binding the contract event 0x620d77127ee7d32615eb70c371bd279ee1c4f664cb785673bdec31a49898f03a.
//
// Solidity: event NewUpgradeProposal(uint64 proposalID, string title, string desc, string url, uint32 deadline, address factory, address pairLogic)
func (_Gov *GovFilterer) ParseNewUpgradeProposal(log types.Log) (*GovNewUpgradeProposal, error) {
	event := new(GovNewUpgradeProposal)
	if err := _Gov.contract.UnpackLog(event, "NewUpgradeProposal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovNewVoteIterator is returned from FilterNewVote and is used to iterate over the raw logs and unpacked data for NewVote events raised by the Gov contract.
type GovNewVoteIterator struct {
	Event *GovNewVote // Event containing the contract specifics and raw log

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
func (it *GovNewVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovNewVote)
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
		it.Event = new(GovNewVote)
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
func (it *GovNewVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovNewVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovNewVote represents a NewVote event raised by the Gov contract.
type GovNewVote struct {
	ProposalID uint64
	Voter      common.Address
	Opinion    uint8
	VoteAmt    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewVote is a free log retrieval operation binding the contract event 0xac44cb0e78b2dd1c33e8aa0f21a40ac00cc0b08f64cd62163ec42809a58e750d.
//
// Solidity: event NewVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) FilterNewVote(opts *bind.FilterOpts) (*GovNewVoteIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "NewVote")
	if err != nil {
		return nil, err
	}
	return &GovNewVoteIterator{contract: _Gov.contract, event: "NewVote", logs: logs, sub: sub}, nil
}

// WatchNewVote is a free log subscription operation binding the contract event 0xac44cb0e78b2dd1c33e8aa0f21a40ac00cc0b08f64cd62163ec42809a58e750d.
//
// Solidity: event NewVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) WatchNewVote(opts *bind.WatchOpts, sink chan<- *GovNewVote) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "NewVote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovNewVote)
				if err := _Gov.contract.UnpackLog(event, "NewVote", log); err != nil {
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

// ParseNewVote is a log parse operation binding the contract event 0xac44cb0e78b2dd1c33e8aa0f21a40ac00cc0b08f64cd62163ec42809a58e750d.
//
// Solidity: event NewVote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) ParseNewVote(log types.Log) (*GovNewVote, error) {
	event := new(GovNewVote)
	if err := _Gov.contract.UnpackLog(event, "NewVote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovRevoteIterator is returned from FilterRevote and is used to iterate over the raw logs and unpacked data for Revote events raised by the Gov contract.
type GovRevoteIterator struct {
	Event *GovRevote // Event containing the contract specifics and raw log

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
func (it *GovRevoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovRevote)
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
		it.Event = new(GovRevote)
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
func (it *GovRevoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovRevoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovRevote represents a Revote event raised by the Gov contract.
type GovRevote struct {
	ProposalID uint64
	Voter      common.Address
	Opinion    uint8
	VoteAmt    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevote is a free log retrieval operation binding the contract event 0x2054d4cdf4c6c9a38470bc7846d035a1dc962fd95a136c48fca4b36b01fb388b.
//
// Solidity: event Revote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) FilterRevote(opts *bind.FilterOpts) (*GovRevoteIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "Revote")
	if err != nil {
		return nil, err
	}
	return &GovRevoteIterator{contract: _Gov.contract, event: "Revote", logs: logs, sub: sub}, nil
}

// WatchRevote is a free log subscription operation binding the contract event 0x2054d4cdf4c6c9a38470bc7846d035a1dc962fd95a136c48fca4b36b01fb388b.
//
// Solidity: event Revote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) WatchRevote(opts *bind.WatchOpts, sink chan<- *GovRevote) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "Revote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovRevote)
				if err := _Gov.contract.UnpackLog(event, "Revote", log); err != nil {
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

// ParseRevote is a log parse operation binding the contract event 0x2054d4cdf4c6c9a38470bc7846d035a1dc962fd95a136c48fca4b36b01fb388b.
//
// Solidity: event Revote(uint64 proposalID, address voter, uint8 opinion, uint112 voteAmt)
func (_Gov *GovFilterer) ParseRevote(log types.Log) (*GovRevote, error) {
	event := new(GovRevote)
	if err := _Gov.contract.UnpackLog(event, "Revote", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovTallyResultIterator is returned from FilterTallyResult and is used to iterate over the raw logs and unpacked data for TallyResult events raised by the Gov contract.
type GovTallyResultIterator struct {
	Event *GovTallyResult // Event containing the contract specifics and raw log

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
func (it *GovTallyResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovTallyResult)
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
		it.Event = new(GovTallyResult)
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
func (it *GovTallyResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovTallyResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovTallyResult represents a TallyResult event raised by the Gov contract.
type GovTallyResult struct {
	ProposalID uint64
	Pass       bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTallyResult is a free log retrieval operation binding the contract event 0x709c5e70cd3dcf1f39400585b4da49bf555425a933cbafa65e5746f81bf3ca4c.
//
// Solidity: event TallyResult(uint64 proposalID, bool pass)
func (_Gov *GovFilterer) FilterTallyResult(opts *bind.FilterOpts) (*GovTallyResultIterator, error) {

	logs, sub, err := _Gov.contract.FilterLogs(opts, "TallyResult")
	if err != nil {
		return nil, err
	}
	return &GovTallyResultIterator{contract: _Gov.contract, event: "TallyResult", logs: logs, sub: sub}, nil
}

// WatchTallyResult is a free log subscription operation binding the contract event 0x709c5e70cd3dcf1f39400585b4da49bf555425a933cbafa65e5746f81bf3ca4c.
//
// Solidity: event TallyResult(uint64 proposalID, bool pass)
func (_Gov *GovFilterer) WatchTallyResult(opts *bind.WatchOpts, sink chan<- *GovTallyResult) (event.Subscription, error) {

	logs, sub, err := _Gov.contract.WatchLogs(opts, "TallyResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovTallyResult)
				if err := _Gov.contract.UnpackLog(event, "TallyResult", log); err != nil {
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

// ParseTallyResult is a log parse operation binding the contract event 0x709c5e70cd3dcf1f39400585b4da49bf555425a933cbafa65e5746f81bf3ca4c.
//
// Solidity: event TallyResult(uint64 proposalID, bool pass)
func (_Gov *GovFilterer) ParseTallyResult(log types.Log) (*GovTallyResult, error) {
	event := new(GovTallyResult)
	if err := _Gov.contract.UnpackLog(event, "TallyResult", log); err != nil {
		return nil, err
	}
	return event, nil
}
