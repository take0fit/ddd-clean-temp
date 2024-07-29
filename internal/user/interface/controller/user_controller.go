package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/usecase"
	"github.com/take0fit/ddd-clean-temp/internal/user/interface/presenter/response"
	"net/http"
)

type UserController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) *UserController {
	return &UserController{uu: uu}
}

func (ctrl *UserController) GetUsers(ctx *gin.Context) {
	users, err := ctrl.uu.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usersResponse := response.NewUsersResponse(users)

	ctx.JSON(http.StatusOK, usersResponse)
}
