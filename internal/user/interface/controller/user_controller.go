package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/service"
	"net/http"
)

type UserController struct {
	service service.UserService
}

func NewUserController(s service.UserService) *UserController {
	return &UserController{service: s}
}

func (ctrl *UserController) GetUsers(ctx *gin.Context) {
	users, err := ctrl.service.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
