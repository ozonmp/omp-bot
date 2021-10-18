package group

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LocationGroupCommander) CommandNotImplemented(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Command "+inputMessage.Text+" not implemented")
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("HelpLocationCommander.Help: error sending reply message to chat - %v", err)
	}
}
