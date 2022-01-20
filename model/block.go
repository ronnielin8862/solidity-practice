package model

import "time"

type Block struct {
	id               int64 `xorm:"not null pk comment('后台用户id') BIGINT"`
	Difficulty       string
	ExtraData        string
	GasLimit         string
	GasUsed          string
	Hash             string
	LogsBloom        string
	Miner            string
	MixHash          string
	Nonce            string
	Number           string
	ParentHash       string
	ReceiptsRoot     string
	Sha3Uncles       string
	Size             string
	StateRoot        string
	Timestamp        time.Time
	TotalDifficulty  string
	Transactions     string
	TransactionsRoot string
	Uncles           string
}
