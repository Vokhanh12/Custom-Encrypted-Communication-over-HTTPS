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

func (h *LoginUserHandler) Handle(ctx context.Context, cmd LoginUserCommand) (*ds.LoginResponseDTO, error) {
	user, err := h.UserRepo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if !h.Hasher.Compare(cmd.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := h.TokenGen.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &ds.LoginResponseDTO{Token: token}, nil
}
