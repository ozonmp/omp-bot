package ticket

import (
	"encoding/json"
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TicketCommander) Edit(inputMessage *tgbotapi.Message) {
	splitArgs := strings.SplitN(inputMessage.CommandArguments(), ", ", 2)
	if len(splitArgs) != 2 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Ticket id and edited value must be provided"))
		c.bot.Send(msg)

		return
	}

	ticketId, err := strconv.ParseUint(splitArgs[0], 10, 64)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, `Failed to parse ticket id.`)
		c.bot.Send(msg)

		return
	}

	parsedTicket := travel.Ticket{}
	err = json.Unmarshal([]byte(splitArgs[1]), &parsedTicket)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf(`Failed to parse ticket "%v"`, splitArgs[1]))
		c.bot.Send(msg)

		return
	}

	err = c.ticketService.Update(ticketId, parsedTicket)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("User with id '%v' was edited successfully.", ticketId),
	)
	c.bot.Send(msg)
}
