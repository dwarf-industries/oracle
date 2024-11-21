package interfaces

import (
	"crypto/ecdsa"
)

type WalletService interface {
	GetWallet(password *string) (*ecdsa.PrivateKey, error)
	SetWallet(wallet *string, password *string) bool
	ActiveWallet() (*ecdsa.PrivateKey, error)
	SignMessage(message []byte) ([]byte, error)
	VerifySignature(message []byte, signature []byte, expectedAddress string) (bool, error)
}
