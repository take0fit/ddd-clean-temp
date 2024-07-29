package repository

import (
	"context"
	"errors"
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/entity"
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/repository"
	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/valueobject"
	"github.com/take0fit/ddd-clean-temp/internal/auth/infrastructure/db"
	"github.com/take0fit/ddd-clean-temp/internal/auth/infrastructure/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *db.DB
}

var _ repository.AuthRepository = (*AuthRepositoryImpl)(nil)

func NewAuthRepository(db *db.DB) repository.AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (r *AuthRepositoryImpl) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(db.TxKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return r.db.Session(&gorm.Session{})
}

func (r *AuthRepositoryImpl) FindByEmail(ctx context.Context, email valueobject.Email) (*entity.AuthInfo, error) {
	var authUser model.AuthUser
	if err := r.conn(ctx).Where("email = ?", email).First(&authUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	authInfo := &entity.AuthInfo{
		UserID:       int(authUser.UserID),
		Email:        email,
		PasswordHash: authUser.PasswordHash,
	}
	return authInfo, nil
}

func (r *AuthRepositoryImpl) SaveAuthInfo(ctx context.Context, authInfo *entity.AuthInfo) error {
	authUser := &model.AuthUser{
		UserID:       uint(authInfo.UserID),
		Email:        string(authInfo.Email),
		PasswordHash: authInfo.PasswordHash,
	}

	return r.conn(ctx).Save(authUser).Error
}
