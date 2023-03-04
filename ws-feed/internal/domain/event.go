package domain

import "github.com/google/uuid"

type NewPost struct {
	PostID   uuid.UUID
	AuthorID uuid.UUID
}
