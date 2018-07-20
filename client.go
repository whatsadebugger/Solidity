package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	acc := common.HexToAddress("0x6e03377a41a708612046c665b30a8653533e05Fb")

	fmt.Println("connected successfully")
	balance, err := client.BalanceAt(context.Background(), acc, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("balance:", balance.String())
}
