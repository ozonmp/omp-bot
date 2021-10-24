package return1

//maybe rename file?

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Return1Commander interface { //did not get why do we need it, but it is said to add this
	Commander

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type Return1Service interface {
	Describe(return1ID uint64) (*exchange.Return1, error)
	List(cursor uint64, limit uint64) ([]exchange.Return1, error)
	Create(exchange.Return1) (uint64, error)
	Update(return1ID uint64, return1 exchange.Return1) error
	Remove(return1ID uint64) (bool, error)
}

type Return1CommanderImpl struct {
	bot     *tgbotapi.BotAPI
	service Return1Service
}

func NewReturn1Commander(bot *tgbotapi.BotAPI, service Return1Service) Return1Commander {
	return &Return1CommanderImpl{
		bot:     bot,
		service: service,
	}
}

func (c *Return1CommanderImpl) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.ListCallback(callback, callbackPath)
	default:
		log.Printf("Return1CommanderImpl: HandleCallback: unknown callback [%s]", callbackPath.CallbackName)
	}

}

const (
	helpCommand   = "help"
	getCommand    = "get"
	listCommand   = "list"
	deleteCommand = "delete"
	newCommand    = "new"
	editCommand   = "edit"
)

func (c *Return1CommanderImpl) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case helpCommand:
		c.Help(message)
	case getCommand:
		c.Get(message)
	case listCommand:
		c.List(message)
	case deleteCommand:
		c.Delete(message)
	case newCommand:
		c.New(message)
	case editCommand:
		c.Edit(message)
	default:
		log.Printf("Return1CommanderImpl: HandleCommand: unknown command [%s]", commandPath.CommandName)
	}
}
