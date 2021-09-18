package main

import (
	"solidity2/pkg/blockChain"
)

func main() {

	addr := "0x2910543af39aba0cd09dbb2d50200b3e800a63d2"
	blockChain.CheckEthAddress(addr)

	// contractAddr := "0xe41d2489571d322189246dafa5ebde1f4699f498" //true
	contractAddr := "0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4" //false
	blockChain.CheckEthContract(contractAddr)
}
