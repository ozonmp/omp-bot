package workplace

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/business/workplace"
	"log"
)

type WorkplaceCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type BusinessWorkplaceCommander struct {
	bot              *tgbotapi.BotAPI
	workplaceService *service.DummyWorkplaceService
}

func NewWorkplaceCommander(bot *tgbotapi.BotAPI,) *BusinessWorkplaceCommander {
	var workplaceService = service.NewDummyWorkplaceService()

	return &BusinessWorkplaceCommander{
		bot:              bot,
		workplaceService: workplaceService,
	}
}

func (c *BusinessWorkplaceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("BusinessWorkplaceCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *BusinessWorkplaceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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



