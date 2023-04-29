package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, Response{
		Status:  status,
		Message: "success",
		Data:    data,
	})
}

func Fail(ctx *gin.Context, status int, errorMessage string) {
	ctx.JSON(status, Response{
		Status:  status,
		Message: errorMessage,
		Data:    nil,
	})
}
