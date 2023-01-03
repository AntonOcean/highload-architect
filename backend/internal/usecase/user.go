package usecase

import (
	"context"

	"github.com/google/uuid"

	"kek/internal/domain"
)

func (uc uc) CreateUser(ctx context.Context, user *domain.User) error {
	return uc.serviceRepo.CreateUser(ctx, user)
}

func (uc uc) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	return uc.serviceRepo.GetUserByID(ctx, userID)
}
