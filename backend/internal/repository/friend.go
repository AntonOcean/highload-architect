package repository

import (
	"context"

	"kek/internal/domain"

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

func (rw rw) GetFriendsWithUserID(ctx context.Context, userID uuid.UUID) ([]*domain.User, error) {
	// TODO u.last_login > последний месяц
	rows, err := rw.store.Query(
		ctx,
		`SELECT 
			u.id,
			u.first_name,
			u.last_name,
			u.age,
			u.gender,
			u.biography,
			u.city,
			u.password
		FROM friends f
		JOIN users u ON f.user_id = u.id
		WHERE f.friend_id = $1
		ORDER BY u.last_login DESC;`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for rows.Next() {
		u := domain.User{}

		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Gender, &u.Biography, &u.City, &u.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}
