package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/fzc/oneswap/conf"
	"github.com/ethereum/go-ethereum/fzc/uniswap/core/abi/pair"
	"github.com/ethereum/go-ethereum/fzc/uniswap/periphery/abi/router02"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/net/context"
	"math/big"
	"strings"
	"time"
)

var (
	uniswapRouter2Address = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	decimalsBigInt18      = big.NewInt(1000000000000000000)
)
var (
	WETHTokenAddress = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")

	YFVTokenAddress   = common.HexToAddress("")
	YFVETHPairAddress = common.HexToAddress("")

	OnesTokenTestAddress   = common.HexToAddress("0x97144b7e76d012f4f6450e7ff29bc0f5f01f6a0f")
	OnesETHTestPairAddress = common.HexToAddress("0x14FA97ea8792C84c8696074464629724a3cAFdCc")
)

type UniswapRouterParams struct {
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

func (pm *ProtocolManager) uniswap(txs []*types.Transaction) {
	for _, tx := range txs {
		if tx == nil || tx.To() == nil {
			continue
		}
		if err := pm.preCheckUniswap(tx); err != nil {
			log.Warn("fzc preCheckUniswap error : " + err.Error())
		}
	}
}

func (pm *ProtocolManager) preCheckUniswap(tx *types.Transaction) error {
	if tx.To().Hex() == uniswapRouter2Address.Hex() {
		routerParams, ok := pm.DecodeInputData(tx.Data())
		if !ok || len(routerParams.Path) != 2 {
			return nil
		}
		// ------------------------------------- add pair -------------------------------------
		if routerParams.Path[0].Hex() == WETHTokenAddress.Hex() && routerParams.Path[1].Hex() == OnesTokenTestAddress.Hex() {
			if err := pm.swapOnesETHTest(tx, routerParams); err != nil {
				return err
			}
		}

		//if routerParams.Path[0].Hex() == WETHTokenAddress.Hex() && routerParams.Path[1].Hex() == YFVTokenAddress.Hex() {
		//	if err := pm.swapYFVETH(tx, routerParams); err != nil {
		//		return err
		//	}
		//}
	}
	return nil
}

func (pm *ProtocolManager) swapOnesETHTest(tx *types.Transaction, params UniswapRouterParams) error {
	buyPath := []common.Address{WETHTokenAddress, OnesTokenTestAddress}
	sellPath := []common.Address{OnesTokenTestAddress, WETHTokenAddress}
	pairAddress := OnesETHTestPairAddress
	valueLimit := big.NewInt(1e8)
	if err := pm.TransferETHSwap(tx, params, valueLimit, pairAddress, buyPath, sellPath); err != nil {
		return err
	}
	return nil
}

func (pm *ProtocolManager) swapYFVETH(tx *types.Transaction, params UniswapRouterParams) error {
	buyPath := []common.Address{WETHTokenAddress, YFVTokenAddress}
	sellPath := []common.Address{YFVTokenAddress, WETHTokenAddress}
	pairAddress := YFVETHPairAddress
	valueLimit := new(big.Int).Mul(big.NewInt(10), decimalsBigInt18)
	if err := pm.TransferETHSwap(tx, params, valueLimit, pairAddress, buyPath, sellPath); err != nil {
		return err
	}
	return nil
}

func (pm *ProtocolManager) TransferETHSwap(tx *types.Transaction, params UniswapRouterParams, valueLimit *big.Int,
	pairAddress common.Address, buyPath, sellPath []common.Address) error {

	// ------------------------------------- valueLimit -------------------------------------
	if tx.Value().Cmp(valueLimit) < 0 {
		return nil
	}
	// ------------------------------------- SimulateSwapETH -------------------------------------
	c := new(big.Int).Sub(tx.Value(), new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(tx.Gas()))))
	if c.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}
	d := params.AmountOutMin
	x, y, swap, err := pm.SimulateSwapETH(pairAddress, buyPath, c, d)
	if err != nil {
		return fmt.Errorf("SimulateSwapETH error %v", err)
	} else {
		if !swap {
			return fmt.Errorf("SimulateSwapETH profit too low")
		}
	}
	// ------------------------------------- SwapExactETHForTokens -------------------------------------
	err = pm.SwapExactETHForTokens(tx.GasPrice(), params.Deadline, tx.Gas(), buyPath, x, y)
	if err != nil {
		return fmt.Errorf("swap first tx error %v", err)
	}
	// ------------------------------------- SwapExactTokensForETH -------------------------------------
	for i := 0; i < 5; i++ {
		err := pm.SwapExactTokensForETH(tx.GasPrice(), params.Deadline, tx.Gas(), sellPath, y, x)
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			log.Warn("fzc swap third tx error error %v" + err.Error())
			continue
		} else {
			break
		}
	}
	return nil
}

