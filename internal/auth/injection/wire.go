//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/take0fit/ddd-clean-temp/internal/auth/application/usecase"
	"github.com/take0fit/ddd-clean-temp/internal/auth/infrastructure/db"
	"github.com/take0fit/ddd-clean-temp/internal/auth/infrastructure/repository"
	"github.com/take0fit/ddd-clean-temp/internal/auth/interface/controller"
)

func InitializeAuthController() *controller.AuthController {
	wire.Build(
		db.GetDB,
		wire.Bind(
			new(usecase.Transaction),
			new(*db.DB),
		),
		repository.NewAuthRepository,
		usecase.NewAuthUsecase,
		controller.NewAuthController,
	)

	return nil
}
