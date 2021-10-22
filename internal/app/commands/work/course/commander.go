package course

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/work/course"
)

type WorkCourseCommander struct {
	bot           *tgbotapi.BotAPI
	courseService *course.Service
	cursor        uint64
	limit         uint64
}

func NewWorkCourseCommander(bot *tgbotapi.BotAPI) *WorkCourseCommander {
	courseService := course.NewService()

	return &WorkCourseCommander{
		bot:           bot,
		courseService: courseService,
	}
}

func (c *WorkCourseCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("WorkCourseCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WorkCourseCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg,0,3)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
