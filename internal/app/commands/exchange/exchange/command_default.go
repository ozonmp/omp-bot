package exchange

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubdomainCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"To see the list of commands type /help__exchange__exchange",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
