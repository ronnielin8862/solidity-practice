package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"solidity2/pkg/btc/getUTXO"
	"solidity2/pkg/btc/omni"
	"solidity2/pkg/btc/transfer"
)

type UTXOStruct struct {
	TxID          string  `json:"TxID"`
	Vout          int     `json:"Vout"`
	Address       string  `json:"Address"`
	Account       string  `json:"Account"`
	RedeemScript  string  `json:"RedeemScript"`
	Amount        float64 `json:"Amount"`
	Confirmations int     `json:"Confirmations"`
	Spendable     bool    `json:"Spendable"`
	ScriptPubKey string `json:"ScriptPubKey"`
}


func main()  {

	client := omni.ConnectOmniNode()

	addrFrom := "2N5adyFFHJ1SvY9ZKZ5fUrWy5YQvaVo5BLs"
	addrTo := "2NB2E14JRpifwJipiAyKnLbxuyPJrP7jWJG"
	var transferAmount float64 = 10
	var fee float64 = 1
	prvKey := "cUXXv3QBadYcUKnQDfv28LA5tHA96KUb6Din7z29LWYwwzWjcaZo"

	err := transfer.SendAddressToAddress(addrFrom,addrTo,transferAmount,fee, client.Client, prvKey)
	
	if err != nil {
		fmt.Println("some err...")
		panic(err)
	}
}


func NewTx() (*wire.MsgTx, error) {
	return wire.NewMsgTx(wire.TxVersion), nil
}


func CreateTx(privKey string, destination string, amount int64, client *rpcclient.Client) (string, error) {

	var utxoStruct UTXOStruct

	_, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return "", err
	}

	UTXO, _ := getUTXO.GetUnspentByAddress("2N5adyFFHJ1SvY9ZKZ5fUrWy5YQvaVo5BLs",client)
	fmt.Printf("%+v\n", UTXO)

	for _ ,v := range UTXO {
		utxoStruct.TxID = v.TxID
		utxoStruct.Amount = v.Amount
		utxoStruct.ScriptPubKey = v.ScriptPubKey
	}
		if int64(utxoStruct.Amount) < amount {
			return "", fmt.Errorf("the balance of the account is not sufficient")
		}

		destinationAddr, err := btcutil.DecodeAddress(destination, &chaincfg.TestNet3Params)
		if err != nil {
			return "", err
		}

		destinationAddrByte, err := txscript.PayToAddrScript(destinationAddr)
		if err != nil {
			return "", err
		}

		redeemTx, err := NewTx()
		if err != nil {
			return "", err
		}

		utxoHash, err := chainhash.NewHashFromStr(utxoStruct.TxID)
		if err != nil {
			return "", err
		}

		outPoint := wire.NewOutPoint(utxoHash, 1)

		txIn := wire.NewTxIn(outPoint, nil, nil)
		redeemTx.AddTxIn(txIn)

		redeemTxOut := wire.NewTxOut(amount, destinationAddrByte)
		redeemTx.AddTxOut(redeemTxOut)

		finalRawTx, err := SignTx(privKey, utxoStruct.ScriptPubKey, redeemTx)

		return finalRawTx, nil
}

func SignTx(privKey string, pkScript string, redeemTx *wire.MsgTx) (string, error) {

	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return "", err
	}

	sourcePKScript, err := hex.DecodeString(pkScript)
	if err != nil {
		return "", nil
	}

	signature, err := txscript.SignatureScript(redeemTx, 0, sourcePKScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return "", nil
	}

	redeemTx.TxIn[0].SignatureScript = signature

	var signedTx bytes.Buffer
	redeemTx.Serialize(&signedTx)

	hexSignedTx := hex.EncodeToString(signedTx.Bytes())

	return hexSignedTx, nil
}