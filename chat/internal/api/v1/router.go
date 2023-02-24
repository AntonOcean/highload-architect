package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"chat/internal/usecase"
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

func (rH RouterHandler) RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("dialog/:receiver_user_id", rH.CreateMessage)
	router.GET("dialog/:receiver_user_id", rH.GetMessages)
}
