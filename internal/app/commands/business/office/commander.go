package office

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/office"
)

type OfficeCommander struct {
	bot           *tgbotapi.BotAPI
	officeService *office.DummyOfficeService
}

func NewOfficeCommander(
	bot *tgbotapi.BotAPI,
) *OfficeCommander {
	officeService := office.NewDummyOfficeService()

	return &OfficeCommander{
		bot:           bot,
		officeService: officeService,
	}
}

func (c *OfficeCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("OfficeCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *OfficeCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Help(msg)
	}
}
