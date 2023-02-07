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
	CreateToken(ctx context.Context, userID uuid.UUID) (string, error)
	GetTokenData(ctx context.Context, token string) (*domain.Claims, error)
}

type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUsersByPrefix(ctx context.Context, firstName, lastName string) ([]*domain.User, error)
}

type Friend interface {
	CreateFriend(ctx context.Context, userID, friendID uuid.UUID) error
	DeleteFriend(ctx context.Context, userID, friendID uuid.UUID) error
}

type ServiceUsecase interface {
	Auth
	User
	Friend
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
