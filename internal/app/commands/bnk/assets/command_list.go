package assets

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *AssetsCommander) List(inputMessage *tgbotapi.Message) {

	c.assetsService.CurrentPage = 0
	outputMsgText := "Список всех активов: \n\n"

	assets := c.assetsService.List()
	for _, p := range assets {
		outputMsgText += p.String() + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: 21,
	})

	callbackPath := path.CallbackPath{
		Domain:       "bnk",
		Subdomain:    "assets",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Previous page", callbackPath.String()),
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.List: error sending reply message to chat - %v", err)
	}
}
