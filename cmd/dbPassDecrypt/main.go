package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"plugin"
)

var (
	pfn    *plugin.Plugin
	aeskey = "yM82GwpkL2Bvxc3R"
)

func main (){
	decryptStr , _ := Decrypt("bz7w/gJTai3UXS9NHqCuxA==")
	fmt.Println("decryptStr = " , decryptStr)
}

func Decrypt(encryptCode string) (string, error) {
	if true {
		return aesDecrypt(encryptCode, aeskey), nil
	}

	it, err := pfn.Lookup("TbDecrypt")
	if err != nil {
		panic(err)
	}

	decryptStr, err := it.(func(string) (string, error))(encryptCode)
	if err != nil {
		panic(err)
	}

	return decryptStr, nil
}

func aesDecrypt(cryted string, key string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)
	orig = pKCS7UnPadding(orig)
	return string(orig)
}

func pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}