package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			// abort
			c.Abort()
			response.ErrorResponse(c, response.ErrInvalidToken)
			return
		} else {
			c.Next()
		}
	}
}