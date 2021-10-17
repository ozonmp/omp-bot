package travel

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/command/travel/ticket"
)

type TicketCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type TicketCallbacker interface {
	CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
}

func NewTicketCommander(bot Sender) TicketCommander {
	return ticket.NewTicketCommander(bot)
}

func NewTicketCallbacker(bot Sender) TicketCallbacker {
	return ticket.NewTicketCommander(bot)
}
