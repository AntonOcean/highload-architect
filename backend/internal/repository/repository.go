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
}

type ServiceRepository interface {
	User
}

func New(dbpool *pgxpool.Pool) ServiceRepository {
	return rw{
		store: dbpool,
	}
}
