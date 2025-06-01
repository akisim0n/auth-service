package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type UserRepository struct {
	user.UnimplementedUserV1Server
	DBPool *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DBPool: db,
	}
}

func (r *UserRepository) Get(ctx context.Context, request *user.GetRequest) (*user.GetResponse, error) {

	selectBuilder := sq.Select("name", "email", "role", "created_at", "updated_at").
		PlaceholderFormat(sq.Dollar).
		From("users").
		Where(sq.Eq{"id": request.Id})

	var retData *user.GetResponse

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while building query:", err))
	}
	if err = r.DBPool.QueryRow(ctx, query, args...).Scan(&retData); err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return retData, nil
}

func (r *UserRepository) Create(ctx context.Context, request *user.CreateRequest) (*user.CreateResponse, error) {

	log.Println("Trying to create user with role - ", request.GetRole())

	insertBuilder := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "role", "created_at").
		Values(request.GetName(), request.GetEmail(), request.GetRole(), time.Now()).
		Suffix("RETURNING id")

	var retData *user.CreateResponse
	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while building query:", err))
	}

	err = r.DBPool.QueryRow(ctx, query, args...).Scan(&retData)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return retData, nil
}
