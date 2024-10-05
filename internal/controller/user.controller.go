package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/internal/service"
	"github.com/longln/go-ecommerce-backend-api/pkg/response"
)



type UserController struct {
	userService service.UserService
}


func NewUserController() *UserController {
	return &UserController{
		userService: *service.NewUserService(),
	}
}

func (uc *UserController) GetUserInforController(c *gin.Context) {
	fmt.Println("Handler")
	data := uc.userService.GetUserInfoService()
	response.SuccessResponse(c, 20001, data)
}


