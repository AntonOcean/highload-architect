package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"chat/internal/domain"
)

type rw struct {
	store *pgxpool.Pool
}

type Ping interface {
	Ping(ctx context.Context) error
}

type Chat interface {
	CreateMessage(ctx context.Context, msg *domain.Message) error
	GetMessages(ctx context.Context, senderID, receiverID uuid.UUID) ([]*domain.Message, error)
}

type ServiceRepository interface {
	Ping
	Chat
}

func New(dbpool *pgxpool.Pool) ServiceRepository {
	return rw{
		store: dbpool,
	}
}

func (rw rw) Ping(ctx context.Context) error {
	return rw.store.Ping(ctx)
}
