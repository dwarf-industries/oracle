package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"

	"oracle/interfaces"
)

type WalletService struct {
	PasswordManager interfaces.PasswordManager
	RpcService      interfaces.RpcService
	activeWallet    *ecdsa.PrivateKey
}

func (w *WalletService) NewWallet() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), crypto.NewKeccakState())
}

func (w *WalletService) SetWallet(wallet *string, password *string) (ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(*wallet)
	if err != nil {

		return ecdsa.PrivateKey{}, err
	}

	keyBytes := crypto.FromECDSA(privateKey)
	w.PasswordManager.Encrypt(keyBytes, []byte(*password))

	return *privateKey, nil
}

func (w *WalletService) GetWallet(password *string) (*ecdsa.PrivateKey, error) {
	wallet, err := w.PasswordManager.LoadFromFile("oracle", []byte(*password))
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

func (w *WalletService) SignMessage(message []byte) ([]byte, error) {
	privateKey, err := w.ActiveWallet()
	if err != nil {
		return nil, fmt.Errorf("failed to access active wallet: %v", err)
	}

	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	return signature, nil
}

func (w *WalletService) GetAddressForPrivateKey(key *ecdsa.PrivateKey) string {
	publicKey := key.PublicKey

	xBytes := publicKey.X.Bytes()
	yBytes := publicKey.Y.Bytes()

	xPadded := make([]byte, 32-len(xBytes))
	yPadded := make([]byte, 32-len(yBytes))
	xPadded = append(xPadded, xBytes...)
	yPadded = append(yPadded, yBytes...)

	publicKeyBytes := append(xPadded, yPadded...)

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(publicKeyBytes)
	publicKeyHash := hasher.Sum(nil)

	address := publicKeyHash[len(publicKeyHash)-20:]
	addressHex := hex.EncodeToString(address)
	return addressHex
}

func (w *WalletService) VerifySignature(message []byte, signature []byte, expectedAddress string) (bool, error) {
	if len(signature) != 65 {
		return false, fmt.Errorf("invalid signature length: expected 65 bytes, got %d", len(signature))
	}

	_, _, v := signature[:32], signature[32:64], signature[64]
	if v < 27 {
		v += 27
	}

	if v != 27 && v != 28 {
		return false, fmt.Errorf("invalid recovery ID: %d", v)
	}

	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))

	publicKey, err := crypto.SigToPub(hash.Bytes(), append(signature[:64], v-27))
	if err != nil {
		return false, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*publicKey).Hex()

	return recoveredAddress == expectedAddress, nil
}

func (w *WalletService) GetBalance(wallet *common.Address) big.Int {
	client := w.RpcService.GetClient()

	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to get latest block, aborting!")
		return *big.NewInt(0)
	}
	walletBalance, err := client.BalanceAt(context.Background(), *wallet, big.NewInt(int64(latestBlock)))

	if err != nil {
		fmt.Println("Failed to get latest block, aborting!")
		return *big.NewInt(0)
	}

	return *walletBalance
}
