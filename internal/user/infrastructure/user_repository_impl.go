package infrastructure

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/common/db"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/entity"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/repository"
)

type UserRepositoryImpl struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) repository.UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	result := r.db.Find(&users)
	return users, result.Error
}
