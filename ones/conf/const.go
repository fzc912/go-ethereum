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
	Network = Local

	Mainnet = "mainnet"
	Ropsten = "ropsten"
	Local   = "localhost"
)

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

// ------------------------------ account ------------------------------
var (
	// 0x0000000000000000000000000000000000000000
	onesTokenAddressMainnet   = common.HexToAddress("0x0B342C51d1592C41068d5D4b4DA4A68C0a04d5A4")
	govAddressMainnet         = common.HexToAddress("0xF0825577c259aA94728310875368F905aFc57C4F")
	pairLogicAddressMainnet   = common.HexToAddress("0xEeFa8Ca24dd1D573882277b917720953e999734D")
	factoryAddressMainnet     = common.HexToAddress("0x5eD3C9089Ed0355bc77CF439Dc2eD28c4054C8c4")
	routerAddressMainnet      = common.HexToAddress("0xEEE21cf8762a87817868039F119e57a7FeC65074")
	onesETHPairAddressMainnet = common.HexToAddress("")

	onesTokenAddressRopsten   = common.HexToAddress("")
	govAddressRopsten         = common.HexToAddress("")
	pairLogicAddressRopsten   = common.HexToAddress("")
	factoryAddressRopsten     = common.HexToAddress("")
	routerAddressRopsten      = common.HexToAddress("")
	onesETHPairAddressRopsten = common.HexToAddress("")

	onesTokenAddressLocalhost   = common.HexToAddress("0x8645C8Eb1FfccB7cb81BC8f5Ea9b28122A8f9D40")
	govAddressLocalhost         = common.HexToAddress("0x06B08aE196bddeA3743D3555496cD49A87C0C858")
	pairLogicAddressLocalhost   = common.HexToAddress("0x030AF041526aE6581B171F1F42201DeE624C0DE8")
	factoryAddressLocalhost     = common.HexToAddress("0xAAe9e544138C6F581E4276b876824b007359a86F")
	routerAddressLocalhost      = common.HexToAddress("0x36c3BD2F646169cceAD5090853E5A1F6763A838d")
	onesETHPairAddressLocalhost = common.HexToAddress("0x45146Adc74E121447E02E340D9674F51279839a3")

	GainerAddress = common.HexToAddress("0x9725B749909F5b6A9570bb15C5EAB5EaBFFf8BA7")
	GainerKey     = GetGainerKey()
)

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

// ------------------------------ wallet ------------------------------
const (
	WalletPassword = "Test1234"
	// 0xb3683757e5A22d444b7825f6890A3c29E1a3b33a
	OnesOwnerWalletPath = "./ones/wallet_onesOwner/UTC--2020-08-30T07-13-37.046817000Z--b3683757e5a22d444b7825f6890a3c29e1a3b33a"
	// 0x9725B749909F5b6A9570bb15C5EAB5EaBFFf8BA7
	GainerWalletPath = "./ones/wallet_gainer/UTC--2020-08-30T07-13-38.386459000Z--9725b749909f5b6a9570bb15c5eab5eabfff8ba7"

	// 0x24A30f6A5C81275B46E32416277d4Fc779d1Ec57
	TesterWalletPath = "./ones/wallet_tester/UTC--2020-08-30T07-13-39.535697000Z--24a30f6a5c81275b46e32416277d4fc779d1ec57"
)

type RouterParams struct {
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
