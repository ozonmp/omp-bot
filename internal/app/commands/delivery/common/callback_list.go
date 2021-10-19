package common

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *DummyCommonCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("DummyCommonCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	outputMsgText := "***Here are deliveries***:\n\n"

	cursor := parsedData.Offset
	deliveries, _ := c.commonService.List(uint64(cursor), uint64(pageSize))
	for _, delivery := range deliveries {
		outputMsgText += delivery.String() + "\n\n"
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	msg.ParseMode = "markdown"
	generatePaginationButtons(cursor, c, &msg)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
