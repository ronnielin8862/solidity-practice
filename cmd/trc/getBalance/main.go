package main

import (
	"encoding/json"
	"fmt"
	"github.com/JFJun/trx-sign-go/grpcs"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
	"solidity2/pkg/trx/httpClient"
)


var myAcc = "TUPnueXTxPWto2bRTevxmMfnupNfzri7M7"
var myAcc1 = "TAAgudMidPhriddmEDJEmVrS5UJyWA75gg"
var samAcc= "TFU3d1TKMvmXcAHyRGfVToyfNtAfdUmH5g"


func main() {
	//getBalance()
	//getBalance1()
	getBalanceFromHttp()
}

//可用
func getBalance(){
	url := "52.53.189.99:50051"
	c := client.NewGrpcClient(url)
	if err := c.Start(grpc.WithInsecure()); err != nil {
		fmt.Println("建立連結失敗", url)
	}
	acc , err := c.GetAccount(myAcc1) // ok

	fmt.Println(acc , err)
}

//可用
func getBalanceFromHttp(){
	httpClient.GetBalance(samAcc, "00000000020e931413765b06f89f534484721d9ffb273413c58e84fad8f91ab5" , "31981230")
}

//不可用
//https://github.com/JFJun/trx-sign-go/blob/master/test/tx_test.go
func getBalance1() {
	c, err := grpcs.NewClient("52.53.189.99:50051")
	if err != nil {
		fmt.Println("err = " ,err)
	}

	acc, err := c.GetTrxBalance(samAcc)
	if err != nil {
		fmt.Println("err = " ,err)
	}
	d, _ := json.Marshal(acc)
	fmt.Println(string(d))
	fmt.Println(acc.GetBalance())

}