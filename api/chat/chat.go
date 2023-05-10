package chat

import (
	"u-future-api/bot"

	"github.com/gin-gonic/gin"
)

type Chat struct {
	b *bot.Bot
}

func New(b *bot.Bot) *Chat {
	return &Chat{b}
}

func (cc *Chat) Send(ctx *gin.Context) {

}
