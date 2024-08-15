package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			// stop
			response.ErrorResponse(c, response.ErrCodeTokenInvalid)
			// can't authorize --> abort
			c.Abort()
		}
		c.Next()
	}
}