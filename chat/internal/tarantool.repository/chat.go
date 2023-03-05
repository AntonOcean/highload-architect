package tarantoolrepository

import (
	"context"
	"time"

	"github.com/google/uuid"

	"chat/internal/domain"
)

type Msg struct {
	ID         string
	SenderID   string
	ReceiverID string
	Text       string
	Created    int64
}

func (t t) CreateMessage(ctx context.Context, msg *domain.Message) error {
	_, err := t.store.Call("insert_chat_message", []interface{}{
		msg.ID.String(), msg.SenderID.String(), msg.ReceiverID.String(), msg.Text, msg.Created.Unix(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (t t) GetMessages(ctx context.Context, senderID, receiverID uuid.UUID) ([]*domain.Message, error) {
	data := make([]*Msg, 0)

	err := t.store.CallTyped("select_chats", []interface{}{senderID.String(), receiverID.String()}, &data)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 && data[0].ID == "" {
		return nil, nil
	}

	msgs := make([]*domain.Message, len(data))

	for i := range data {
		msgs[i] = &domain.Message{
			ID:         uuid.MustParse(data[i].ID),
			SenderID:   uuid.MustParse(data[i].SenderID),
			ReceiverID: uuid.MustParse(data[i].ReceiverID),
			Text:       data[i].Text,
			Created:    time.Unix(data[i].Created, 0),
		}
	}

	return msgs, nil
}
