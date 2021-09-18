package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"solidity2/pkg/blockChain"
)

func main() {
	// TODO: 用本地節點永遠是０餘額，但是ＣＭＤ卻可以正常查詢到。  原因待確認
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/c1f49e60480e47379556cc03ced9bddf")
	// client, err := ethclient.Dial("https://localhost:30303")
	if err != nil {
		fmt.Println(err)
	}

	account := common.HexToAddress("0x459CB0A51D053716D7ddBd6a1c04A7CaCc7a4909")

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("從wei轉成eth單位", blockChain.WeiToEther(balance))
}
