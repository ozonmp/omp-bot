package subject

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) Get(inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()

	subjectID, err := c.parseIdOrSendError(commandArgs, chatID, "to get subject by id")
	if err != nil {
		return
	}

	subject, err := c.subjectService.Describe(subjectID)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Cannot describe subject by id, id: %v, err: %v", subjectID, err),
		)
		c.sendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		chatID,
		subject.String(),
	)
	c.sendMessage(msg)
}
