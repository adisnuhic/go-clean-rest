package repositories

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// IAccountRepository represents the account repository contract
type IAccountRepository interface {
	Register(user *models.User) (*models.User, *apperror.AppError)
}
