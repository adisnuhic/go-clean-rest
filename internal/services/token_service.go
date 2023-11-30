package services

import (
	"time"

	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/adisnuhic/go-clean/pkg/log"
	"github.com/adisnuhic/go-clean/pkg/utils"
	"github.com/rs/xid"
)

const (
	// TokenTypeRefreshToken - refresh token
	TokenTypeRefreshToken = 1

	// TokenTypePasswordResetToken - password reset token
	TokenTypePasswordResetToken = 2

	// TokenTypePollResponseValidationToken - email reset token
	TokenTypeEmailResetToken = 3

	// RefreshTokenDuration holds duration value in minutes for refresh token
	RefreshTokenDuration = 43200 // 30 days

	// ResetPasswordTokenDuration holds duration value in minutes for refresh token
	ResetPasswordTokenDuration = 1440 // 1 day

	// ResetEmailTokenDuration holds duration value in minutes for email reset token
	ResetEmailTokenDuration = 1440 // 1 day
)

// ITokenService represents token service contract
type ITokenService interface {
	CreateRefreshToken(userID uint64, email string) (*models.Token, *apperror.AppError)
	GetByToken(token string) (*models.Token, *apperror.AppError)
	CreateResetPasswordToken(userID uint64, email string) (*models.Token, *apperror.AppError)
	CreateEmailResetToken(userID uint64) (*models.Token, *apperror.AppError)
	Invalidate(t *models.Token) *apperror.AppError
}

type tokenService struct {
	Logger     log.ILogger
	Repository repositories.ITokenRepository
}

// NewTokenService -
func NewTokenService(logger log.ILogger, repo repositories.ITokenRepository) ITokenService {
	return &tokenService{
		Logger:     logger,
		Repository: repo,
	}
}

// CreateRefreshToken creates refresh token in DB
func (svc tokenService) CreateRefreshToken(userID uint64, email string) (*models.Token, *apperror.AppError) {
	expireAt := time.Now().UTC().Add(time.Minute * time.Duration(RefreshTokenDuration))
	token := &models.Token{}
	token.UserID = &userID
	token.TokenTypeID = TokenTypeRefreshToken
	token.Token = xid.New().String()
	token.ExpiresAt = expireAt

	return svc.Repository.CreateToken(token)
}

// CreateResetPasswordToken creates reset password token
func (svc tokenService) CreateResetPasswordToken(userID uint64, email string) (*models.Token, *apperror.AppError) {
	expireAt := time.Now().UTC().Add(time.Minute * time.Duration(ResetPasswordTokenDuration))
	myCode, err := utils.GenerateNumericCode(6)
	if err != nil {
		return nil, apperror.New(ecode.ErrGeneratingTokenCode, ecode.ErrGeneratingTokenMsg, ecode.ErrGeneratingTokenMsg)
	}

	token := &models.Token{}
	token.UserID = &userID
	token.TokenTypeID = TokenTypePasswordResetToken
	token.Token = xid.New().String()
	token.ExpiresAt = expireAt
	token.Code = myCode
	return svc.Repository.CreateToken(token)
}

// CreateEmailResetToken creates email rest token for provided userID
func (svc tokenService) CreateEmailResetToken(userID uint64) (*models.Token, *apperror.AppError) {
	expireAt := time.Now().UTC().Add(time.Minute * time.Duration(ResetEmailTokenDuration))
	myCode, err := utils.GenerateNumericCode(6)
	if err != nil {
		return nil, apperror.New(ecode.ErrGeneratingTokenCode, ecode.ErrGeneratingTokenMsg, ecode.ErrGeneratingTokenMsg)
	}

	token := &models.Token{}
	token.UserID = &userID
	token.TokenTypeID = TokenTypeEmailResetToken
	token.Token = xid.New().String()
	token.ExpiresAt = expireAt
	token.Code = myCode
	return svc.Repository.CreateToken(token)
}

// Invalidate token
func (svc tokenService) Invalidate(t *models.Token) *apperror.AppError {
	t.ExpiresAt = time.Now()
	return svc.Repository.Update(t)
}

// GetByToken returns token for provided token string
func (svc tokenService) GetByToken(token string) (*models.Token, *apperror.AppError) {
	return svc.Repository.GetByToken(token)
}
