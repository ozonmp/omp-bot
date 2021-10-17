package office

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const ListLimit = 2

func (c *OfficeCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Entity list page 1: \n\n"

	entities, err := c.officeService.List(0, ListLimit)

	for _, e := range entities {
		outputMsgText += e.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: ListLimit,
		Limit:  ListLimit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "business",
		Subdomain:    "office",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("OfficeCommander.List: error sending reply message to chat - %v", err)
	}
}
