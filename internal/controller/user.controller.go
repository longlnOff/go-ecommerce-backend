package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-ecommerce-backend/internal/service"
	"github.com/longln/go-ecommerce-backend/pkg/response"
)

// type UserController struct {
// 	userService *service.UserService
// }

// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// if err != nil {
// 	fmt.Println("---> My Handler")
// 	response.ErrorResponse(c, response.ErrCodeParamInvalid)
// 	// }
// 	// return response.SuccessResponse(c, response.ErrCodeSuccess, []string{"longln", "hello"})
// }

// Note: gin.H is a map string


// Interface version
type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}


func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}