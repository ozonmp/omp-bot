package employee

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *DemoSubdomainCommander) Create(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, c.subdomainService.Create(args))

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
