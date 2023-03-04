package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"ws-feed/internal/config"
	"ws-feed/internal/repository"

	"go.uber.org/zap"

	"ws-feed/internal/adapters"
	"ws-feed/internal/domain"
)

type Auth interface {
	GetTokenData(ctx context.Context, token string) (*domain.Claims, error)
}

type Feed interface {
	UpdateFeed(ctx context.Context, postID uuid.UUID) error
}

type Client interface {
	NewClient(ctx context.Context, userID uuid.UUID, conn *websocket.Conn)
}

type ServiceUsecase interface {
	Feed
	Auth
	Client
}

type uc struct {
	backend *adapters.BackendService
	hub     repository.ServiceRepository
	logger  *zap.Logger
	jwt     *config.Jwt
}

func New(b *adapters.BackendService, h repository.ServiceRepository, logger *zap.Logger, jwt *config.Jwt) ServiceUsecase {
	return &uc{
		backend: b,
		hub:     h,
		logger:  logger,
		jwt:     jwt,
	}
}
