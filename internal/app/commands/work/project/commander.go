package project

import (
	"github.com/ozonmp/omp-bot/internal/service/work/project"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type ProjectCommander struct {
	bot              *tgbotapi.BotAPI
	projectService *project.Service
}

func NewProjectCommander(bot *tgbotapi.BotAPI) *ProjectCommander {
	projectService := project.NewService()

	return &ProjectCommander{
		bot:              bot,
		projectService: projectService,
	}
}

func (c *ProjectCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("projectCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ProjectCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
	default:
		c.Default(msg)
	}
}
