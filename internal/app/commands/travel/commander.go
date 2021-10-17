package travel

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/travel/ticket"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TravelCommander struct {
	bot             Sender
	ticketCommander Commander
}

func NewTravelCommander(
	bot Sender,
) *TravelCommander {
	return &TravelCommander{
		bot: bot,
		// subdomainCommander
		ticketCommander: ticket.NewTravelTicketCommander(bot),
	}
}

func (c *TravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "ticket":
		c.ticketCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("TravelCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *TravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "ticket":
		c.ticketCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("TravelCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
