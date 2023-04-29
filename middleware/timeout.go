package middleware

import (
	"net/http"
	"time"
	"u-future-api/util/exception"
	"u-future-api/util/response"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(second int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(second)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			response.Fail(c, http.StatusRequestTimeout, exception.ErrTimeout.Error())
		}),
	)
}
