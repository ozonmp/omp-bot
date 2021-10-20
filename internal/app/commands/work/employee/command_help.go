package employee

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoSubdomainCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__employee - help\n"+
			"/list__work__employee - list work employees",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
