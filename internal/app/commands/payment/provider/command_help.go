package provider

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PaymentProviderCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__payment__provider - shows this help\n"+
			"/get__payment__provider - get provider by id\n"+
			"/list__payment__provider - get all provider\n"+
			"/new__payment__provider - create new payment provider\n"+
			"/delete__payment__provider - deletes provider by id\n"+
			"/edit__payment__provider - edit provider attributes",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Help: error sending reply message to chat - %v", err)
	}
}
