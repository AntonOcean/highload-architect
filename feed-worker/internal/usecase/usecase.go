package usecase

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"feed-worker/internal/adapters"
	"feed-worker/internal/domain"
	"feed-worker/internal/repository"
)

type Feed interface {
	UpdateFeedByUserID(ctx context.Context, userID uuid.UUID, newPosts []*domain.Post) error
	GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error)
}

type Post interface {
	PostHandler(ctx context.Context, msg domain.NewPost) error
}

type Friend interface {
	FriendHandler(ctx context.Context, msg domain.NewFriend) error
}

type ServiceUsecase interface {
	Post
	Friend
	Feed
}

type uc struct {
	serviceRepo repository.ServiceRepository
	backend     *adapters.BackendService
	logger      *zap.Logger
}

func New(mr repository.ServiceRepository, b *adapters.BackendService, logger *zap.Logger) ServiceUsecase {
	return &uc{
		serviceRepo: mr,
		backend:     b,
		logger:      logger,
	}
}
