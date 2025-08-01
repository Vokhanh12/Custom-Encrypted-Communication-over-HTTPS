//go:build wireinject

package user

import (
	"myapp/internal/user/application/commands"
	"myapp/internal/user/domain"
	"myapp/internal/user/infrastructure/repositories"
	"myapp/internal/user/interface/grpc"

	"github.com/google/wire"
)

func InitializeUserHandlers() *grpc.UserHandler {
	wire.Build(
		repositories.NewGormUserRepository,
		wire.Bind(new(domain.UserRepository), new(*repositories.GormUserRepository)),

		commands.NewLoginUserHandler,

		grpc.NewUserHandler,
	)
	return nil
}
