package main

import (
	"log"
	"solidity2/pkg/btc/omni"
)

func main() {

	client := omni.ConnectOmniNode()

	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}

