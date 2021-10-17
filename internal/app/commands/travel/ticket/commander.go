package ticket

import (
	"log"

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
	bot Sender,
) *TravelTicketCommander {

	return &TravelTicketCommander{
		bot:              bot,
		ticketCommander:  travel.NewTicketCommander(bot),
		ticketCallbacker: travel.NewTicketCallbacker(bot),
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
