package provider

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *PaymentProviderCommander) List(inputMessage *tgbotapi.Message) {
	var outputMsgText string
	entCnt := c.providerService.EntitiesCount()
	pageCnt := entCnt / limit

	if entCnt%limit != 0 {
		pageCnt++
	}

	providers := c.providerService.List(0, limit)
	for _, p := range providers {
		outputMsgText += fmt.Sprintf("%s\n", c.providerService.ShortDescription(&p))
	}
	outputMsgText += fmt.Sprintf("<1/%d>", pageCnt)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "payment",
		Subdomain:    "provider",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.List: error sending reply message to chat - %v", err)
	}
}
