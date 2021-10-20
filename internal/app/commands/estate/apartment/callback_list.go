package apartment

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
}

func (c *DummyApartmentCommander) CallbackList(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) (resp tgbotapi.MessageConfig, err error) {
	parsedData := CallbackListData{}
	err = json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		return
	}
	resp, err = c.prepareListResponse(callback.Message.Chat.ID, parsedData.Offset)
	return
}
