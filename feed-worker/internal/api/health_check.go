package api

import (
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"

	"feed-worker/internal/api/v1/formatter"
)

type healthCheckHandler struct {
	db *redis.Client
}

func newHealthCheckHandler(db *redis.Client) *healthCheckHandler {
	return &healthCheckHandler{
		db: db,
	}
}

func (h *healthCheckHandler) Health(c *gin.Context) {
	if err := h.db.Ping(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, formatter.NewError("error ping DB", http.StatusInternalServerError))

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (*healthCheckHandler) Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
