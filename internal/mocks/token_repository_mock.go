package mocks

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/stretchr/testify/mock"
)

type TokenRepositoryMock struct {
	mock.Mock
}

// CreateToken mocks the CreateToken method of the ITokenRepository interface
func (m *TokenRepositoryMock) CreateToken(token *models.Token) (*models.Token, *apperror.AppError) {
	args := m.Called(token)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	return model.(*models.Token), nil
}

// GetByToken mocks the GetByToken method of the ITokenRepository interface
func (m *TokenRepositoryMock) GetByToken(token string) (*models.Token, *apperror.AppError) {
	args := m.Called(token)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	return model.(*models.Token), nil
}

// Update mocks the Update method of the ITokenRepository interface
func (m *TokenRepositoryMock) Update(token *models.Token) *apperror.AppError {
	args := m.Called(token)

	err := args.Get(0)

	if err != nil {
		return err.(*apperror.AppError)
	}

	return nil
}
