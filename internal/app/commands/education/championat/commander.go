package championat

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/championat"
)

type ChampionatCommander struct {
	bot               *tgbotapi.BotAPI
	championatService *championat.DummyChampionatService
}

func NewChampionatCommander(
	bot *tgbotapi.BotAPI,
) *ChampionatCommander {
	championatService := championat.NewDummyChampionatService()

	return &ChampionatCommander{
		bot:               bot,
		championatService: championatService,
	}
}

func (c *ChampionatCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ChampionatCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ChampionatCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
