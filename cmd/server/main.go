package main

import (
	"context"
	"fmt"
	user "github.com/akisim0n/auth-service/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const port = 50051

type userServer struct {
	user.UnimplementedUserV1Server
}

func (server *userServer) Get(ctx context.Context, request *user.GetRequest) (*user.GetResponse, error) {
	log.Printf("Received: %v", request.GetId())

	retInfo := &user.GetResponse{
		Id:        request.GetId(),
		Name:      "DAK",
		Email:     "D@AK.com",
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}

	return retInfo, nil
}

func main() {
	lis, lesErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if lesErr != nil {
		log.Fatalf("failed to listen: %v", lesErr)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	user.RegisterUserV1Server(server, &userServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
