package common

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/delivery/common"
	"log"
)

type CommonCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type DummyCommonCommander struct {
	bot           *tgbotapi.BotAPI
	commonService service.CommonService
}

func NewCommonCommander(bot *tgbotapi.BotAPI, service service.CommonService) DummyCommonCommander {
	return DummyCommonCommander{
		bot,
		service,
	}
}

func (c *DummyCommonCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DummyCommonCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *DummyCommonCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
	}
}
