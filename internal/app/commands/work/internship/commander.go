package internship

import (
	"log"

	"github.com/VYBel/omp-bot/internal/app/commands/work/internship"
	"github.com/VYBel/omp-bot/internal/service/work/internship"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type WorkInternshipCommander struct {
	bot               *tgbotapi.BotAPI
	internshipService *internship.Service
}

func NewDemoSubdomainCommander(
	bot *tgbotapi.BotAPI,
) *WorkInternshipCommander {
	internshipService := internship.NewService()

	return &WorkInternshipCommander{
		bot:               bot,
		internshipService: internshipService,
	}
}

func (c *WorkInternshipCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("WorkInternshipCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WorkInternshipCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
