package grpc

import (
	"context"

	userv1 "myapp/api/user/v1"
	"myapp/internal/user/application/dtos"
	"myapp/internal/user/application/mappers"
	"myapp/internal/user/application/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userv1.UnimplementedUserServiceServer
	loginUsecase *usecases.LoginUserUsecase
}

// Handshake implements userv1.UserServiceServer.
func (h *UserHandler) Handshake(context.Context, *userv1.HandshakeRequest) (*userv1.HandshakeResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedUserServiceServer implements userv1.UserServiceServer.
func (h *UserHandler) mustEmbedUnimplementedUserServiceServer() {
	panic("unimplemented")
}

func NewUserHandler(loginUsecase *usecases.LoginUserUsecase) *UserHandler {
	return &UserHandler{loginUsecase: loginUsecase}
}

func (h *UserHandler) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
	dto := dtos.LoginRequestDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	cmd := mappers.MapLoginRequestToCommand(dto)

	result, err := h.loginUsecase.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	resDTO := mappers.MapLoginResultToResponseDTO(result)
	return &userv1.LoginResponse{
		AccessToken:  resDTO.AccessToken,
		RefreshToken: resDTO.RefreshToken,
	}, nil
}
