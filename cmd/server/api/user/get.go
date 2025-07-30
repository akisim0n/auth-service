package user

import (
	"context"
	serviceConventer "github.com/akisim0n/auth-service/cmd/server/converter"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *user.GetRequest) (*user.GetResponse, error) {
	newUser, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	retUser := serviceConventer.ToUserFromService(newUser)
	return &user.GetResponse{
		User: retUser,
	}, err
}
