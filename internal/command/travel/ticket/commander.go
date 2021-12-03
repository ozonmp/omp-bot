package ticket

import (
	"context"

	"github.com/ozonmp/omp-bot/internal/model/commander"

	trv_ticket_facade "github.com/ozonmp/trv-ticket-facade/pkg/trv-ticket-facade"

	"github.com/ozonmp/omp-bot/internal/service/travel"
	"github.com/ozonmp/omp-bot/internal/service/travel/ticket"
	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
)

type TicketCommander struct {
	bot           commander.Sender
	ticketService travel.TicketService
}

func NewTicketCommander(
	ctx context.Context,
	apiClient trv_ticket_api.TravelTicketApiServiceClient,
	facadeClient trv_ticket_facade.TravelTicketFacadeServiceClient,
	bot commander.Sender,
) *TicketCommander {
	ticketService := ticket.NewTravelTicketService(ctx, apiClient, facadeClient)

	return &TicketCommander{
		bot:           bot,
		ticketService: ticketService,
	}
}
