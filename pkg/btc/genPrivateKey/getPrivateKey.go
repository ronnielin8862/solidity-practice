package genPrivateKey

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func ByteString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02X", b[i])
	}
	return s
}

func b58encode(b []byte) (s string) {

	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	fmt.Println("Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)")
	bcpy := append([]byte{ver}, b...)
	fmt.Println(ByteString(bcpy))
	fmt.Println("=======================")

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 Hash #1 */
	fmt.Println("SHA-256 hash on the extended RIPEMD-160 result")
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	fmt.Println(ByteString(hash1))
	fmt.Println("=======================")

	/* SHA256 Hash #2 */
	fmt.Println("SHA-256 hash on the result of the previous SHA-256 hash")
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	fmt.Println(ByteString(hash2))
	fmt.Println("=======================")

	/* Append first four bytes of hash */
	fmt.Println("Take the first 4 bytes of the second SHA-256 hash. This is the address checksum")
	fmt.Println(ByteString(hash2[0:4]))
	fmt.Println("=======================")

	fmt.Println("Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.")
	bcpy = append(bcpy, hash2[0:4]...)
	fmt.Println(ByteString(bcpy))
	fmt.Println("=======================")

	/* Encode base58 string */
	s = b58encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	fmt.Println("Convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format")
	fmt.Println(s)
	fmt.Println("=======================")

	return s
}


func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}



//const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey []byte
	PublicKey  []byte
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() (address string) {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	/* Convert the public key to bytes */
	pubBytes := w.PublicKey

	/* SHA256 Hash */
	fmt.Println("SHA-256 hashing on the public key")
	sha256H := sha256.New()
	sha256H.Reset()
	sha256H.Write(pubBytes)
	pubHash1 := sha256H.Sum(nil)
	fmt.Println(ByteString(pubHash1))
	fmt.Println("=======================")

	/* RIPEMD-160 Hash */
	fmt.Println("RIPEMD-160 hashing on the result of SHA-256")
	ripemd160H := ripemd160.New()
	ripemd160H.Reset()
	ripemd160H.Write(pubHash1)
	pubHash2 := ripemd160H.Sum(nil)
	fmt.Println(ByteString(pubHash2))
	fmt.Println("=======================")
	/* Convert hash bytes to base58 check encoded sequence */
	//const version = byte(0x00)
	address = b58checkencode(0x00, pubHash2)

	return address
}

func convert( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,",")
}

const privateKeyBytesLen = 32

func newKeyPair() ([]byte, []byte) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256() , rand.Reader)
	fmt.Println("private key = ", privateKey)
	if err != nil {
		log.Panic(err)
	}
	d := privateKey.D.Bytes()
	b := make([]byte, 0, privateKeyBytesLen)
	priKet := paddedAppend(privateKeyBytesLen, b, d)
	fmt.Println("priKet = ", priKet)
	fmt.Println("priKet = ", convert(priKet))
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	return priKet, pubKey
}

// ToWIF converts a Bitcoin private key to a Wallet Import Format string.
func ToWIF(private []byte) (wif string) {
	/* Convert bytes to base-58 check encoded string with version 0x80 */
	wif = b58checkencode(0x80, private)

	return wif
}