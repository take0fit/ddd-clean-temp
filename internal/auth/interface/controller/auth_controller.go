package controller

import (
	"github.com/take0fit/ddd-clean-temp/internal/auth/application/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/internal/auth/application/usecase"
)

type AuthController struct {
	au usecase.AuthUsecase
}

func NewAuthController(au usecase.AuthUsecase) *AuthController {
	return &AuthController{au: au}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthController) LoginHandler(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	inputLoginDTO, err := dto.NewInputLoginDTO(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId, err := c.au.Login(ctx, inputLoginDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_id": userId})
}
