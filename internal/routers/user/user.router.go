package user

import "github.com/gin-gonic/gin"


type UserRouter struct {}



func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// public
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")

	}


	// private
	userRouterPrivate := r.Group("/user")

	// userRouterPrivate.Use(Limiter())				// rate miter
	// userRouterPrivate.Use(Authentication())			// use jwt
	// userRouterPrivate.Use(Permission())				// similar to authorization
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.POST("/otp")

	}
}