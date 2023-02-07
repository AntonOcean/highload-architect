package repository

import (
	"context"

	"github.com/google/uuid"
)

func (rw rw) CreateFriend(ctx context.Context, userID, friendID uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO 
			friends (user_id, friend_id)
		VALUES ($1, $2) ON CONFLICT DO NOTHING;`,
		userID, friendID,
	); err != nil {
		return err
	}

	return nil
}

func (rw rw) DeleteFriend(ctx context.Context, userID, friendID uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`DELETE FROM 
			friends
		WHERE user_id=$1 AND friend_id=$2`,
		userID, friendID,
	); err != nil {
		return err
	}

	return nil
}
