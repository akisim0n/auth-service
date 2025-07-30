package user

import (
	"context"
	serviceConventer "github.com/akisim0n/auth-service/cmd/server/converter"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *user.CreateRequest) (*user.CreateResponse, error) {

	userId, err := i.userService.Create(ctx, serviceConventer.ToServiceFromUserData(req.GetData()))
	if err != nil {
		return nil, err
	}

	return &user.CreateResponse{
		Id: userId,
	}, nil
}
