package interfaces

import (
	"crypto/ecdsa"
)

type WalletService interface {
	NewWallet() (*ecdsa.PrivateKey, error)
	GetWallet(password *string) (*ecdsa.PrivateKey, error)
	SetWallet(wallet *string, password *string) bool
	ActiveWallet() (*ecdsa.PrivateKey, error)
	GetAddressForPrivateKey(key *ecdsa.PrivateKey) string
	SignMessage(message []byte) ([]byte, error)
	VerifySignature(message []byte, signature []byte, expectedAddress string) (bool, error)
}
