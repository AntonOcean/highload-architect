package usecase

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"kek/internal/adapters"
	amqp "kek/internal/amqp"

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
	GetFriendsWithUserID(ctx context.Context, userID uuid.UUID) ([]*domain.User, error)
}

type Post interface {
	CreatePost(ctx context.Context, text string, authorID uuid.UUID) (*domain.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.Post, error)
	UpdatePost(ctx context.Context, text string, postID, userID uuid.UUID) (*domain.Post, error)
	DeletePostByID(ctx context.Context, postID, userID uuid.UUID) error
	GetPostsByAuthorID(ctx context.Context, authorID uuid.UUID) ([]*domain.Post, error)
}

type Feed interface {
	GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error)
}

type ServiceUsecase interface {
	Auth
	User
	Friend
	Post
	Feed
}

type uc struct {
	serviceRepo repository.ServiceRepository
	queue       *amqp.Publisher
	feedWorker  *adapters.FeedWorkerService
	logger      *zap.Logger
	jwt         *config.Jwt
}

func New(mr repository.ServiceRepository, q *amqp.Publisher,
	w *adapters.FeedWorkerService, logger *zap.Logger, jwt *config.Jwt) ServiceUsecase {
	return &uc{
		serviceRepo: mr,
		queue:       q,
		feedWorker:  w,
		logger:      logger,
		jwt:         jwt,
	}
}
