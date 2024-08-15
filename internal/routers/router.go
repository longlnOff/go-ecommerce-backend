package routers

// import (
// 	"fmt"

// 	"github.com/gin-gonic/gin"
// 	"github.com/longln/go-ecommerce-backend/internal/controller"
// 	"github.com/longln/go-ecommerce-backend/internal/middleware"
// )

// func AA() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before AA")
// 		c.Next()
// 		fmt.Println("After AA")
// 	}
// }

// func BB() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("Before BB")
// 		c.Next()
// 		fmt.Println("After BB")
// 	}
// }

// func CC(c * gin.Context) {
// 	fmt.Println("Before CC")
// 	c.Next()
// 	fmt.Println("After CC")
// }

// func NewRouter() *gin.Engine {
// 	r := gin.Default()

// 	// use middleware
// 	r.Use(middleware.AuthenMiddleware(), BB(), CC)

// 	v1 := r.Group("/v1/2024")
// 	{
// 	  v1.GET("/ping", controller.NewPongController().Pong)	// --> /v1/2024/ping
// 	  v1.GET("/user", controller.NewUserController().GetUserByID)
// 	}

// 	return r
// }



