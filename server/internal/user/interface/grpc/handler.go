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
	loginUsecase     *usecases.LoginUserUsecase
	handshakeUsecase *usecases.HandshakeUsecase
}

func NewUserHandler(
	loginUsecase *usecases.LoginUserUsecase,
	handshakeUsecase *usecases.HandshakeUsecase,
) *UserHandler {
	return &UserHandler{
		loginUsecase:     loginUsecase,
		handshakeUsecase: handshakeUsecase,
	}
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

func (h *UserHandler) Handshake(ctx context.Context, req *userv1.HandshakeRequest) (*userv1.HandshakeResponse, error) {
	dto := dtos.HandshakeRequestDTO{
		ClientPublicKey: req.ClientPublicKey,
	}

	cmd := mappers.MapHandshakeRequestToCommand(dto)

	result, err := h.handshakeUsecase.Execute(ctx, cmd)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resDTO := mappers.MapHandshakeResultToResponseDTO(result)
	return &userv1.HandshakeResponse{
		ServerPublicKey:      resDTO.ServerPublicKey,
		EncryptedSessionData: resDTO.EncryptedSessionData,
		SessionId:            resDTO.SessionID,
	}, nil
}
