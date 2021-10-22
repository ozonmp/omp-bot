package estate

import (
	estate "github.com/ozonmp/omp-bot/internal/app/commands/estate/rent"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/estate/rent"
	"github.com/ozonmp/omp-bot/internal/service/estate/rent/storage"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type EventHandler struct {
	sub string
	commander *estate.Commander
}

func NewEventHandler(bot *tgbotapi.BotAPI) *EventHandler {
	memStorage := storage.NewMemoryStorage()
	svc := rent.NewService(memStorage)

	return &EventHandler{
		commander: estate.NewRentCommander(bot, svc),
	}
}

func (h *EventHandler) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	if h.commander.IsBelongsSubdomain(callbackPath.Subdomain) {
		h.commander.HandleCallback(callback, callbackPath)
	} else {
		log.Printf("RentCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (h *EventHandler) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	if h.commander.IsBelongsSubdomain(commandPath.Subdomain) {
		h.commander.HandleCommand(message, commandPath)
	} else {
		log.Printf("RentCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}

