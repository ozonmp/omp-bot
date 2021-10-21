package exchange

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/exchange/exchange"
)

type ExchangeCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type SubdomainCommander struct {
	bot             *tgbotapi.BotAPI
	exchangeService *exchange.DummyExchangeService
}

func NewExchangeCommander(bot *tgbotapi.BotAPI, service *exchange.DummyExchangeService) *SubdomainCommander {
	return &SubdomainCommander{
		bot:             bot,
		exchangeService: service,
	}
}

func (c *SubdomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *SubdomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("SubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}
