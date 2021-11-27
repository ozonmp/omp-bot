package travel

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/command/travel/ticket"
	trv_ticket_api "github.com/ozonmp/trv-ticket-api/pkg/trv-ticket-api"
)

type TicketCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type TicketCallbacker interface {
	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

func NewTicketCommander(
	ctx context.Context,
	client trv_ticket_api.TravelTicketApiServiceClient,
	bot Sender) TicketCommander {
	return ticket.NewTicketCommander(ctx, client, bot)
}

func NewTicketCallbacker(
	ctx context.Context,
	client trv_ticket_api.TravelTicketApiServiceClient,
	bot Sender) TicketCallbacker {
	return ticket.NewTicketCommander(ctx, client, bot)
}
