package user

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/wire"
)

type UserRouter struct{

}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()
	// public router
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := r.Group("/user")
	{
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authentication())
		// userRouterPrivate.Use(Authorization())

		userRouterPrivate.GET("/info")

	}
}