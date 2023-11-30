package mocks

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

// GetByID mocks the GetByID method of the IUserRepository interface
func (m *UserRepositoryMock) GetByID(id uint64) (*models.User, *apperror.AppError) {
	args := m.Called(id)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	return model.(*models.User), nil
}

// GetByEmail mocks the GetByEmail method of the IUserRepository interface
func (m *UserRepositoryMock) GetByEmail(email string) (*models.User, *apperror.AppError) {
	args := m.Called(email)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, err.(*apperror.AppError)
	}

	return model.(*models.User), nil
}
