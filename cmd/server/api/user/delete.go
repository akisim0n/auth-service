package user

import (
	"context"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *user.DeleteRequest) (*emptypb.Empty, error) {
	if err := i.userService.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return nil, nil
}
