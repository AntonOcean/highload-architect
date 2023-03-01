package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (uc uc) UpdateFeed(ctx context.Context, postID uuid.UUID) error {
	post, err := uc.backend.GetPostByID(ctx, postID)
	if err != nil {
		return err
	}

	userIDs, err := uc.backend.GetFriendWithUserID(ctx, post.AuthorID)
	if err != nil {
		return err
	}

	err = uc.hub.SendPostToClients(userIDs, post)
	if err != nil {
		return err
	}

	return nil
}
