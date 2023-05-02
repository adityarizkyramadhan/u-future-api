package controller

import (
	"net/http"
	"u-future-api/api/student/usecase"
	"u-future-api/models"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
)

type Student struct {
	us *usecase.Student
}

func New(us *usecase.Student) *Student {
	return &Student{us}
}

func (cs *Student) Login(ctx *gin.Context) {
	var input models.StudentLogin
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := cs.us.Login(&input)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}

func (cs *Student) Register(ctx *gin.Context) {
	var input models.StudentRegister
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := cs.us.Register(&input)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, gin.H{
		"token": token,
	})
}

func (cs *Student) Mount(student *gin.RouterGroup) {
	student.POST("login", cs.Login)
	student.POST("register", cs.Register)
}
