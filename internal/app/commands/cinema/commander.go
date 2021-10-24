package cinema

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/cinema/purchase"
	"github.com/ozonmp/omp-bot/internal/app/path"
	purchaseService "github.com/ozonmp/omp-bot/internal/service/cinema/purchase"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CinemaCommander struct {
	bot              *tgbotapi.BotAPI
	purchaseCommander Commander
}

func NewCinemaCommander(
	bot *tgbotapi.BotAPI,
) *CinemaCommander {
	return &CinemaCommander{
		bot:              bot,
		purchaseCommander: purchase.NewPurchaseCommander(bot, purchaseService.NewDummyPurchaseService()),
	}
}

func (c *CinemaCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "purchase":
		c.purchaseCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CinemaCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CinemaCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "purchase":
		c.purchaseCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CinemaCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
