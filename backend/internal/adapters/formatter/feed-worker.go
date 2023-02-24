package formatter

import (
	"time"

	"github.com/google/uuid"

	"kek/internal/domain"
)

type PostResponse struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	AuthorID string    `json:"author_id"`
	Created  time.Time `json:"created"`
}

func (p *PostResponse) ToDomain() *domain.Post {
	return &domain.Post{
		ID:       uuid.MustParse(p.ID),
		Text:     p.Text,
		AuthorID: uuid.MustParse(p.AuthorID),
		Created:  p.Created,
	}
}

func ToDomainPostList(posts []PostResponse) []*domain.Post {
	r := make([]*domain.Post, len(posts))
	for i := range posts {
		r[i] = posts[i].ToDomain()
	}

	return r
}
