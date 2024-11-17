package services

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"oracle/interfaces"
)

type WalletService struct {
	PasswordManager interfaces.PasswordManager
	activeWallet    *ecdsa.PrivateKey
}

func (w *WalletService) SetWallet(wallet *string, password *string) bool {
	privateKey, err := crypto.HexToECDSA(*wallet)
	if err != nil {
		fmt.Println(err)
		return false
	}

	keyBytes := crypto.FromECDSA(privateKey)
	w.PasswordManager.Encrypt(keyBytes, []byte(*password))

	return true
}

func (w *WalletService) GetWallet(password *string) (*ecdsa.PrivateKey, error) {
	wallet, err := w.PasswordManager.LoadFromFile("password", []byte(*password))
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(string(wallet))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	w.activeWallet = privateKey
	return privateKey, nil
}

func (w *WalletService) ActiveWallet() (*ecdsa.PrivateKey, error) {
	if w.activeWallet == nil {
		return nil, fmt.Errorf("wallet is not unlocked")
	}

	return w.activeWallet, nil
}
