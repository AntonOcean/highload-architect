package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"ws-feed/internal/api/v1/formatter"
)

func (rH RouterHandler) GetFeed(c *gin.Context) {
	ctx := c.Request.Context()

	userID, ok := c.Get("user_id")
	if !ok {
		_ = c.Error(formatter.ErrInvalidData)
		return
	}

	ws, err := rH.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		_ = c.Error(err)
		return
	}

	rH.ucService.NewClient(ctx, userID.(uuid.UUID), ws)
}
