package access

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/acc/access"
)

type AccAccessCommander struct {
	bot           *tgbotapi.BotAPI
	accessService *access.Service
}

func NewAccAccessCommander(
	bot *tgbotapi.BotAPI,
) *AccAccessCommander {
	return &AccAccessCommander{
		bot:           bot,
		accessService: access.NewService(),
	}
}

func (c *AccAccessCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("AccAccessCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *AccAccessCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
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
