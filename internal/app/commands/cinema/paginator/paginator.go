package paginator

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const defaultPageLength = 10

var callbackPathNext = path.CallbackPath{
	Domain:       "cinema",
	Subdomain:    "film",
	CallbackName: "list",
	CallbackData: "next",
}

var callbackPathPrev = path.CallbackPath{
	Domain:       "cinema",
	Subdomain:    "film",
	CallbackName: "list",
	CallbackData: "prev",
}

var keyBoard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Next", callbackPathNext.String()),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Prev", callbackPathPrev.String()),
	),
)

type CinemaPaginator struct {
	CurrentPage int
	PageLength int
	Keyboard tgbotapi.InlineKeyboardMarkup
}

func NewCinemaPaginator() *CinemaPaginator {
	return &CinemaPaginator{CurrentPage: 0, PageLength: defaultPageLength, Keyboard: keyBoard}
}
