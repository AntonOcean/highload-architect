package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"ws-feed/internal/usecase"
)

type RouterHandler struct {
	ucService usecase.ServiceUsecase
	logger    *zap.Logger
	upgrader  *websocket.Upgrader
}

func NewRouter(i usecase.ServiceUsecase, logger *zap.Logger) RouterHandler {
	return RouterHandler{
		ucService: i,
		logger:    logger,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (rH RouterHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("feed", rH.GetFeed)
}
