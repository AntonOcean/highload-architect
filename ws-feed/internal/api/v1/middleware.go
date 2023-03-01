package v1

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"ws-feed/internal/domain"
)

func (rH RouterHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" || !strings.HasPrefix(auth, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimSpace(strings.Split(auth, "Bearer")[1])

		data, err := rH.ucService.GetTokenData(context.Background(), token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if data.TokenType != string(domain.Access) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", data.UserID)
	}
}
