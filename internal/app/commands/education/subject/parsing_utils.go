package subject

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubjectCommander) parseIdOrSendError(strID string, chatID int64, IDdescription string) (uint64, error) {
	id, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("Couldn't parse id %v: %v", IDdescription, err),
		)
		c.sendMessage(msg)
		return 0, err
	}

	return id, nil
}

func (c *SubjectCommander) validateArgumentsCountOrSendError(arguments []string, expectedLen int, chatID int64, argsDescription string) bool {
	if len(arguments) != expectedLen {
		msg := tgbotapi.NewMessage(
			chatID,
			fmt.Sprintf("expected %v arguments: %v, but received %v: %v", expectedLen, argsDescription, len(arguments), arguments),
		)
		c.sendMessage(msg)
		return false
	}

	return true
}
