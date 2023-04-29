package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"
	"u-future-api/util/exception"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(3 * 24 * time.Hour).Unix(),
	})
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			response.Fail(c, http.StatusUnauthorized, exception.ErrMissingBearerToken.Error())
			return
		}
		bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
		token, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			response.Fail(c, http.StatusUnauthorized, err.Error())
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := uint(claims["id"].(float64))
			c.Set("id", userId)
			c.Next()
		} else {
			response.Fail(c, http.StatusUnauthorized, "token invalid")
			return
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(os.Getenv("SECRET_KEY")), nil
}
