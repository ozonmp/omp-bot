package course

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *WorkCourseCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "next":
		c.cursor += c.limit
		c.List(callback.Message)
	case "prev":
		c.cursor -= c.limit
		if c.cursor < 0 {
			c.cursor = 0
		}
		c.List(callback.Message)
	}
}
