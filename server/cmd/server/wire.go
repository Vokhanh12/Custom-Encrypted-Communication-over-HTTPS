//go:build wireinject

package main

import (
	"myapp/internal/user/application"
	"myapp/internal/user/domain"
	"myapp/internal/user/infrastructure/repository"

	"github.com/google/wire"
)

func InitializeLoginUsecase() *application.LoginUserUsecase {
	wire.Build(
		repository.NewGormUserRepository,
		wire.Bind(new(domain.UserRepository), new(*repository.GormUserRepository)),
		application.NewLoginUserUsecase,
	)
	return nil
}
