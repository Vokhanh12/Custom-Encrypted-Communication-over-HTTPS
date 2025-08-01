package commands

import (
	"context"
	"errors"
	ds "myapp/internal/user/application/dtos"
	"myapp/internal/user/domain"
)

type LoginUserCommand struct {
	Email    string
	Password string
}

type LoginUserHandler struct {
	UserRepo domain.UserRepository
}

func NewLoginUserHandler(repo domain.UserRepository) *LoginUserHandler {
	return &LoginUserHandler{UserRepo: repo}
}

func (h *LoginUserHandler) Handle(ctx context.Context, cmd LoginUserCommand) (*ds.LoginResponseDTO, error) {
	user, err := h.UserRepo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	return &ds.LoginResponseDTO{Token: "TOKEN"}, nil
}
