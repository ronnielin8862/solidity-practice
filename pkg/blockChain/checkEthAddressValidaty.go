package blockChain

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckEthAddress(addr string) {
	regexp := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("address is valid: %v \n", regexp.MatchString(addr))
}

func CheckEthContract(addr string) {
	contractAddr := common.HexToAddress(addr)
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/b904b0ebc4cb4a6692b47ef9147cafd8")
	if err != nil {
		fmt.Printf("Dial err %v ", err)
	}

	bytecode, err := client.CodeAt(context.Background(), contractAddr, nil)
	if err != nil {
		fmt.Printf("CodeAt err %v ", err)
	}

	fmt.Printf("is contract: %v\n", len(bytecode) > 0)

}
