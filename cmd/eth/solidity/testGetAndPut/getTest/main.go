package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

	"solidity2/solidityRawData/testGetAndPut"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/c1f49e60480e47379556cc03ced9bddf")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xe0Cf5653598BA9B2eF83299f7F62abbE8E2A5E3a")
	instance, err := testGetAndPut.NewTestGetAndPut(address, client)
	if err != nil {
		log.Fatal(err)
	}

	number, err := instance.Retrieve(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(number) // "1.0"

}