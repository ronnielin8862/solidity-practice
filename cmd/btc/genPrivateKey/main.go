package main

import (
	"solidity2/pkg/btc/genPrivateKey"
)

// Deprecated: 採用 genPrivateKey2
func main() {
	wallet := genPrivateKey.NewWallet()

	wallet.GetAddress()
}