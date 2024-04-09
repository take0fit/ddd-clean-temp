package response

import "github.com/take0fit/ddd-clean-temp/internal/user/application/dto"

type UsersResponse []*UserResponse

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewUsersResponse(users dto.OutputUsers) *UsersResponse {
	var ur UsersResponse

	for _, user := range users {
		ur = append(ur, &UserResponse{
			Id:   user.Id,
			Name: user.Name,
		})
	}

	return &ur
}
