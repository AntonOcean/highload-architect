package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"ws-feed/internal/domain"
)

func (uc uc) NewClient(ctx context.Context, userID uuid.UUID, conn *websocket.Conn) {
	client := domain.NewClient(userID, conn)

	uc.hub.RegisterNewClient(client)
}
