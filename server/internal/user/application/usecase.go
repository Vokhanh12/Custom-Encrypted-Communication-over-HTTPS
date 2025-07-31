package application

import (
	"context"
	"errors"
	"myapp/internal/user/domain"
)

type LoginUserUsecase struct {
	Repo domain.UserRepository
}

func NewLoginUserUsecase(repo domain.UserRepository) *LoginUserUsecase {
	return &LoginUserUsecase{Repo: repo}
}

func (u *LoginUserUsecase) Execute(ctx context.Context, dto LoginRequestDTO) (*LoginResponseDTO, error) {
	user, err := u.Repo.FindByEmail(ctx, dto.Email)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	if !user.CheckPassword(dto.Password) {
		return nil, errors.New("invalid password")
	}

	return &LoginResponseDTO{Token: "mock_token_for_" + user.ID}, nil
}
