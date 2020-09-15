#!/usr/bin/env bash

rm -rf abi
mkdir abi
solc --abi uniswap-v2-core/contracts/interfaces/IERC20.sol uniswap-v2-core/contracts/interfaces/IUniswapV2Callee.sol uniswap-v2-core/contracts/interfaces/IUniswapV2ERC20.sol uniswap-v2-core/contracts/interfaces/IUniswapV2Factory.sol uniswap-v2-core/contracts/interfaces/IUniswapV2Pair.sol -o abi --optimize

cd abi

mkdir erc20
abigen  --abi=IERC20.abi   --pkg=erc20 --out=erc20/erc20.go
mkdir callee
abigen --abi=IUniswapV2Callee.abi --pkg=callee --out=callee/callee.go
mkdir uniswap_erc20
abigen --abi=IUniswapV2ERC20.abi --pkg=uniswap_erc20 --out=uniswap_erc20/uniswap_erc20.go
mkdir factory
abigen --abi=IUniswapV2Factory.abi --pkg=factory --out=factory/factory.go
mkdir pair
abigen --abi=IUniswapV2Pair.abi --pkg=pair --out=pair/pair.go