package service

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/entity"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/repository"
)

type userService struct {
	repo repository.UserRepository
	tx   Transaction
}

func NewUserService(
	repo repository.UserRepository,
	tx Transaction,
) UserService {
	return &userService{
		repo: repo,
		tx:   tx,
	}
}

func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return s.repo.GetAll(ctx)
}
