package usecase

import (
	"context"

	"github.com/google/uuid"

	"chat/internal/domain"

	"go.uber.org/zap"

	"chat/internal/config"
	"chat/internal/repository"
)

type Auth interface {
	GetTokenData(ctx context.Context, token string) (*domain.Claims, error)
}

type Chat interface {
	CreateMessage(ctx context.Context, senderID, receiverID uuid.UUID, text string) (*domain.Message, error)
	GetMessages(ctx context.Context, senderID, receiverID uuid.UUID) ([]*domain.Message, error)
}

type ServiceUsecase interface {
	Auth
	Chat
}

type uc struct {
	serviceRepo repository.ServiceRepository
	logger      *zap.Logger
	jwt         *config.Jwt
}

func New(mr repository.ServiceRepository, logger *zap.Logger, jwt *config.Jwt) ServiceUsecase {
	return &uc{
		serviceRepo: mr,
		logger:      logger,
		jwt:         jwt,
	}
}
