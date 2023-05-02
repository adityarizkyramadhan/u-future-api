package response

import "github.com/gin-gonic/gin"

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, status int, data any) {
	meta := Meta{
		Status:  status,
		Message: "success",
	}
	response := Response{
		Meta: meta,
		Data: data,
	}
	ctx.JSON(status, response)
}

func Fail(ctx *gin.Context, status int, errorMessage string) {
	meta := Meta{
		Status:  status,
		Message: errorMessage,
	}
	response := Response{
		Meta: meta,
		Data: nil,
	}
	ctx.JSON(status, response)
}
