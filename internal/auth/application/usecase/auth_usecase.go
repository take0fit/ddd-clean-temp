package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/internal/auth/application/dto"
)

type AuthUsecase interface {
	Login(*gin.Context, *dto.InputLoginDTO) (int, error)
}
