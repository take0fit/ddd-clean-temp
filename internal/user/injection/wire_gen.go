// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/take0fit/ddd-clean-temp/internal/user/application/usecase"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/db"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/repository"
	"github.com/take0fit/ddd-clean-temp/internal/user/interface/controller"
)

// Injectors from wire.go:

func InitializeUserController() *controller.UserController {
	dbDB := db.GetDB()
	userRepository := repository.NewUserRepository(dbDB)
	userUsecase := usecase.NewUserUsecase(userRepository, dbDB)
	userController := controller.NewUserController(userUsecase)
	return userController
}
