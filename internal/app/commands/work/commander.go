package work

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/work/project"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type SubCommander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Commander struct {
	bot                *tgbotapi.BotAPI
	projectCommander SubCommander
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot: bot,
		// subdomainCommander
		projectCommander: project.NewProjectCommander(bot),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "project":
		c.projectCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "project":
		c.projectCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
