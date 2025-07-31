package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "myapp/api/user"
	user_grpc "myapp/internal/user/interface/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	loginUsecase := InitializeLoginUsecase()
	userHandler := user_grpc.NewUserHandler(loginUsecase)

	pb.RegisterUserServiceServer(s, userHandler)

	log.Println("gRPC server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
