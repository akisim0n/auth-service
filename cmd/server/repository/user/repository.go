package user

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/akisim0n/auth-service/cmd/server/models"
	"github.com/akisim0n/auth-service/cmd/server/repository"
	"github.com/akisim0n/auth-service/cmd/server/repository/user/converter"
	repoModel "github.com/akisim0n/auth-service/cmd/server/repository/user/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	tableName = "users"
)

type repo struct {
	DBPool *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{
		DBPool: db,
	}
}

func (r *repo) Get(ctx context.Context, id int64) (*models.User, error) {

	selectBuilder := sq.Select("id", "name", "email", "role_id", "created_at", "coalesce(updated_at, now()) as updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{"u.id": id})

	var newUser repoModel.User

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while building query:", err))
	}
	if err = r.DBPool.QueryRow(ctx, query, args...).
		Scan(
			&newUser.Id,
			&newUser.Data.Name,
			&newUser.Data.Email,
			&newUser.Data.Role,
			&newUser.CreatedAt,
			&newUser.UpdatedAt,
		); err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return converter.FromRepoToUser(&newUser), nil
}

func (r *repo) Create(ctx context.Context, data *models.UserData) (int64, error) {

	/*	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.GetPassword()), bcrypt.DefaultCost)
		switch {
		case errors.Is(err, bcrypt.ErrPasswordTooLong):
			return nil, errors.New("password too long")
		case err != nil:
			return nil, errors.New("error during password hashing")
		}

		if compareErr := bcrypt.CompareHashAndPassword(hashPassword, []byte(request.GetPasswordConfirm())); compareErr != nil {
			switch {
			case errors.Is(compareErr, bcrypt.ErrMismatchedHashAndPassword):
				return nil, errors.New("password incorrect")
			case errors.Is(compareErr, bcrypt.ErrHashTooShort):
				return nil, errors.New("password too short")
			default:
				return nil, errors.New(fmt.Sprint("Password hashing failed"))
			}
		}*/

	insertBuilder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("name", "surname", "email", "age", "role_id", "password").
		Values(data.Name, data.Surname, data.Email, data.Age, data.Role, data.Password).
		Suffix("RETURNING id")

	var retId int64
	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, errors.New(fmt.Sprint("Error while building query:", err))
	}

	err = r.DBPool.QueryRow(ctx, query, args...).Scan(&retId)
	if err != nil {
		return 0, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return retId, nil
}

func (r *repo) Update(ctx context.Context, id int64, data *models.UserData) error {

	updateBuilder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set("name", data.Name).
		Set("surname", data.Surname).
		Set("email", data.Email).
		Set("age", data.Age).
		Set("role_id", data.Role).
		Set("password", data.Password).
		Set("updated_at", timestamppb.Now()).
		Where(sq.Eq{"id": id})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return errors.New(fmt.Sprint("Error while building query:", err))
	}

	_, err = r.DBPool.Exec(ctx, query, args...)
	if err != nil {
		return errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {

	deleteBuilder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return errors.New(fmt.Sprint("Error while building query:", err))
	}

	_, err = r.DBPool.Exec(ctx, query, args...)
	if err != nil {
		return errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return nil
}
