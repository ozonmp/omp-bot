package ticket

import (
	"context"

	"github.com/ozonmp/omp-bot/internal/service/travel"
	"github.com/ozonmp/omp-bot/internal/service/travel/ticket"
	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
)

type TicketCommander struct {
	bot           Sender
	ticketService travel.TicketService
}

func NewTicketCommander(
	ctx context.Context,
	client trv_ticket_api.TravelTicketApiServiceClient,
	bot Sender,
) *TicketCommander {
	ticketService := ticket.NewTravelTicketService(ctx, client)

	return &TicketCommander{
		bot:           bot,
		ticketService: ticketService,
	}
}
