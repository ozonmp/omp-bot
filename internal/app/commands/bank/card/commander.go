package card

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/bank/card"
	"log"
)

type CardCommander interface {
	Help(inMsg *tgbotapi.Message)
	Get(inMsg  *tgbotapi.Message)
	List(inMsg *tgbotapi.Message)
	Delete(inMsg *tgbotapi.Message)

	//New(inMsg  *tgbotapi.Message)
	//Edit(inMsg *tgbotapi.Message)
}

type DummyCardCommander struct {
	bot         *tgbotapi.BotAPI
	cardService service.CardService
}

func NewCardCommander(
	bot     *tgbotapi.BotAPI,
	service service.CardService) *DummyCardCommander {
	return &DummyCardCommander {
		bot:         bot,
		cardService: service,
	}
}

func (c *DummyCardCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	//case "list":
		//c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DummyCardCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DummyCardCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	default:
		log.Printf("DummyCardCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}