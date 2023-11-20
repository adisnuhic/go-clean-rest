package repositories

import (
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/golobby/container/pkg/container"
)

// Bind controllers to IoC (dependency injection) container
func Init(c container.Container) {

	// Bind user controller
	c.Singleton(func() repositories.IUserRepository {
		return NewMySQLUserRepository(mysql.Connection())
	})

}
