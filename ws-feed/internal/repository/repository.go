package repository

import (
	"github.com/google/uuid"

	"ws-feed/internal/domain"
)

type rw struct {
	store *Hub
}

type Client interface {
	RegisterNewClient(client *domain.Client)
	SendPostToClients(userIDs []uuid.UUID, post *domain.Post) error
}

type ServiceRepository interface {
	Client
}

func New() ServiceRepository {
	return rw{
		store: newHub(),
	}
}
