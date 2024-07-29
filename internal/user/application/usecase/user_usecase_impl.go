package usecase

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/dto"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/repository"
)

type UserUsecaseImpl struct {
	repo repository.UserRepository
	tx   Transaction
}

func NewUserUsecase(
	repo repository.UserRepository,
	tx Transaction,
) UserUsecase {
	return &UserUsecaseImpl{
		repo: repo,
		tx:   tx,
	}
}

func (s *UserUsecaseImpl) GetAllUsers(ctx context.Context) (dto.OutputUsers, error) {
	userEntityList, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputUsers(userEntityList)
}
