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

	"feed-worker/internal/adapters"
	"feed-worker/internal/amqp"
	"feed-worker/internal/api"
	"feed-worker/internal/config"
	"feed-worker/internal/repository"
	"feed-worker/internal/usecase"
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

	rds, err := repository.Connect(cfg)
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		if rds != nil {
			_ = rds.Close()
		}
	}()

	uc := usecase.New(
		repository.New(rds),
		adapters.NewBackendService(
			adapters.NewRestyClient(cfg.Backend.HOST, cfg.Backend.ConnectionTimeout),
		),
		logger,
	)

	httpAPIrouter := api.New(
		logger,
		rds,
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
			logger.Error("feed worker: failed to start consume", zap.Error(err))
			quit <- syscall.SIGTERM
		}
	}()

	logger.Info("feed worker: ready to accept msg")

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

	logger.Info("feed worker: successfully stopped")
}
