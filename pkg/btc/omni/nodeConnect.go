package omni

import (
	"github.com/btcsuite/btcd/rpcclient"
)

var client *Client
var btcAddressMap = make(map[string]bool)

func ConnectOmniNode() *Client {
	btcNodeUrl := "54.169.143.240:8888"
	//btcNodeUrl := "54.254.229.156:8332" //正式節點
	user := "admin"
	pass := "admin"
	var connCfg = &rpcclient.ConnConfig{
		Host:         btcNodeUrl,
		User:         user,
		Pass:         pass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	var err error
	client, err = NewClient(connCfg, nil)
	if err != nil {
		panic("与节点建立连接")
	}

	return client
}