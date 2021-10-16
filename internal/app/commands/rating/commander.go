package rating

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/rating/customer"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type RatingCommander struct {
	bot               *tgbotapi.BotAPI
	customerCommander Commander
}

func NewRatingCommander(
	bot *tgbotapi.BotAPI,
) *RatingCommander {
	return &RatingCommander{
		bot: bot,
		// customerCommander
		customerCommander: customer.NewCustomerCommander(bot),
	}
}

func (c *RatingCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "customer":
		c.customerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("RatingCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *RatingCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "customer":
		c.customerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("RatingCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
