package services

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
)

// IUserService represents the user's service/usecase contract
type IUserService interface {
	GetByID(id uint64) (*models.User, error)
}

type userService struct {
	Repo repositories.IUserRepository
}

// NewUserService -
func NewUserService(repo repositories.IUserRepository) IUserService {
	return userService{
		Repo: repo,
	}
}

// GetByID returns user for provided ID
func (svc userService) GetByID(id uint64) (*models.User, error) {
	return svc.Repo.GetByID(id)
}
