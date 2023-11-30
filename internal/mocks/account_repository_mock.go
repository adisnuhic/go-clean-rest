package mocks

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMock struct {
	mock.Mock
}

// Register mocks the Register method of the IAccountRepository interface
func (m *AccountRepositoryMock) Register(user *models.User) (*models.User, *apperror.AppError) {
	args := m.Called(user)

	model := args.Get(0)
	err := args.Get(1)

	if err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateUserCode, err.(error).Error(), ecode.ErrUnableToCreateUserMsg)
	}

	return model.(*models.User), nil
}
