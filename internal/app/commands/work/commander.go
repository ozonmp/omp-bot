package work

import (
	"github.com/ozonmp/omp-bot/internal/app/commands/work/intern"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WorkCommander struct {
	bot             *tgbotapi.BotAPI
	internCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *WorkCommander {
	return &WorkCommander{
		bot:             bot,
		internCommander: intern.NewWorkInternCommander(bot),
	}
}

func (c *WorkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "intern":
		c.internCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *WorkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "intern":
		c.internCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
