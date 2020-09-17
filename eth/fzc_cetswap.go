package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/fzc/conf"
	"github.com/ethereum/go-ethereum/fzc/oneswap/ones/abi/router"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
	"time"
)

func (pm *ProtocolManager) cetswap(txs []*types.Transaction) {
	for _, tx := range txs {
		if tx == nil || tx.To() == nil {
			continue
		}
		if err := pm.preCheckCetswap(tx); err != nil {
			log.Warn("fzc preCheckCetswap error ", err)
		}
	}
}

func (pm *ProtocolManager) preCheckCetswap(tx *types.Transaction) error {
	if tx.To().Hex() == pm.address[conf.RouterAddress].Hex() {
		routerParams, ok := pm.DecodeCetswapInputData(tx.Data())
		if !ok || routerParams.Money.Hex() != common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() ||
			routerParams.Stock.Hex() != pm.address[conf.OnesTokenAddress].Hex() || !routerParams.IsOnlySwap{
			return nil
		}

		expectMinStock := new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18)) // 1w cet
		ethValue := new(big.Int).Mul(big.NewInt(5), big.NewInt(1e17))          // 210 eth
		amountOutMin := new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18))   // 1w cet
		if routerParams.AmountStockDesired.Cmp(expectMinStock) < 0 {
			return fmt.Errorf("routerParmas amount stocke desired too low : %+v", routerParams)
		}
		for i := 0; i < 60; i++ {
			if err := pm.SwapSpesETHToCET(conf.GainerKey, pm.address[conf.RouterAddress], pm.address[conf.OnesETHPairAddress], conf.GainerAddress, ethValue, amountOutMin, tx.GasPrice()); err != nil {
				log.Warn("fzc SwapSpesETHToCET error ", err)
				time.Sleep(350 * time.Millisecond)
				continue
			}
			break
		}
	}
	return nil
}

func (pm *ProtocolManager) DecodeCetswapInputData(data []byte) (conf.OneswapRouterAddLiquidityParams, bool) {
	//log.Info("fzc input data : " + hex.EncodeToString(data))
	var routerParams conf.OneswapRouterAddLiquidityParams
	if len(data) < 4 {
		return routerParams, false
	}

	routerABIJson := fmt.Sprintf(`[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"stockAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"moneyAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"liquidity","type":"uint256"}],"name":"AddLiquidity","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"address","name":"stock","type":"address"},{"indexed":false,"internalType":"address","name":"money","type":"address"},{"indexed":false,"internalType":"bool","name":"isOnlySwap","type":"bool"}],"name":"PairCreated","type":"event"},{"inputs":[{"internalType":"address","name":"stock","type":"address"},{"internalType":"address","name":"money","type":"address"},{"internalType":"bool","name":"isOnlySwap","type":"bool"},{"internalType":"uint256","name":"amountStockDesired","type":"uint256"},{"internalType":"uint256","name":"amountMoneyDesired","type":"uint256"},{"internalType":"uint256","name":"amountStockMin","type":"uint256"},{"internalType":"uint256","name":"amountMoneyMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountStock","type":"uint256"},{"internalType":"uint256","name":"amountMoney","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"bool","name":"isBuy","type":"bool"},{"internalType":"address","name":"pair","type":"address"},{"internalType":"uint256","name":"prevKey","type":"uint256"},{"internalType":"uint256","name":"price","type":"uint256"},{"internalType":"uint32","name":"id","type":"uint32"},{"internalType":"uint256","name":"stockAmount","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"limitOrder","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"pair","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountStockMin","type":"uint256"},{"internalType":"uint256","name":"amountMoneyMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountStock","type":"uint256"},{"internalType":"uint256","name":"amountMoney","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapToken","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"}]`)
	routerABI, err := abi.JSON(strings.NewReader(routerABIJson))
	if err != nil {
		log.Warn("fzc abi.JSON routerABIJson error " + err.Error())
		return routerParams, false
	}

	method, err := routerABI.MethodById(data[0:4])
	if err != nil || (method.Name != "addLiquidity") {
		return routerParams, false
	}

	err = method.Inputs.Unpack(&routerParams, data[4:])
	if err != nil {
		log.Warn("fzc unpack " + method.Name + " error " + err.Error())
		return routerParams, false
	}
	//log.Info("fzc " + method.Name + " unpack success")
	return routerParams, true
}

func (pm *ProtocolManager) SwapSpesETHToCET(key *keystore.Key, routerAddr, pairAddr, toAddr common.Address, ethValue, amountOutMin, gasPrice *big.Int) error {
	instance, err := router.NewRouter(routerAddr, pm.infuraAPI)
	if err != nil {
		return fmt.Errorf("NewRouter error %v", err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	auth.Nonce = nil
	auth.Value = ethValue
	auth.GasLimit = uint64(400000) // 40w
	auth.GasPrice = new(big.Int).Sub(gasPrice, big.NewInt(1))
	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())

	tx, err := instance.SwapToken(auth, common.Address{}, ethValue, amountOutMin, []common.Address{pairAddr}, toAddr, deadline)
	if err != nil {
		return fmt.Errorf("SwapSpesETHToCET error %v", err)
	}
	log.Info("fzc SwapSpesETHToCET tx hash : %v", tx.Hash().Hex())
	return nil
}
