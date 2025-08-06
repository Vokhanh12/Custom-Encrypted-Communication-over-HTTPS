// infrastructure/repositories/gorm_repository.go
package repositories

import (
	"context"

	"myapp/internal/user/domain"
)

type GormRepository struct {
}

func NewGormRepository() domain.UserRepository {
	return &GormRepository{}
}

func (r *GormRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	// var user domain.User
	// err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	// if err != nil {
	// 	return nil, err
	// }
	return &domain.User{
		ID: "ID", Email: "EMAIL", Password: "PASSWORD",
	}, nil
}
