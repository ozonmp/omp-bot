package customer

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/general_errors"
)

func (c *CustomerCommander) Delete(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseInt(args, 10, 64)
	if err != nil {
		return general_errors.NewUserError("wrong args: " + args)
	}

	isRemoved, err := c.customerService.Remove(uint64(idx))
	if err != nil {
		return fmt.Errorf("fail to get product with idx %d: %w", idx, err)
	}

	if isRemoved {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			args+" was removed",
		)
		_, err := c.bot.Send(msg)
		return err
	}
	return nil
}
