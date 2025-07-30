package main

import (
	"context"
	"fmt"
	"github.com/akisim0n/auth-service/cmd/server/converter"
	"github.com/akisim0n/auth-service/cmd/server/database"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	userRepo "github.com/akisim0n/auth-service/cmd/server/repository/user"
	"github.com/akisim0n/auth-service/cmd/server/service"
	userServ "github.com/akisim0n/auth-service/cmd/server/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
)

type server struct {
	user.UnimplementedUserV1Server
	userService service.UserService
}

func (s *server) Create(ctx context.Context, req *user.CreateRequest) (*user.CreateResponse, error) {

	userId, err := s.userService.Create(ctx, converter.ToServiceFromUserData(req.GetData()))
	if err != nil {
		return nil, err
	}

	return &user.CreateResponse{
		Id: userId,
	}, nil
}

func (s *server) Update(ctx context.Context, req *user.UpdateRequest) (*emptypb.Empty, error) {

	if err := s.userService.Update(ctx, req.GetId(), converter.ToServiceFromUserData(req.GetData())); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *server) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
	newUser, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	retUser := converter.ToUserFromService(newUser)
	return &user.GetResponse{
		User: retUser,
	}, err
}

func (s *server) Delete(ctx context.Context, req *user.DeleteRequest) (*emptypb.Empty, error) {
	if err := s.userService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return nil, nil
}

func main() {

	ctx := context.Background()

	lis, lesErr := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("SERVER_PORT_OUT")))
	if lesErr != nil {
		log.Fatalf("failed to listen: %v", lesErr)
	}

	dbPool, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if pingErr := dbPool.Ping(ctx); pingErr != nil {
		log.Fatalf("failed to ping database: %v", pingErr)
	}
	defer dbPool.Close()

	repo := userRepo.NewUserRepository(dbPool)
	serv := userServ.NewService(repo)

	s := grpc.NewServer()
	reflection.Register(s)
	user.RegisterUserV1Server(s, &server{userService: serv})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
