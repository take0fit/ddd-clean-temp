package dto

import "github.com/take0fit/ddd-clean-temp/internal/user/domain/entity"

type OutputUsers []*OutputUser

type OutputUser struct {
	Id   int
	Name string
}

func NewOutputUsers(users []*entity.User) OutputUsers {
	outputUsers := make([]*OutputUser, len(users))
	for i, user := range users {
		outputUsers[i] = &OutputUser{
			Id:   user.Id,
			Name: user.Name,
		}
	}

	return outputUsers
}
