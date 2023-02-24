package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID       uuid.UUID
	Text     string
	AuthorID uuid.UUID
	Created  time.Time
}
