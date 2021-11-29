package ticket

import (
	"context"
	"log"

	trv_ticket_facade "github.com/ozonmp/trv-ticket-facade/pkg/trv-ticket-facade"

	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/command/travel"
)

type TravelTicketCommander struct {
	bot              Sender
	ticketCommander  travel.TicketCommander
	ticketCallbacker travel.TicketCallbacker
}

func NewTravelTicketCommander(
	ctx context.Context,
	apiClient trv_ticket_api.TravelTicketApiServiceClient,
	facadeClient trv_ticket_facade.TravelTicketFacadeServiceClient,
	bot Sender,
) *TravelTicketCommander {

	return &TravelTicketCommander{
		bot:              bot,
		ticketCommander:  travel.NewTicketCommander(ctx, apiClient, facadeClient, bot),
		ticketCallbacker: travel.NewTicketCallbacker(ctx, apiClient, facadeClient, bot),
	}
}

func (c *TravelTicketCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.ticketCallbacker.CallbackList(callback, callbackPath)
	default:
		log.Printf("TravelTicketCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TravelTicketCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.ticketCommander.Help(msg)
	case "list":
		c.ticketCommander.List(msg)
	case "get":
		c.ticketCommander.Get(msg)
	case "delete":
		c.ticketCommander.Delete(msg)
	case "new":
		c.ticketCommander.New(msg)
	case "edit":
		c.ticketCommander.Edit(msg)
	default:
		log.Printf("TravelTicketCommander.HandleCommand: unknown callback name: %s", commandPath.CommandName)
	}
}
