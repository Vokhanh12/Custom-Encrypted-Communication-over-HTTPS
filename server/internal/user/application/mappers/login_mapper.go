package mappers

import (
	"myapp/internal/user/application/commands"
	"myapp/internal/user/application/dtos"
)

func MapLoginDTOToCommand(dto dtos.LoginRequestDTO) commands.LoginUserCommand {
	return commands.LoginUserCommand{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
