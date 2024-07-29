//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/usecase"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/db"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/repository"
	"github.com/take0fit/ddd-clean-temp/internal/user/interface/controller"
)

func InitializeUserController() *controller.UserController {
	wire.Build(
		db.GetDB,
		wire.Bind(
			new(usecase.Transaction),
			new(*db.DB),
		),
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		controller.NewUserController,
	)

	return nil
}
