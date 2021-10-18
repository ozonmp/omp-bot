package internship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *WorkInternshipCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "next":
		c.CursorNextPage()
		c.List(callback.Message)
	case "prev":
		c.CursorPrevPage()
		c.List(callback.Message)
	}
}

func (c *WorkInternshipCommander) CursorReset() {
	c.cursor = 0
}

func (c *WorkInternshipCommander) CursorNextPage() {
	c.cursor += c.limit
}

func (c *WorkInternshipCommander) CursorPrevPage() {
	c.cursor -= c.limit
	if c.cursor < 0 {
		c.cursor = 0
	}
}
