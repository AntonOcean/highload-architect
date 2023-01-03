package formatter_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"kek/internal/api/v1/formatter"
)

func TestUsecase_AuthParseUserID(t *testing.T) {
	t.Run("parse userID", func(t *testing.T) {
		userRequestID := formatter.UserID{ID: "90"}
		userID, err := userRequestID.ToDomain()
		assert.Equal(t, userID, uuid.UUID{})
		assert.True(t, errors.Is(err, formatter.ErrInvalidData))

		userRequestID = formatter.UserID{ID: "        "}
		userID, err = userRequestID.ToDomain()
		assert.Equal(t, userID, uuid.UUID{})
		assert.True(t, errors.Is(err, formatter.ErrInvalidData))

		userRequestID = formatter.UserID{ID: "dd724b0b-8907-41b2-807b-7d359dd77f4c"}
		userID, err = userRequestID.ToDomain()
		assert.Equal(t, err, nil)
		assert.Equal(t, userID.String(), "dd724b0b-8907-41b2-807b-7d359dd77f4c")
	})
}

func TestUsecase_AuthParseAuth(t *testing.T) {
	t.Run("parse auth", func(t *testing.T) {
		passwordRequest := formatter.AuthUser{
			UserID:       formatter.UserID{ID: "dd724b0b-8907-41b2-807b-7d359dd77f4c"},
			UserPassword: formatter.UserPassword{Password: "kek"},
		}

		auth, err := passwordRequest.ToDomain()
		assert.Equal(t, err, nil)
		assert.Equal(t, auth.Password, "kek")
		assert.Equal(t, auth.ID.String(), "dd724b0b-8907-41b2-807b-7d359dd77f4c")

		passwordRequest = formatter.AuthUser{
			UserID:       formatter.UserID{ID: "dd724b0b-8907-41b2-807b-7d359dd77f4c"},
			UserPassword: formatter.UserPassword{Password: " "},
		}

		auth, err = passwordRequest.ToDomain()
		assert.True(t, errors.Is(err, formatter.ErrInvalidData))
		assert.Empty(t, auth)
	})
}
