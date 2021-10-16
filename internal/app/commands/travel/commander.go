package travel

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/travel/schedule"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TravelCommander struct {
	bot               *tgbotapi.BotAPI
	scheduleCommander Commander
}

func NewTravelCommander(
	bot *tgbotapi.BotAPI,
) *TravelCommander {
	return &TravelCommander{
		bot:               bot,
		scheduleCommander: schedule.NewScheduleCommander(bot),
	}
}

func (c *TravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "schedule":
		c.scheduleCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("TravelCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *TravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "schedule":
		c.scheduleCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("TravelCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
