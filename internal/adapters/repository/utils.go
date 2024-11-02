package repository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func EncryptString(plaintext *string, keyString string) (*string, error) {
	// See: https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	if plaintext == nil {
		return nil, nil
	}

	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cipher: %w", err)
	}

	plaintextBytes := []byte(*plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("failed to read iv: %w", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintextBytes)

	encoded := hex.EncodeToString(ciphertext)
	return &encoded, nil
}

func DecryptString(ciphertextHex *string, keyString string) (*string, error) {
	// See: https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865
	if ciphertextHex == nil {
		return nil, nil
	}
	key := []byte(keyString)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cipher: %w", err)
	}

	ciphertext, err := hex.DecodeString(*ciphertextHex)
	if err != nil || len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	plaintextBytes := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintextBytes, ciphertext)

	decoded := string(plaintextBytes)
	return &decoded, nil
}
