package service

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/dto"
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

func (s *userService) GetAllUsers(ctx context.Context) (dto.OutputUsers, error) {
	userEntityList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputUsers(userEntityList), nil
}
