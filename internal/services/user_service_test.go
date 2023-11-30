package services

import (
	"errors"
	"testing"
	"time"

	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/mocks"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

func TestGetByID_Success(t *testing.T) {
	// create mock repo
	mockRepo := &mocks.UserRepositoryMock{}

	// expected return values
	expectedUser := &models.User{
		ID:          uint64(1),
		FirstName:   "First name",
		LastName:    "Last name",
		Email:       "test@test.com",
		IsConfirmed: true,
		AcceptedTos: true,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	// Set up expectations on the mock
	mockRepo.On("GetByID", uint64(1)).Return(expectedUser, nil)

	// Create a userService instance with the mocked repository
	userService := NewUserService(nil, mockRepo)

	// Call the GetByID method on the userService
	user, err := userService.GetByID(uint64(1))

	// Assert that the user and error returned by the service match the expectations
	assert.Empty(t, err)
	assert.Equal(t, expectedUser, user)

	// Ensure that the GetByID method of the mock was called with the correct arguments
	mockRepo.AssertCalled(t, "GetByID", uint64(1))
}

func TestGetByID_Error(t *testing.T) {
	// create mock repo
	mockRepo := &mocks.UserRepositoryMock{}

	// expected return values
	err := errors.New("not found")
	expectedError := apperror.New(ecode.ErrUnableToGetUserCode, err.Error(), ecode.ErrUnableToGetUserMsg)

	// Set up expectations on the mock
	mockRepo.On("GetByID", uint64(2)).Return(nil, expectedError)

	// Create a userService instance with the mocked repository
	userService := NewUserService(nil, mockRepo)

	// Call the GetByID method on the userService
	user, userErr := userService.GetByID(uint64(2))

	// Assert that the user and error returned by the service match the expectations
	assert.Empty(t, user)
	assert.Equal(t, expectedError, userErr)

	// Ensure that the GetByID method of the mock was called with the correct arguments
	mockRepo.AssertCalled(t, "GetByID", uint64(2))
}
