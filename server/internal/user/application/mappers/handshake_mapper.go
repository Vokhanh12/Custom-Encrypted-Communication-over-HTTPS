package mappers

import (
	"myapp/internal/user/application/commands"
	"myapp/internal/user/application/dtos"
)

func MapHandshakeRequestToCommand(dto dtos.HandshakeRequestDTO) commands.HandshakeCommand {
	return commands.HandshakeCommand{
		ClientPublicKey: dto.ClientPublicKey,
	}
}

func MapHandshakeResultToResponseDTO(result *commands.HandshakeResult) dtos.HandshakeResponseDTO {
	return dtos.HandshakeResponseDTO{
		ServerPublicKey:      result.ServerPublicKey,
		EncryptedSessionData: result.EncryptedSessionData,
		SessionID:            result.SessionID,
	}
}
