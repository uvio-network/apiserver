package signer

import (
    "crypto/ecdsa"
    // "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "context"
)

func (r *run) sendTransaction() {
    // Connect to the Ethereum client
    client, err := ethclient.Dial("https://sepolia-rollup.arbitrum.io/rpc")
    if err != nil {
        log.Fatalf("Failed to connect to the Ethereum client: %v", err)
    }

    // Load the keystore file
    keystoreFile := "./tmp/UTC--2024-08-19T15-49-12.517663095Z--d8d85fdccc0e63e4d934a69bad45c7aea8c45a50"
    keyJSON, err := ioutil.ReadFile(keystoreFile)
    if err != nil {
        log.Fatalf("Failed to read the keystore file: %v", err)
    }

    // Decrypt the key
    passphrase := "secret"
    key, err := keystore.DecryptKey(keyJSON, passphrase)
    if err != nil {
        log.Fatalf("Failed to decrypt the key: %v", err)
    }

    privateKey := key.PrivateKey

    // Use the private key to sign and send a transaction
    sendTransaction(client, privateKey)
}

func sendTransaction(client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatalf("Failed to get nonce: %v", err)
    }

    value := big.NewInt(1000000000000000) // in wei (0.001 ETH)
    gasLimit := uint64(30000)              // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatalf("Failed to get gas price: %v", err)
    }

    toAddress := common.HexToAddress("0x318d0059efE546b5687FA6744aF4339391153981")
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatalf("Failed to get network ID: %v", err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatalf("Failed to sign transaction: %v", err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatalf("Failed to send transaction: %v", err)
    }

    fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
