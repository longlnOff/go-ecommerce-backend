package manager

import "github.com/gin-gonic/gin"


type ManagerRouter struct {}



func (pr *ManagerRouter) InitUserManagerRouter(r *gin.RouterGroup) {
	// public
	managerRouterPublic := r.Group("/admin")
	{
		managerRouterPublic.POST("/login")
	}


	// private
	// managerRouterPrivate := r.Group("/admin/user")

	// userRouterPrivate.Use(Limiter())				// rate miter
	// userRouterPrivate.Use(Authentication())			// use jwt
	// userRouterPrivate.Use(Permission())				// similar to authorization
	{
		// managerRouterPrivate.POST("/active_user")

	}
}