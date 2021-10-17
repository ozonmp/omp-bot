package internship

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/work/internship"
)

type WorkCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type WorkInternshipCommander struct {
	bot               *tgbotapi.BotAPI
	internshipService internship.WorkService
}

func NewWorkInternshipCommander(
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
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
