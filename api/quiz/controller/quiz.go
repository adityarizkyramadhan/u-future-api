package controller

import (
	"net/http"
	"u-future-api/api/quiz/usecase"
	"u-future-api/middleware"
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
	name := ctx.Query("title")
	quiz, err := qc.qu.FindByName(name)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, quiz)
}

func (qc *Quiz) IsUserAttemptQuiz(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	result, err := qc.qu.SearchTestUser(id)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{
		"already_taken": result,
	})
}

func (qc *Quiz) Mount(quiz *gin.RouterGroup) {
	quiz.GET("/status-quiz", middleware.ValidateJWToken(), qc.IsUserAttemptQuiz)
	quiz.GET("/question", middleware.ValidateJWToken(), qc.FindByName)
}
