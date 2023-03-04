package formatter

import (
	"time"

	"github.com/google/uuid"

	"ws-feed/internal/domain"
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

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Biography string `json:"biography"`
	City      string `json:"city"`
}

func (u *UserResponse) ToDomain() uuid.UUID {
	return uuid.MustParse(u.ID)
}

func ToDomainUserIDList(users []UserResponse) []uuid.UUID {
	r := make([]uuid.UUID, len(users))
	for i := range users {
		r[i] = users[i].ToDomain()
	}

	return r
}
