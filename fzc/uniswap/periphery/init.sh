#!/usr/bin/env bash

rm -rf abi
mkdir abi
solc --abi uniswap-v2-periphery/contracts/interfaces/IERC20.sol uniswap-v2-periphery/contracts/interfaces/IUniswapV2Migrator.sol uniswap-v2-periphery/contracts/interfaces/IUniswapV2Router02.sol uniswap-v2-periphery/contracts/interfaces/IWETH.sol -o abi --optimize

cd abi

mkdir  erc20
abigen  --abi=IERC20.abi   --pkg=erc20 --out=erc20/erc20.go
mkdir migrator
abigen  --abi=IUniswapV2Migrator.abi   --pkg=migrator --out=migrator/migrator.go
mkdir router02
abigen  --abi=IUniswapV2Router02.abi   --pkg=router02 --out=router02/router02.go
mkdir weth
abigen  --abi=IWETH.abi   --pkg=weth --out=weth/weth.go