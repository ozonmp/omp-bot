package task

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/task"
)

const maxElemListPerPage = 5

type TaskCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TaskStruct struct {
	bot         *tgbotapi.BotAPI
	taskService *task.DummyTaskService
}

func NewTaskCommander(bot *tgbotapi.BotAPI) TaskCommander {

	taskService := task.NewDummyTaskService()

	return &TaskStruct{
		bot:         bot,
		taskService: taskService,
	}
}

func (c *TaskStruct) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("TaskCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TaskStruct) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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

func (c *TaskStruct) SendMessage(msg tgbotapi.MessageConfig) {

	_, err := c.bot.Send(msg)
	if err != nil {
		fmt.Printf("Send messag error %s", err)
	}

}
