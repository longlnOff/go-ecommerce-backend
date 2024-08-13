package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
	  v1.GET("/ping", controller.NewPongController().Pong)	// --> /v1/2024/ping
	  v1.GET("/user", controller.NewUserController().GetUserByID)
	}

	return r
}



