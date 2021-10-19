package verification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ServiceVerificationCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__service__verification — list of commands\n"+
			"/get__service__verification [ID] — get a item\n"+
			"/list__service__verification — list of items\n"+
			"/delete__service__verification [ID] — delete item\n"+
			"/new__service__verification — create new item\n"+
			"/edit__service__verification [ID Name] — edit item\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
