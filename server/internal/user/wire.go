//go:build wireinject
// +build wireinject

package user

import (
	"myapp/internal/user/application/usecases"
	"myapp/internal/user/infrastructure/repositories"
	"myapp/internal/user/interface/grpc"

	"github.com/google/wire"
)

func InitializeUserHandler() (*grpc.UserHandler, error) {
	wire.Build(
		repositories.NewGormRepository,
		usecases.NewLoginUserUsecase,
		usecases.NewHandshakeUsecase,
		grpc.NewUserHandler,
	)
	return &grpc.UserHandler{}, nil
}
