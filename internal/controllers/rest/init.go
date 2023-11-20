package rest

import (
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/golobby/container/pkg/container"
)

var (
	userSvc services.IUserService
)

// Bind controllers to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&userSvc)

	// Bind user controller
	c.Singleton(func() IUserController {
		return NewUserController(userSvc)
	})

}
