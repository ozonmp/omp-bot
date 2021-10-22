package point

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/loyalty/point"
)

type PointCommander struct {
	bot           *tgbotapi.BotAPI
	pointService *point.DummyPointService
}

func NewPointCommander(
	bot *tgbotapi.BotAPI,
) *PointCommander {
	pointService := point.NewDummyPointService()

	return &PointCommander{
		bot:           bot,
		pointService: pointService,
	}
}

func (c *PointCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "prev":
		c.CallbackList(callback, callbackPath)
	case "next":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PointCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PointCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}