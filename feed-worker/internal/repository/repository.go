package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"feed-worker/internal/domain"
)

type rw struct {
	store *redis.Client
}

type Feed interface {
	GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error)
	SetFeedForUser(ctx context.Context, userID uuid.UUID, posts []*domain.Post) error
}

type ServiceRepository interface {
	Feed
}

func New(rds *redis.Client) ServiceRepository {
	return rw{
		store: rds,
	}
}
