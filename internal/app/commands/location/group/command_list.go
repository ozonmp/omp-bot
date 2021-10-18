package group

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
	"strconv"
)

type CallbackData struct {
	PageSize   int `json:"pageSize"`
	StartIndex int `json:"startIndex"`
}

func (c *LocationGroupCommander) List(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	pageSize, err := strconv.Atoi(args)
	if err != nil {
		pageSize = 3
	}
	products := c.subdomainService.List(0, pageSize)
	msgString := "Here all the products:\n\n"
	for _, p := range products {
		msgString += p.String()
	}
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgString,
	)
	serializedData, _ := json.Marshal(CallbackData{
		PageSize: pageSize,
		StartIndex: pageSize,
	})
	callbackPath := path.CallbackPath{
		Domain:       "location",
		Subdomain:    "group",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}
	if pageSize < c.subdomainService.Size() && pageSize > 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
			),
		)
	}
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("LocationGroupCommander.List: error sending reply message to chat - %v", err)
	}
}
