package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v6"
)

type Jwt struct {
	ExpirationMinutes int `env:"JWT_EXPIRATION,notEmpty"`
	Expiration        time.Duration
	SignKey           string `env:"JWT_KEY,notEmpty"`
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

	Jwt
}

func (cfg *Config) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s target_session_attrs=read-only sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, cfg.Postgres.User, cfg.Postgres.Password,
	)
}

func Read() (*Config, error) {
	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	config.HTTPAPI.Addr = fmt.Sprintf(":%s", config.HTTPAPI.Addr)
	config.Jwt.Expiration = time.Duration(config.Jwt.ExpirationMinutes)

	return setStaticSettings(&config), nil
}

func setStaticSettings(cfg *Config) *Config {
	cfg.HTTPAPI.ServerShutdownTimeout = 10 * time.Second

	cfg.Postgres.ConnectionTimeout = 5 * time.Second

	return cfg
}
