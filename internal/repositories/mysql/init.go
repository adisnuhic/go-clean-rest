package repositories

import (
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/golobby/container/pkg/container"
)

// Bind repositories to IoC (dependency injection) container
func Init(c container.Container) {

	// Bind user repository
	c.Singleton(func() repositories.IUserRepository {
		return NewMySQLUserRepository(mysql.Connection())
	})

	// Bind token repository
	c.Singleton(func() repositories.ITokenRepository {
		return NewMySQLTokenRepository(mysql.Connection())
	})

	// Bind auth provider repository
	c.Singleton(func() repositories.IAuthProviderRepository {
		return NewMySQLAuthProviderRepository(mysql.Connection())
	})

	// Bind account repository
	c.Singleton(func() repositories.IAccountRepository {
		return NewMySQLAccountRepository(mysql.Connection())
	})

}
