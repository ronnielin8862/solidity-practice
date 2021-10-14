package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/protobuf/proto"

	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/shopspring/decimal"
	//"github.com/fbsobreira/gotron-sdk/pkg/common"
	//"google.golang.org/grpc"
)

var (
	myAcc = "TUPnueXTxPWto2bRTevxmMfnupNfzri7M7"
	samAcc= "TFU3d1TKMvmXcAHyRGfVToyfNtAfdUmH5g"
	tronClient *client.GrpcClient
)


func main(){
	transfer()
}

func transfer(){

	privateKey := "86540d31f37a88994e7fde229ba82657dd6369ff4f098d9129819afc94051ad0"

	amount , _ := decimal.NewFromString("50")

	feeLimit := decimal.New(40, tronCommon.AmountDecimalPoint).IntPart()

	tronTx, err := CreateTronTx(myAcc, samAcc, "", amount, feeLimit)

	if err != nil {
		panic("创建转账交易")
	}

	rawData, err := proto.Marshal(tronTx.GetTransaction().GetRawData())

	RawDataHex := hex.EncodeToString(rawData)

	signTronTx(RawDataHex, privateKey)

}

func signTronTx(rawDataHex string, privateKeyHex string ){
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(err)
	}
	rawDataBytes, err := hex.DecodeString(rawDataHex)
	if err != nil {
		panic(err)
	}

	h256h := sha256.New()
	h256h.Write(rawDataBytes)
	hash := h256h.Sum(nil)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		panic(err)
	}

	sigStr := hex.EncodeToString(signature)
	panic("签名字符串"+ sigStr)
	transaction := &core.TransactionRaw{}
	if err = proto.Unmarshal(rawDataBytes, transaction); err != nil {
		panic(err)
		return
	}
	tx := &core.Transaction{RawData: transaction, Signature: [][]byte{signature}}
	marshal, err := proto.Marshal(tx)
	if err != nil {
		panic(err)
	}
	panic("签名后的交易"+ hex.EncodeToString(marshal))
}

// CreateTronTx 创建交易
func CreateTronTx(from, to, contract string, amount decimal.Decimal, feeLimit int64) (*api.TransactionExtention, error) {
	if len(contract) == 0 {
		// trx乘以10^6得到sun的数量
		transferNum := amount.Mul(decimal.New(1, tronCommon.AmountDecimalPoint)).IntPart()
		transactionExt, err := tronClient.Transfer(from, to, transferNum)
		if err != nil {
			fmt.Println("创建转账交易")
			return nil, err
		}
		return transactionExt, nil
	}
	decimals, err := tronClient.TRC20GetDecimals(contract)
	if err != nil {
		fmt.Println("查询%s小数位", contract)
		return nil, err
	}
	// trc20代币数量乘以小数位
	transferNum := amount.Mul(decimal.New(1, int32(decimals.Int64()))).BigInt()
	transactionExt, err := tronClient.TRC20Send(from, to, contract, transferNum, feeLimit)
	if err != nil {
		fmt.Println("创建转账交易")
		return nil, err
	}
	return transactionExt, nil
}

//func tronClient() {
//		url := "52.53.189.99:50051"
//		c := client.NewGrpcClient(url)
//		if err := c.Start(grpc.WithInsecure()); err != nil || c == nil{
//			panic("建立連結失敗 : "+ url)
//		}
//}

//func transfer(){
//	url := "52.53.189.99:50051"
//	c := client.NewGrpcClient(url)
//	if err := c.Start(grpc.WithInsecure()); err != nil || c == nil{
//		panic("建立連結失敗 : "+ url)
//	}
//
//	tx, err := c.Transfer(myAcc,samAcc,20)
//	if err != nil {
//		fmt.Println(111)
//	}
//	signTx, err := sign.SignTransaction(tx.Transaction, "86540d31f37a88994e7fde229ba82657dd6369ff4f098d9129819afc94051ad0")
//	if err != nil {
//		fmt.Println(222)
//	}
//	result , err := c.Broadcast(signTx)
//	if err != nil {
//		fmt.Println(333)
//	}
//	fmt.Println(common.BytesToHexString(tx.GetTxid()))
//	fmt.Println("result = " ,result)
//}
//
//func Test_TransferTrx(t *testing.T) {
//	c, err := grpcs.NewClient("54.168.218.95:50051")
//	if err != nil {
//		t.Fatal(err)
//	}
//	tx, err := c.Transfer("TFTGMfp7hvDtt4fj3vmWnbYsPSmw5EU8oX", "TVwt3HTg6PjP5bbb5x1GtSvTe1J5FYM2BT", 10000)
//	if err != nil {
//		fmt.Println(111)
//		t.Fatal(err)
//	}
//	signTx, err := sign.SignTransaction(tx.Transaction, "")
//	if err != nil {
//		fmt.Println(222)
//		t.Fatal(err)
//	}
//	err = c.BroadcastTransaction(signTx)
//	if err != nil {
//		fmt.Println(333)
//		t.Fatal(err)
//	}
//	fmt.Println(common.BytesToHexString(tx.GetTxid()))
//
//}