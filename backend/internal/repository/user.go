package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"kek/internal/api/v1/formatter"

	"github.com/google/uuid"

	"kek/internal/domain"
)

func (rw rw) CreateUser(ctx context.Context, u *domain.User) error {
	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO 
			users (id, first_name, last_name, age, gender, biography, city, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		u.ID, u.FirstName, u.LastName, u.Age, u.Gender, u.Biography, u.City, u.Password,
	); err != nil {
		return err
	}

	return nil
}

func (rw rw) GetUserByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	u := domain.User{}

	if err := rw.store.QueryRow(
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
		FROM users u 
		WHERE u.id = $1`, userID,
	).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.Gender, &u.Biography, &u.City, &u.Password); err != nil {
		if err == pgx.ErrNoRows {
			return nil, formatter.ErrNotFound
		}

		return nil, err
	}

	return &u, nil
}
