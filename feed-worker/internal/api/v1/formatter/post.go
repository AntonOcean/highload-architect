package formatter

import (
	"fmt"
	"time"

	"feed-worker/internal/domain"
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

type GetPost struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	AuthorID string    `json:"author_id"`
	Created  time.Time `json:"created"`
}

func CreatePostResp(post *domain.Post) *GetPost {
	if post == nil {
		return nil
	}

	return &GetPost{
		ID:       post.ID.String(),
		Text:     post.Text,
		AuthorID: post.AuthorID.String(),
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
