package usecase

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/internal/auth/application/dto"

	"github.com/take0fit/ddd-clean-temp/internal/auth/domain/repository"
)

type AuthUsecaseImpl struct {
	ar repository.AuthRepository
	tx Transaction
}

func NewAuthUsecase(ar repository.AuthRepository, tx Transaction) AuthUsecase {
	return &AuthUsecaseImpl{
		ar: ar,
		tx: tx,
	}
}

func (s *AuthUsecaseImpl) Login(ctx *gin.Context, inputLoginDTO *dto.InputLoginDTO) (int, error) {
	authInfo, err := s.ar.FindByEmail(ctx, inputLoginDTO.Email)
	if err != nil {
		return 0, err
	}

	if !authInfo.VerifyPassword(inputLoginDTO.Password) {
		return 0, errors.New("invalid password")
	}

	return authInfo.UserID, nil
}
