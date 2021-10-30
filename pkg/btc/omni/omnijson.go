package omni

import (
	"encoding/json"
	"github.com/btcsuite/btcd/btcjson"
)

type getTransactionCmd struct {
	Txid string
}

type Transaction struct {
	TxId             string `json:"txid"`
	Fee              string `json:"fee"`
	SendingAddress   string `json:"sendingaddress"`   // 发送方比特币地址
	ReferenceAddress string `json:"referenceaddress"` // 作为参照的比特币地址
	IsMine           bool   `json:"ismine"`
	Version          int    `json:"version"`
	TypeInt          int    `json:"type_int"` // 0是转账类型
	Type             string `json:"type"`     // "Simple send"是转账类型
	PropertyId       int    `json:"propertyid"`
	Divisible        bool   `json:"divisible"`
	Amount           string `json:"amount"`
	Valid            bool   `json:"valid"`
	BlockHash        string `json:"blockhash"`
	BlockTime        int64  `json:"blocktime"`
	PositionInBlock  int    `json:"positioninblock"`
	Block            int64  `json:"block"`
	Confirmations    int64  `json:"confirmations"`
}

// 查询omni余额的参数
type getBalanceCmd struct {
	Address    string `json:"address"`    // 地址，字符串，必需
	PropertyId int64  `json:"propertyid"` // 资产ID，数值，必需
}

type listBlockTransactionsCmd struct {
	Index int64
}

type GetBalance struct {
	Balance  string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen   string `json:"frozen"`
}

func init() {
	flags := btcjson.UsageFlag(0)
	btcjson.MustRegisterCmd("omni_gettransaction", (*getTransactionCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_listblocktransactions", (*listBlockTransactionsCmd)(nil), flags)
	btcjson.MustRegisterCmd("omni_getbalance", (*getBalanceCmd)(nil), flags)
}

type JsonRpcResponse struct {
	Result json.RawMessage   `json:"result"`
	Error  *btcjson.RPCError `json:"error"`
}
