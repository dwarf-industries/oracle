package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/crypto"
)

type VerificationService struct {
	activeChallenges map[string][]byte
}

func (v *VerificationService) GenerateChallenge(expected []byte) ([]byte, error) {
	challenge := make([]byte, 256)
	_, err := io.ReadFull(rand.Reader, challenge)
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	hash := sha256.Sum256(expected)
	challengeID := hex.EncodeToString(hash[:])

	v.activeChallenges[challengeID] = challenge

	return challenge, nil
}

func (v *VerificationService) VerifyChallange(message []byte, signature []byte, expectedAddress string) (*string, error) {
	hash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)))
	publicKey, err := crypto.SigToPub(hash.Bytes(), signature[:65])
	if err != nil {
		return nil, fmt.Errorf("failed to recover public key: %v", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*publicKey).Hex()
	if recoveredAddress != expectedAddress {
		return nil, fmt.Errorf("signature doesn't match")
	}

	hash = sha256.Sum256([]byte(expectedAddress))
	challengeID := hex.EncodeToString(hash[:])
	_, exists := v.activeChallenges[challengeID]
	if !exists {
		return nil, errors.New("no active challenge found for this certificate")
	}

	delete(v.activeChallenges, challengeID)
	return &challengeID, nil
}

func (v *VerificationService) Init() {
	v.activeChallenges = make(map[string][]byte)
}
