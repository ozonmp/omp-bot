package subject

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) Delete(inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID
	commandArgs := inputMessage.CommandArguments()
	subjectID, err := c.parseIdOrSendError(commandArgs, chatID, "to delete subject by")
	if err != nil {
		return
	}

	_, err = c.subjectService.Remove(subjectID)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("cannot delete subject by id %v: %v", subjectID, err),
		)
		c.sendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		chatID,
		"Successfully removed subject",
	)
	c.sendMessage(msg)
}
