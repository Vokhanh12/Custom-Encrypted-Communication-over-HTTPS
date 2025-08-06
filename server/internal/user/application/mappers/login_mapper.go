package mappers

import (
	"myapp/internal/user/application/commands"
	"myapp/internal/user/application/dtos"
)

func MapLoginRequestToCommand(dto dtos.LoginRequestDTO) commands.LoginCommand {
	return commands.LoginCommand{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func MapLoginResultToResponseDTO(result *commands.LoginResult) dtos.LoginResponseDTO {
	return dtos.LoginResponseDTO{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
	}
}
