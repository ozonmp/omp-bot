package apartment

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *DummyApartmentCommander) CallbackList(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) (resp tgbotapi.MessageConfig, err error) {
	return
}
