package tarantoolrepository

import (
	"context"

	"github.com/tarantool/go-tarantool"

	"chat/internal/repository"
)

type t struct {
	store *tarantool.Connection
}

func NewTarantool(db *tarantool.Connection) repository.ServiceRepository {
	return t{
		store: db,
	}
}

func (t t) Ping(ctx context.Context) error {
	_, err := t.store.Ping()
	return err
}
