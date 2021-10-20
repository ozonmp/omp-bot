package provider

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PaymentProviderCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__payment__provider - shows this help\n"+
			"/list__payment__provider - get all providers,\n"+
			"/get__payment__provider <id>- get provider by id\n"+
			"/new__payment__provider <jsonString> - create new provider\n"+
			"/delete__payment__provider <id> - deletes provider by id\n"+
			"/edit__payment__provider <jsonString>- edit provider attributes",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Help: error sending reply message to chat - %v", err)
	}
}
