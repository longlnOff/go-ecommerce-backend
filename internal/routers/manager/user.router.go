package manager

import "github.com/gin-gonic/gin"


type UserRouter struct {}



func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// public
	// userRouterPublic := r.Group("/admin/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")

	// }


	// private
	userRouterPrivate := r.Group("/admin/user")

	// userRouterPrivate.Use(Limiter())				// rate miter
	// userRouterPrivate.Use(Authentication())			// use jwt
	// userRouterPrivate.Use(Permission())				// similar to authorization
	{
		userRouterPrivate.POST("/active_user")

	}
}