package test

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/test"
)

type EducationTestCommander struct {
	bot         *tgbotapi.BotAPI
	testService *test.Service
}

func NewEducationTestCommander(
	bot *tgbotapi.BotAPI,
) *EducationTestCommander {
	testService := test.NewService()

	return &EducationTestCommander{
		bot:         bot,
		testService: testService,
	}
}

func (c *EducationTestCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("EducationTestCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *EducationTestCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
