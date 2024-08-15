package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{

}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// public router
	// userRouterPublic := r.Group("/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.POST("/otp")
	// }

	// private router
	userRouterPrivate := r.Group("/admin/user")
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authentication())
		// userRouterPrivate.Use(Authorization())
	{
		userRouterPrivate.POST("/active_user")
	}
}