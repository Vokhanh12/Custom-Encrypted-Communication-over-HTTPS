package grpc

import (
	"context"
	pb "myapp/api/user"
	"myapp/internal/user/application"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	Usecase *application.LoginUserUsecase
}

func NewUserHandler(u *application.LoginUserUsecase) *UserHandler {
	return &UserHandler{Usecase: u}
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	dto := application.LoginRequestDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	res, err := h.Usecase.Execute(ctx, dto)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{Token: res.Token}, nil
}
