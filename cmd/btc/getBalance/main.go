package main
//
//import (
//	"github.com/btcsuite/btcd/chaincfg"
//
//	"github.com/btcsuite/btcutil"
//	"log"
//)
//
//func main() {
//	// create new client instance
//	client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
//		HTTPPostMode: true,
//		DisableTLS:   true,
//		Host:         "127.0.0.1:8332",
//		User:         "rpcUsername",
//		Pass:         "rpcPassword",
//	}, nil)
//	if err != nil {
//		log.Fatalf("error creating new btc client: %v", err)
//	}
//
//	// list accounts
//	accounts, err := client.ListAccounts()
//	if err != nil {
//		log.Fatalf("error listing accounts: %v", err)
//	}
//	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
//	for label, amount := range accounts {
//		log.Printf("%s: %s", label, amount)
//	}
//
//	// prepare a sendMany transaction
//	receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
//	if err != nil {
//		log.Fatalf("address receiver1 seems to be invalid: %v", err)
//	}
//	receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
//	if err != nil {
//		log.Fatalf("address receiver2 seems to be invalid: %v", err)
//	}
//	receivers := map[btcutil.Address]btcutil.Amount{
//		receiver1: 42,  // 42 satoshi
//		receiver2: 100, // 100 satoshi
//	}
//
//	// create and send the sendMany tx
//	txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
//	if err != nil {
//		log.Fatalf("error sendMany: %v", err)
//	}
//	log.Printf("sendMany completed! tx sha is: %s", txSha.String())
//}