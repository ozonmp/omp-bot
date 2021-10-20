package order

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OrderCommander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("Unknown command from [%s]: %s", inputMessage.From.UserName, inputMessage.Text)
	c.Help(inputMessage)
}
