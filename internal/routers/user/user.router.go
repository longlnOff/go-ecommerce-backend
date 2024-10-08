package user

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/internal/wire"
)


type UserRouter struct {}





func (pr *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController, err := wire.InitUserRouterHandler()
	_ = userController
	if err != nil {
		panic(err)
	}
	// public
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")

	}

	// private
	userRouterPrivate := r.Group("/user/private")

	// userRouterPrivate.Use(Limiter())				// rate miter
	// userRouterPrivate.Use(Authentication())			// use jwt
	// userRouterPrivate.Use(Permission())				// similar to authorization
	{
		userRouterPrivate.GET("/info")
		userRouterPrivate.POST("/otp")

	}
}