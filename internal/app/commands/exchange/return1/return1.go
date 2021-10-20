package return1

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/exchange/return1"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Return1Commander interface {
	Commander

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type Return1CommanderImpl struct {
	bot     *tgbotapi.BotAPI
	service service.DummyReturn1Service
}

func NewReturn1Commander(bot *tgbotapi.BotAPI, service service.DummyReturn1Service) Return1Commander {

	return &Return1CommanderImpl{
		bot:     bot,
		service: service,
	}
}

func (c *Return1CommanderImpl) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	//tobedone
}

func (c *Return1CommanderImpl) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "get":
		c.Get(message)
	case "list":
		c.List(message)
	case "delete":
		c.Delete(message)
	case "new":
		c.New(message)
	case "edit":
		c.Edit(message)
	default:
		log.Printf("Return1CommanderImpl: HandleCommand: unknown command [%s]", commandPath.CommandName)
	}
}
