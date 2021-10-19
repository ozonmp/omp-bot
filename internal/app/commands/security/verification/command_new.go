package verification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SecurityVerificationCommander) New(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"new operation does not support yet",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SecurityVerificationCommander.New: error sending reply message to chat - %v", err)
	}
}
