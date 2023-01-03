package formatter_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"

	"kek/internal/api/v1/formatter"
	"kek/internal/domain"
)

func TestUsecase_UserParseCreateUser(t *testing.T) {
	t.Run("parse user", func(t *testing.T) {
		userRequest := formatter.CreateUser{
			User: formatter.User{
				FirstName: "User",
				LastName:  "User",
				Age:       30,
				Gender:    "female",
				Biography: "-",
				City:      "Moscow",
			},
			UserPassword: formatter.UserPassword{
				Password: "kek",
			},
		}

		user, err := userRequest.ToDomain()
		assert.Equal(t, err, nil)
		assert.Equal(t, user.FirstName, "User")
		assert.Equal(t, user.LastName, "User")
		assert.Equal(t, user.Age, 30)
		assert.Equal(t, user.Gender, domain.Female)
		assert.Equal(t, user.Biography, "-")
		assert.Equal(t, user.City, "Moscow")
		assert.NotEmpty(t, user.Password)
		assert.NotEqual(t, user.Password, "kek")
		assert.Equal(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("kek")), nil)
		assert.NotEqual(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("kek1")), nil)
		assert.NotEmpty(t, user.ID)

		userRequest = formatter.CreateUser{
			User: formatter.User{
				FirstName: "",
				LastName:  "User",
				Age:       30,
				Gender:    "female",
				Biography: "-",
				City:      "Moscow",
			},
			UserPassword: formatter.UserPassword{
				Password: "kek",
			},
		}

		user, err = userRequest.ToDomain()
		assert.True(t, errors.Is(err, formatter.ErrInvalidData))
		assert.Empty(t, user)

		userRequest = formatter.CreateUser{
			User: formatter.User{
				FirstName: "User",
				LastName:  "User",
				Age:       -30,
				Gender:    "female",
				Biography: "-",
				City:      "Moscow",
			},
			UserPassword: formatter.UserPassword{
				Password: "kek",
			},
		}

		user, err = userRequest.ToDomain()
		assert.True(t, errors.Is(err, formatter.ErrInvalidData))
		assert.Empty(t, user)
	})
}
