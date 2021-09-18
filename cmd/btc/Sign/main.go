package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcutil"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
	"strings"
)

type SignData struct {
	PrivateKey string `json:"private_key" binding:"required"`
	RawData    string `json:"raw_data" binding:"required"`
}

func main(){
	//可以用genPrivateKey產生出來的值貼過來
	signData := SignData{"L1X8qm7VeutRezhYRzZvSxz6ZVKHZEL5Zg98TKucizGnz674TfFG","RawData"}
	BtcSign(signData)
}

func BtcSign(signData SignData) {

	decodeCheck, _ := common.DecodeCheck(signData.PrivateKey)
	priKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), decodeCheck)
	rawDataBytes, _ := hex.DecodeString(signData.RawData)
	doubleHash := chainhash.DoubleHashB(rawDataBytes)
	sign, err := SignBtc(doubleHash, priKey)
	if err != nil {
		fmt.Println("has something error")
	}

	pub, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		fmt.Println("has something error")
	}
	fmt.Println("公鑰：", hex.EncodeToString(pubKey.SerializeCompressed()))
	fmt.Println("地址：", pub.EncodeAddress())
	signOk, err := VerifySignBtc(doubleHash, sign, pubKey.ToECDSA())
	if err != nil {
		fmt.Println("has something error")
	}
	if signOk {
		fmt.Println("Done!")
	} else {
		fmt.Println("has something error")
	}
}


func SignBtc(text []byte, prv *btcec.PrivateKey) (string, error) {
	signature, err := prv.Sign(text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature.Serialize()), nil
}

func VerifySignBtc(text []byte, signature string, pubKey *ecdsa.PublicKey) (bool, error) {
	if strings.HasPrefix(signature, "0x") {
		signature = signature[2:]
	}
	signBytes, err := hex.DecodeString(signature)
	if err != nil {
		fmt.Println("has something error")
	}
	parseSignature, err := btcec.ParseSignature(signBytes, btcec.S256())
	if err != nil {
		return false, err
	}
	verify := parseSignature.Verify(text, (*btcec.PublicKey)(pubKey))
	return verify, nil
}