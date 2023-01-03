package usecase

import (
	"context"
	"errors"
	"time"

	"kek/internal/api/v1/formatter"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"kek/internal/domain"
)

func (uc uc) AuthUser(ctx context.Context, userID uuid.UUID, password string) (string, error) {
	user, err := uc.serviceRepo.GetUserByID(ctx, userID)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", formatter.ErrInvalidData
		}

		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * uc.jwt.Expiration).Unix(),
			Issuer:    "backend.auth.service",
		},
		UserID:    userID,
		TokenType: string(domain.Access),
	})

	return token.SignedString([]byte(uc.jwt.SignKey))
}
