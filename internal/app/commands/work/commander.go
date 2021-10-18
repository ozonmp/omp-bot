package work

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/work/internship"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type workCommander struct {
	bot                 *tgbotapi.BotAPI
	internshipCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *workCommander {
	return &workCommander{
		bot:                 bot,
		internshipCommander: internship.NewWorkInternshipCommander(bot),
	}
}

func (c *workCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "internship":
		c.internshipCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("workCommander.HandleCallback: unknown internship - %s", callbackPath.Subdomain)
	}
}

func (c *workCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "internship":
		c.internshipCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("workCommander.HandleCommand: unknown internship - %s", commandPath.Subdomain)
	}
}
