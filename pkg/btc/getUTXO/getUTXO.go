package getUTXO

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"

)

//根据address获取未花费的tx
func GetUnspentByAddress(address string, client *rpcclient.Client) (unspents []btcjson.ListUnspentResult, err error) {
	btcAdd, err := btcutil.DecodeAddress(address, &chaincfg.RegressionNetParams)
	if err != nil {
		return nil, err
	}
	adds := [1]btcutil.Address{btcAdd}
	unspents, err = client.ListUnspentMinMaxAddresses(1, 999999, adds[:])
	if err != nil {
		return nil, err
	}
	return unspents, nil
}
