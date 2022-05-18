package subject

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/education/subject"
)

func (c *SubjectCommander) New(inputMessage *tgbotapi.Message) {
	chatID := inputMessage.Chat.ID
	commandArgs := strings.Split(inputMessage.CommandArguments(), " ")

	if !c.validateArgumentsCountOrSendError(
		commandArgs,
		3,
		chatID,
		"<owner_id> <subject_id> <title>",
	) {
		return
	}

	ownerID, err := c.parseIdOrSendError(commandArgs[0], chatID, "of owner")
	if err != nil {
		return
	}
	subjectID, err := c.parseIdOrSendError(commandArgs[1], chatID, "of subject")
	if err != nil {
		return
	}

	title := commandArgs[2]

	newSubject := subject.NewSubject(0, ownerID, subjectID, title)
	newID, err := c.subjectService.Create(*newSubject)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Couldn't create new subject: %v", err),
		)
		c.sendMessage(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		chatID,
		fmt.Sprintf("Successfully created new subject, ID is %v", newID),
	)
	c.sendMessage(msg)
}
