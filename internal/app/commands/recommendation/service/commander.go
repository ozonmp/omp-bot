package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/recommendation/service"
)

type RecommendationServiceCommander struct {
	bot            *tgbotapi.BotAPI
	serviceService service.ServiceService
}

func NewRecommendationServiceCommander(
	bot *tgbotapi.BotAPI,
) *RecommendationServiceCommander {
	serviceService := service.NewDummyServiceService(service.AllEntities)

	return &RecommendationServiceCommander{
		bot:            bot,
		serviceService: serviceService,
	}
}

func (c *RecommendationServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("RecommendationServiceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RecommendationServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
