package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"math/big"
	"strings"

)

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func main(){
	privateKeyHex, privateKey, err := GenerateRandomPrivateKey()
	if err != nil {
		panic("生成随机私钥")
	}

	fmt.Println("privateKeyHex = " , privateKeyHex)
	fmt.Println("privateKey = " , privateKey)

	publicKey := privateKey.Public()
	fmt.Println("publicKey = " , publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("私钥格式转换")
	}

	fmt.Println("publicKeyECDSA = " , publicKeyECDSA)

	addr := PubToTronAddress(*publicKeyECDSA)

	fmt.Println("add = " , addr)

}

// GenerateRandomPrivateKey 生成随机的私钥
func GenerateRandomPrivateKey() (string, *ecdsa.PrivateKey, error) {
	uuidStr := uuid.New().String() + uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	privateKey, err := crypto.HexToECDSA(uuidStr)
	if err != nil {
		panic("生成私钥")
	}
	return uuidStr, privateKey, nil
}

// PubToTronAddress 公钥推出tron地址
func PubToTronAddress(pub ecdsa.PublicKey) string {
	eth := strings.ToLower(crypto.PubkeyToAddress(pub).String())
	hexString := "41" + eth[2:]
	address, _ := FromHexAddress(hexString)
	return address
}

// FromHexAddress 41 ---- > T
func FromHexAddress(hexAddress string) (string, error) {
	addrByte, err := hex.DecodeString(hexAddress)
	if err != nil {
		return "", err
	}

	sha := sha256.New()
	sha.Write(addrByte)
	shaStr := sha.Sum(nil)

	sha2 := sha256.New()
	sha2.Write(shaStr)
	shaStr2 := sha2.Sum(nil)

	addrByte = append(addrByte, shaStr2[:4]...)

	return string(Base58Encode(addrByte)), nil
}

// Base58Encode Base58编码
func Base58Encode(input []byte) []byte {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	reverseBytes(result)
	return result
}

// ReBytes ReverseBytes 翻转字节
func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}