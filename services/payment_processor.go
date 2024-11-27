package services

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"oracle/interfaces"
	"oracle/models"
)

type PaymentProcessor struct {
	WalletService    interfaces.WalletService
	expectedPayments map[string]bool
}

func (p *PaymentProcessor) GeneratePaymentRequest(dataSize int) models.PaymentRequest {
	amount := big.NewInt(int64(dataSize))
	wallet, err := p.WalletService.ActiveWallet()
	if err != nil {
		panic("wallet doesn't exist")
	}

	publicAddress := p.WalletService.GetAddressForPrivateKey(wallet)

	_, nodePaymentIdentifier, err := p.generatePaymentID(wallet)
	if err != nil {
		panic("failed to generate node payment identifier")
	}

	hashInput := fmt.Sprintf("%d-%s-%s", dataSize, nodePaymentIdentifier, publicAddress)
	hash := sha256.Sum256([]byte(hashInput))

	paymentID := hex.EncodeToString(hash[:])
	p.expectedPayments[paymentID] = false

	return models.PaymentRequest{
		PaymentID: paymentID,
		Amount:    amount,
	}
}

func (p *PaymentProcessor) VerifyPayment(expectedPaymentID string) (bool, error) {
	isPaid := p.expectedPayments[expectedPaymentID]
	if isPaid {
		delete(p.expectedPayments, expectedPaymentID)
		return false, nil
	}

	delete(p.expectedPayments, expectedPaymentID)
	return true, fmt.Errorf("paymentID mismatch")
}

func (p *PaymentProcessor) generatePaymentID(nodePrivateKey *ecdsa.PrivateKey) (string, string, error) {
	timestamp := time.Now().UnixNano()
	rawID := fmt.Sprintf("PAY-%d", timestamp)

	hash := sha256.Sum256([]byte(rawID))

	signature, err := crypto.Sign(hash[:], nodePrivateKey)
	if err != nil {
		return "", "", err
	}

	return rawID, hex.EncodeToString(signature), nil
}
