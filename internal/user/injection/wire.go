//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/google/wire"
	"github.com/take0fit/ddd-clean-temp/internal/common/db"
	"github.com/take0fit/ddd-clean-temp/internal/user/application/service"
	"github.com/take0fit/ddd-clean-temp/internal/user/infrastructure"
	"github.com/take0fit/ddd-clean-temp/internal/user/interface/controller"
)

func InitializeUserController() *controller.UserController {
	wire.Build(
		db.GetDB,
		wire.Bind(
			new(service.Transaction),
			new(*db.DB),
		),
		infrastructure.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return nil
}
