package repositories

import "github.com/adisnuhic/go-clean/internal/models"

// IUserRepository represents the user's repository contract
type IUserRepository interface {
	GetByID(id uint64) (*models.User, error)
}
