package repositories

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// IAuthProviderRepository represents the auth provider's repository contract
type IAuthProviderRepository interface {
	GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError)
	GetByUserIDProviderID(userID uint64, provider string) (*models.AuthProvider, *apperror.AppError)
	Update(auth *models.AuthProvider) *apperror.AppError
	Create(auth *models.AuthProvider) *apperror.AppError
}
