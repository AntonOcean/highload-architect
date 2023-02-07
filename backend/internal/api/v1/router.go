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

func (rH RouterHandler) RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("friend", rH.CreateFriend)
	router.DELETE("friend/:id", rH.DeleteFriendByID)
}

// router.GET("post", rH.GetPosts)
// router.POST("post", rH.CreatePost)
// router.PUT("post/:id", rH.EditPost)
// router.GET("post/:id", rH.GetPostByID)
// router.DELETE("post/:id", rH.DeletePostByID)
