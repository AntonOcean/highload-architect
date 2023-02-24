package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	amqp "kek/internal/amqp"
	"kek/internal/domain"

	"kek/internal/api/v1/formatter"

	"github.com/google/uuid"
)

func (uc uc) CreateFriend(ctx context.Context, userID, friendID uuid.UUID) error {
	_, err := uc.serviceRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, formatter.ErrNotFound) {
			return formatter.ErrInvalidData
		}

		return err
	}

	_, err = uc.serviceRepo.GetUserByID(ctx, friendID)
	if err != nil {
		if errors.Is(err, formatter.ErrNotFound) {
			return formatter.ErrInvalidData
		}

		return err
	}

	err = uc.serviceRepo.CreateFriend(ctx, userID, friendID)
	if err != nil {
		return err
	}

	go func() {
		err := uc.queue.Push(&amqp.Message{
			Data: domain.NewFriend{
				FriendID: friendID,
				UserID:   userID,
			},
			Timestamp: time.Now(),
			Key:       string(amqp.FriendEvent),
		})
		if err != nil {
			uc.logger.Error(fmt.Sprintf("error push msg new friend: %+v", err))
		}
	}()

	return nil
}

func (uc uc) DeleteFriend(ctx context.Context, userID, friendID uuid.UUID) error {
	return uc.serviceRepo.DeleteFriend(ctx, userID, friendID)
}

func (uc uc) GetFriendsWithUserID(ctx context.Context, userID uuid.UUID) ([]*domain.User, error) {
	return uc.serviceRepo.GetFriendsWithUserID(ctx, userID)
}
