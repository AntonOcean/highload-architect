package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"kek/internal/usecase"
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
	router.POST("/login", rH.AuthUser)

	router.GET("/user", rH.GetUsers)
	router.POST("/user", rH.CreateUser)
	router.GET("/user/:id", rH.GetUserByID)
}
