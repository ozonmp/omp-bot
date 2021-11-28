package film

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/paginator"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const defaultCallbackMessage = "Sorry, i cant understand this callback"

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *CinemaFilmCommander) CallbackList(ctx context.Context, callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	p := paginator.NewPaginator(callbackPath.CallbackData)
	switch p.Direction {
	case "next":
		c.List(ctx, callback.Message, p)
		return
	case "prev":
		c.List(ctx, callback.Message, p)
		return
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		defaultCallbackMessage,
	)
	_ = c.sendMessage(msg)
}