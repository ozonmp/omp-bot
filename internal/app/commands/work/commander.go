package internship

import (
	"log"

	"github.com/VYBel/omp-bot/internal/app/commands/work/internship"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WorkCommander struct {
	bot                 *tgbotapi.BotAPI
	internshipCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *WorkCommander {
	return &WorkCommander{
		bot: bot,
		// subdomainCommander
		internshipCommander: internship.NewWorkInternshipCommander(bot),
	}
}

func (c *WorkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Internship {
	case "internship":
		c.internshipCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown internship - %s", callbackPath.Internship)
	}
}

func (c *WorkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Internship {
	case "internship":
		c.internshipCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown internship - %s", callbackPath.Internship)
	}
}
