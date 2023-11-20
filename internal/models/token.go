package models

import "time"

// Token represents data struct
type Token struct {
	ID          uint64    `json:"id"`
	UserID      uint64    `json:"user_id"`
	Token       string    `json:"token"`
	TokenTypeID string    `json:"token_type_id"`
	Code        string    `json:"code"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TokenService represents token's service contract
type TokenService interface {
	GetByID(id uint64) (*Token, error)
}

// TokenRepository represents the token's repository contract
type TokenRepository interface {
	GetByID(id uint64) (*Token, error)
}
