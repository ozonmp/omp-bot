package subject

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/subject"
)

func (c *SubjectCommander) Edit(inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID
	commandArgs := strings.Split(inputMessage.CommandArguments(), " ")

	if !c.validateArgumentsCountOrSendError(
		commandArgs,
		4,
		chatID,
		"<id> <owner_id> <subject_id> <title>",
	) {
		return
	}

	subjectID, err := c.parseIdOrSendError(commandArgs[0], chatID, "to edit subject by")
	if err != nil {
		return
	}
	ownerID, err := c.parseIdOrSendError(commandArgs[1], chatID, "of owner")
	if err != nil {
		return
	}
	educationID, err := c.parseIdOrSendError(commandArgs[2], chatID, "of subject")
	if err != nil {
		return
	}

	newTitle := commandArgs[3]
	newSubject := subject.NewSubject(subjectID, ownerID, educationID, newTitle)

	err = c.subjectService.Update(subjectID, *newSubject)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Cannot edit subject by id %v: %v", subjectID, err),
		)
		c.sendMessage(msg)
		return
	}
	msg := tgbotapi.NewMessage(chatID, "Successfully updated subject")
	c.sendMessage(msg)
}
