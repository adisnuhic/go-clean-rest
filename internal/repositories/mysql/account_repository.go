package repositories

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

type accountRepository struct {
	Store mysql.Store
}

func NewMySQLAccountRepository(s mysql.Store) repositories.IAccountRepository {
	return &accountRepository{
		Store: s,
	}
}

// Register user
func (repo accountRepository) Register(user *models.User) (*models.User, *apperror.AppError) {

	if err := repo.Store.Create(&user).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateUserCode, err.Error(), ecode.ErrUnableToCreateUserMsg)
	}

	return user, nil
}
