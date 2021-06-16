package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
)

var (
	// we're running on testnet3
	testnet3Parameters = &chaincfg.TestNet3Params
)

func main() {
	fmt.Printf("mas.s62 pset03 - utxohunt\n")

	// Task #1 make an address pair
	// Call AddressFrom PrivateKey() to make a keypair
	address, _ := AddressFromPrivateKey()

	fmt.Printf("Task#1 Address: %s\n", address)
	//mnMEoorvojs62PziUdsjJtPwdiDDJY4yo4

	// Task #2 make a transaction
	// Call EZTxBuilder to make a transaction
	tx1 := EZTxBuilder()
	fmt.Printf("Task#2 transaction hex: %s\n", TxToHex(tx1))
	//https://testnet.smartbit.com.au/tx/001d1b95fba51001e2acdbc62ea5c109a51372f988c2c0433ff8d1c4b793bf2f
	// ./bitcoin-cli sendrawtransaction "$tx1"

	// task 3, call OpReturnTxBuilder() the same way EZTxBuilder() was used
	tx2 := OpReturnTxBuilder()
	fmt.Printf("Task#3 transaction hex: %s\n", TxToHex(tx2))
	//https://testnet.smartbit.com.au/tx/c879a3cab45dc4d233a39e8ca8bf87fc44f9b88017b36e863354a9a0d331039d
	// ./bitcoin-cli sendrawtransaction "$tx2"

}