func (pm *ProtocolManager) SwapExactETHForTokens(gasPrice, deadline *big.Int, gasLimit uint64, path []common.Address, x, y *big.Int) error {
	router, err := router02.NewRouter02(uniswapRouter2Address, pm.infuraAPI)
	if err != nil {
		return fmt.Errorf("SwapExactETHForTokens NewRouter error %v", err)
	}

	auth := bind.NewKeyedTransactor(conf.GainerKey.PrivateKey)
	auth.Nonce = nil
	auth.Value = new(big.Int).Add(x, new(big.Int).Mul(big.NewInt(int64(gasLimit)), new(big.Int).Add(gasPrice, big.NewInt(1))))
	auth.GasLimit = gasLimit
	auth.GasPrice = new(big.Int).Add(gasPrice, big.NewInt(1))
	//amountOutMin := new(big.Int).Div(new(big.Int).Mul(y, big.NewInt(95)), big.NewInt(100))
	tx, err := router.SwapExactETHForTokensSupportingFeeOnTransferTokens(auth, y, path, conf.GainerAddress, deadline)
	if err != nil {
		return fmt.Errorf("SwapExactETHForTokensSupportingFeeOnTransferTokens error %v", err)
	}
	log.Info("fzc SwapExactETHForTokens tx hash : " + tx.Hash().Hex())
	return nil
}

func (pm *ProtocolManager) SwapExactTokensForETH(gasPrice, deadline *big.Int, gasLimit uint64, path []common.Address, y, x *big.Int) error {
	router, err := router02.NewRouter02(uniswapRouter2Address, pm.infuraAPI)
	if err != nil {
		return fmt.Errorf("NewRouter error %v", err)
	}
	auth := bind.NewKeyedTransactor(conf.GainerKey.PrivateKey)
	auth.Nonce = nil
	auth.Value = nil
	auth.GasLimit = gasLimit
	auth.GasPrice = new(big.Int).Sub(gasPrice, big.NewInt(1))
	ethAmountOutMin := new(big.Int).Div(new(big.Int).Mul(x, big.NewInt(95)), big.NewInt(100))
	tx, err := router.SwapExactTokensForETHSupportingFeeOnTransferTokens(auth, y, ethAmountOutMin, path, conf.GainerAddress, deadline)
	if err != nil {
		return fmt.Errorf("SwapExactTokensForETHSupportingFeeOnTransferTokens error %v", err)
	}
	log.Info("fzc SwapExactTokensForETH tx hash : " + tx.Hash().Hex())
	return nil
}

func (pm *ProtocolManager) DecodeInputData(data []byte) (UniswapRouterParams, bool) {
	//log.Info("fzc input data : " + hex.EncodeToString(data))
	var routerParams UniswapRouterParams
	if len(data) < 4 {
		return routerParams, false
	}

	router02ABIJson := fmt.Sprintf(`[{"inputs":[],"name":"WETH","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"amountADesired","type":"uint256"},{"internalType":"uint256","name":"amountBDesired","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amountTokenDesired","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"addLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"},{"internalType":"uint256","name":"liquidity","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"factory","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountIn","outputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"reserveIn","type":"uint256"},{"internalType":"uint256","name":"reserveOut","type":"uint256"}],"name":"getAmountOut","outputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsIn","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"}],"name":"getAmountsOut","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"reserveA","type":"uint256"},{"internalType":"uint256","name":"reserveB","type":"uint256"}],"name":"quote","outputs":[{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"pure","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidity","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETH","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"removeLiquidityETHSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermit","outputs":[{"internalType":"uint256","name":"amountToken","type":"uint256"},{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountTokenMin","type":"uint256"},{"internalType":"uint256","name":"amountETHMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityETHWithPermitSupportingFeeOnTransferTokens","outputs":[{"internalType":"uint256","name":"amountETH","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"},{"internalType":"uint256","name":"liquidity","type":"uint256"},{"internalType":"uint256","name":"amountAMin","type":"uint256"},{"internalType":"uint256","name":"amountBMin","type":"uint256"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"internalType":"bool","name":"approveMax","type":"bool"},{"internalType":"uint8","name":"v","type":"uint8"},{"internalType":"bytes32","name":"r","type":"bytes32"},{"internalType":"bytes32","name":"s","type":"bytes32"}],"name":"removeLiquidityWithPermit","outputs":[{"internalType":"uint256","name":"amountA","type":"uint256"},{"internalType":"uint256","name":"amountB","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapETHForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactETHForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForETHSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"amountOutMin","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapExactTokensForTokensSupportingFeeOnTransferTokens","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactETH","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"amountOut","type":"uint256"},{"internalType":"uint256","name":"amountInMax","type":"uint256"},{"internalType":"address[]","name":"path","type":"address[]"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"deadline","type":"uint256"}],"name":"swapTokensForExactTokens","outputs":[{"internalType":"uint256[]","name":"amounts","type":"uint256[]"}],"stateMutability":"nonpayable","type":"function"}]`)
	routerABI02, err := abi.JSON(strings.NewReader(router02ABIJson))
	if err != nil {
		return routerParams, false
	}

	method, err := routerABI02.MethodById(data[0:4])
	if err != nil || (method.Name != "swapExactETHForTokens" && method.Name != "swapExactETHForTokensSupportingFeeOnTransferTokens") {
		return routerParams, false
	}

	err = method.Inputs.Unpack(&routerParams, data[4:])
	if err != nil {
		log.Warn("fzc " + method.Name + " unpack error")
		return routerParams, false
	}
	//log.Info("fzc " + method.Name + " unpack success")
	return routerParams, true
}

