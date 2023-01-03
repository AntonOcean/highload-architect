package formatter

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"kek/internal/domain"
)

type TokenResp struct {
	Token string `json:"token" example:"jwt token"`
}

func CreateTokenResp(token string) *TokenResp {
	return &TokenResp{Token: token}
}

type UserID struct {
	ID string `json:"id" example:"dd724b0b-8907-41b2-807b-7d359dd77f4c" binding:"required"`
}

type UserPassword struct {
	Password string `json:"password" example:"P@ssW0rD" binding:"required"`
}

type AuthUser struct {
	UserID
	UserPassword
}

func (a *AuthUser) ToDomain() (*domain.UserAuth, error) {
	userID, err := a.UserID.ToDomain()
	if err != nil {
		return nil, err
	}

	password := strings.TrimSpace(a.Password)
	if password == "" {
		return nil, fmt.Errorf("%w %s", ErrInvalidData, "Поле <Пароль> не может быть пустым")
	}

	return &domain.UserAuth{
		ID:       userID,
		Password: password,
	}, nil
}

func (u *UserID) ToDomain() (uuid.UUID, error) {
	userStringID := strings.TrimSpace(u.ID)
	if userStringID == "" {
		return uuid.UUID{}, fmt.Errorf("%w %s", ErrInvalidData, "Поле <ИД> не может быть пустым")
	}

	userID, err := uuid.Parse(userStringID)

	if err != nil {
		fmt.Println(err)
		return uuid.UUID{}, fmt.Errorf("%w %s", ErrInvalidData, "Поле <ИД> не корректно")
	}

	return userID, nil
}
