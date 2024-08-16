//go:build wireinject
package wire

import (
	"github.com/google/wire"
	"github.com/longln/go-ecommerce-backend/internal/controller"
	"github.com/longln/go-ecommerce-backend/internal/repo"
	"github.com/longln/go-ecommerce-backend/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}