package group

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/product/group"
)

const groupsPerPage = 3

type GroupCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type ProductGroupCommander struct {
	bot          *tgbotapi.BotAPI
	groupService group.GroupService
}

func NewProductGroupCommander(bot *tgbotapi.BotAPI) *ProductGroupCommander {
	groupService := group.NewDummyGroupService()

	return &ProductGroupCommander{
		bot:          bot,
		groupService: groupService,
	}
}

func (c *ProductGroupCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ProductGroupCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ProductGroupCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	}
}
