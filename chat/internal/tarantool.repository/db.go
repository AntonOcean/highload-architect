package tarantoolrepository

import (
	"fmt"

	"github.com/tarantool/go-tarantool"

	"chat/internal/config"
)

func ConnectTarantool(cfg *config.Config) (*tarantool.Connection, error) {
	connectionString := cfg.TarantoolDSN()

	conn, err := tarantool.Connect(connectionString, tarantool.Opts{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect tarantool: %w", err)
	}

	_, err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping tarantool: %w", err)
	}

	return conn, nil
}
