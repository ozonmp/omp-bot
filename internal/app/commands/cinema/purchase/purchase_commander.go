package purchase

//maybe rename file?

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/cinema"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type PurchaseCommander interface { //did not get why do we need it, but it is said to add this
	Commander

	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type PurchaseService interface {
	Describe(purchaseID uint64) (*cinema.Purchase, error)
	List(cursor uint64, limit uint64) ([]cinema.Purchase, error)
	Create(cinema.Purchase) (uint64, error)
	Update(purchaseID uint64, purchase cinema.Purchase) error
	Remove(purchaseID uint64) (bool, error)
}

type PurchaseCommanderImpl struct {
	bot     *tgbotapi.BotAPI
	service PurchaseService
}

func NewPurchaseCommander(bot *tgbotapi.BotAPI, service PurchaseService) PurchaseCommander {
	return &PurchaseCommanderImpl{
		bot:     bot,
		service: service,
	}
}

func (c *PurchaseCommanderImpl) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.ListCallback(callback, callbackPath)
	default:
		log.Printf("PurchaseCommanderImpl: HandleCallback: unknown callback [%s]", callbackPath.CallbackName)
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

func (c *PurchaseCommanderImpl) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
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
		log.Printf("PurchaseCommanderImpl: HandleCommand: unknown command [%s]", commandPath.CommandName)
	}
}
