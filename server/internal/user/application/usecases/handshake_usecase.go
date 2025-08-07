package usecases

import (
	"context"
	"crypto/aes"
	"crypto/ecdh"
	"crypto/rand"
	"encoding/base64"
	"time"

	"myapp/internal/user/application/commands"
	"myapp/internal/user/domain"
	"myapp/pkg/crypto"
)

type HandshakeUsecase struct {
	UserRepo domain.UserRepository
}

func NewHandshakeUsecase(repo domain.UserRepository) *HandshakeUsecase {
	return &HandshakeUsecase{
		UserRepo: repo,
	}
}

func (u *HandshakeUsecase) Execute(ctx context.Context, cmd commands.HandshakeCommand) (*commands.HandshakeResult, error) {
	curve := ecdh.P521()

	// 🔐 Tạo key pair server
	serverPriv, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	// 📥 Parse client public key
	clientPub, err := crypto.ParsePublicKeyFromBase64(cmd.ClientPublicKey)
	if err != nil {
		return nil, err
	}

	// 🔐 Derive shared secret
	sharedSecret, err := serverPriv.ECDH(clientPub)
	if err != nil {
		return nil, err
	}

	// 🔐 Sinh AES key, IV, salt
	aesKey := make([]byte, 32) // AES-256
	iv := make([]byte, aes.BlockSize)
	salt := make([]byte, 16)

	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	// 🧱 Chuẩn bị session data để mã hóa
	sessionInfo := crypto.SessionInfo{
		AESKey:     aesKey,
		IV:         iv,
		Salt:       salt,
		Expiration: time.Now().Add(15 * time.Minute),
	}

	// 🔐 Mã hóa sessionInfo bằng sharedSecret
	encryptedData, err := crypto.EncryptSessionInfo(sharedSecret, sessionInfo)
	if err != nil {
		return nil, err
	}

	// 🧾 Tạo session ID (simple UUID string)
	sessionID := crypto.GenerateSessionID()

	// 🧬 Encode public key
	serverPubKey := serverPriv.PublicKey()
	serverPubKeyBytes := serverPubKey.Bytes()
	serverPubKeyBase64 := base64.StdEncoding.EncodeToString(serverPubKeyBytes)

	return &commands.HandshakeResult{
		ServerPublicKey:      serverPubKeyBase64,
		EncryptedSessionData: base64.StdEncoding.EncodeToString(encryptedData),
		SessionID:            sessionID,
	}, nil
}
