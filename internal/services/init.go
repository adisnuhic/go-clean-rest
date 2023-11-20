package services

import (
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/golobby/container/pkg/container"
)

var (
	userRepo repositories.IUserRepository
)

// Bind controllers to IoC (dependency injection) container
func Init(c container.Container) {

	// Resolve dependencies and return concrete type of given abstractions
	c.Make(&userRepo)

	// Bind user controller
	c.Singleton(func() IUserService {
		return NewUserService(userRepo)
	})

}
