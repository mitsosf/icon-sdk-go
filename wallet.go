package iconsdk

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"hash"
	"log"
)

type Wallet struct {
	PrivateKey    string
	PublicKey     string
	PublicAddress string
}

func NewWallet(existingKey *string) *Wallet {
	var privateKeyBytes []byte
	var publicKeyBytes []byte
	var pubAddressHash hash.Hash
	var privateKey *ecdsa.PrivateKey
	var err error

	if existingKey != nil {
		privateKeyBytes, err := hex.DecodeString(*existingKey)
		if err != nil {
			log.Fatalf("Failed to decode hex string: %v", err)
		}
		privateKey, err = crypto.ToECDSA(privateKeyBytes)
		if err != nil {
			return nil
		}
	} else {
		privateKey, err = crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
	}

	privateKeyBytes = crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes = crypto.FromECDSAPub(publicKeyECDSA)

	pubAddressHash = sha3.New256()
	pubAddressHash.Write(publicKeyBytes[1:])

	return &Wallet{
		PrivateKey:    hexutil.Encode(privateKeyBytes)[2:],
		PublicKey:     hexutil.Encode(publicKeyBytes)[4:],
		PublicAddress: "hx" + hexutil.Encode(pubAddressHash.Sum(nil)[12:])[2:],
	}
}
