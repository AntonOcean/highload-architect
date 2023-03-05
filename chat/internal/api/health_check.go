package api

import (
	"net/http"

	"chat/internal/repository"

	"github.com/gin-gonic/gin"

	"chat/internal/api/v1/formatter"
)

type healthCheckHandler struct {
	db repository.ServiceRepository
}

func newHealthCheckHandler(db repository.ServiceRepository) *healthCheckHandler {
	return &healthCheckHandler{
		db: db,
	}
}

func (h *healthCheckHandler) Health(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusOK, gin.H{})
	}

	if err := h.db.Ping(c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, formatter.NewError("error ping DB", http.StatusInternalServerError))

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (*healthCheckHandler) Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
