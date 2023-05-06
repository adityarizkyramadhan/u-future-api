package controller

import (
	"net/http"
	"u-future-api/api/school/usecase"
	"u-future-api/models"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
)

type School struct {
	us *usecase.School
}

func New(us *usecase.School) *School {
	return &School{us}
}

func (cs *School) FindPagination(ctx *gin.Context) {
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	name := ctx.Query("name")
	data, err := cs.us.FindPagination(limit, page, name)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, data)
}

func (cs *School) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := cs.us.FindById(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, data)
}

func (cs *School) Create(ctx *gin.Context) {
	var input models.SchoolInput
	if err := ctx.BindJSON(&input); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if err := cs.us.Create(&input, ""); err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusCreated, gin.H{
		"school": input.Name,
	})
}
func (cs *School) Mount(school *gin.RouterGroup) {
	school.GET("search", cs.FindPagination)
	school.GET("single/:id", cs.FindById)
	school.POST("", cs.Create)
}
