package main

import (
	"context"
	"fmt"
	"github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const port = 50051

type userServer struct {
	user_v1.UnimplementedUserV1Server
}

func (server *userServer) Get(ctx context.Context, request *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	log.Printf("Received: %v", request.GetId())

	retInfo := &user_v1.GetResponse{
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
	user_v1.RegisterUserV1Server(server, &userServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
