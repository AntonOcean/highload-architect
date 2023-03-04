package api

import (
	"time"

	"ws-feed/internal/api/v1/formatter"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	v1 "ws-feed/internal/api/v1"
	"ws-feed/internal/usecase"
)

func New(
	log *zap.Logger,
	ucService usecase.ServiceUsecase,
) *gin.Engine {
	router := newGINRouter(log)

	registerHealthCheck(router)

	routerHandler := v1.NewRouter(ucService, log)

	api := router.Group("/api/v1", routerHandler.AuthMiddleware())
	routerHandler.RegisterRoutes(api)

	return router
}

func newGINRouter(log *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(ginzap.RecoveryWithZap(log, true))
	router.Use(ginzap.Ginzap(log, time.RFC3339, true))

	router.Use(formatter.HandleErrors)

	return router
}

func registerHealthCheck(router *gin.Engine) {
	handler := newHealthCheckHandler()

	router.GET("/health", handler.Health)
	router.GET("/ready", handler.Ready)
}
