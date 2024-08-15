package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{

}

func (ar *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {
	// public router
	adminRouterPublic := r.Group("/admin")
	{
		adminRouterPublic.POST("/login")
		// userRouterPublic.POST("/otp")
	}

	// private router
	adminRouterPrivate := r.Group("/admin/user")
	{
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authentication())
		// userRouterPrivate.Use(Authorization())

		adminRouterPrivate.POST("/active_user")

	}
}