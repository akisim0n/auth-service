package converter

import (
	servModel "github.com/akisim0n/auth-service/cmd/server/models"
	"github.com/akisim0n/auth-service/cmd/server/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toUserFromService(user *servModel.User) *user_v1.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &user_v1.User{
		Id:        user.Id,
		Data:      toUserDataFromService(user.Data),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func toUserDataFromService(data servModel.UserData) *user_v1.UserData {
	return &user_v1.UserData{
		Name:    data.Name,
		Surname: data.Surname,
		Email:   data.Email,
		Age:     data.Age,
		Role:    toRole(data.Role),
	}
}

func toRole(role servModel.UserRole) user_v1.Role {
	return user_v1.Role(role)
}
