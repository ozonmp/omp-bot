package subject

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) sendMessage(msg tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat %v: %v", msg.ChatID, err)
	}
}

func (c *SubjectCommander) editMessage(msg tgbotapi.EditMessageTextConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("Error sending reply message to chat %v: %v", msg.ChatID, err)
	}
}
