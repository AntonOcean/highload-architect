package repository

import (
	"context"

	"github.com/google/uuid"

	"chat/internal/domain"
)

func (rw rw) CreateMessage(ctx context.Context, msg *domain.Message) error {
	if msg == nil {
		return nil
	}

	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO 
			chats (id, sender_id, receiver_id, text)
		VALUES ($1, $2, $3, $4);`,
		msg.ID, msg.SenderID, msg.ReceiverID, msg.Text,
	); err != nil {
		return err
	}

	return nil
}

func (rw rw) GetMessages(ctx context.Context, senderID, receiverID uuid.UUID) ([]*domain.Message, error) {
	rows, err := rw.store.Query(
		ctx,
		`SELECT 
			c.id,
			c.sender_id,
			c.receiver_id,
			c.text,
			c.created
		FROM chats c
		WHERE c.sender_id=$1 AND c.receiver_id=$2
		ORDER BY c.created DESC;`,
		senderID, receiverID,
	)
	if err != nil {
		return nil, err
	}

	var msgs []*domain.Message

	for rows.Next() {
		p := domain.Message{}

		err := rows.Scan(&p.ID, &p.SenderID, &p.ReceiverID, &p.Text, &p.Created)

		if err != nil {
			return nil, err
		}

		msgs = append(msgs, &p)
	}

	return msgs, nil
}
