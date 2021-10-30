package omni

import (
	"bytes"
	"encoding/json"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*rpcclient.Client
	config *rpcclient.ConnConfig
}

func NewClient(config *rpcclient.ConnConfig, ntfnHandlers *rpcclient.NotificationHandlers) (*Client, error) {
	btcClient, err := rpcclient.New(config, ntfnHandlers)
	if err != nil {
		return nil, err
	}
	return &Client{Client: btcClient, config: config}, nil
}

// OmniGetTransaction 获取指定Omni交易的详细信息
func (c *Client) OmniGetTransaction(txId string) (*Transaction, error) {
	cmd := &getTransactionCmd{Txid: txId}
	result := new(Transaction)
	err := c.sendCmd(cmd, result)
	return result, err
}

// OmniGetBalance 查询代币余额
func (c *Client) OmniGetBalance(address string, propertyId int64) (GetBalance, error) {
	cmd := &getBalanceCmd{
		Address:    address,
		PropertyId: propertyId,
	}
	result := GetBalance{}
	err := c.sendCmd(cmd, &result)
	return result, err
}

// OmniListBlockTransactions 列出指定区块内的所有omni交易id
func (c *Client) OmniListBlockTransactions(index int64) ([]string, error) {
	cmd := &listBlockTransactionsCmd{Index: index}
	result := make([]string, 0)
	err := c.sendCmd(cmd, &result)
	return result, err
}

func (c *Client) sendCmd(cmd interface{}, result interface{}) error {
	rpcVersion := btcjson.RpcVersion1

	id := c.NextID()
	marshalledJSON, err := btcjson.MarshalCmd(rpcVersion, id, cmd)
	if err != nil {
		panic("序列化参数")
	}

	protocol := "http"
	if !c.config.DisableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host
	bodyReader := bytes.NewReader(marshalledJSON)
	httpReq, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		panic("创建post请求")
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")
	for key, value := range c.config.ExtraHeaders {
		httpReq.Header.Set(key, value)
	}
	httpReq.SetBasicAuth(c.config.User, c.config.Pass)
	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		panic("发送post请求")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("读取body数据")
	}
	respObj := new(JsonRpcResponse)
	err = json.Unmarshal(body, respObj)
	if err != nil {
		return err
	}
	if respObj.Error != nil {
		return respObj.Error
	}
	return json.Unmarshal(respObj.Result, result)
}
