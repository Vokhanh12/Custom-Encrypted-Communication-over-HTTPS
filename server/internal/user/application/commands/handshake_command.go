package commands

type HandshakeCommand struct {
	ClientPublicKey string
}

type HandshakeResult struct {
	ServerPublicKey      string
	EncryptedSessionData string
	SessionID            string
}
