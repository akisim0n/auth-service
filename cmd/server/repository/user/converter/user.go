package converter

import (
	servModel "github.com/akisim0n/auth-service/cmd/server/models"
	repoModel "github.com/akisim0n/auth-service/cmd/server/repository/user/models"
)

func FromRepoToUser(user *repoModel.User) *servModel.User {
	return &servModel.User{
		Id:        user.Id,
		Data:      fromRepoToUserData(user.Data),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func fromRepoToUserData(data repoModel.UserData) servModel.UserData {
	return servModel.UserData{
		Name:     data.Name,
		Surname:  data.Surname,
		Email:    data.Email,
		Age:      data.Age,
		Password: data.Password,
		Role:     toRole(data.Role),
	}
}

func toRole(role repoModel.UserRole) servModel.UserRole {
	return servModel.UserRole(role)
}
