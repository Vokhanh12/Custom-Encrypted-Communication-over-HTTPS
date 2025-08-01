package usecases

import (
	"context"
	"errors"
	d "myapp/internal/user/application/dtos"
	"myapp/internal/user/domain"
)

type HandshakeUsecase struct {
	Repo domain.UserRepository
}

func NewHandshakeUsecase(repo domain.UserRepository) *HandshakeUsecase {
	return &HandshakeUsecase{Repo: repo}
}

func (u *HandshakeUsecase) Execute(ctx context.Context, dto d.LoginRequestDTO) (*d.LoginResponseDTO, error) {
	user, err := u.Repo.FindByEmail(ctx, dto.Email)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	if !user.CheckPassword(dto.Password) {
		return nil, errors.New("invalid password")
	}

	return &d.LoginResponseDTO{Token: "mock_token_for_" + user.ID}, nil
}
