package user

import (
	"context"
	"github.com/akisim0n/auth-service/cmd/server/models"
)

func (s *serv) Get(ctx context.Context, id int64) (*models.User, error) {
	user, err := s.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
