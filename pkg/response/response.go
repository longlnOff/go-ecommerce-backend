package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type ResponseData struct {
	Code int 		`json:"code"`
	Message string 	`json:"message"`
	Data any 		`json:"data"` 
}

func SuccessResponse(c *gin.Context, code int, data any) {
	c.JSON(
		http.StatusOK,
		ResponseData{
			Code: code,
			Message: Msg[code],
			Data: data,
		},
	)
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(
		http.StatusInternalServerError,
		ResponseData{
			Code: code,
			Message: Msg[code],
			Data: nil,
		},
	)
}