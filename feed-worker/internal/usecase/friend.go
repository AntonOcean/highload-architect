package usecase

import (
	"context"

	"feed-worker/internal/domain"
)

func (uc uc) FriendHandler(ctx context.Context, msg domain.NewFriend) error {
	posts, err := uc.backend.GetPostsByAuthorID(ctx, msg.FriendID)
	if err != nil {
		return err
	}

	err = uc.UpdateFeedByUserID(ctx, msg.UserID, posts)
	if err != nil {
		return err
	}

	return nil
}
