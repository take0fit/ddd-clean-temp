package repository

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/entity"
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/valueobject"
)

type AuthRepository interface {
	FindByEmail(context.Context, valueobject.Email) (*entity.AuthInfo, error)
}
