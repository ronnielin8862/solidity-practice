package transfer

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"log"
	"solidity2/pkg/btc/getUTXO"
)

var addressPubley *btcutil.AddressPubKey

//转账
//addrForm来源地址，addrTo去向地址
//transfer 转账金额
//fee 小费
func SendAddressToAddress(addrFrom, addrTo string, transfer, fee float64, client *rpcclient.Client, prvKey string) error {

	UTXO, err := getUTXO.GetUnspentByAddress(addrFrom, client)
	if err != nil {
		panic(err)
	}
	//各种参数声明 可以构建为内部小对象
	outsu := float64(0)                 //unspent单子相加
	feesum := fee                       //交易费总和
	totalTran := transfer + feesum      //总共花费
	var pkscripts [][]byte              //txin签名用script
	tx := wire.NewMsgTx(wire.TxVersion) //构造tx

	for _, v := range UTXO {
		if v.Amount == 0 {
			continue
		}
		if outsu < totalTran {
			outsu += v.Amount
			{
				//txin输入-------start-----------------
				hash, _ := chainhash.NewHashFromStr(v.TxID)
				outPoint := wire.NewOutPoint(hash, v.Vout)
				txIn := wire.NewTxIn(outPoint, nil, nil)

				tx.AddTxIn(txIn)

				//设置签名用script
				txinPkScript, err := hex.DecodeString(v.ScriptPubKey)
				if err != nil {
					return err
				}
				pkscripts = append(pkscripts, txinPkScript)
			}
		} else {
			break
		}
	}

	if outsu < totalTran {
		panic("餘額不足")
	}
	// 输出1, 给form----------------找零-------------------
	addrf, err := btcutil.DecodeAddress(addrFrom, &chaincfg.RegressionNetParams)
	if err != nil {
		return err
	}
	pkScriptf, err := txscript.PayToAddrScript(addrf)
	if err != nil {
		return err
	}
	baf := int64((outsu - totalTran) * 1e8)
	tx.AddTxOut(wire.NewTxOut(baf, pkScriptf))
	//输出2，给to
	addrt, err := btcutil.DecodeAddress(addrTo, &chaincfg.RegressionNetParams)
	if err != nil {
		return err
	}
	pkScriptt, err := txscript.PayToAddrScript(addrt)
	if err != nil {
		return err
	}
	bat := int64(transfer * 1e8)
	tx.AddTxOut(wire.NewTxOut(bat, pkScriptt))
	//-------------------输出填充end------------------------------
	fmt.Println("tx:",tx,"; prvKey:",prvKey,"; pkscripts:",pkscripts)
	err = sign(tx, prvKey, pkscripts) //签名
	if err != nil {
		return err
	}
	//广播
	txHash, err := client.SendRawTransaction(tx, false)
	if err != nil {
		return err
	}

	fmt.Println("Transaction successfully signed")
	fmt.Println(txHash.String())
	return nil
}


//签名
//privkey的compress方式需要与TxIn的
func sign(tx *wire.MsgTx, privKey string, pkScripts [][]byte) error {
	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return err
	}
	/* lookupKey := func(a btcutil.Address) (*btcec.PrivateKey, bool, error) {
	    return wif.PrivKey, false, nil
	} */
	for i, _ := range tx.TxIn {
		script, err := txscript.SignatureScript(tx, i, pkScripts[i], txscript.SigHashAll, wif.PrivKey, false)
		//script, err := txscript.SignTxOutput(&chaincfg.RegressionNetParams, tx, i, pkScripts[i], txscript.SigHashAll, txscript.KeyClosure(lookupKey), nil, nil)
		if err != nil {
			return err
		}
		tx.TxIn[i].SignatureScript = script
		vm, err := txscript.NewEngine(pkScripts[i], tx, i,
			txscript.StandardVerifyFlags, nil, nil, -1)
		if err != nil {
			return err
		}
		err = vm.Execute()
		if err != nil {
			return err
		}
		log.Println("Transaction successfully signed")
	}
	return nil
}