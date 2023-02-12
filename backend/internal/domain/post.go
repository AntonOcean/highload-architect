package domain

import "github.com/google/uuid"

type Post struct {
	ID       uuid.UUID
	Text     string
	AuthorID uuid.UUID
}
