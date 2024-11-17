package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

type PasswordManager struct {
}

func (p *PasswordManager) Encrypt(plaintext, key []byte) (bool, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return false, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return false, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	os.WriteFile("oracle", ciphertext, 0600)

	return true, nil
}

func (p *PasswordManager) LoadFromFile(filename string, password []byte) ([]byte, error) {
	file, err := os.ReadFile(filename)

	if err != nil {
		return []byte{}, err
	}

	return p.decrypt(password, file)
}

func (p *PasswordManager) decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
