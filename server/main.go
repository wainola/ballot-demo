package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	testNet := os.Getenv("URL")
	accountAddress := os.Getenv("ACCOUNT")
	conn, err := ethclient.Dial(testNet)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the Ganache network")

	address := common.HexToAddress(accountAddress)

	balance, err := conn.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)
}
