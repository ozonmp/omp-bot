package business

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/business/workplace"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessCommander struct {
	bot                *tgbotapi.BotAPI
	workplaceCommander Commander
}

func NewBusinessCommander(bot *tgbotapi.BotAPI,) *BusinessCommander {
	return &BusinessCommander{
		bot: bot,
		// workplaceCommander
		workplaceCommander: workplace.NewWorkplaceCommander(bot),
	}
}

func (c *BusinessCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "workplace":
		c.workplaceCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("BusinessCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *BusinessCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "workplace":
		c.workplaceCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("BusinessCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
