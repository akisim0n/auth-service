package main

import (
	"context"
	"fmt"
	"github.com/akisim0n/auth-service/cmd/server/database"
	"github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"github.com/akisim0n/auth-service/cmd/server/repository/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

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

	repo := user.NewUserRepository(dbPool)

	server := grpc.NewServer()
	reflection.Register(server)
	user_v1.RegisterUserV1Server(server, repo)

	log.Printf("server listening at %v", lis.Addr())

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
