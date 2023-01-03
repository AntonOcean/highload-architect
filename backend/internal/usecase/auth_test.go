package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"kek/internal/usecase"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"golang.org/x/crypto/bcrypt"

	"kek/internal/config"
	"kek/internal/domain"
	"kek/internal/mocks"
)

//nolint:funlen // ok
func TestUsecase_AuthUser(t *testing.T) {
	t.Run("auth user", func(t *testing.T) {
		mockServiceRepo := mocks.ServiceRepository{}
		userID := uuid.New()

		secretKey := "key"
		pwdRight := "my_password"
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwdRight), bcrypt.DefaultCost)
		if err != nil {
			assert.Nil(t, err)
			return
		}

		password := string(hashPwd)

		mockServiceRepo.On("GetUserByID", mock.AnythingOfType("*context.emptyCtx"), userID).Return(
			func(ctx context.Context, userID uuid.UUID) *domain.User {
				return &domain.User{
					ID:        userID,
					FirstName: "User",
					LastName:  "User",
					Age:       30,
					Gender:    "Female",
					Biography: "-",
					City:      "Moscow",
					Password:  password,
				}
			},
			func(ctx context.Context, userID uuid.UUID) error {
				return nil
			},
		)

		logger := zaptest.NewLogger(t)

		uc := usecase.New(
			&mockServiceRepo,
			logger,
			&config.Jwt{
				ExpirationMinutes: 10,
				Expiration:        time.Duration(10),
				SignKey:           secretKey,
			},
		)
		at(time.Unix(0, 0), func() {
			token, err := uc.AuthUser(context.Background(), userID, pwdRight)
			assert.Equal(t, err, nil)
			assert.NotEmpty(t, token)

			tokenObj, err := jwt.ParseWithClaims(token, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(secretKey), nil
			})

			if err != nil {
				assert.Nil(t, err)
				return
			}

			assert.Equal(t, tokenObj.Valid, true)

			data, ok := tokenObj.Claims.(*domain.Claims)
			if !ok {
				assert.Equal(t, ok, true)
				return
			}

			assert.Equal(t, data.UserID, userID)
			assert.Equal(t, data.TokenType, string(domain.Access))
			assert.Equal(t, data.Issuer, "backend.auth.service")
			assert.Greater(t, data.ExpiresAt, time.Now().Unix())
		})
	})
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}

	f()

	jwt.TimeFunc = time.Now
}
