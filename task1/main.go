package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <infura_api_key> <block_number>")
		fmt.Println("Example: go run main.go YOUR_INFURA_API_KEY 1000000")
		os.Exit(1)
	}

	apiKey := os.Args[1]
	blockNumberArg := os.Args[2]

	// 解析区块号
	blockNumber, err := strconv.ParseInt(blockNumberArg, 10, 64)
	if err != nil {
		log.Fatal("Invalid block number: ", err)
	}

	// 连接到Sepolia测试网络
	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", apiKey)
	fmt.Println(infuraURL)
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client: ", err)
	}
	defer client.Close()

	// 检查连接
	_, err = client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to get chain ID: ", err)
	}
	fmt.Println("Successfully connected to Sepolia network")

	// 查询指定区块号的区块信息
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Fatal("Failed to get block: ", err)
	}

	// 输出区块信息
	printBlockInfo(block)
}

func printBlockInfo(block *types.Block) {
	fmt.Println("\n=== Block Information ===")
	fmt.Printf("Block Number: %d\n", block.Number())
	fmt.Printf("Block Hash: %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash: %s\n", block.ParentHash().Hex())
	fmt.Printf("Block Time: %d\n", block.Time())
	fmt.Printf("Transactions Count: %d\n", len(block.Transactions()))
	fmt.Printf("Gas Used: %d\n", block.GasUsed())
	fmt.Printf("Gas Limit: %d\n", block.GasLimit())
	fmt.Printf("Base Fee: %d\n", block.BaseFee())
	fmt.Printf("Difficulty: %d\n", block.Difficulty())
	fmt.Printf("Nonce: %d\n", block.Nonce())
	fmt.Printf("Size: %d bytes\n", block.Size())
	fmt.Printf("Coinbase: %s\n", block.Coinbase().Hex())

	// 显示交易哈希列表
	if len(block.Transactions()) > 0 {
		fmt.Println("\n=== Transaction Hashes ===")
		for i, tx := range block.Transactions() {
			fmt.Printf("Tx %d: %s\n", i, tx.Hash().Hex())
		}
	}
}
