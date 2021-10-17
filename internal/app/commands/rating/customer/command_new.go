package customer

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/rating"
	"github.com/ozonmp/omp-bot/internal/service/rating/customer"
)

const argumentDelimiter = " "

func (c *CustomerCommander) New(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	fields := strings.Split(args, argumentDelimiter)
	if len(fields) != 2 {
		return rating.NewUserError(fmt.Sprintf("Required 2 fields, got %d", len(fields)))
	}

	customer := customer.NewCustomer(fields[0], fields[1])
	idx, err := c.customerService.Create(customer)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Customer %s with index %d was successfully created!", customer, idx),
	)

	_, err = c.bot.Send(msg)
	return err
}
