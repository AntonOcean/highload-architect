package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCheckHandler struct {
}

func newHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (*healthCheckHandler) Ready(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
