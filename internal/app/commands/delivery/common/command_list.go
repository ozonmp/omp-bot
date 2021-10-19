package common

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c DummyCommonCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := "***Here are deliveries***:\n\n"

	cursor := 0
	deliveries, _ := c.commonService.List(uint64(cursor), uint64(pageSize))
	for _, delivery := range deliveries {
		outputMsgText += delivery.String() + "\n\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)
	msg.ParseMode = "markdown"
	generatePaginationButtons(cursor, &c, &msg)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyCommonCommander.List: error sending reply message to chat - %v", err)
	}
}
