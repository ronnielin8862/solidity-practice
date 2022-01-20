package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"solidity2/response"

	"strconv"
	"strings"
)

type Client struct {
	rpcClient *rpc.Client
	EthClient *ethclient.Client
}

func Connect(host string) (*Client, error) {
	rpcClient, err := rpc.Dial(host)
	if err != nil {
		return nil, err
	}
	ethClient := ethclient.NewClient(rpcClient)
	return &Client{rpcClient, ethClient}, nil
}

func (ec *Client) GetBlockNumber(ctx context.Context) (*big.Int, error) {
	var result hexutil.Big
	err := ec.rpcClient.CallContext(ctx, &result, "eth_blockNumber")
	return (*big.Int)(&result), err
}

func main() {
	client, err := Connect("https://data-seed-prebsc-2-s3.binance.org:8545")
	if err != nil {
		fmt.Println(err.Error())
	}

	////取得最新區塊號
	//blockNumber, err := client.GetBlockNumber(context.TODO())
	//fmt.Println("blockNumber = ", blockNumber)
	//
	////以txHash取得交易
	//hash := common.HexToHash("0xaaa18d0adfd480a59780dd4ce89056f6d1b5ce547ae0fc9c5a843aada7141284")
	//fmt.Println("hash = ", hash)
	//client.GetTransactionByHash(context.TODO(), hash)
	//
	////以txHash取得交易 - rpc
	//client.GetTransactionByHashAndRpc(context.TODO(), "0x8ea74292c1ccd7f9f0d923c1973df375d5f46419d5d9d39624c3c7ab4a9ff84f")

	//以區塊號 取得區塊
	client.GetBlockByNumber(context.TODO(), "0xD608F0")
}

func (ec *Client) GetBlockByNumber(ctx context.Context, num string) {
	var block response.BlockNumResponse
	err := ec.rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", num, false)
	if err != nil {
		fmt.Println("error ", err)
		return
	}
	block.Difficulty = hexaNumberToInteger(block.Difficulty)
	block.GasLimit = hexaNumberToInteger(block.GasLimit)
	block.GasUsed = hexaNumberToInteger(block.GasUsed)
	block.Nonce = hexaNumberToInteger(block.Nonce)
	block.Number = hexaNumberToInteger(block.Number)
	block.Size = hexaNumberToInteger(block.Size)
	block.Timestamp = hexaNumberToInteger(block.Timestamp)
	block.TotalDifficulty = hexaNumberToInteger(block.TotalDifficulty)

	//global.Db.InsertOne()
	fmt.Println("block = ", block.Difficulty)
}

func (ec *Client) GetTransactionByHash(ctx context.Context, txHash common.Hash) {
	tx, isPending, _ := ec.EthClient.TransactionByHash(context.TODO(), txHash)
	fmt.Println("tx, isPending = ", tx, isPending)
	fmt.Println(tx.Hash())
	fmt.Println(tx.Gas())
	fmt.Println(tx.Data())
	fmt.Println(tx.To())
}

func (ec *Client) GetTransactionByHashAndRpc(ctx context.Context, txHash string) {
	var result response.TransactionResponse
	err := ec.rpcClient.CallContext(ctx, &result, "eth_getTransactionByHash", txHash)
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	blockNum, _ := strconv.ParseUint(hexaNumberToInteger(result.BlockNumber), 16, 64)
	value, _ := strconv.ParseUint(hexaNumberToInteger(result.Value), 16, 64)
	fmt.Println("block num , value  = ", blockNum, value)
}

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
