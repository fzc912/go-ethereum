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

func (pm *ProtocolManager) oneswap(txs []*types.Transaction) {
	for _, tx := range txs {
		if tx == nil || tx.To() == nil || ExistTxHash[tx.Hash().Hex()] {
			continue
		}
		if err := pm.preCheckOneswap(tx); err != nil {
			log.Warn("fzc preCheckOneswap error " + err.Error())
		}
	}
}

func (pm *ProtocolManager) preCheckOneswap(tx *types.Transaction) error {
	onesTokenAddr := common.HexToAddress("0x0B342C51d1592C41068d5D4b4DA4A68C0a04d5A4")
	if tx.To().Hex() == pm.address[conf.RouterAddress].Hex() {
		routerParams, ok := pm.DecodeOneswapInputData(tx.Data())
		if !ok ||
			routerParams.Stock.Hex() != onesTokenAddr.Hex() ||
			routerParams.Money.Hex() != common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
			return nil
		}

		expectMinStockAll := new(big.Int).Mul(big.NewInt(700000), big.NewInt(1e18)) // 70w ones
		ethValueAll := new(big.Int).Mul(big.NewInt(210), big.NewInt(1e18))          // 210 eth
		amountOutMinAll := new(big.Int).Mul(big.NewInt(100000), big.NewInt(1e18))   // 10w ones

		expectMaxStockHalf := new(big.Int).Mul(big.NewInt(60000), big.NewInt(1e18))  // 60w ones
		expectMinStockHalf := new(big.Int).Mul(big.NewInt(300000), big.NewInt(1e18)) // 30w ones
		ethValueHalf := new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))          // 100 eth
		amountOutMinHalf := new(big.Int).Mul(big.NewInt(50000), big.NewInt(1e18))    // 5w ones

		var pairAddr common.Address
		if routerParams.IsOnlySwap {
			pairAddr = common.HexToAddress("0x1025F2334E88a4a0E93E5E2Ff4a4E67317682B60")
		} else {
			pairAddr = common.HexToAddress("0x7E8dC7C0feF7bB84304CC5c2DACF5E583198CA4f")
		}

		// what
		if routerParams.AmountStockDesired.Cmp(expectMinStockHalf) < 0 {
			return fmt.Errorf("routerParmas amount stocke desired too low : %+v", routerParams)
		}

		// all stock
		if routerParams.AmountStockDesired.Cmp(expectMinStockAll) > 0 {
			log.Info("routerParams stock all amount ok! : " + routerParams.AmountStockDesired.String())
			for i := 0; i < 60; i++ {
				if err := pm.SwapSpesETHToOnes(conf.GainerKey, pm.address[conf.RouterAddress], pairAddr, conf.GainerAddress, ethValueAll, amountOutMinAll, tx.GasPrice()); err != nil {
					log.Warn("fzc SwapSpesETHToOnes error " + err.Error())
					time.Sleep(350 * time.Millisecond)
					continue
				}
				ExistTxHash[tx.Hash().Hex()] = true
				break
			}
		}

		// half stock
		if routerParams.AmountStockDesired.Cmp(expectMinStockHalf) > 0 && routerParams.AmountStockDesired.Cmp(expectMaxStockHalf) < 0 {
			log.Info("routerParams stock half amount ok! : " + routerParams.AmountStockDesired.String())
			for i := 0; i < 60; i++ {
				if err := pm.SwapSpesETHToOnes(conf.GainerKey, pm.address[conf.RouterAddress], pairAddr, conf.GainerAddress, ethValueHalf, amountOutMinHalf, tx.GasPrice()); err != nil {
					log.Warn("fzc SwapSpesETHToOnes error " + err.Error())
					time.Sleep(350 * time.Millisecond)
					continue
				}
				ExistTxHash[tx.Hash().Hex()] = true
				break
			}
		}
	}
	return nil
}

func (pm *ProtocolManager) SwapSpesETHToOnes(key *keystore.Key, routerAddr, pairAddr, toAddr common.Address, ethValue, amountOutMin, gasPrice *big.Int) error {
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
		return fmt.Errorf("SwapSpesETHToOnes error %v", err)
	}
	log.Info("fzc SwapSpesETHToOnes tx hash : " + tx.Hash().Hex())
	return nil
}

func (pm *ProtocolManager) DecodeOneswapInputData(data []byte) (conf.OneswapRouterAddLiquidityParams, bool) {
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
