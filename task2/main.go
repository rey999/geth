package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到Sepolia测试网络
	infuraURL := "https://sepolia.infura.io/v3/f20466115d654fbf9e7b59b22d75fd86" // 替换为你的Infura项目ID
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal("无法连接到以太坊客户端:", err)
	}
	defer client.Close()

	fmt.Println("成功连接到Sepolia网络")

	// 检查连接
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("无法获取网络ID:", err)
	}
	fmt.Printf("网络ID: %s\n", chainID.String())

	// 发送方私钥 (这是一个示例私钥，请替换为你自己的私钥)
	privateKey, err := crypto.HexToECDSA("xxx") // 私钥不包含0x前缀
	if err != nil {
		log.Fatal("无法解析私钥:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法获取公钥")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("发送方地址: %s\n", fromAddress.Hex())

	// 接收方地址
	toAddress := common.HexToAddress("xxx") // 替换为实际接收方地址

	// 获取发送方nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("无法获取nonce:", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)

	// 设置交易参数
	value := big.NewInt(100000000000000) // 0.00001 ETH (单位为wei)
	gasLimit := uint64(21000)            // 转账的标准gas限制

	// 建议的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("无法获取建议的gas价格:", err)
	}
	fmt.Printf("建议的Gas价格: %s\n", gasPrice.String())

	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("无法签名交易:", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("无法发送交易:", err)
	}

	fmt.Printf("交易已发送! 交易哈希: %s\n", signedTx.Hash().Hex())
}
