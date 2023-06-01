package main

import (
	"eth-demo/scrds"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

const (
	// your private key
	privateKey = ""
	// your address
	yourAddress = ""
	// your signature
	sig             = "26fbea64c0f45ff298ad3fe9e653f25e4a663cf3b5d52840f8b09a043080902d78bf9f8265bca90b540dab80d8530a020e7a8cffe18de893a675f08e0ebe03d21c"
	chainId         = 421613
	contractAddress = "0x20287418E3C8a6aeEdAE3E2Fc0463EC534481b94"
	rpcUrl          = "https://goerli-rollup.arbitrum.io/rpc"
)

func main() {
	conn, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sCrds, err := scrds.NewSCRDS(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatal(err)
	}

	// read
	nonce, err := sCrds.ExchangerNonces(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, common.HexToAddress(yourAddress))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nonce)

	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// invoke
	auth, err := bind.NewKeyedTransactorWithChainID(pk, new(big.Int).SetInt64(chainId))
	if err != nil {
		log.Fatal(err)
	}

	myAddr := crypto.PubkeyToAddress(pk.PublicKey)
	fmt.Println(myAddr.Hex())

	tx, err := sCrds.Exchange(&bind.TransactOpts{
		From: auth.From,
		//Nonce:     nil,
		Signer: auth.Signer,
		//Value:     nil,
		//GasPrice:  nil,
		//GasFeeCap: nil,
		//GasTipCap: nil,
		//GasLimit:  0,
		//Context:   nil,
		//NoSend:    false,
	}, big.NewInt(100), common.Hex2Bytes(sig))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash())
}
