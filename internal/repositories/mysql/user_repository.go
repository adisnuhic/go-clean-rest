package repositories

import (
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
)

type mySqlUserRepository struct {
	DB mysql.Store
}

// NewMySQLUserRepository -
func NewMySQLUserRepository(db mysql.Store) repositories.IUserRepository {
	return mySqlUserRepository{
		DB: db,
	}
}

// GetByID returns User for provided ID
func (repo mySqlUserRepository) GetByID(id uint64) (*models.User, error) {
	u := &models.User{
		ID:          1,
		FirstName:   "Adis",
		LastName:    "Nuhic",
		Email:       "adis@test.com",
		IsConfirmed: true,
		AcceptedTos: true,
	}

	return u, nil
}
