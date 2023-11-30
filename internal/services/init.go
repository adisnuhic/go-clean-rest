package services

import (
	"github.com/adisnuhic/go-clean/config"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/golobby/container/pkg/container"
)

var (
	tokenSvc ITokenService
	authSvc  IAuthService

	accountRepo      repositories.IAccountRepository
	userRepo         repositories.IUserRepository
	authProviderRepo repositories.IAuthProviderRepository
	tokenRepo        repositories.ITokenRepository
)

// Bind services to IoC (dependency injection) container
func Init(c container.Container, logger log.ILogger) {

	// Resolve dependencies and return concrete type of given abstractions (for repos)
	c.Make(&accountRepo)
	c.Make(&userRepo)
	c.Make(&authProviderRepo)
	c.Make(&tokenRepo)

	// Bind token service
	c.Singleton(func() ITokenService {
		return NewTokenService(logger, tokenRepo)
	})

	// Bind auth service
	c.Singleton(func() IAuthService {
		return NewAuthService(logger, config.Load().JWTConf.Secret)
	})

	// Bind user service
	c.Singleton(func() IUserService {
		return NewUserService(logger, userRepo)
	})

	// Bind auth provider service
	c.Singleton(func() IAuthProviderService {
		return NewAuthProviderService(logger, config.Load().JWTConf.Secret, authProviderRepo)
	})

	// Resolve dependencies and return concrete type of given abstractions (for services)
	c.Make(&tokenSvc)
	c.Make(&authSvc)

	// Bind account service
	c.Singleton(func() IAccountService {
		return NewAccountService(logger, tokenSvc, authSvc, accountRepo, userRepo, authProviderRepo, tokenRepo)
	})

}
