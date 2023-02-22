package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/google/uuid"

	"feed-worker/internal/domain"
)

type Post struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	AuthorID string    `json:"author_id"`
	Created  time.Time `json:"created"`
}

func ToRedisList(p []*domain.Post) []*Post {
	r := make([]*Post, len(p))

	for i := range p {
		r[i] = &Post{
			ID:       p[i].ID.String(),
			Text:     p[i].Text,
			AuthorID: p[i].AuthorID.String(),
			Created:  p[i].Created,
		}
	}

	return r
}

func ToDomainList(p []*Post) []*domain.Post {
	r := make([]*domain.Post, len(p))

	for i := range p {
		r[i] = &domain.Post{
			ID:       uuid.MustParse(p[i].ID),
			Text:     p[i].Text,
			AuthorID: uuid.MustParse(p[i].AuthorID),
			Created:  p[i].Created,
		}
	}

	return r
}

func (rw rw) GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error) {
	data := make([]*Post, 0)

	val, err := rw.store.Get(ctx, userID.String()).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, fmt.Errorf("can't get by key: %w", err)
	}

	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, fmt.Errorf("can't  unmarshal from json redis: %w", err)
	}

	return ToDomainList(data), nil
}

func (rw rw) SetFeedForUser(ctx context.Context, userID uuid.UUID, posts []*domain.Post) error {
	data, err := json.Marshal(ToRedisList(posts))
	if err != nil {
		return fmt.Errorf("can't marshal to json redis: %w", err)
	}

	err = rw.store.Set(ctx, userID.String(), data, 0).Err()
	if err != nil {
		return fmt.Errorf("can't set by key: %w", err)
	}

	return nil
}
