package signer

import (
	"fmt"
	"log"

	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func (r *run) generateKeystore() {
	privKeyHex := "372baf8d01e760d1124571a57b139d9fbb8b557d74b1b25282f71f8415b45fc0"
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}

	privateKey, err := crypto.ToECDSA(privKeyBytes)
	if err != nil {
		log.Fatalf("Failed to convert to ECDSA private key: %v", err)
	}

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.ImportECDSA(privateKey, password)
	if err != nil {
		log.Fatalf("Failed to import ECDSA private key: %v", err)
	}

	fmt.Println("Account address:", account.Address.Hex())
}
