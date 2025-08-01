package usecases

import (
	"context"
	"myapp/internal/user/application/commands"
	ds "myapp/internal/user/application/dtos"
)

type LoginUserUsecase struct {
	Handler *commands.LoginUserHandler
}

func (u *LoginUserUsecase) Login(ctx context.Context, dto ds.LoginRequestDTO) (*ds.LoginResponseDTO, error) {
	cmd := commands.LoginUserCommand{
		Email:    dto.Email,
		Password: dto.Password,
	}

	return u.Handler.Handle(ctx, cmd)
}

func NewLoginUserUsecase(handler *commands.LoginUserHandler) *LoginUserUsecase {
	return &LoginUserUsecase{
		Handler: handler,
	}
}
