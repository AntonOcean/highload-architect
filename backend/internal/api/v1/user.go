package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kek/internal/api/v1/formatter"
)

// CreateUser godoc
// @Summary Зарегистрировать пользователя
// @Description Регистрация нового пользователя
// @tags users
// @Accept json
// @Param user body formatter.CreateUser true "Даннные пользователя"
// @Produce json
// @Success 201 {object} formatter.GetUser "Успешная регистрация"
// @Failure 400 "Невалидные данные"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/user [post].
func (rH RouterHandler) CreateUser(c *gin.Context) {
	var request formatter.CreateUser

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&request)
	if err != nil {
		_ = c.Error(fmt.Errorf("%w %v", formatter.ErrInvalidData, err))
		return
	}

	domainRequest, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = rH.ucService.CreateUser(ctx, domainRequest)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, formatter.CreateUserResp(domainRequest))
}

// GetUserByID godoc
// @Description Получить анкету пользователя по ИД
// @Summary Получить анкету пользователя по ИД
// @tags users
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} formatter.GetUser "Успешное получение анкеты пользователя"
// @Failure 400 "Невалидные данные"
// @Failure 404 "Анкета не найдена"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/user/{id} [get].
func (rH RouterHandler) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.UserID{
		ID: c.Param("id"),
	}

	userID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	user, err := rH.ucService.GetUserByID(ctx, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreateUserResp(user))
}
