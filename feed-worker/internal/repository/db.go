package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"feed-worker/internal/config"
)

func Connect(cfg *config.Config) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Hosts,
		DB:   0,
	})

	if _, err := rds.Ping(context.Background()).Result(); err != nil {
		return nil, fmt.Errorf("failed to ping keydb cluster: %w", err)
	}

	return rds, nil
}
