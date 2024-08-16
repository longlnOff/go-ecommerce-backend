package manager

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/wire"
)

type UserRouter struct{

}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	// public router
	// // this is non-dependency
	// urepo := repo.NewUserRepository()
	// us := service.NewUserService(urepo)
	// userHandlerNonDepend := controller.NewUserController(us)
	// _ = userHandlerNonDepend
	// WIRE golang - Dependency Injection (DI)
	// Inversion of Control (IoC)
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := r.Group("/admin/user")
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authentication())
		// userRouterPrivate.Use(Authorization())
	{
		userRouterPrivate.POST("/active_user")
	}
}