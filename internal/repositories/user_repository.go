package repositories

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// IUserRepository represents the user repository contract
type IUserRepository interface {
	GetByID(id uint64) (*models.User, error)
	GetByEmail(email string) (*models.User, *apperror.AppError)
}
