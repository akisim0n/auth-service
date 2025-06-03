package repository

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	user "github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	selectBuilder := sq.Select("u.id", "u.name", "u.email", "u.role_id", "u.created_at", "coalesce(u.updated_at, now())").
		PlaceholderFormat(sq.Dollar).
		From("users u").
		Where(sq.Eq{"u.id": request.GetId()})

	var id int64
	var name, email string
	var createdAt, updatedAt time.Time
	var role user.Role

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while building query:", err))
	}
	if err = r.DBPool.QueryRow(ctx, query, args...).Scan(&id, &name, &email, &role, &createdAt, &updatedAt); err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return &user.GetResponse{
		Id:        id,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: timestamppb.New(createdAt),
		UpdatedAt: timestamppb.New(updatedAt),
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, request *user.CreateRequest) (*user.CreateResponse, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.GetPassword()), bcrypt.DefaultCost)
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
	}

	insertBuilder := sq.Insert("users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "role_id", "password").
		Values(request.GetName(), request.GetEmail(), request.GetRole(), hashPassword).
		Suffix("RETURNING id")

	var retData user.CreateResponse
	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while building query:", err))
	}

	err = r.DBPool.QueryRow(ctx, query, args...).Scan(&retData.Id)
	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return &retData, nil
}

func (r *UserRepository) Update(ctx context.Context, request *user.UpdateRequest) (*emptypb.Empty, error) {

	updateBuilder := sq.Update("users").
		PlaceholderFormat(sq.Dollar).
		Set("name", request.GetName()).
		Set("email", request.GetEmail()).
		Where(sq.Eq{"id": request.GetId()})

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, errors.New(fmt.Sprint("Error while building query:", err))
	}

	_, err = r.DBPool.Exec(ctx, query, args...)
	if err != nil {
		return &emptypb.Empty{}, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return &emptypb.Empty{}, nil
}

func (r *UserRepository) Delete(ctx context.Context, request *user.DeleteRequest) (*emptypb.Empty, error) {

	deleteBuilder := sq.Delete("users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": request.GetId()})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return &emptypb.Empty{}, errors.New(fmt.Sprint("Error while building query:", err))
	}

	_, err = r.DBPool.Exec(ctx, query, args...)
	if err != nil {
		return &emptypb.Empty{}, errors.New(fmt.Sprint("Error while executing query:", err))
	}

	return &emptypb.Empty{}, nil
}
