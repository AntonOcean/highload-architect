package formatter

import (
	"time"

	"ws-feed/internal/domain"
)

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
