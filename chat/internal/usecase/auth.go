package usecase

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt"

	"chat/internal/domain"
)

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
