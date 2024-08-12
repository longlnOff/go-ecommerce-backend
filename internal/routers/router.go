package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1/2024")
	{
	  v1.GET("/ping", pong)	// --> /v1/2024/ping
	  v1.GET("/ping1", pong)
	  v1.GET("/ping2", pong)
	
	}

	return r
}



func pong (c *gin.Context) {
  name := c.DefaultQuery("name", "longln")
  uid := c.Query("uid")
  c.JSON(http.StatusOK, gin.H{
	"message": "pong",
	"name": name,
	"uid": uid,
  })
}