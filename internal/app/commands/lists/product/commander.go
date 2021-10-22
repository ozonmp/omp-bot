package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/lists/product"
	"log"
)

type ProductCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type ProductCommanderImpl struct {
	bot      *tgbotapi.BotAPI
	service  service.ProductService
	pageSize uint64
}

func (c *ProductCommanderImpl) New(inputMsg *tgbotapi.Message) {
	log.Printf("unimplemented")
}

func (c *ProductCommanderImpl) Edit(inputMsg *tgbotapi.Message) {
	log.Printf("unimplemented")
}

func (c *ProductCommanderImpl) Default(inputMsg *tgbotapi.Message) {
	log.Printf("unimplemented")
}

func (c *ProductCommanderImpl) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ProductCommanderImpl.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ProductCommanderImpl) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
		c.Default(msg)
	}
}

func NewProductCommander(bot *tgbotapi.BotAPI, service service.ProductService) *ProductCommanderImpl {
	return &ProductCommanderImpl{bot, service, 10}
}

var _ ProductCommander = (*ProductCommanderImpl)(nil)
