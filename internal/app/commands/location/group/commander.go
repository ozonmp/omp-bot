package group

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/location/group"
	"log"
)

type LocationGroupCommander struct {
	bot              *tgbotapi.BotAPI
	subdomainService *group.LocationGroupService
}

func (c *LocationGroupCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("LocationGroupCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func NewGroupCommander(bot *tgbotapi.BotAPI) *LocationGroupCommander {
	service := group.NewLocationGroupService()

	return &LocationGroupCommander{
		bot:              bot,
		subdomainService: service,
	}
}

func (c *LocationGroupCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
		//	case "new":
		//		c.New(msg)
		//	case "edit":
		//		c.Edit(msg)
	default:
		c.CommandNotImplemented(msg)
	}
}
