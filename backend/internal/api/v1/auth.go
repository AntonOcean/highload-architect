package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kek/internal/api/v1/formatter"
)

// AuthUser godoc
// @Summary Аутентификация пользователя
// @Description Упрощенный процесс аутентификации путем передачи идентификатор пользователя
// @Description и получения токена для дальнейшего прохождения авторизации
// @tags auth
// @Accept json
// @Param user body formatter.AuthUser true "ИД/пароль пользователя"
// @Produce json
// @Success 200 {object} formatter.TokenResp "Успешная аутентификация"
// @Failure 400 "Невалидные данные"
// @Failure 404 "Пользователь не найден"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/login [post].
func (rH RouterHandler) AuthUser(c *gin.Context) {
	var request formatter.AuthUser

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

	token, err := rH.ucService.AuthUser(ctx, domainRequest.ID, domainRequest.Password)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreateTokenResp(token))
}
