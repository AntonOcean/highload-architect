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

	"kek/internal/adapters"
	amqp "kek/internal/amqp"

	"go.uber.org/zap"

	_ "kek/docs"
	"kek/internal/api"
	"kek/internal/config"
	"kek/internal/repository"
	"kek/internal/usecase"
)

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

	dbPool, err := repository.Connect(cfg)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		if dbPool != nil {
			dbPool.Close()
		}
	}()

	publisher, err := amqp.BuildPublisher(cfg)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		if publisher != nil {
			_ = publisher.Close()
		}
	}()

	uc := usecase.New(
		repository.New(dbPool),
		publisher,
		adapters.NewFeedWorkerService(
			adapters.NewRestyClient(cfg.FeedWorker.HOST, cfg.FeedWorker.ConnectionTimeout),
		),
		logger,
		&cfg.Jwt,
	)

	httpAPIrouter := api.New(
		logger,
		dbPool,
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
