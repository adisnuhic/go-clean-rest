package repositories

import (
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// ITokenRepository represents the token repository contract
type ITokenRepository interface {
	CreateToken(token *models.Token) (*models.Token, *apperror.AppError)
	GetByToken(token string) (*models.Token, *apperror.AppError)
	Update(token *models.Token) *apperror.AppError
}
