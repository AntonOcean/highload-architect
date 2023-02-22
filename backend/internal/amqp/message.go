package amqp

import (
	"time"
)

type KeyType string

const (
	PostEvent   KeyType = "new-post"
	FriendEvent KeyType = "new-friend"
)

type Message struct {
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
	Key       string      `json:"key"`
}
