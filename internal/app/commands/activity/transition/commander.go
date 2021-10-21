package transition

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/activity/transition"
)

type ActivityTransitionCommander struct {
	bot               *tgbotapi.BotAPI
	transitionService *transition.DummyTransitionService
}

func NewActivityTransitionCommander(
	bot *tgbotapi.BotAPI,
) *ActivityTransitionCommander {
	transitionService := transition.NewDummyTransitionService()

	return &ActivityTransitionCommander{
		bot:               bot,
		transitionService: transitionService,
	}
}

func (c *ActivityTransitionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ActivityTransitionCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ActivityTransitionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "create":
		c.Create(msg)
	case "update":
		c.Update(msg)
	default:
		c.Default(msg)
	}
}
