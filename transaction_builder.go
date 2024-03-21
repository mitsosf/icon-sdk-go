package iconsdk

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"strconv"
	"time"
)

type TransactionBuilder struct {
	transaction Transaction
}

func NewTransactionBuilder(iconService *IconService) *TransactionBuilder {
	return &TransactionBuilder{
		transaction: *NewTransaction(iconService),
	}
}

func (t *TransactionBuilder) Method(method string) *TransactionBuilder {
	t.transaction.SetMethod(method)
	return t
}

func (t *TransactionBuilder) BlockHeight(height string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["height"] = height
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) BlockHash(hash string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["hash"] = hash
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Address(address string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["address"] = address
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) TxHash(txHash string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["txHash"] = txHash
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) From(from string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["from"] = from
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) To(to string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["to"] = to
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Value(value string) *TransactionBuilder {
	if value[:2] != "0x" {
		value = icxToHex(value, nil)
	}
	params := make(map[string]interface{})
	params["value"] = value
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Version(version string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["version"] = version
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Nid(nid string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["nid"] = nid
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Nonce(nonce string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["nonce"] = nonce
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) StepLimit(stepLimit string) *TransactionBuilder {
	params := make(map[string]interface{})
	params["stepLimit"] = stepLimit
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Timestamp() *TransactionBuilder {
	timestamp := "0x" + strconv.FormatInt(time.Now().UnixNano()/int64(time.Microsecond), 16)

	params := make(map[string]interface{})
	params["timestamp"] = timestamp
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Message(message string) *TransactionBuilder {
	// Timestamp in microseconds, then convert to hex
	params := make(map[string]interface{})
	params["dataType"] = "message"
	params["data"] = "0x" + hex.EncodeToString([]byte(message))
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Call(call_params map[string]interface{}) *TransactionBuilder {
	// Timestamp in microseconds, then convert to hex
	params := make(map[string]interface{})
	params["dataType"] = "call"
	params["data"] = call_params
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Sign(wallet Wallet) *TransactionBuilder {
	serializedTransaction, _ := serializeTransaction(t.transaction.data["params"], true)
	serializedTransactionBytes, err := hex.DecodeString(serializedTransaction)
	if err != nil {
		log.Fatalf("Failed to decode serialized transaction: %v", err)
	}

	privateKeyBytes, err := hex.DecodeString(wallet.PrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Sign the transaction
	sig, err := secp256k1.Sign(serializedTransactionBytes, privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	// Optionally, append recovery ID if needed, similar to Rust's implementation

	// Encode the signature in Base64
	signatureBase64 := base64.StdEncoding.EncodeToString(sig)

	params := make(map[string]interface{})
	params["signature"] = signatureBase64
	t.transaction.SetParams(params)
	return t
}

func (t *TransactionBuilder) Build() Transaction {
	return t.transaction
}
