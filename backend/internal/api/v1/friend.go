package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kek/internal/api/v1/formatter"
)

// CreateFriend godoc
// @Summary Добавить друга
// @Description Для пользователя с ИД(user_id) из Authorization добавляем друга с ИД из body
// @tags friends
// @Accept json
// @Param user body formatter.DomainID true "ИД потенциального друга"
// @Produce json
// @Success 201 "Пользователь успешно указал своего друга"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/friend [post].
func (rH RouterHandler) CreateFriend(c *gin.Context) {
	var request formatter.DomainID

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&request)
	if err != nil {
		_ = c.Error(fmt.Errorf("%w %v", formatter.ErrInvalidData, err))
		return
	}

	friendID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	err = rH.ucService.CreateFriend(ctx, userID.(uuid.UUID), friendID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// DeleteFriendByID godoc
// @Summary Удалить друга из друзей
// @Description Для пользователя с ИД(user_id) из Authorization удаляем друга с ИД из query params
// @tags friends
// @Accept json
// @Param id path string true "ИД удаляемого друга"
// @Produce json
// @Success 200 "Пользователь успешно удалил из друзей пользователя"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/friend/{id} [post].
func (rH RouterHandler) DeleteFriendByID(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("id"))

	friendID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	err = rH.ucService.DeleteFriend(ctx, userID.(uuid.UUID), friendID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetFriendsByUserID godoc
// @Summary Получить друзей пользователя с ИД user_id
// @Description Получить друзей пользователя с ИД user_id
// @tags friends,admin
// @Accept json
// @Produce json
// @Param user_id path string true "user ID"
// @Success 200 []formatter.GetUser "Успешно получен пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/admin/user/:user_id/friend [get].
func (rH RouterHandler) GetFriendsByUserID(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("user_id"))

	userID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	users, err := rH.ucService.GetFriendsWithUserID(ctx, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreateUserListResp(users))
}
