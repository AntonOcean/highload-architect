package domain

import "github.com/google/uuid"

type NewFriend struct {
	UserID   uuid.UUID
	FriendID uuid.UUID
}

type NewPost struct {
	PostID   uuid.UUID
	AuthorID uuid.UUID
}
