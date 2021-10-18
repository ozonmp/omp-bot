package product

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/recommendation/product"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
//Help(inputMsg *tgbotapi.Message)
//Get(inputMsg *tgbotapi.Message)
//List(inputMsg *tgbotapi.Message)
//Delete(inputMsg *tgbotapi.Message)

//New(inputMsg *tgbotapi.Message)    // return error not implemented
//Edit(inputMsg *tgbotapi.Message)   // return error not implemented
}

type ProductCommander struct {
	bot *tgbotapi.BotAPI
	service service.Service
}

func NewProductCommander(bot *tgbotapi.BotAPI, service service.Service) *ProductCommander {
   return &ProductCommander{bot : bot, service: service}
}
func (c *ProductCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	//case "list":
	//	c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (commander *ProductCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	log.Printf(commandPath.CommandName)
	switch commandPath.CommandName {
	case "help":
		commander.Help(msg)
	//case "list":
	//	c.List(msg)
	//case "get":
	//	c.Get(msg)
	//default:
	//	c.Default(msg)
	}
}