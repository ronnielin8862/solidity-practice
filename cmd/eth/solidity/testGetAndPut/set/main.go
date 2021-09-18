package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xe0Cf5653598BA9B2eF83299f7F62abbE8E2A5E3a")
	instance, err := testGetAndPut.NewTestGetAndPut(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Store(auth, big.NewInt(333))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx :", tx)
	fmt.Println("tx sent: ", tx.Hash().Hex())

	result, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result :",result )

}