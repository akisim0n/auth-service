package user

import (
	"context"
	"github.com/akisim0n/auth-service/cmd/server/models"
)

func (s *serv) Create(ctx context.Context, data *models.UserData) (int64, error) {

	id, err := s.userRepo.Create(ctx, data)
	if err != nil {
		return 0, err
	}

	return id, nil
}
