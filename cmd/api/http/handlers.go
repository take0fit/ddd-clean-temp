package http

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/take0fit/ddd-clean-temp/cmd/api/middleware"
	authInjection "github.com/take0fit/ddd-clean-temp/internal/auth/injection"
	userInjection "github.com/take0fit/ddd-clean-temp/internal/user/injection"
)

func RegisterHandlers(e *gin.Engine, publicKey *rsa.PublicKey) {
	root := e.Group("/api/v1")

	{
		RegisterAuthenticationHandlers(root, publicKey)
		RegisterUserHandlers(root, publicKey)
	}
}

func RegisterAuthenticationHandlers(root *gin.RouterGroup, publicKey *rsa.PublicKey) {
	auth := authInjection.InitializeAuthController()

	session := root.Group("/login").Use()
	{
		session.POST("/", auth.LoginHandler)
	}
}

func RegisterUserHandlers(router *gin.RouterGroup, publicKey *rsa.PublicKey) {
	user := userInjection.InitializeUserController()

	users := router.Group("/users").Use(middleware.AuthMiddleware(publicKey))
	{
		users.GET("", user.GetUsers)
	}
}
