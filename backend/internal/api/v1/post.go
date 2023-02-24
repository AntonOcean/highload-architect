package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kek/internal/api/v1/formatter"
)

// CreatePost godoc
// @Summary Создать пост
// @Description Для пользователя с ИД(user_id) из Authorization создаем пост
// @tags posts
// @Accept json
// @Param text body formatter.Text true "Текст поста, не более 500 символов"
// @Produce json
// @Success 201 formatter.GetPost "Успешно создан пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/post [post].
func (rH RouterHandler) CreatePost(c *gin.Context) {
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

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	post, err := rH.ucService.CreatePost(ctx, text, userID.(uuid.UUID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, formatter.CreatePostResp(post))
}

// GetPostByID godoc
// @Summary Получить пост по ИД
// @Description Получить пост по ИД
// @tags posts,admin
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 formatter.GetPost "Успешно получен пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/post/{id} [get].
func (rH RouterHandler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("id"))

	postID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	post, err := rH.ucService.GetPostByID(ctx, postID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreatePostResp(post))
}

// EditPost godoc
// @Summary Редактировать пост по ИД
// @Description Редактировать пост по ИД может только автор
// @tags posts
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param text body formatter.Text true "Текст поста, не более 500 символов"
// @Success 200 formatter.GetPost "Успешно изменен пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 403 "Вы не автор"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/post/{id} [put].
func (rH RouterHandler) EditPost(c *gin.Context) {
	var requestBody formatter.Text

	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("id"))

	postID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	err = c.ShouldBindJSON(&requestBody)
	if err != nil {
		_ = c.Error(fmt.Errorf("%w %v", formatter.ErrInvalidData, err))
		return
	}

	text, err := requestBody.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	post, err := rH.ucService.UpdatePost(ctx, text, postID, userID.(uuid.UUID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreatePostResp(post))
}

// DeletePost godoc
// @Summary Удалить пост по ИД
// @Description Удалить пост по ИД может только автор
// @tags posts
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 "Успешно удален пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 403 "Вы не автор"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/post/{id} [delete].
func (rH RouterHandler) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("id"))

	postID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	err = rH.ucService.DeletePostByID(ctx, postID, userID.(uuid.UUID))
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// GetPostsByUserID godoc
// @Summary Получить посты пользователя с ИД user_id
// @Description Получить посты пользователя с ИД user_id
// @tags posts,admin
// @Accept json
// @Produce json
// @Param user_id path string true "user ID"
// @Success 200 []formatter.GetPost "Успешно получен пост"
// @Failure 400 "Невалидные данные"
// @Failure 401 "Неавторизованный доступ"
// @Failure 500 {object} formatter.Error "Ошибка сервера"
// @Failure 503 {object} formatter.Error "Ошибка сервера"
// @Header 500,503 {integer} Retry-After "Время, через которое еще раз нужно сделать запрос"
// @Router /api/v1/admin/user/:user_id/post [get].
func (rH RouterHandler) GetPostsByUserID(c *gin.Context) {
	ctx := c.Request.Context()

	request := formatter.DomainIDType(c.Param("user_id"))

	userID, err := request.ToDomain()
	if err != nil {
		_ = c.Error(err)
		return
	}

	posts, err := rH.ucService.GetPostsByAuthorID(ctx, userID)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, formatter.CreatePostListResp(posts))
}
