package usecase

import (
	"context"
	"errors"

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

	return uc.serviceRepo.CreateFriend(ctx, userID, friendID)
}

func (uc uc) DeleteFriend(ctx context.Context, userID, friendID uuid.UUID) error {
	return uc.serviceRepo.DeleteFriend(ctx, userID, friendID)
}
