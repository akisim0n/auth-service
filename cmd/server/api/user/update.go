package user

import (
	"context"
	serviceConventer "github.com/akisim0n/auth-service/cmd/server/converter"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *user.UpdateRequest) (*emptypb.Empty, error) {

	if err := i.userService.Update(ctx, req.GetId(), serviceConventer.ToServiceFromUserData(req.GetData())); err != nil {
		return nil, err
	}
	return nil, nil
}
