package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"chat/internal/api/v1/formatter"
)

// CreateMessage godoc
// @Summary Отправить сообщение в диалог
// @Description От с ИД(user_id) из Authorization отправляем сообщение пользователю с ИД в урле
// @tags chat
// @Accept json
// @Param text body formatter.Text true "Текст сообщения, не более 200 символов"
// @Produce json
// @Success 201 formatter.GetMessage "Успешно отправлено сообщение"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/dialog/{receiver_user_id} [post].
func (rH RouterHandler) CreateMessage(c *gin.Context) {
	var request formatter.Text

	ctx := c.Request.Context()

	err := c.ShouldBindJSON(&request)
	if err != nil {
		_ = c.Error(fmt.Errorf("%w %v", formatter.ErrInvalidData, err))
		return
	}

	text, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	requestID := formatter.DomainIDType(c.Param("receiver_user_id"))

	receiverID, err := requestID.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	senderID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	msg, err := rH.ucService.CreateMessage(ctx, senderID.(uuid.UUID), receiverID, text)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, formatter.CreateChatResp(msg))
}

// GetMessages godoc
// @Summary Получить диалог
// @Description Получить диалог с пользователем ИД из урла
// @tags chat
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 []formatter.GetMessage "Диалог между двумя пользователями"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/dialog/{receiver_user_id} [get].
func (rH RouterHandler) GetMessages(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("receiver_user_id"))

	receiverID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	senderID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	msgs, err := rH.ucService.GetMessages(ctx, senderID.(uuid.UUID), receiverID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreateChatListResp(msgs))
}
