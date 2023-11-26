package services

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// IAuthProviderService represents auth provider service contract
type IAuthProviderService interface {
	GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError)
	GetByUserIDProviderID(userID uint64, provider string) (*models.AuthProvider, *apperror.AppError)
	Create(auth *models.AuthProvider) *apperror.AppError
	Update(auth *models.AuthProvider) *apperror.AppError
}

type authProviderService struct {
	JWTSecret  string
	Repository repositories.IAuthProviderRepository
}

// NewAuthProviderService -
func NewAuthProviderService(secret string, repo repositories.IAuthProviderRepository) IAuthProviderService {
	return &authProviderService{
		JWTSecret:  secret,
		Repository: repo,
	}
}

// GetByUserID return auth provider for provided userID
func (svc authProviderService) GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError) {
	return svc.Repository.GetByUserID(id)
}

// GetByUserIDProviderID return auth provider for provided userID and provider
func (svc authProviderService) GetByUserIDProviderID(userID uint64, provider string) (*models.AuthProvider, *apperror.AppError) {
	return svc.Repository.GetByUserIDProviderID(userID, provider)
}

// Creates auth data
func (svc authProviderService) Create(auth *models.AuthProvider) *apperror.AppError {
	return svc.Repository.Create(auth)
}

// Creates auth data
func (svc authProviderService) Update(auth *models.AuthProvider) *apperror.AppError {
	return svc.Repository.Update(auth)
}
