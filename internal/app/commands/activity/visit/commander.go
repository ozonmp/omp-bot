package visit

import (
	service "github.com/ozonmp/omp-bot/internal/service/activity/visit"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type VisitCommander interface {
	Help(inputMessage *tgbotapi.Message)
	Get(inputMessage *tgbotapi.Message)
	List(inputMessage *tgbotapi.Message)
	Delete(inputMessage *tgbotapi.Message)

	New(inputMessage *tgbotapi.Message)  // return error not implemented
	Edit(inputMessage *tgbotapi.Message) // return error not implemented

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

func NewVisitCommander(bot *tgbotapi.BotAPI, visitService service.VisitService) *VisitCommanderStruct {
	return &VisitCommanderStruct{
		bot:          bot,
		visitService: visitService,
	}
}

type VisitCommanderStruct struct {
	bot          *tgbotapi.BotAPI
	visitService service.VisitService
}

func (c *VisitCommanderStruct) Send(chatID int64, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("VisitCommanderStruct.Get: error sending reply message to chat - %v", err)
	}

	return msg
}

func (c *VisitCommanderStruct) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("VisitCommanderStruct.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *VisitCommanderStruct) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
