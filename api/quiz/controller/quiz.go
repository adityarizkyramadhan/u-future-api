package controller

import (
	"net/http"
	"u-future-api/api/quiz/usecase"
	"u-future-api/middleware"
	"u-future-api/models"
	"u-future-api/util/exception"
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
	if name == "SectionOne" {
		quiz, err := qc.qu.FindByName(name)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, quiz)
	} else if name == "SectionTwo" {
		id := ctx.MustGet("id").(string)
		quiz, err := qc.qu.SectionTwoQuiz(id)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, quiz)
	} else if name == "SectionThree" {
		id := ctx.MustGet("id").(string)
		quiz, err := qc.qu.SectionThreeQuiz(id)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, quiz)
	} else {
		response.Fail(ctx, http.StatusBadRequest, exception.ErrNoQuery.Error())
	}
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

func (qc *Quiz) AttemptQuiz(ctx *gin.Context) {
	name := ctx.Query("title")
	id := ctx.MustGet("id").(string)
	if name == "SectionOne" {
		var data []models.InputQuizString
		if err := ctx.Bind(&data); err != nil {
			response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		}
		if err := qc.qu.UpdateResult(name, data, id); err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, gin.H{
			"quiz_attempt": name,
		})
	} else if name == "SectionTwo" {
		var data []models.InputQuizInteger
		if err := ctx.Bind(&data); err != nil {
			response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		}
		if err := qc.qu.UpdateResult(name, data, id); err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, gin.H{
			"quiz_attempt": name,
		})
	} else if name == "SectionThree" {
		var data []models.InputQuizInteger
		if err := ctx.Bind(&data); err != nil {
			response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		}
		if err := qc.qu.UpdateResult(name, data, id); err != nil {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		response.Success(ctx, http.StatusOK, gin.H{
			"quiz_attempt": name,
		})
	}
}

func (qc *Quiz) Mount(quiz *gin.RouterGroup) {
	quiz.GET("status-quiz", middleware.ValidateJWToken(), qc.IsUserAttemptQuiz)
	quiz.GET("question", middleware.ValidateJWToken(), qc.FindByName)
	quiz.POST("attempt-quiz", middleware.ValidateJWToken(), qc.AttemptQuiz)
}
