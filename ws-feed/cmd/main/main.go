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

	"ws-feed/internal/adapters"
	"ws-feed/internal/amqp"
	"ws-feed/internal/api"
	"ws-feed/internal/config"
	"ws-feed/internal/repository"
	"ws-feed/internal/usecase"
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

	uc := usecase.New(
		adapters.NewBackendService(
			adapters.NewRestyClient(cfg.Backend.HOST, cfg.Backend.ConnectionTimeout),
		),
		repository.New(),
		logger,
		&cfg.Jwt,
	)

	httpAPIrouter := api.New(
		logger,
		uc,
	)

	httpServer := http.Server{
		Addr:              cfg.HTTPAPI.Addr,
		Handler:           httpAPIrouter,
		ReadHeaderTimeout: time.Minute * 1,
	}

	consumer, err := amqp.BuildConsumer(cfg)
	if err != nil {
		logger.Error(err.Error())
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

	go func() {
		if err := consumer.StartConsume(logger, uc); err != nil {
			logger.Error("ws feed: failed to start consume", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	logger.Info("ws feed: ready to accept msg")

	<-quit

	ctxGIN, cancelGIN := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelGIN()

	if err := httpServer.Shutdown(ctxGIN); err != nil {
		logger.Error("http server: forced to shutdown", zap.Error(err))
	}

	logger.Info("http server: successfully stopped")

	ctxConsumer, cancelConsumer := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelConsumer()

	if err := consumer.Close(ctxConsumer); err != nil {
		logger.Error("http server: forced to shutdown", zap.Error(err))
	}

	logger.Info("ws feed: successfully stopped")
}
