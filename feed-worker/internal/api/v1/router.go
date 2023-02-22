package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"feed-worker/internal/usecase"
)

type RouterHandler struct {
	ucService usecase.ServiceUsecase
	logger    *zap.Logger
}

func NewRouter(i usecase.ServiceUsecase, logger *zap.Logger) RouterHandler {
	return RouterHandler{
		ucService: i,
		logger:    logger,
	}
}

func (rH RouterHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("user/:user_id/feed", rH.GetFeed)
}
