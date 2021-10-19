package work

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/work/employee"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WorkCommander struct {
	bot               *tgbotapi.BotAPI
	employeeCommander Commander
}

func NewDemoCommander(
	bot *tgbotapi.BotAPI,
) *WorkCommander {
	return &WorkCommander{
		bot:               bot,
		employeeCommander: employee.NewDemoSubdomainCommander(bot),
	}
}

func (c *WorkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "employee":
		c.employeeCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown employee - %s", callbackPath.Subdomain)
	}
}

func (c *WorkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "employee":
		c.employeeCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCommand: unknown employee - %s", commandPath.Subdomain)
	}
}
