package services

import (
	c "crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
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
	challenge := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, challenge)
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	hash := sha256.Sum256(expected)
	challengeID := hex.EncodeToString(hash[:])

	v.activeChallenges[challengeID] = challenge

	return challenge, nil
}

func (v *VerificationService) GetChallenge(certificate []byte) []byte {
	hash := sha256.Sum256(certificate)
	challengeID := hex.EncodeToString(hash[:])
	return v.activeChallenges[challengeID]
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

func (v *VerificationService) VerifyCertificateChallange(certificate []byte, signedChallenge []byte) (*string, error) {
	cert, err := x509.ParseCertificate(certificate)
	if err != nil {
		return nil, errors.New("invalid certificate")
	}

	publicKey, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain a valid RSA public key")
	}

	hash := sha256.Sum256(certificate)
	challengeID := hex.EncodeToString(hash[:])
	challenge, exists := v.activeChallenges[challengeID]
	if !exists {
		return nil, errors.New("no active challenge found for this certificate")
	}

	challengeHash := sha256.Sum256(challenge)
	err = rsa.VerifyPKCS1v15(publicKey, c.SHA256, challengeHash[:], signedChallenge)
	if err != nil {
		return nil, errors.New("challenge verification failed")
	}
	fmt.Println("Challenge verified successfully for:", cert.Subject.CommonName)
	delete(v.activeChallenges, challengeID)
	return &challengeID, nil
}

func (v *VerificationService) Init() {
	v.activeChallenges = make(map[string][]byte)
}
