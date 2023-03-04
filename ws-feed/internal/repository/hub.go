package repository

import (
	"github.com/google/uuid"

	"ws-feed/internal/domain"
)

// Hub TODO в будущем мб стоит разбить на группы в случае каких-то популярных блоггеров.
type Hub struct {
	clients    map[uuid.UUID]*domain.Client
	register   chan *domain.Client
	unregister chan *domain.Client
	broadcast  chan []byte
}

func newHub() *Hub {
	hub := &Hub{
		register:   make(chan *domain.Client),
		unregister: make(chan *domain.Client),
		clients:    make(map[uuid.UUID]*domain.Client),
		broadcast:  make(chan []byte),
	}

	go hub.run()

	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.ID] = client

		case client := <-h.unregister:
			if _, ok := h.clients[client.ID]; ok {
				delete(h.clients, client.ID)
				close(client.SendCh)
			}
		}
	}
}
