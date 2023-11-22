package repositories

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

type authProviderRepository struct {
	Store mysql.Store
}

// NewAuthProviderRepository -
func NewAuthProviderRepository(s mysql.Store) repositories.IAuthProviderRepository {
	return &authProviderRepository{
		Store: s,
	}
}

// GetByUserID return auth provider for provided ID
func (repo authProviderRepository) GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError) {
	model := &models.AuthProvider{}

	if err := repo.Store.Where("user_id = ?", id).Find(model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchAuthCode, err.Error(), ecode.ErrUnableToFetchAuthMsg)
	}
	return model, nil
}

// GetByUserIDProviderID return auth provider for provided userID and provider
func (repo authProviderRepository) GetByUserIDProviderID(userID uint64, provider string) (*models.AuthProvider, *apperror.AppError) {
	model := &models.AuthProvider{}

	if err := repo.Store.Where("user_id = ?", userID).Where("provider = ?", provider).Find(model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchAuthCode, err.Error(), ecode.ErrUnableToFetchAuthMsg)
	}
	return model, nil
}

// Update updates auth data
func (repo authProviderRepository) Update(auth *models.AuthProvider) *apperror.AppError {
	tx := repo.Store.Exec("UPDATE auth_providers SET uid = ? WHERE user_id = ? AND provider = ?", auth.UID, auth.UserID, auth.Provider)

	if tx.Error != nil {
		return apperror.New(ecode.ErrUnableToSaveAuthCode, tx.Error.Error(), ecode.ErrUnableToSaveAuthMsg)
	}

	return nil
}

// Create creates auth data
func (repo authProviderRepository) Create(auth *models.AuthProvider) *apperror.AppError {
	if err := repo.Store.Create(&auth).Error; err != nil {
		return apperror.New(ecode.ErrUnableToCreateAuthCode, err.Error(), ecode.ErrUnableToCreateAuthMsg)
	}

	return nil
}
