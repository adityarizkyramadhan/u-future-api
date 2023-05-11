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
		return
	}
	response.Success(ctx, http.StatusOK, analisis)
}

func (cj *Jurusan) GetComparation(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	compareOne := ctx.Query("compareOne")
	compareTwo := ctx.Query("compareTwo")
	compareDataOne, err := cj.uj.GetComparationData(compareOne, id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	compareDataTwo, err := cj.uj.GetComparationData(compareTwo, id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{
		"data_one": compareDataOne,
		"data_two": compareDataTwo,
		"analysis": "Analisis OPENAI",
	})
}

func (cj *Jurusan) Mount(jurusan *gin.RouterGroup) {
	jurusan.GET("predict", middleware.ValidateJWToken(), cj.GetAnalisis)
	jurusan.GET("compare", middleware.ValidateJWToken(), cj.GetComparation)
}
