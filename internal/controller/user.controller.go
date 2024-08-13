package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/service"
	"github.com/longln/go-ecommerce-backend/pkg/response"
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
	// if err != nil {
		// return response.ErrorResponse(c, response.ErrCodeParamInvalid)
	// }
	// return response.SuccessResponse(c, response.ErrCodeSuccess, []string{"longln", "hello"})
}

// Note: gin.H is a map string