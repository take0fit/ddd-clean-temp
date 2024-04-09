package repository

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/beautician/domain/entity"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
}
