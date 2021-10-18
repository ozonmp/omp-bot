package ground

import (
	"errors"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Get(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		return errors.New(fmt.Sprintf("Wrong args %s. \nErr: %v", args, err))
	}

	product, err := c.groundService.Describe(uint64(idx))
	if err != nil {
		return fmt.Errorf("Fail to get auto with idx %d: \n%w", idx, err)
	}

	c.Send(
		inputMessage.Chat.ID,
		product.String(),
	)
	return nil
}
