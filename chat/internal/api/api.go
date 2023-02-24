package api

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	v1 "chat/internal/api/v1"
	"chat/internal/api/v1/formatter"
	"chat/internal/usecase"
)

func New(
	log *zap.Logger,
	db *pgxpool.Pool,
	ucService usecase.ServiceUsecase,
) *gin.Engine {
	router := newGINRouter(log)

	registerSwagger(router)
	registerHealthCheck(router, db)

	routerHandler := v1.NewRouter(ucService, log)

	apiAuth := router.Group("/api/v1", routerHandler.AuthMiddleware())
	routerHandler.RegisterAuthRoutes(apiAuth)

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

func registerHealthCheck(router *gin.Engine, db *pgxpool.Pool) {
	handler := newHealthCheckHandler(db)

	router.GET("/health", handler.Health)
	router.GET("/ready", handler.Ready)
}

func registerSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
