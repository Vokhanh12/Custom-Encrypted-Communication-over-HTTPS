package dtos

type HandshakeRequestDTO struct {
	ClientPublicKey string
}

type HandshakeResponseDTO struct {
	ServerPublicKey      string
	EncryptedSessionData string
	SessionID            string
}
