package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"

	"solidity2/solidityRawData/testGetAndPut"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/c1f49e60480e47379556cc03ced9bddf")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("7b4481bec252c72bbf3a21c312b2dfef8fd983bf0eb579d39efd84040a7315ce")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//auth := bind.NewKeyedTransactor(privateKey)  改用下面新的api
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(3))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//input := "1.0"
	//address, tx, instance, err := testGetAndPut.DeployTestGetAndPut(auth, client, input)
	address, tx, instance, err := testGetAndPut.DeployTestGetAndPut(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("address = ", address.Hex())
	fmt.Println("tx = ", tx.Hash().Hex())

	_ = instance
}