package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/fzc/oneswap/conf"
	"github.com/ethereum/go-ethereum/fzc/oneswap/ones/oneswap/router"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"strings"
	"time"
)

func (pm *ProtocolManager) ones(txs []*types.Transaction) {
	for _, tx := range txs {
		if err := pm.readySwap(tx); err != nil {
			log.Warn("spesSwapETHToToken error ", err)
		}
	}
}

func (pm *ProtocolManager) readySwap(tx *types.Transaction) error {
	if tx.To().Hex() == pm.address[conf.RouterAddress].Hex() {
		def := fmt.Sprintf(`[{"inputs":[{"internalType":"address","name":"_factory","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"stockAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"moneyAmount","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"liquidity","type":"uint256"}],"name":"AddLiquidity","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"address","name":"stock","type":"address"},{"indexed":false,"internalType":"address","name":"money","type":"address"},{"indexed":false,"internalType":"bool","name":"isOnlySwap","type":"bool"}],"name":"PairCreated","type":"event"},{"inputs":[{"internalType":"address","name":"stock","type":"address"},{"internalType":"address","name":"money","type":"address"},{"internalType":"bool","name":"isOnlySwap","type":"bool"},{"internalType":"uint256","name":"amountStockDesired","type":"uint256"},{"internalType":"uint256","name":"amountMoneyDesired","type":"uint256"},{"internalType":"uint256","name":"amountStockMin","type":"uint256"},{"internalType":"uint256","name":"amountMoneyMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountStock","type":"uint256"},{"internalType":"uint256","name":"amountMoney","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bool","name":"isBuy","type":"bool"},{"internalType":"address","name":"pair","type":"address"},{"internalType":"uint256","name":"prevKey","type":"uint256"},{"internalType":"uint256","name":"price","type":"uint256"},{"internalType":"uint32","name":"id","type":"uint32"},{"internalType":"uint256","name":"stockAmount","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"limitOrder","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"pair","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountStockMin","type":"uint256"},{"internalType":"uint256","name":"amountMoneyMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountStock","type":"uint256"},{"internalType":"uint256","name":"amountMoney","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapToken","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"}]`)
		routerABI, err := abi.JSON(strings.NewReader(def))
		if err != nil {
			return fmt.Errorf("panic abi json router abi error %v", err)
		}
		var routerParams conf.RouterParams
		err = routerABI.Unpack(&routerParams, "addLiquidity", tx.Data())
		if err != nil {
			return fmt.Errorf("warning abi unpack addLiquidity method error %v", err)
		}

		if routerParams.Stock.Hex() == pm.address[conf.OnesTokenAddress].Hex() &&
			routerParams.Money.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
			decimals18 := "000000000000000000"
			expectMinLiquidity := new(big.Int)
			expectMinLiquidity.SetString("800000"+decimals18, 10) // 80w ones

			ethValue := new(big.Int)
			ethValue.SetString("200"+decimals18, 10) // 200 eth
			amountOutMin := new(big.Int)
			amountOutMin.SetString("100000"+decimals18, 10) // 10w ones

			gasPrice := new(big.Int).Sub(tx.GasPrice(), big.NewInt(1))
			log.Info("get gas price ", tx.GasPrice())
			log.Info("set gas price ", gasPrice)
			if routerParams.AmountStockDesired.Cmp(expectMinLiquidity) >= 0 {
				for i := 0; i < 60; i++ {
					err = pm.SpesSwapETHToOnes(conf.GainerKey, gasPrice,
						pm.address[conf.OnesTokenAddress], pm.address[conf.RouterAddress], pm.address[conf.OnesETHPairAddress], ethValue, amountOutMin)
					time.Sleep(350 * time.Millisecond)
					if err != nil {
						log.Warn("SpesSwapETHToOnes error ", err)
					}
				}
			}
		}
	}
	return nil
}

func (pm *ProtocolManager) SpesSwapETHToOnes(key *keystore.Key, gasPrice *big.Int, tokenAddr, routerAddr, pairAddr common.Address, ethValue, amountOutMin *big.Int) error {
	instance, err := router.NewRouter(routerAddr, pm.infuraAPI)
	if err != nil {
		return fmt.Errorf("NewRouter error %v", err)
	}

	auth := bind.NewKeyedTransactor(key.PrivateKey)
	// todo cost nonce
	auth.Nonce = nil
	auth.Value = ethValue
	auth.GasLimit = uint64(200000) // 20w gas
	auth.GasPrice = gasPrice

	deadline := big.NewInt(time.Now().Add(10 * time.Minute).Unix())
	tx, err := instance.SwapToken(auth, common.Address{}, ethValue, amountOutMin, []common.Address{pairAddr}, conf.GainerAddress, deadline)
	if err != nil {
		return fmt.Errorf("SwapToken error %v", err)
	}
	log.Info("SpesSwapETHToOnes tx hash : ", tx.Hash().Hex())
	return nil
}
