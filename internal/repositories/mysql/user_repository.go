package repositories

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

type mySqlUserRepository struct {
	Store mysql.Store
}

// NewMySQLUserRepository -
func NewMySQLUserRepository(s mysql.Store) repositories.IUserRepository {
	return mySqlUserRepository{
		Store: s,
	}
}

// GetByID returns User for provided ID
func (repo mySqlUserRepository) GetByID(id uint64) (*models.User, error) {
	u := &models.User{}

	tx := repo.Store.Where("id = ?", id).Find(u)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return u, nil
}

// GetByEmail returns user for provided email
func (repo mySqlUserRepository) GetByEmail(email string) (*models.User, *apperror.AppError) {
	model := new(models.User)

	if err := repo.Store.Where("email = ?", email).Find(&model).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToFetchUserCode, err.Error(), ecode.ErrUnableToFetchUserMsg)
	}
	return model, nil
}
