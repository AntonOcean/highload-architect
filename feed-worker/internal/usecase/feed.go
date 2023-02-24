package usecase

import (
	"context"

	"github.com/google/uuid"

	"feed-worker/internal/domain"
)

const FeedSize = 1000

func (uc uc) UpdateFeedByUserID(ctx context.Context, userID uuid.UUID, newPosts []*domain.Post) error {
	oldPosts, err := uc.serviceRepo.GetFeedByUserID(ctx, userID, FeedSize, 0)
	if err != nil {
		return err
	}

	err = uc.serviceRepo.SetFeedForUser(ctx, userID, Merge(oldPosts, newPosts))
	if err != nil {
		return err
	}

	return nil
}

func Merge(left, right []*domain.Post) (result []*domain.Post) {
	result = make([]*domain.Post, len(left)+len(right))

	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0].Created.After(right[0].Created) {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	if len(result) > FeedSize {
		result = result[:FeedSize]
	}

	return result
}

func (uc uc) GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*domain.Post, error) {
	return uc.serviceRepo.GetFeedByUserID(ctx, userID, limit, offset)
}
