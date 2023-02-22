package api

import (
	"time"

	"feed-worker/internal/api/v1/formatter"

	"github.com/go-redis/redis/v8"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	v1 "feed-worker/internal/api/v1"
	"feed-worker/internal/usecase"
)

func New(
	log *zap.Logger,
	db *redis.Client,
	ucService usecase.ServiceUsecase,
) *gin.Engine {
	router := newGINRouter(log)

	registerSwagger(router)
	registerHealthCheck(router, db)

	routerHandler := v1.NewRouter(ucService, log)

	api := router.Group("/api/v1/admin")
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

func registerHealthCheck(router *gin.Engine, db *redis.Client) {
	handler := newHealthCheckHandler(db)

	router.GET("/health", handler.Health)
	router.GET("/ready", handler.Ready)
}

func registerSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
