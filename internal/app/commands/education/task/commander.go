package task

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/education/task"
)

type TaskCommander struct {
	bot         *tgbotapi.BotAPI
	taskService *task.DummyTaskService
}

func NewTaskCommander(bot *tgbotapi.BotAPI) *TaskCommander {

	taskService := task.NewDummyTaskService()

	return &TaskCommander{
		bot:         bot,
		taskService: taskService,
	}
}

func (c *TaskCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("TaskCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TaskCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "create":
		c.Create(msg)
	case "update":
		c.Update(msg)
	case "remove":
		c.Remove(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}

func (c *TaskCommander) SendMessage(msg tgbotapi.MessageConfig) {

	_, err := c.bot.Send(msg)
	if err != nil {
		fmt.Printf("Send messag error %s", err)
	}

}
