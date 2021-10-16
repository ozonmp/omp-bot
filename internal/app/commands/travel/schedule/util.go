package schedule

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func newPagesKeyboard(prev, next []byte) tgbotapi.InlineKeyboardMarkup {
	prevCallback := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "schedule",
		CallbackName: "list",
		CallbackData: string(prev),
	}
	nextCallback := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "schedule",
		CallbackName: "list",
		CallbackData: string(next),
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<", prevCallback.String()),
			tgbotapi.NewInlineKeyboardButtonData(">", nextCallback.String()),
		),
	)
}
