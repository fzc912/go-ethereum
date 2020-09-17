package conf

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/hpcloud/tail/util"
	"io/ioutil"
	"math/big"
)

// ------------------------------ network ------------------------------
const (
	Network = Mainnet

	Mainnet = "mainnet"
	Ropsten = "ropsten"
	Local   = "localhost"
)

// ------------------------------ host ------------------------------
const (
	infuraProjectID  = "8315fd240807477ba6a81671d1897f71"
	infuraAPIMainnet = "https://mainnet.infura.io/v3/" + infuraProjectID
	infuraWSMainnet  = "wss://mainnet.infura.io/ws/v3/" + infuraProjectID
	infuraAPIRopsten = "https://ropsten.infura.io/v3/" + infuraProjectID
	infuraWSRopsten  = "wss://ropsten.infura.io/ws/v3/" + infuraProjectID
	apiLocalhost     = "http://localhost:8545"
	wsLocalhost      = "ws://localhost:8546"
)

// ------------------------------ wallet ------------------------------
const (
	WalletPassword   = "Test1234"
	GainerWalletPath = "./wallet_gainer/UTC--2020-08-30T07-13-38.386459000Z--9725b749909f5b6a9570bb15c5eab5eabfff8ba7" // 0x9725B749909F5b6A9570bb15C5EAB5EaBFFf8BA7
)

// ------------------------------ account ------------------------------
var (
	// 0x0000000000000000000000000000000000000000
	onesTokenAddressMainnet   = common.HexToAddress("")
	govAddressMainnet         = common.HexToAddress("0xF0825577c259aA94728310875368F905aFc57C4F")
	pairLogicAddressMainnet   = common.HexToAddress("0x4e3E0852D664b671222384d654d789B2434a9446")
	factoryAddressMainnet     = common.HexToAddress("0x5eD3C9089Ed0355bc77CF439Dc2eD28c4054C8c4")
	routerAddressMainnet      = common.HexToAddress("0xEEE21cf8762a87817868039F119e57a7FeC65074")
	onesETHPairAddressMainnet = common.HexToAddress("")

	onesTokenAddressRopsten   = common.HexToAddress("")
	govAddressRopsten         = common.HexToAddress("")
	pairLogicAddressRopsten   = common.HexToAddress("")
	factoryAddressRopsten     = common.HexToAddress("")
	routerAddressRopsten      = common.HexToAddress("")
	onesETHPairAddressRopsten = common.HexToAddress("")

	onesTokenAddressLocalhost   = common.HexToAddress("")
	govAddressLocalhost         = common.HexToAddress("")
	pairLogicAddressLocalhost   = common.HexToAddress("")
	factoryAddressLocalhost     = common.HexToAddress("")
	routerAddressLocalhost      = common.HexToAddress("")
	onesETHPairAddressLocalhost = common.HexToAddress("")

	OneswapBigAddress = common.HexToAddress("0xf430caaed56e84351337d710d80d0f125055E8FA")
	GainerAddress     = common.HexToAddress("0x9725B749909F5b6A9570bb15C5EAB5EaBFFf8BA7")
	GainerKey         = GetGainerKey()
)

func GetConfig() (map[string]string, map[string]common.Address) {
	host := make(map[string]string)
	address := make(map[string]common.Address)
	switch Network {
	case Mainnet:
		host[APIHost] = infuraAPIMainnet
		host[WSHost] = infuraWSMainnet
		address[OnesTokenAddress] = onesTokenAddressMainnet
		address[GovAddress] = govAddressMainnet
		address[FactoryAddress] = factoryAddressMainnet
		address[RouterAddress] = routerAddressMainnet
		address[PairLogicAddress] = pairLogicAddressMainnet
		address[OnesETHPairAddress] = onesETHPairAddressMainnet
	case Ropsten:
		host[APIHost] = infuraAPIRopsten
		host[WSHost] = infuraWSRopsten
		address[OnesTokenAddress] = onesTokenAddressRopsten
		address[GovAddress] = govAddressRopsten
		address[FactoryAddress] = factoryAddressRopsten
		address[RouterAddress] = routerAddressRopsten
		address[PairLogicAddress] = pairLogicAddressRopsten
		address[OnesETHPairAddress] = onesETHPairAddressRopsten
	case Local:
		host[APIHost] = apiLocalhost
		host[WSHost] = wsLocalhost
		address[OnesTokenAddress] = onesTokenAddressLocalhost
		address[GovAddress] = govAddressLocalhost
		address[FactoryAddress] = factoryAddressLocalhost
		address[RouterAddress] = routerAddressLocalhost
		address[PairLogicAddress] = pairLogicAddressLocalhost
		address[OnesETHPairAddress] = onesETHPairAddressLocalhost
	default:
		panic("network error : " + Network)
	}
	return host, address
}

func GetGainerKey() *keystore.Key {
	// get gainer private key
	keyBytes, err := ioutil.ReadFile(GainerWalletPath)
	if err != nil {
		util.Fatal("WalletImportKs read ks file error : %v", err)
	}
	key, err := keystore.DecryptKey(keyBytes, WalletPassword)
	if err != nil {
		util.Fatal("GetWalletKey DecryptKey error : %v", err)
	}
	log.Info("GetGainerKey success")
	return key
}

type OneswapRouterAddLiquidityParams struct {
	Stock              common.Address
	Money              common.Address
	IsOnlySwap         bool
	AmountStockDesired *big.Int
	AmountMoneyDesired *big.Int
	AmountStockMin     *big.Int
	AmountMoneyMin     *big.Int
	To                 common.Address
	Deadline           *big.Int
}

const (
	APIHost = "api_host"
	WSHost  = "ws_host"

	OnesTokenAddress   = "ones_token_address"
	GovAddress         = "gov_address"
	FactoryAddress     = "factory_address"
	RouterAddress      = "router_address"
	PairLogicAddress   = "pair_logic_address"
	OnesETHPairAddress = "ones_eth_pair_address"
)
