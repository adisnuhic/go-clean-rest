package services

import (
	"time"

	"github.com/adisnuhic/go-clean/internal/ecode"
	"github.com/adisnuhic/go-clean/pkg/apperror"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// IAuthService represents auth service contract
type IAuthService interface {
	GeneratePasswordHash(password string) (string, *apperror.AppError)
	ComparePasswordHash(password string, hash string) bool
	GenerateAccessToken(userID uint64, email string) (string, *apperror.AppError)
}

type authService struct {
	JWTSecret string
}

// NewAuthService -
func NewAuthService(jwtSecret string) IAuthService {
	return &authService{
		JWTSecret: jwtSecret,
	}
}

// GeneratePasswordHash generates hash for provided password
func (authService) GeneratePasswordHash(password string) (string, *apperror.AppError) {
	hashByte, errHashByte := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHashByte != nil {
		return "", apperror.New(ecode.ErrUnableToGenerateHashCode, errHashByte.Error(), ecode.ErrUnableToGenerateHashMsg)
	}

	return string(hashByte), nil
}

// GeneratePasswordHash generates hash for provided password
func (svc authService) ComparePasswordHash(password string, hash string) bool {
	errPassword := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return errPassword == nil
}

// GenerateAccessToken generates access token
func (svc authService) GenerateAccessToken(userID uint64, email string) (string, *apperror.AppError) {
	jwtSecret := []byte(svc.JWTSecret)
	expirationTime := time.Now().UTC().Add(3000 * time.Minute)

	type CustomClaims struct {
		ID    uint64 `json:"id"`
		Email string `json:"email"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := CustomClaims{
		userID,
		email,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// JWT implementation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", apperror.New(ecode.ErrUnableToGenerateAccessTokenCode, err.Error(), ecode.ErrUnableToGenerateAccessTokenMsg)
	}

	return tokenString, nil
}
