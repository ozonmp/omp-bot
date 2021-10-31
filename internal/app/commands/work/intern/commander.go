package intern

import (
	"github.com/ozonmp/omp-bot/internal/service/work/intern"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type WorkInternCommander struct {
	bot           *tgbotapi.BotAPI
	internService *intern.InternService
}

func NewWorkInternCommander(
	bot *tgbotapi.BotAPI,
) *WorkInternCommander {
	internService := intern.NewService()

	return &WorkInternCommander{
		bot:           bot,
		internService: internService,
	}
}

func (c *WorkInternCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("WorkInternCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WorkInternCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
