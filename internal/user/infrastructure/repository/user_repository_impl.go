package repository

import (
	"context"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/entity"
	"github.com/take0fit/ddd-clean-temp/internal/user/domain/repository"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/db"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *db.DB
}

func NewUserRepository(db *db.DB) repository.UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(db.TxKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return r.db.Session(&gorm.Session{})
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	result := r.conn(ctx).Find(&users)
	return users, result.Error
}
