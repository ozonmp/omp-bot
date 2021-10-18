package ground

import (
	"errors"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Delete(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 64)
	if err != nil {
		return errors.New("Wrong args: " + args)
	}

	isRemoved, err := c.groundService.Remove(idx)
	if err != nil {
		return fmt.Errorf("Fail to get auto with idx %d: %w", idx, err)
	}

	if isRemoved {
		c.Send(
			inputMessage.Chat.ID,
			args+" was removed",
		)
	}
	return nil
}
