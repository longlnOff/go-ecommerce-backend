package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend-api/internal/service"
	"github.com/longln/go-ecommerce-backend-api/internal/vo"
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

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid)
		return
	}
	fmt.Printf("Email params: %v", params.Email)
	code := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, code, nil)
}



