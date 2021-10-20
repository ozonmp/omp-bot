package apartment

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const pageSize = 5

func (c *DummyApartmentCommander) prepareListResponse(chatID int64, offset uint64) (
	resp tgbotapi.MessageConfig,
	err error,
) {
	apartments, err := c.service.List(offset, pageSize)
	if err != nil {
		return
	}
	var outputMsgText string
	for _, a := range apartments {
		outputMsgText += a.String()
		outputMsgText += "\n"
	}

	if len(apartments) < pageSize {
		// we reached end of the list so pagination is not needed

		resp = tgbotapi.NewMessage(chatID, outputMsgText)
	} else {
		// setup pagination

		var serializedData []byte
		serializedData, err = json.Marshal(CallbackListData{
			Offset: offset + pageSize,
		})
		if err != nil {
			return
		}

		callbackPath := path.CallbackPath{
			Domain:       "estate",
			Subdomain:    "apartment",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		resp = tgbotapi.NewMessage(chatID, outputMsgText)
		resp.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}
	return
}

func (c *DummyApartmentCommander) List(inputMessage *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error) {
	resp, err = c.prepareListResponse(inputMessage.Chat.ID, 0)
	return
}
