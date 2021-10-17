package ticket

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
	"log"
)

func (c *TicketCommander) New(inputMessage *tgbotapi.Message) {
	parsedTicket := travel.Ticket{}
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &parsedTicket)
	if err != nil {
		log.Printf("Failed to parse ticket %v", inputMessage.CommandArguments())

		return
	}

	newTicketId, err := c.ticketService.Create(parsedTicket)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
		c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("New user with id '%v' was created successfully.", newTicketId),
	)
	c.bot.Send(msg)
}
