package grpc

import (
	"context"

	pb "myapp/api/user"
	"myapp/internal/user/application/commands"
	"myapp/internal/user/application/queries"
	"myapp/internal/user/domain/dto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer

	RegisterHandler       *commands.RegisterUserHandler
	LoginHandler          *commands.LoginUserHandler
	GetUserByEmailHandler *queries.GetUserByEmailHandler
}

func NewUserHandler(
	register *commands.RegisterUserHandler,
	login *commands.LoginUserHandler,
	getUserByEmail *queries.GetUserByEmailHandler,
) *UserHandler {
	return &UserHandler{
		RegisterHandler:       register,
		LoginHandler:          login,
		GetUserByEmailHandler: getUserByEmail,
	}
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	dtoReq := dto.RegisterRequestDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	err := h.RegisterHandler.Handle(ctx, dtoReq)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RegisterResponse{Success: true}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	dtoReq := dto.LoginRequestDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	dtoRes, err := h.LoginHandler.Handle(ctx, dtoReq)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &pb.LoginResponse{Token: dtoRes.Token}, nil
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	query := dto.GetUserByEmailQueryDTO{
		Email: req.GetEmail(),
	}

	userDTO, err := h.GetUserByEmailHandler.Handle(ctx, query)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.GetUserByEmailResponse{
		Id:    userDTO.ID,
		Email: userDTO.Email,
		Name:  userDTO.Name,
	}, nil
}
