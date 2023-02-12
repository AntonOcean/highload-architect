package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"kek/internal/api/v1/formatter"
	"kek/internal/domain"
)

func (rw rw) CreatePost(ctx context.Context, post *domain.Post) error {
	if post == nil {
		return nil
	}

	if _, err := rw.store.Exec(
		ctx,
		`INSERT INTO 
			posts (id, author_id, text)
		VALUES ($1, $2, $3);`,
		post.ID, post.AuthorID, post.Text,
	); err != nil {
		return err
	}

	return nil
}

func (rw rw) GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.Post, error) {
	p := domain.Post{}

	if err := rw.store.QueryRow(
		ctx,
		`SELECT 
			p.id,
			p.text,
			p.author_id
		FROM posts p 
		WHERE p.id = $1`, postID,
	).Scan(&p.ID, &p.Text, &p.AuthorID); err != nil {
		if err == pgx.ErrNoRows {
			return nil, formatter.ErrNotFound
		}

		return nil, err
	}

	return &p, nil
}

func (rw rw) UpdatePost(ctx context.Context, text string, postID uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`UPDATE 
			posts SET text=$2, updated=now()
		WHERE id=$1;`,
		postID, text,
	); err != nil {
		return err
	}

	return nil
}

func (rw rw) DeletePostByID(ctx context.Context, postID uuid.UUID) error {
	if _, err := rw.store.Exec(
		ctx,
		`DELETE FROM 
			posts
		WHERE id=$1`,
		postID,
	); err != nil {
		return err
	}

	return nil
}
