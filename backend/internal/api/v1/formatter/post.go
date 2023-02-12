package formatter

import (
	"fmt"
	"strings"

	"kek/internal/domain"
)

type Text struct {
	Text string `json:"text" example:"Lorem ipsum dolor sit amet" binding:"required"`
}

func (t *Text) ToDomain() (string, error) {
	text := strings.TrimSpace(t.Text)
	if text == "" {
		return "", fmt.Errorf("%w %s", ErrInvalidData, "Поле <Текст> не может быть пустым")
	}

	// валидация именно на размер
	if len(text) > 500 {
		return "", fmt.Errorf("%w %s", ErrInvalidData, "Поле <Текст> не может быть больше 500 символов")
	}

	return text, nil
}

type GetPost struct {
	DomainID
	Text
	AuthorID DomainIDType `json:"author_id"`
}

func CreatePostResp(post *domain.Post) *GetPost {
	if post == nil {
		return nil
	}

	return &GetPost{
		DomainID: DomainID{ID: DomainIDType(post.ID.String())},
		Text:     Text{Text: post.Text},
		AuthorID: DomainIDType(post.AuthorID.String()),
	}
}
