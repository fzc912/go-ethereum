#!/usr/bin/env bash
set -e

rm -rf abi
rm -r contracts
cd ~/GolandProjects/oneswap_contract_ethereum
git pull
cd -
cp -r ~/GolandProjects/oneswap_contract_ethereum/contracts .

mkdir abi
solc --abi contracts/interfaces/IOneSwapToken.sol  contracts/interfaces/IOneSwapFactory.sol contracts/interfaces/IOneSwapRouter.sol contracts/interfaces/IOneSwapPair.sol -o abi --optimize

cd abi

mkdir token
abigen  --abi=IOneswapToken.abi  --pkg=token --out=token/OneswapToken.go
mkdir factory
abigen  --abi=IOneswapFactory.abi --pkg=factory --out=factory/OneswapFactory.go
mkdir router
abigen  --abi=IOneswapRouter.abi  --pkg=router --out=router/OneswapRouter.go
mkdir pool
abigen  --abi=IOneswapPool.abi --pkg=pool --out=pool/OneswapPool.go