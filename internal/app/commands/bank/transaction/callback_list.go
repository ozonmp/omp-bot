package transaction

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *BankTransactionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackData {
	case "next":
		c.CursorNextPage()
		c.List(callback.Message)
	case "prev":
		c.CursorPrevPage()
		c.List(callback.Message)
	}
}

func (c *BankTransactionCommander) CursorReset() {
	c.cursor = 0
}

func (c *BankTransactionCommander) CursorNextPage() {
	c.cursor += c.limit
}

func (c *BankTransactionCommander) CursorPrevPage() {
	if c.limit > c.cursor {
		c.cursor = 0
	} else {
		c.cursor -= c.limit
	}
}
