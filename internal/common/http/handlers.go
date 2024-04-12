package http

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	userInjection "github.com/take0fit/ddd-clean-temp/internal/user/injection"
	"github.com/take0fit/ddd-clean-temp/pkg/middleware"
)

func RegisterHandlers(e *gin.Engine, publicKey *rsa.PublicKey) {
	root := e.Group("/api/v1")

	{
		RegisterUserHandlers(root, publicKey)
	}
}

func RegisterUserHandlers(router *gin.RouterGroup, publicKey *rsa.PublicKey) {
	user := userInjection.InitializeUserController()

	users := router.Group("/users").Use(middleware.AuthMiddleware(publicKey))
	{
		users.GET("", user.GetUsers)
	}
}
