package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/global"
	"github.com/longln/go-ecommerce-backend/internal/routers"
)


func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Sever.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()	// don't write log
	}

	// moddlewares
	r.Use()	// logger
	r.Use()	// cross
	r.Use()	// limiter global

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("checkStatus")	// tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}

	return r
}



