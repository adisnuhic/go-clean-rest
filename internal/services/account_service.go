package services

import (
	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/adisnuhic/go-clean/pkg/apperror"
)

// IAccountService represents account service contract
type IAccountService interface {
	Login(email string, password string) (*models.User, string, string, *apperror.AppError)
	AuthenticateUser(user *models.User) (string, string, *apperror.AppError)
}

type accountService struct {
	TokenService ITokenService
	AuthService  IAuthService

	Repo             repositories.IAccountRepository
	UserRepo         repositories.IUserRepository
	AuthProviderRepo repositories.IAuthProviderRepository
	TokenRepo        repositories.ITokenRepository
}

// NewAccountService -
func NewAccountService(tokenSvc ITokenService, authSvc IAuthService, repo repositories.IAccountRepository, userRepo repositories.IUserRepository, authProviderRepo repositories.IAuthProviderRepository, tokenRepo repositories.ITokenRepository) IAccountService {
	return &accountService{
		TokenService:     tokenSvc,
		AuthService:      authSvc,
		Repo:             repo,
		UserRepo:         userRepo,
		AuthProviderRepo: authProviderRepo,
		TokenRepo:        tokenRepo,
	}
}

// Login user
func (svc accountService) Login(email string, password string) (*models.User, string, string, *apperror.AppError) {
	user, errUser := svc.UserRepo.GetByEmail(email)

	if errUser != nil {
		return nil, "", "", apperror.New(ecode.ErrLoginFailedCode, errUser.Cause, ecode.ErrLoginFailedMsg)
	}

	// get user data
	authProvider, errAuth := svc.AuthProviderRepo.GetByUserID(user.ID)
	if errAuth != nil {
		return nil, "", "", errAuth
	}

	// check if hash matches
	if ok := svc.AuthService.ComparePasswordHash(password, authProvider.UID); !ok {
		return nil, "", "", apperror.New(ecode.ErrLoginFailedCode, ecode.ErrLoginFailedMsg, ecode.ErrLoginFailedMsg)
	}

	// authenticate user
	accessToken, refreshToken, err := svc.AuthenticateUser(user)
	if err != nil {
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

// AuthenticateUser authenticates user
func (svc accountService) AuthenticateUser(user *models.User) (string, string, *apperror.AppError) {

	// access token
	accessToken, errToken := svc.AuthService.GenerateAccessToken(user.ID, user.Email)
	if errToken != nil {
		return "", "", errToken
	}

	// refresh token
	refreshToken, errRefreshToken := svc.TokenService.CreateRefreshToken(user.ID, user.Email)
	if errRefreshToken != nil {
		return "", "", errRefreshToken
	}

	return accessToken, refreshToken.Token, nil
}
