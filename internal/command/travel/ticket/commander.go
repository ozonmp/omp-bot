package ticket

import (
	"github.com/ozonmp/omp-bot/internal/service/travel"
)

type TicketCommander struct {
	bot           Sender
	ticketService travel.TicketService
}

func NewTicketCommander(
	bot Sender,
) *TicketCommander {
	ticketService := travel.NewTravelTicketService()

	return &TicketCommander{
		bot:           bot,
		ticketService: ticketService,
	}
}
