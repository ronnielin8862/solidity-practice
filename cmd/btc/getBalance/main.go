package main

import (
	"fmt"
	"solidity2/pkg/btc/getUTXO"
	"solidity2/pkg/btc/omni"
)

type UTXOStruct struct {
	TxID          string  `json:"TxID"`
	Vout          uint32    `json:"Vout"`
	Address       string  `json:"Address"`
	Account       string  `json:"Account"`
	RedeemScript  string  `json:"RedeemScript"`
	Amount        float64 `json:"Amount"`
	Confirmations int64     `json:"Confirmations"`
	Spendable     bool    `json:"Spendable"`
	ScriptPubKey string `json:"ScriptPubKey"`
}

func main() {
	var utxo UTXOStruct

	// create new client instance
	client := omni.ConnectOmniNode()

	UTXO, _ := getUTXO.GetUnspentByAddress("2N5adyFFHJ1SvY9ZKZ5fUrWy5YQvaVo5BLs",client.Client)

	for _, v := range UTXO {
		utxo.Address = v.Address
		utxo.Account = v.Account
		utxo.Amount = v.Amount
		utxo.Confirmations = v.Confirmations
		utxo.RedeemScript = v.RedeemScript
		utxo.Spendable = v.Spendable
		utxo.Vout = v.Vout
	}

	fmt.Println("UTXO 0-1 = " , utxo.Amount)

	fmt.Printf("%+v\n", UTXO)
}