package business

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/travel"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessCommander struct {
	bot             *tgbotapi.BotAPI
	travelCommander Commander
}

func NewBusinessCommander(
	bot *tgbotapi.BotAPI,
) *BusinessCommander {
	return &BusinessCommander{
		bot: bot,
		// businessCommander
		travelCommander: travel.NewBusinessTravelCommander(bot),
	}
}

func (c *BusinessCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Travel {
	case "travel":
		c.travelCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown travel - %s", callbackPath.Travel)
	}
}

func (c *BusinessCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Travel {
	case "travel":
		c.travelCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown travel - %s", commandPath.Travel)
	}
}
