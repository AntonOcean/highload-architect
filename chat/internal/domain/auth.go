package domain

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenType string

const Access TokenType = "access"

type Claims struct {
	jwt.StandardClaims
	UserID    uuid.UUID `json:"user_id"`
	TokenType string    `json:"token_type"`
}
