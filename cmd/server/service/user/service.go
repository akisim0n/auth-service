package user

import (
	"github.com/akisim0n/auth-service/cmd/server/repository"
	"github.com/akisim0n/auth-service/cmd/server/service"
)

type serv struct {
	userRepo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) service.UserService {
	return &serv{
		userRepo: userRepo,
	}
}
