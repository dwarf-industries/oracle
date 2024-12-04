package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"

	"oracle/contracts"
	"oracle/interfaces"
	"oracle/models"
)

type PaymentProcessor struct {
	WalletService    interfaces.WalletService
	RpcService       interfaces.RpcService
	expectedPayments map[string]bool
}

func (p *PaymentProcessor) GeneratePaymentRequest(dataSize int) models.PaymentRequest {

	rand, err := rand.Int(rand.Reader, big.NewInt(int64(10000000)))
	if err != nil {
		return models.PaymentRequest{}
	}

	amount := big.NewInt(int64(dataSize))
	wallet, err := p.WalletService.ActiveWallet()
	if err != nil {
		panic("wallet doesn't exist")
	}

	publicAddress := p.WalletService.GetAddressForPrivateKey(wallet)

	hashInput := fmt.Sprintf("%d-%s-%s", dataSize, publicAddress, rand)
	hashBytes := sha3.Sum256([]byte(hashInput))
	paymentID := hex.EncodeToString(hashBytes[:])

	p.expectedPayments[paymentID] = false
	return models.PaymentRequest{
		PaymentID: paymentID,
		Amount:    amount,
	}
}

func (p *PaymentProcessor) VerifyPayment(expectedPaymentID string) (bool, error) {

	contractAddress := common.HexToAddress(os.Getenv("PAYMENT_LEDGER"))
	contract, err := contracts.NewPaymentLedger(contractAddress, p.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return false, err
	}
	wallet, err := p.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Wallet locked, aborting!")
		return false, err
	}

	transactionOps, err := p.WalletService.NewTransactor(wallet)
	if err != nil {
		fmt.Println("Failed to generate transaction options for active wallet, aborting")
		return false, err
	}

	isPaid, err := contract.PaymentProcessed(&bind.CallOpts{
		From: transactionOps.From,
	}, expectedPaymentID)

	if err != nil {
		return isPaid, err
	}

	delete(p.expectedPayments, expectedPaymentID)
	return isPaid, nil
}

func (p *PaymentProcessor) Init() {
	p.expectedPayments = make(map[string]bool)
}
