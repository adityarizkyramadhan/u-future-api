package controller

import (
	"net/http"
	"u-future-api/api/quiz/usecase"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
)

type Quiz struct {
	qu *usecase.Quiz
}

func New(qu *usecase.Quiz) *Quiz {
	return &Quiz{qu}
}

func (qc *Quiz) FindByName(ctx *gin.Context) {
	name := ctx.Param("name")
	quiz, err := qc.qu.FindByName(name)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, quiz)
}
