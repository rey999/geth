package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	counter "github.com/rey999/geth"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/f20466115d654fbf9e7b59b22d75fd86")
	if err != nil {
		log.Fatal("Failed to connect to Sepolia:", err)
	}

	privateKey, err := crypto.HexToECDSA("dfa734e3d1ef518b767b0f1c4efc42cd961a882a8daa3fa68188d876d41895fd")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 构造授权对象
	chainID := big.NewInt(11155111) // Sepolia
	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	contractAddress := common.HexToAddress("0xe162CEBCE98C74A75D5C971AA925E98d1bcc3AF3")
	instance, err := counter.NewCounter(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 调用 increment
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Increment Tx Hash:", tx.Hash().Hex())

	//调用 Increment
	instance.Increment(auth)
	instance.Increment(auth)
	instance.Increment(auth)
	instance.Increment(auth)

	// 调用 getCount
	count, err := instance.GetCount(&bind.CallOpts{
		Pending: false,
		From:    fromAddress,
		Context: context.Background(),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Counter Value:", count)
}
