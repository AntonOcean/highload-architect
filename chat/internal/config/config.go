package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Jwt struct {
	SignKey string `env:"JWT_KEY,notEmpty"`
}

type Config struct {
	LogLevel string `env:"LOG_LEVEL,notEmpty"`

	HTTPAPI struct {
		Addr                  string `env:"ADDR,notEmpty"`
		ServerShutdownTimeout time.Duration
	}

	Postgres struct {
		Host              string `env:"POSTGRES_HOST,notEmpty"`
		Port              string `env:"POSTGRES_PORT,notEmpty"`
		User              string `env:"POSTGRES_USER,notEmpty"`
		Password          string `env:"POSTGRES_PASSWORD,notEmpty"`
		Database          string `env:"POSTGRES_DB,notEmpty"`
		ConnectionTimeout time.Duration
	}

	Tarantool struct {
		Host              string `env:"TARANTOOL_HOST"`
		Port              string `env:"TARANTOOL_PORT"`
		ConnectionTimeout time.Duration
	}

	Jwt
}

func (cfg *Config) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, cfg.Postgres.User, cfg.Postgres.Password,
	)
}

func (cfg *Config) TarantoolDSN() string {
	return fmt.Sprintf("%s:%s",
		cfg.Tarantool.Host, cfg.Tarantool.Port,
	)
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

	cfg.Postgres.ConnectionTimeout = 5 * time.Second

	cfg.Tarantool.ConnectionTimeout = 5 * time.Second

	return cfg
}
