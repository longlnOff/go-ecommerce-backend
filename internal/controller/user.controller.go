package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/service"
)


type UserController struct{
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetInforUserService(),
	})
}