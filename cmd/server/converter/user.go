package converter

import (
	"database/sql"
	servModel "github.com/akisim0n/auth-service/cmd/server/models"
	"github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *servModel.User) *user_v1.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &user_v1.User{
		Id:        user.Id,
		Data:      ToUserDataFromService(user.Data),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserDataFromService(data servModel.UserData) *user_v1.UserData {
	return &user_v1.UserData{
		Name:    data.Name,
		Surname: data.Surname,
		Email:   data.Email,
		Age:     data.Age,
		Role:    ToRoleFromService(data.Role),
	}
}

func ToRoleFromService(role servModel.UserRole) user_v1.Role {
	return user_v1.Role(role)
}

func ToServiceFromUser(user *user_v1.User) *servModel.User {

	var updatedAt sql.NullTime

	if user.UpdatedAt != nil {
		updatedAt = sql.NullTime{
			Time:  user.UpdatedAt.AsTime(),
			Valid: true,
		}
	}

	return &servModel.User{
		Id:        user.Id,
		Data:      *ToServiceFromUserData(user.Data),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: updatedAt,
	}
}

func ToServiceFromUserData(user *user_v1.UserData) *servModel.UserData {
	return &servModel.UserData{
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Age:     user.Age,
		Role:    ToServiceFromRole(user.Role),
	}
}

func ToServiceFromRole(role user_v1.Role) servModel.UserRole {
	return servModel.UserRole(role)
}
