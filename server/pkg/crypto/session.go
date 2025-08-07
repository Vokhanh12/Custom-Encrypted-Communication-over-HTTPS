package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/gob"
	"time"

	"github.com/google/uuid"
)

type SessionInfo struct {
	AESKey     []byte
	IV         []byte
	Salt       []byte
	Expiration time.Time
}

func GenerateSessionID() string {
	return uuid.NewString()
}

func EncryptSessionInfo(sharedSecret []byte, info SessionInfo) ([]byte, error) {
	// 1. Serialize session info
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(info); err != nil {
		return nil, err
	}
	plaintext := buf.Bytes()

	// 2. Derive key from shared secret
	key := sha256.Sum256(sharedSecret)

	// 3. Encrypt using AES-CBC
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	// Pad plaintext to block size
	plaintext = PKCS7Pad(plaintext, aes.BlockSize)

	ciphertext := make([]byte, len(plaintext))
	iv := make([]byte, aes.BlockSize) // all-zero IV as agreed
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func PKCS7Pad(data []byte, blockSize int) []byte {
	padLen := blockSize - len(data)%blockSize
	pad := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, pad...)
}
