package service

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/entity"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}
