package initialize

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/global"
	"github.com/longln/go-ecommerce-backend-api/internal/routers"
	// "github.com/longln/go-ecommerce-backend-api/internal/controller"
	// "github.com/longln/go-ecommerce-backend-api/internal/middlewares"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()		// write log
	} else {
		gin.SetMode(gin.ReleaseMode)

		r = gin.New()			// don't write log
	}


	// middlewares
	// r.Use(middlewares.Cors())
	// r.Use(middlewares.Logger())
	// r.Use(middlewares.RateLimiter())			// global rate limiter

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("v1/2024")
	{
		mainGroup.GET("checkStatus")

	}
	{
		managerRouter.InitUserManagerRouter(mainGroup)
		managerRouter.InitUserRouter(mainGroup)
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	return r
}