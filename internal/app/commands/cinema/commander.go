package cinema

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/seat"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewCinemaCommander(
	bot *tgbotapi.BotAPI,
) *CinemaCommander {
	return &CinemaCommander{
		bot: bot,
		// subdomainCommander
		subdomainCommander: seat.NewCinemaSeatCommander(bot),
	}
}

func (c *CinemaCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "seat":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CinemaCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "seat":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
