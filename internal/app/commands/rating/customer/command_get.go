package customer

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/general_errors"
)

func (c *CustomerCommander) Get(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		return general_errors.NewUserError(fmt.Sprintf("wrong args %s. Err: %v", args, err))
	}

	product, err := c.customerService.Describe(uint64(idx))
	if err != nil {
		return fmt.Errorf("fail to get product with idx %d: %w", idx, err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		product.Title,
	)

	_, err = c.bot.Send(msg)
	return err
}
