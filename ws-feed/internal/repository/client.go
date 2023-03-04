package repository

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"

	"ws-feed/internal/domain"
)

type Message struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	AuthorID string    `json:"author_id"`
	Created  time.Time `json:"created"`
}

func toMessage(post *domain.Post) *Message {
	return &Message{
		ID:       post.ID.String(),
		Text:     post.Text,
		AuthorID: post.AuthorID.String(),
		Created:  post.Created,
	}
}

func (rw rw) SendPostToClients(userIDs []uuid.UUID, post *domain.Post) error {
	data := toMessage(post)

	msg, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for _, uid := range userIDs {
		client, ok := rw.store.clients[uid]
		if !ok {
			continue
		}

		client.SendCh <- msg
	}

	return nil
}

func (rw rw) RegisterNewClient(client *domain.Client) {
	rw.store.register <- client

	go client.WriteToWS()
	go client.ReadFromWS(rw.store.unregister, rw.store.broadcast)
}
