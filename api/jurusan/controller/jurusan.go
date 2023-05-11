package controller

import (
	"net/http"
	"u-future-api/api/jurusan/usecase"
	"u-future-api/middleware"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
)

type Jurusan struct {
	uj *usecase.Jurusan
}

func New(uj *usecase.Jurusan) *Jurusan {
	return &Jurusan{uj}
}

func (cj *Jurusan) GetAnalisis(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	analisis, err := cj.uj.GetResult(id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
	}
	response.Success(ctx, http.StatusOK, analisis)
}

func (cj *Jurusan) Mount(jurusan *gin.RouterGroup) {
	jurusan.GET("predict", middleware.ValidateJWToken(), cj.GetAnalisis)
}
