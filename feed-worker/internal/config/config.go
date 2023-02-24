package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type RabbitMQ struct {
	DSN   string `env:"RABBITMQ_DSN,notEmpty"`
	Queue string `env:"RABBITMQ_QUEUE,notEmpty"`
}

type Config struct {
	LogLevel string `env:"LOG_LEVEL,notEmpty"`

	HTTPAPI struct {
		Addr                  string `env:"ADDR,notEmpty"`
		ServerShutdownTimeout time.Duration
	}

	Backend struct {
		HOST              string `env:"BACKEND_HOST,notEmpty"`
		ConnectionTimeout time.Duration
	}

	Redis struct {
		Hosts             string `env:"REDIS_HOSTS,notEmpty"`
		ConnectionTimeout time.Duration
	}

	RabbitMQ RabbitMQ
}

func Read() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	config.HTTPAPI.Addr = fmt.Sprintf(":%s", config.HTTPAPI.Addr)

	return setStaticSettings(&config), nil
}

func setStaticSettings(cfg *Config) *Config {
	cfg.HTTPAPI.ServerShutdownTimeout = 10 * time.Second

	cfg.Backend.ConnectionTimeout = 10 * time.Second

	cfg.Redis.ConnectionTimeout = 5 * time.Second

	return cfg
}
