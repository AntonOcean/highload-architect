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

	token, err := uc.CreateToken(ctx, userID)
	if err != nil {
		return "", err
	}

	err = uc.serviceRepo.SetLastLoginUser(ctx, userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc uc) CreateToken(ctx context.Context, userID uuid.UUID) (string, error) {
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

func (uc uc) GetTokenData(ctx context.Context, token string) (*domain.Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(uc.jwt.SignKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenObj.Valid {
		return nil, errors.New("token invalid")
	}

	data, ok := tokenObj.Claims.(*domain.Claims)
	if !ok {
		return nil, errors.New("token claims not parse")
	}

	return data, nil
}
