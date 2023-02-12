package formatter

import (
	"fmt"
	"strings"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"

	"kek/internal/domain"
)

type User struct {
	FirstName string `json:"name" example:"Ivan" binding:"required"`
	LastName  string `json:"last_name" example:"Ivanov" binding:"required"`
	Age       int    `json:"age" example:"42" binding:"required"`
	Gender    string `json:"gender" example:"male" binding:"required"`
	Biography string `json:"biography" example:"-" binding:"required"`
	City      string `json:"city" example:"Moscow" binding:"required"`
}

type GetUser struct {
	DomainID
	User
}

type CreateUser struct {
	User
	UserPassword
}

type SearchUser struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}

func (u *CreateUser) ToDomain() (*domain.User, error) {
	errStrings := make([]string, 0)

	firstName := strings.TrimSpace(u.FirstName)
	if firstName == "" {
		errStrings = append(errStrings, "Поле <Имя> не может быть пустым")
	}

	lastName := strings.TrimSpace(u.LastName)
	if lastName == "" {
		errStrings = append(errStrings, "Поле <Фамилия> не может быть пустым")
	}

	age := u.Age
	if age <= 0 {
		errStrings = append(errStrings, "Поле <Возраст> должно быть больше 0")
	}

	gender := domain.GenderType(strings.TrimSpace(u.Gender))
	switch gender {
	case domain.Male, domain.Female, domain.Other:
	default:
		errStrings = append(errStrings, "Поле <Пол> может быть: мужской, женский, другое")
	}

	biography := strings.TrimSpace(u.Biography)
	if biography == "" {
		errStrings = append(errStrings, "Поле <Интересы> не может быть пустым")
	}

	city := strings.TrimSpace(u.City)
	if city == "" {
		errStrings = append(errStrings, "Поле <Город> не может быть пустым")
	}

	password := strings.TrimSpace(u.Password)
	if password == "" {
		errStrings = append(errStrings, "Поле <Пароль> не может быть пустым")
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	if len(errStrings) > 0 {
		return nil, fmt.Errorf("%w: %s", ErrInvalidData, strings.Join(errStrings, "; "))
	}

	password = string(hashPwd)

	return &domain.User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		Gender:    gender,
		Biography: biography,
		City:      city,
		Password:  password,
	}, nil
}

func CreateUserResp(user *domain.User) *GetUser {
	if user == nil {
		return nil
	}

	return &GetUser{
		DomainID: DomainID{
			ID: DomainIDType(user.ID.String()),
		},
		User: User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Age:       user.Age,
			Gender:    string(user.Gender),
			Biography: user.Biography,
			City:      user.City,
		},
	}
}

func CreateUserListResp(users []*domain.User) []*GetUser {
	response := make([]*GetUser, len(users))

	for i := range users {
		response[i] = CreateUserResp(users[i])
	}

	return response
}

func (q *SearchUser) ToDomain() error {
	errStrings := make([]string, 0)

	q.LastName = strings.TrimSpace(q.LastName)
	if q.LastName == "" {
		errStrings = append(errStrings, "Поле <Фамилия> не может быть пустым")
	}

	q.FirstName = strings.TrimSpace(q.FirstName)
	if q.FirstName == "" {
		errStrings = append(errStrings, "Поле <Имя> не может быть пустым")
	}

	if len(errStrings) > 0 {
		return fmt.Errorf("%w: %s", ErrInvalidData, strings.Join(errStrings, "; "))
	}

	return nil
}