func (pm *ProtocolManager) SimulateSwapETH(pairAddress common.Address, path []common.Address, c, d *big.Int) (
	*big.Int, *big.Int, bool, error) {
	// ------------------------------------- get reserves  -------------------------------------
	var a, b, x *big.Int
	pairInstance, err := pair.NewPair(pairAddress, pm.infuraAPI)
	if err != nil {
		return nil, nil, false, fmt.Errorf("NewPair error %v", err)
	}
	reserves, err := pairInstance.GetReserves(nil)
	if err != nil {
		return nil, nil, false, fmt.Errorf("GetReserves error %v", err)
	}
	token0, err := pairInstance.Token0(nil)
	if err != nil {
		return nil, nil, false, fmt.Errorf("token0 error %v", err)
	}
	if token0.Hex() == WETHTokenAddress.Hex() {
		a = reserves.Reserve0
		b = reserves.Reserve1
	} else {
		a = reserves.Reserve1
		b = reserves.Reserve0
	}

	// ------------------------------------- get eth balance  -------------------------------------
	balance, err := pm.infuraAPI.BalanceAt(context.Background(), conf.GainerAddress, nil)
	if err != nil {
		return nil, nil, false, fmt.Errorf("BalanceAt error %v", err)
	}

	// ------------------------------------- get x eth amount in  -------------------------------------
	if d.Cmp(big.NewInt(0)) > 0 {
		tmp1 := new(big.Int).Sub(c, new(big.Int).Mul(big.NewInt(2), a))
		tmp2 := new(big.Int).Mul(new(big.Int).Mul(new(big.Int).Mul(a, b), c), big.NewInt(4))
		tmp3 := new(big.Int).Div(tmp2, d)
		tmp4 := new(big.Int).Add(new(big.Int).Mul(c, c), tmp3)
		tmp5 := new(big.Int).Sqrt(tmp4)
		tmp6 := new(big.Int).Add(tmp1, tmp5)
		x = new(big.Int).Div(tmp6, big.NewInt(2))
	} else if d.Cmp(big.NewInt(0)) == 0 {
		x = balance
	} else {
		return nil, nil, false, fmt.Errorf("d < 0 %v", d)
	}
	if balance.Cmp(x) < 1 {
		x = balance
	}

	// (a + x + c - x3) * (b - y - d + y) = a * b
	// x3 = a + x + c - [(a * b) / (b - y - d + y)]
	// ------------------------------------- get profit  -------------------------------------
	tmp8 := new(big.Int).Add(new(big.Int).Add(a, x), c)
	tmp9 := new(big.Int).Div(new(big.Int).Mul(a, b), new(big.Int).Sub(b, d))
	x3 := new(big.Int).Sub(tmp8, tmp9)
	profit := new(big.Int).Sub(x3, x)
	cost := new(big.Int).Div(new(big.Int).Mul(x, big.NewInt(6)), big.NewInt(1000))
	if profit.Cmp(cost) < -1 {
		return nil, nil, false, nil
	}

	// ------------------------------------- get y token amount out  -------------------------------------
	router, err := router02.NewRouter02(uniswapRouter2Address, pm.infuraAPI)
	if err != nil {
		return nil, nil, false, fmt.Errorf("NewRouter02 error %v", err)
	}
	amountsOut, err := router.GetAmountsOut(nil, x, path)
	if err != nil {
		return nil, nil, false, fmt.Errorf("GetAmountsOut error %v", err)
	}

	return x, amountsOut[len(amountsOut)-1], true, nil
}
