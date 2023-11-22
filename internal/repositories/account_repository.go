package repositories

import "github.com/adisnuhic/go-clean/internal/models"

// IAccountRepository represents the account's repository contract
type IAccountRepository interface {
	Login(email string, password string) (*models.User, error)
}
