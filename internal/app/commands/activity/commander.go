package activity

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/activity/visit"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	serviceVisit "github.com/ozonmp/omp-bot/internal/service/activity/visit"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ActivityCommander struct {
	bot                *tgbotapi.BotAPI
	VisitCommandStruct Commander
}

func NewActivityCommander(bot *tgbotapi.BotAPI) *ActivityCommander {
	visitCommand := serviceVisit.NewDummyVisitService()

	return &ActivityCommander{
		bot:                bot,
		VisitCommandStruct: visit.NewVisitCommander(bot, visitCommand),
	}
}

func (c *ActivityCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	switch callbackPath.Subdomain {
	case "visit":
		c.VisitCommandStruct.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ActivityCommander.HandleCallback: unknown visit - %s", callbackPath.Subdomain)
	}
}

func (c *ActivityCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "visit":
		c.VisitCommandStruct.HandleCommand(msg, commandPath)
	default:
		log.Printf("ActivityCommander.HandleCommand: unknown visit - %s", commandPath.Subdomain)
	}
}
