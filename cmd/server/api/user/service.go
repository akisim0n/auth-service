package user

import (
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"github.com/akisim0n/auth-service/cmd/server/service"
)

type Implementation struct {
	user.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
