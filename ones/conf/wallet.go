package conf

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
)

type WalletAccount struct {
	Account accounts.Account
	Key     *keystore.Key
}

func InitWallet() {
	_ = WalletCreateKs("./wallet_onesOwner", WalletPassword)
	_ = WalletCreateKs("./wallet_gainer", WalletPassword)
	_ = WalletCreateKs("./wallet_tester", WalletPassword)
}

func WalletCreateKs(backupPath, password string) accounts.Account {
	// create new keystore
	ks := keystore.NewKeyStore(backupPath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Panicf("WalletCreateKs NewAccount : %v", err)
	}
	log.Printf("WalletCreateKs create ks success : %v", account.Address.Hex())
	return account
}

func WalletImportKs(filePath, password, tmpPath string) WalletAccount {
	// get keystore, importer account
	ks := keystore.NewKeyStore(tmpPath, keystore.StandardScryptN, keystore.StandardScryptP)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("WalletImportKs ReadFile : %v", err)
	}
	account, err := ks.Import(bytes, password, password)
	if err != nil {
		log.Panicf("WalletImportKs Import : %v", err)
	}
	log.Printf("WalletImportKs import keystore success, account : %v" + account.Address.Hex())

	// get private key, import key
	keyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("WalletImportKs read ks file error : %v", err)
	}
	key, err := keystore.DecryptKey(keyBytes, password)
	if err != nil {
		log.Panicf("GetWalletKey DecryptKey error : %v", err)
	}
	log.Printf("WalletImportKs importer private key success, account : %v", account.Address.Hex())

	return WalletAccount{
		Account: account,
		Key:     key,
	}
}