package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"chat/internal/api"
	"chat/internal/config"
	"chat/internal/repository"
	t "chat/internal/tarantool.repository"
	"chat/internal/usecase"
)

//nolint:nestif,funlen // ok
func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Println("failed to read config:", err.Error())

		return
	}

	logger, _ := zap.NewProduction()
	defer func() {
		_ = logger.Sync()
	}()

	var repo repository.ServiceRepository

	if cfg.Tarantool.Host != "" {
		db, err := t.ConnectTarantool(cfg)
		if err != nil {
			logger.Error(err.Error())
		}

		defer func() {
			if db != nil {
				_ = db.Close()
			}
		}()

		repo = t.NewTarantool(db)
	} else {
		dbPool, err := repository.Connect(cfg)
		if err != nil {
			logger.Error(err.Error())
		}

		defer func() {
			if dbPool != nil {
				dbPool.Close()
			}
		}()

		repo = repository.New(dbPool)
	}

	uc := usecase.New(
		repo,
		logger,
		&cfg.Jwt,
	)

	httpAPIrouter := api.New(
		logger,
		repo,
		uc,
	)

	httpServer := http.Server{
		Addr:              cfg.HTTPAPI.Addr,
		Handler:           httpAPIrouter,
		ReadHeaderTimeout: time.Minute * 1,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("http server: failed to listen and serve", zap.Error(err))

			quit <- syscall.SIGTERM
		}
	}()

	logger.Info("http server: ready to accept requests")

	<-quit

	ctxGIN, cancelGIN := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelGIN()

	if err := httpServer.Shutdown(ctxGIN); err != nil {
		logger.Error("http server: forced to shutdown", zap.Error(err))
	}

	logger.Info("http server: successfully stopped")
}
