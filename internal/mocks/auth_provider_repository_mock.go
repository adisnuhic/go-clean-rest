package mocks

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/stretchr/testify/mock"
)

type AuthProviderRepositoryMock struct {
	mock.Mock
}

// GetByUserID mocks the GetByUserID method of the IAuthProviderRepository interface
func (m *AuthProviderRepositoryMock) GetByUserID(id uint64) (*models.AuthProvider, *apperror.AppError) {
	args := m.Called(id)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	// type assert
	return model.(*models.AuthProvider), nil
}

// GetByUserIDProviderID mocks the GetByUserIDProviderID method of the IAuthProviderRepository interface
func (m *AuthProviderRepositoryMock) GetByUserIDProviderID(userID uint64, provider string) (*models.AuthProvider, *apperror.AppError) {
	args := m.Called(userID, provider)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	return model.(*models.AuthProvider), nil
}

// Update mocks the Update method of the IAuthProviderRepository interface
func (m *AuthProviderRepositoryMock) Update(auth *models.AuthProvider) *apperror.AppError {
	args := m.Called(auth)

	err := args.Get(0)

	if err != nil {
		return err.(*apperror.AppError)
	}

	return nil
}

// Create mocks the Create method of the IAuthProviderRepository interface
func (m *AuthProviderRepositoryMock) Create(auth *models.AuthProvider) *apperror.AppError {
	args := m.Called(auth)

	err := args.Get(0)

	if err != nil {
		return err.(*apperror.AppError)
	}

	return nil
}
