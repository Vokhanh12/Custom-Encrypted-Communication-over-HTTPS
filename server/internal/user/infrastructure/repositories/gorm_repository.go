package repositories

import (
	"context"
	"myapp/internal/user/domain"
)

type GormUserRepository struct{}

func NewGormUserRepository() *GormUserRepository {
	return &GormUserRepository{}
}

func (r *GormUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if email == "test@example.com" {
		return &domain.User{
			ID:       "1",
			Email:    email,
			Password: domain.HashPassword("123456"),
		}, nil
	}
	return nil, nil
}
