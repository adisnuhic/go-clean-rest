package repositories

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/initialize/mysql"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

type tokenRepository struct {
	Store mysql.Store
}

// NewMySQLTokenRepository -
func NewMySQLTokenRepository(s mysql.Store) repositories.ITokenRepository {
	return &tokenRepository{
		Store: s,
	}
}

// CreateToken creates token
func (repo tokenRepository) CreateToken(token *models.Token) (*models.Token, *apperror.AppError) {
	if err := repo.Store.Create(token).Error; err != nil {
		return nil, apperror.New(ecode.ErrUnableToCreateTokenCode, err.Error(), ecode.ErrUnableToCreateTokenMsg)
	}

	return token, nil
}

// GetByToken returns token for provided token string
func (repo tokenRepository) GetByToken(token string) (*models.Token, *apperror.AppError) {
	model := new(models.Token)

	tx := repo.Store.Where("token = ?", token).Find(&model)

	if tx.Error != nil {
		return nil, apperror.New(ecode.ErrUnableToGetTokenCode, tx.Error.Error(), ecode.ErrUnableToGetTokenMsg)
	}

	return model, nil
}

// Update token
func (repo tokenRepository) Update(token *models.Token) *apperror.AppError {
	if err := repo.Store.Save(token).Error; err != nil {
		return apperror.New(ecode.ErrUnableToUpdateTokenCode, err.Error(), ecode.ErrUnableToUpdateTokenMsg)
	}

	return nil
}
