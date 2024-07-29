package usecase

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/dto"
)

type UserUsecase interface {
	GetAllUsers(ctx context.Context) (dto.OutputUsers, error)
}
