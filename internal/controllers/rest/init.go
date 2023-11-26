package rest

import (
	"github.com/adisnuhic/go-clean/internal/services"
	"github.com/golobby/container/pkg/container"
)

var (
	accountSvc      services.IAccountService
	authProviderSvc services.IAuthProviderService
	authSvc         services.IAuthService
	tokenSvc        services.ITokenService
	userSvc         services.IUserService
)

// Bind controllers to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&accountSvc)
	c.Make(&authProviderSvc)
	c.Make(&authSvc)
	c.Make(&tokenSvc)
	c.Make(&userSvc)

	// Bind account controller
	c.Singleton(func() IAccountController {
		return NewAccountController(accountSvc)
	})

	// Bind user controller
	c.Singleton(func() IUserController {
		return NewUserController(userSvc)
	})

	// Bind user controller
	c.Singleton(func() IHealthController {
		return NewHealthController()
	})
}
