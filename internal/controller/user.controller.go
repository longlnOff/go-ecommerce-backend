package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/internal/service"
	"github.com/longln/go-ecommerce-backend-api/pkg/response"
)



type UserController struct {
	userService service.IUserService
}


func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserInforController(c *gin.Context) {
	code := uc.userService.Register("", "")
	response.SuccessResponse(c, code, nil)
}



