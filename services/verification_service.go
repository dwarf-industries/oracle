package services

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

type VerificationService struct {
	activeChallenges map[string][]byte
}

func (v *VerificationService) GenerateChallenge(certificate []byte) ([]byte, error) {
	challenge := make([]byte, 256)
	_, err := io.ReadFull(rand.Reader, challenge)
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	hash := sha256.Sum256(certificate)
	challengeID := hex.EncodeToString(hash[:])

	v.activeChallenges[challengeID] = challenge

	return challenge, nil
}

func (v *VerificationService) VerifyChallange(certificate []byte, signedChallenge []byte) (*string, error) {
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
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, challengeHash[:], signedChallenge)
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
