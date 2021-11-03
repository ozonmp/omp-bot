package assets

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *AssetsCommander) List(inputMessage *tgbotapi.Message) {
	assets := c.assetsService.List(1)

	outputMsgText := "Всего активов " + fmt.Sprintf("%d",c.assetsService.Count())

	msg := tgbotapi.MessageConfig{}

	if c.assetsService.Count() > 10 {
		outputMsgText += ", страница 1 из " +
			fmt.Sprintf("%d", c.assetsService.PageCount()) +
			": \n\n"

		for _, p := range assets {
			outputMsgText += p.String() + "\n"
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

		nextSerializedData, _ := json.Marshal(CallbackListData{
			Page: 2,
		})

		callbackPath := path.CallbackPath{
			Domain:       "bnk",
			Subdomain:    "assets",
			CallbackName: "list",
			CallbackData: string(nextSerializedData),
		}

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	} else {
		outputMsgText += ": \n\n"

		for _, p := range assets {
			outputMsgText += p.String() + "\n"
		}

		msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	}

	_, err := c.Bot.Send(msg)
	if err != nil {
		log.Printf("AssetsCommander.List: error sending reply message to chat - %v", err)
	}
}
