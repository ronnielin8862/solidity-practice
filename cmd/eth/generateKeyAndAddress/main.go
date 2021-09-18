package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//生成私鑰
	privateKey, _ := crypto.GenerateKey()
	fmt.Println("privateKey = ", privateKey)

	//橢圓加密 私鑰
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateKeyBytes = ", privateKeyBytes)

	//生成簽章用的私鑰  轉成１６進位，並刪除前兩位數"0x"
	usingPrivateKey := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println("usingPrivateKey 2 = ", usingPrivateKey)

	//生成public key
	publicKey := privateKey.Public()
	fmt.Println("publicKey = ", publicKey)

	//這步不確定在幹嘛
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fmt.Println("publicKeyECDSA", publicKeyECDSA)

	//公鑰區要需要進行橢圓加密？為何？
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKeyByte = ", publicKeyBytes)

	//生成實際使用的公鑰  去掉0x和 前 2 個字符04，它們總是 EC 前綴，不是必需的。
	usingPublicKey := hexutil.Encode(publicKeyBytes)[4:]
	fmt.Println("usingPublicKey 轉為１６進位 及取四位以後    = ", usingPublicKey)

	//公共地址  公鑰做sha3-256 hash,取最後２０個字節，前面加上0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address = ", address)

}
