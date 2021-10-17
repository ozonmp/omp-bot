package ticket

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TicketCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	ticketId, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf(`Invalid argument "%v"`, args))
		c.bot.Send(msg)

		return
	}

	_, err = c.ticketService.Remove(ticketId)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "User was deleted successfully.")
	c.bot.Send(msg)
}
