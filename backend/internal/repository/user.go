package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"kek/internal/api/v1/formatter"

	"github.com/google/uuid"

	"kek/internal/domain"
)

func (rw rw) CreateUser(ctx context.Context, u *domain.User) error {
	if u == nil {
		return nil
	}

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

func (rw rw) GetUsersByPrefix(ctx context.Context, firstName, lastName string) ([]*domain.User, error) {
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
		FROM users u 
		WHERE u.first_name LIKE $1 AND u.last_name LIKE $2
		ORDER BY u.id;`, // TODO add limit/offset for performance
		fmt.Sprintf("%s%%", firstName), fmt.Sprintf("%s%%", lastName),
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

func (rw rw) SetLastLoginUser(ctx context.Context, userID uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE 
			users SET last_login=now()
		WHERE id=$1`,
		userID,
	); err != nil {
		return err
	}

	return nil
}
