package user

import (
	"context"
	"github.com/akisim0n/auth-service/cmd/server/models"
)

func (s *serv) Update(ctx context.Context, id int64, data *models.UserData) error {

	if err := s.userRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
