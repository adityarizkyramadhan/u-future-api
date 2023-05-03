package controller

import (
	"net/http"
	"u-future-api/api/school/usecase"
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
func (cs *School) Mount(student *gin.RouterGroup) {
	student.GET("search", cs.FindPagination)
	student.GET("single/:id", cs.FindById)
}
