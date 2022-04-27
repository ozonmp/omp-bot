package travel

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/business/travel"
)

type BusinessTravelCommander struct {
	bot           *tgbotapi.BotAPI
	travelService *travel.Service
}

func NewBusinessTravelCommander(
	bot *tgbotapi.BotAPI,
) *BusinessTravelCommander {
	travelService := travel.NewService()

	return &BusinessTravelCommander{
		bot:           bot,
		travelService: travelService,
	}
}

func (c *BusinessTravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BusinessTravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
