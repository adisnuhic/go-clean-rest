package models

import "time"

// Token -
type Token struct {
	ID          uint64    `json:"id"`
	UserID      *uint64   `json:"user_id"`
	Token       string    `json:"token" binding:"required"`
	TokenTypeID uint64    `json:"token_type_id" binding:"required"`
	Code        string    `json:"code"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Token collection
type Tokens []Token
