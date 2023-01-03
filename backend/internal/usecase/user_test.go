package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
	"golang.org/x/crypto/bcrypt"

	"kek/internal/config"
	"kek/internal/domain"
	"kek/internal/mocks"
	"kek/internal/usecase"
)

func TestUsecase_GetUserByID(t *testing.T) {
	t.Run("get user by id", func(t *testing.T) {
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
					Gender:    "female",
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

		user, err := uc.GetUserByID(context.Background(), userID)
		assert.Equal(t, err, nil)
		assert.NotEmpty(t, user)

		assert.Equal(t, user.ID, userID)
		assert.Equal(t, user.FirstName, "User")
		assert.Equal(t, user.LastName, "User")
		assert.Equal(t, user.Gender, domain.Female)
		assert.Equal(t, user.Age, 30)
		assert.Equal(t, user.City, "Moscow")
		assert.Equal(t, user.Biography, "-")
		assert.Equal(t, user.Password, password)
	})
}
