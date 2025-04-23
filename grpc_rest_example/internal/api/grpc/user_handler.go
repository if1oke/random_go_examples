package grpc

import (
	"context"
	"fmt"
	v1 "grpc_rest/pkg/api/grpc/v1"
)

type UserHandler struct {
	v1.UnimplementedUserServiceServer
}

func NewUserHandler() v1.UserServiceServer {
	return &UserHandler{}
}

func (h *UserHandler) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	fmt.Println("Return example user")
	return &v1.GetResponse{Username: "Igor Fedorov"}, nil
}

func (h *UserHandler) Create(ctx context.Context, req *v1.CreateRequest) (*v1.Empty, error) {
	fmt.Printf("Create user with username: %s\n", req.Username)
	return &v1.Empty{}, nil
}
