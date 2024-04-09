package http

import (
	"github.com/gin-gonic/gin"
	userInjection "github.com/take0fit/ddd-clean-temp/internal/user/injection"
)

func RegisterHandlers(e *gin.Engine) {
	root := e.Group("/api/v1")

	{
		RegisterUserHandlers(root)
	}
}

func RegisterUserHandlers(router *gin.RouterGroup) {
	user := userInjection.InitializeUserController()

	users := router.Group("/users")
	{
		users.GET("", user.GetUsers)
	}
}
