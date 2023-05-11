package chat

import (
	"net/http"
	"u-future-api/bot"
	"u-future-api/middleware"
	"u-future-api/util/response"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	b *bot.Bot
}

func New(b *bot.Bot) *Chat {
	return &Chat{b}
}

func (cc *Chat) Send(ctx *gin.Context) {
	text := ctx.Query("message")
	response.Success(ctx, http.StatusOK, gin.H{"text": text})
}

func (cc *Chat) Mount(bot *gin.RouterGroup) {
	bot.GET("send", middleware.ValidateJWToken(), cc.Send)
}
