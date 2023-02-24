package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kek/internal/api/v1/formatter"
)

// GetFeed godoc
// @Summary Получить ленту новостой для пользователя
// @Description Получить ленту новостой для пользователя
// @tags feed
// @Accept json
// @Produce json
// @Param offset query int false "офсет - дефолт 0"
// @Param limit query int false "лимит - дефолт 10"
// @Success 200 []formatter.GetPost "Успешно получены посты друзей"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 403 "Вы не автор"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/feed [get].
func (rH RouterHandler) GetFeed(c *gin.Context) {
	var request formatter.LimitOffset

	ctx := c.Request.Context()

	err := c.ShouldBindQuery(&request)
	if err != nil {
		_ = c.Error(fmt.Errorf("%w %v", formatter.ErrInvalidData, err))
		return
	}

	err = request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	posts, err := rH.ucService.GetFeedByUserID(ctx, userID.(uuid.UUID), request.Limit, request.Offset)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreatePostListResp(posts))
}
