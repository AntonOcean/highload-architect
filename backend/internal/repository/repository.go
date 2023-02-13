package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"kek/internal/domain"
)

type rw struct {
	store *pgxpool.Pool
}

type User interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error)
	GetUsersByPrefix(ctx context.Context, firstName, lastName string) ([]*domain.User, error)
	SetLastLoginUser(ctx context.Context, userID uuid.UUID) error
}

type Friend interface {
	CreateFriend(ctx context.Context, userID, friendID uuid.UUID) error
	DeleteFriend(ctx context.Context, userID, friendID uuid.UUID) error
}

type Post interface {
	CreatePost(ctx context.Context, post *domain.Post) error
	GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.Post, error)
	UpdatePost(ctx context.Context, text string, postID uuid.UUID) error
	DeletePostByID(ctx context.Context, postID uuid.UUID) error
}

type ServiceRepository interface {
	User
	Friend
	Post
}

func New(dbpool *pgxpool.Pool) ServiceRepository {
	return rw{
		store: dbpool,
	}
}
