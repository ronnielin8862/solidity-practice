package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main(){

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
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("from addr = " , fromAddress)


	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	to := accounts.Account{Address: common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (1 eth)

	gasLimit := uint64(210000) // in units
	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	data:= []byte(" ~ 好棒棒耶 ~")

	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to.Address,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})


	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent

}
