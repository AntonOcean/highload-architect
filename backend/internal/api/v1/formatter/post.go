package formatter

import (
	"fmt"
	"strings"
	"time"

	"kek/internal/domain"
)

type LimitOffset struct {
	Limit  int `form:"limit" example:"10"`
	Offset int `form:"offset" example:"100"`
}

func (t *LimitOffset) ToDomain() error {
	if t.Limit < 0 {
		return fmt.Errorf("%w %s", ErrInvalidData, "Поле <Лимит> не может быть меньше 0")
	}

	if t.Limit == 0 {
		t.Limit = 10
	}

	if t.Offset < 0 {
		return fmt.Errorf("%w %s", ErrInvalidData, "Поле <Офсет> не может быть меньше 0")
	}

	return nil
}

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
	Created  time.Time    `json:"created"`
}

func CreatePostResp(post *domain.Post) *GetPost {
	if post == nil {
		return nil
	}

	return &GetPost{
		DomainID: DomainID{ID: DomainIDType(post.ID.String())},
		Text:     Text{Text: post.Text},
		AuthorID: DomainIDType(post.AuthorID.String()),
		Created:  post.Created,
	}
}

func CreatePostListResp(posts []*domain.Post) []*GetPost {
	response := make([]*GetPost, len(posts))

	for i := range posts {
		response[i] = CreatePostResp(posts[i])
	}

	return response
}
