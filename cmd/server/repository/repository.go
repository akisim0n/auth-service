package repository

import (
	"context"
	"github.com/akisim0n/auth-service/cmd/server/models"
)

type UserRepository interface {
	Create(ctx context.Context, data *models.UserData) (int64, error)
	Get(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, data *models.User) error
	Delete(ctx context.Context, id int64) error
}
