package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseSuccessWithDataModel struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseErrorCustomModel struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func ResponseSuccessWithData(c *gin.Context, data interface{}) {
	response := ResponseSuccessWithDataModel{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseSuccess(c *gin.Context) {
	response := ResponseSuccessModel{
		Success: true,
		Message: "Success",
	}

	c.JSON(http.StatusOK, response)
}

func ResponseErrorCustom(c *gin.Context, err interface{}, message string, code int) {
	response := ResponseErrorCustomModel{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(code, response)
}
