package customer

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/rating"
)

func (c *CustomerCommander) Delete(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		return rating.NewUserError("wrong args: " + args)
	}

	isRemoved, err := c.customerService.Remove(idx)
	if err != nil {
		return fmt.Errorf("fail to get customer with idx %d: %w", idx, err)
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
