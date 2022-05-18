package subject

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Unknown command %s, type /help__education_subject for help", inputMessage.Text),
	)
	c.sendMessage(msg)
}
