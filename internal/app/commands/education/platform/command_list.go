package platform

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

func (c *PlatformBaseCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := "All platforms are here: \n\n"

	platforms, err := c.service.List(0, DefaultListLimit)
	if err != nil {
		log.Printf(err.Error())

		return
	}

	for _, p := range platforms {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	if uint64(len(platforms)) == DefaultListLimit {
		serializedData, err := json.Marshal(
			CallbackListData{
				Cursor: DefaultListLimit,
				Limit:  DefaultListLimit,
			},
		)

		if err != nil {
			log.Printf(err.Error())

			return
		}

		callbackPath := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "platform",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(NextButtonText, callbackPath.String()),
			),
		)
	}

	c.sendMessage(msg)
}
