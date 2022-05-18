package subject

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) List(inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID
	msgText, keyboard, err := getPaginatedMessage(c.subjectService, 0, DefaultSubjectPerPage)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Cannot /list command: %v", err),
		)
		c.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, msgText)
	msg.ReplyMarkup = keyboard
	c.sendMessage(msg)
}
