package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	tronCommon "github.com/fbsobreira/gotron-sdk/pkg/common"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/shopspring/decimal"

)

var (
	myAcc = "TUPnueXTxPWto2bRTevxmMfnupNfzri7M7"
	myAcc1 = "TAAgudMidPhriddmEDJEmVrS5UJyWA75gg"
	samAcc= "TFU3d1TKMvmXcAHyRGfVToyfNtAfdUmH5g"
	generateAcc= "TEJAKcwkUPE1NNMkQ8BoWjcgsWsPJbaCnF"
	tronClient *client.GrpcClient
)

// CreateChainTransaction 创建交易对象
type CreateChainTransaction struct {
	From         string `form:"from" json:"from" binding:"required"`
	To           string `form:"to" json:"to" binding:"required"`
	Amount       string `form:"amount" json:"amount" binding:"required"`
	ContractAddr string `form:"contract_addr" json:"contract_addr"`
	FeeLimit     int64  `form:"fee_limit" json:"fee_limit"`
}

// BroadcastTronTx 广播交易
type BroadcastTronTx struct {
	TxID       string   `form:"txID" json:"txID"`
	RawDataHex string   `form:"raw_data_hex" json:"raw_data_hex" binding:"required"`
	Signature  []string `form:"signature" json:"signature" binding:"required"`
}


func main(){
	connectTronNode()
	transfer()
}

//先創建交易
//再廣播交易
func transfer(){
	privateKey := "86540d31f37a88994e7fde229ba82657dd6369ff4f098d9129819afc94051ad0"

	feeLimit := decimal.New(40, tronCommon.AmountDecimalPoint).IntPart()

	param := &CreateChainTransaction{
		From: myAcc1, To: generateAcc, Amount: "1", ContractAddr: "", FeeLimit: feeLimit,
	}

	rawData, err := CreateTronTransaction(param)
	if err != nil {
		panic("CreateTronTransaction err")
	}

	signature, question := signTronTx(rawData[:], privateKey)

	param3 := &BroadcastTronTx{
		RawDataHex: rawData,  Signature: signature, TxID: question,
	}

	BroadcastTronTransaction(param3)

}

func connectTronNode() {
	tronNodeUrl := "52.53.189.99:50051"
	tronClient = client.NewGrpcClient(tronNodeUrl)
	if err := tronClient.Start(grpc.WithInsecure()); err != nil {
		panic("与tron节点建立连接")
	}
}

func CreateTronTransaction(param *CreateChainTransaction) (string, error){
	amount, err := decimal.NewFromString(param.Amount)
	if err != nil {
		panic("参数错误")
	}
	feeLimit := param.FeeLimit
	if feeLimit == 0 {
		// 默认40trx
		feeLimit = decimal.New(40, 6).IntPart()
	}
	tronTx, err := CreateTronTx(param.From, param.To, param.ContractAddr, amount, feeLimit)
	if err != nil {
		panic("创建交易")
	}
	rawData, err := proto.Marshal(tronTx.GetTransaction().GetRawData())
	if err != nil {
		panic("序列化交易")
	}
	rawDataHex := hex.EncodeToString(rawData)
	fmt.Println("RawDataHex: " + hex.EncodeToString(rawData))

	return rawDataHex, err
}


func BroadcastTronTransaction(param *BroadcastTronTx) {

	rawDataBytes, err := hex.DecodeString(param.RawDataHex)
	if err != nil {
		panic("解码交易数据")
	}
	var transaction = &core.TransactionRaw{}
	if err = proto.Unmarshal(rawDataBytes, transaction); err != nil {
		panic("解码交易数据，转换proto对象")
	}
	var signature [][]byte
	for _, s := range param.Signature {
		signBytes, _ := hex.DecodeString(s)
		signature = append(signature, signBytes)
	}
	tx := &core.Transaction{
		RawData:   transaction,
		Signature: signature,
	}

	h256h := sha256.New()
	h256h.Write(rawDataBytes)
	hash := h256h.Sum(nil)
	txId := common.Bytes2Hex(hash)
	chainTransaction, err := tronClient.Broadcast(tx)
	if err != nil || chainTransaction.GetCode() != api.Return_SUCCESS {
		panic("广播交易，code=%d  ")
	}
	fmt.Println("交易完成 ：" + txId)
}

func signTronTx(rawDataHex string, privateKeyHex string ) (sig []string, question string){
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
	fmt.Println("签名字符串"+ sigStr)

	transaction := &core.TransactionRaw{}
	if err = proto.Unmarshal(rawDataBytes, transaction); err != nil {
		panic(err)
		return
	}
	tx := &core.Transaction{RawData: transaction, Signature: [][]byte{signature}}
	marshal, err := proto.Marshal(tx)
	question = hex.EncodeToString(marshal)

	sigStr2StringArray := []string {sigStr[:]}

	fmt.Println("签名后的交易",question)
	return sigStr2StringArray, question
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
		panic("查询%s小数位"+contract)
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