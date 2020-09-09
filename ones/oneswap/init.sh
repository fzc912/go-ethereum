#!/usr/bin/env bash
set -e

rm -rf token factory router pair
rm -rf abi bin
rm -r contracts
cd ~/GolandProjects/oneswap_contract_ethereum
git pull
cd -
cp -r ~/GolandProjects/oneswap_contract_ethereum/contracts .

mkdir abi bin
solc --abi contracts/OneSwapToken.sol contracts/OneSwapFactory.sol contracts/OneSwapRouter.sol contracts/OneSwapPair.sol -o abi --optimize
solc --bin contracts/OneSwapToken.sol contracts/OneSwapFactory.sol contracts/OneSwapRouter.sol contracts/OneSwapPair.sol -o bin --optimize

mkdir token
abigen  --abi=abi/OneswapToken.abi  --bin=bin/OneSwapToken.bin --pkg=token --out=token/OneswapToken.go
mkdir factory
abigen  --abi=abi/OneswapFactory.abi  --bin=bin/OneSwapFactory.bin --pkg=factory --out=factory/OneswapFactory.go
mkdir router
abigen  --abi=abi/OneswapRouter.abi  --bin=bin/OneSwapRouter.bin --pkg=router --out=router/OneswapRouter.go
mkdir pair
abigen  --abi=abi/OneswapPair.abi  --bin=bin/OneSwapPair.bin --pkg=pair --out=pair/OneswapPair.go