package usecase

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"feed-worker/internal/domain"
)

func (uc uc) PostHandler(ctx context.Context, msg domain.NewPost) error {
	post, err := uc.backend.GetPostByID(ctx, msg.PostID)
	if err != nil {
		return err
	}

	friends, err := uc.backend.GetFriendWithUserID(ctx, msg.AuthorID)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	posts := []*domain.Post{post}

	for _, friendID := range friends {
		wg.Add(1)

		go func(fID uuid.UUID) {
			defer wg.Done()

			err = uc.UpdateFeedByUserID(ctx, fID, posts)
			if err != nil {
				uc.logger.Error(fmt.Sprintf("err update feer by user id: %+v", err))
			}
		}(friendID)
	}

	wg.Wait()

	return nil
}
