package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	service "github.com/ozonmp/omp-bot/internal/service/recommendation/product"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type Serializer interface {
	serialize(data string) (recomendation.Product, error)
	deserialize(product *recomendation.Product) (string, error)
}

type ProductCommander struct {
	bot        *tgbotapi.BotAPI
	service    service.Service
	serializer Serializer
}

func NewProductCommander(bot *tgbotapi.BotAPI, service service.Service) *ProductCommander {
	ser := &JsonSerializer{}
	return &ProductCommander{bot: bot, service: service, serializer: ser}
}
func (commander *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list_next":
		commander.CallbackNextList(callback, callbackPath)
	case "list_prev":
		commander.CallbackPrevList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (commander *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	log.Printf(commandPath.CommandName)
	switch commandPath.CommandName {
	case "help":
		commander.Help(msg)
	case "new":
		commander.New(msg)
	case "delete":
		commander.Delete(msg)
	case "list":
		commander.List(msg)
	case "get":
		commander.Get(msg)
	case "edit":
		commander.Edit(msg)
	default:
		commander.Default(msg)
	}
}

func (commander *ProductCommander) Send(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	commander.bot.Send(msg)
}
