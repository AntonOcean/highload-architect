package usecase

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"kek/internal/config"
	"kek/internal/domain"
	"kek/internal/repository"
)

type Auth interface {
	AuthUser(ctx context.Context, userID uuid.UUID, password string) (string, error)
}

type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
}

type ServiceUsecase interface {
	Auth
	User
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
