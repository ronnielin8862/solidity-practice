package main

import (
	"fmt"
	"solidity2/pkg/btc/genPrivateKey"
)

func main(){
	key,_ := genPrivateKey.GenerateSimpleKey()
	fmt.Printf("%+v\n" , key)
}
