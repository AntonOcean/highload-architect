package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"kek/internal/domain"
)

const cacheSize = 1000

func (uc uc) GetFeedByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) (posts []*domain.Post, err error) {
	if offset+limit <= cacheSize {
		ctxl, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
		defer cancel()

		posts, err = uc.feedWorker.GetFeedByUserID(ctxl, userID, limit, offset)
		if err != nil {
			uc.logger.Error(fmt.Sprintf("err in cache: %+v", err))
		}
	}

	if err != nil || offset+limit > cacheSize {
		posts, err = uc.serviceRepo.GetFeedByUserID(ctx, userID, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}
