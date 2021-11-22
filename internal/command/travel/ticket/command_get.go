package ticket

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TicketCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	ticketId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf(`Invalid argument "%v"`, args))
		c.bot.Send(msg)

		return
	}

	ticket, err := c.ticketService.Describe(ticketId)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ticket.String())
	c.bot.Send(msg)
}
