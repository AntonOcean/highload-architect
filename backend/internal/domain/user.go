package domain

import (
	"github.com/google/uuid"
)

type GenderType string

const (
	Male   GenderType = "male"
	Female GenderType = "female"
	Other  GenderType = "other"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int // лучше Birthday time.Time
	Gender    GenderType
	Biography string
	City      string
	Password  string
}
