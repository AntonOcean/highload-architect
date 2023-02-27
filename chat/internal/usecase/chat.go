package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"chat/internal/domain"
)

func (uc uc) CreateMessage(ctx context.Context, senderID, receiverID uuid.UUID, text string) (*domain.Message, error) {
	msg := &domain.Message{
		ID:         uuid.New(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Text:       text,
		Created:    time.Now(),
	}

	err := uc.serviceRepo.CreateMessage(ctx, msg)
	if err != nil {
		return nil, err
	}

	return msg, err
}

func (uc uc) GetMessages(ctx context.Context, senderID, receiverID uuid.UUID) ([]*domain.Message, error) {
	return uc.serviceRepo.GetMessages(ctx, senderID, receiverID)
}
