package services

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/adisnuhic/go-clean/pkg/log"
)

// IUserService represents the user service contract
type IUserService interface {
	GetByID(id uint64) (*models.User, *apperror.AppError)
	GetByEmail(email string) (*models.User, *apperror.AppError)
}

type userService struct {
	Logger log.ILogger
	Repo   repositories.IUserRepository
}

// NewUserService -
func NewUserService(logger log.ILogger, repo repositories.IUserRepository) IUserService {
	return userService{
		Logger: logger,
		Repo:   repo,
	}
}

// GetByID returns user for provided ID
func (svc userService) GetByID(id uint64) (*models.User, *apperror.AppError) {
	return svc.Repo.GetByID(id)
}

// GetByEmail returns user for provided email
func (svc userService) GetByEmail(email string) (*models.User, *apperror.AppError) {
	return svc.Repo.GetByEmail(email)
}
