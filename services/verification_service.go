package services

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

type VerificationService struct {
	activeChallenges map[string][]byte
	mu               sync.Mutex
}

func (v *VerificationService) GenerateChallenge(expected []byte) ([]byte, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

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
	v.mu.Lock()
	defer v.mu.Unlock()

	hash := sha256.Sum256(certificate)
	challengeID := hex.EncodeToString(hash[:])
	return v.activeChallenges[challengeID]
}

func (v *VerificationService) VerifyChallange(message []byte, signature []byte, expectedAddress string) (*string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

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
	v.mu.Lock()
	defer v.mu.Unlock()

	cert, err := x509.ParseCertificate(certificate)
	if err != nil {
		return nil, errors.New("invalid certificate")
	}

	publicKey, ok := cert.PublicKey.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("certificate does not contain a valid Ed25519 public key")
	}

	if len(cert.RawTBSCertificate) == 0 {
		return nil, errors.New("certificate does not contain RawTBSCertificate")
	}

	if !ed25519.Verify(publicKey, cert.RawTBSCertificate, signedChallenge) {
		return nil, errors.New("signature verification failed")
	}
	encode := hex.EncodeToString(certificate)
	return &encode, nil
}

func (v *VerificationService) Init() {
	v.activeChallenges = make(map[string][]byte)
}
