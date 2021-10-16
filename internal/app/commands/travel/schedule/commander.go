package schedule

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/travel/schedule"
	service "github.com/ozonmp/omp-bot/internal/service/travel/schedule"
	"github.com/sirupsen/logrus"
)

const pageSize = 10

type ScheduleCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type TravelScheduleCommander struct {
	bot     *tgbotapi.BotAPI
	service service.ScheduleService
}

func NewScheduleCommander(bot *tgbotapi.BotAPI) *TravelScheduleCommander {
	service := schedule.NewDummyScheduleService()

	return &TravelScheduleCommander{
		bot:     bot,
		service: service,
	}
}

func (c *TravelScheduleCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		logrus.Infof("TravelScheduleCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *TravelScheduleCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "get":
		c.Get(message)
	case "list":
		c.List(message)
	case "delete":
		c.Delete(message)
	case "new":
		c.New(message)
	case "edit":
		c.Edit(message)
	default:
		logrus.Infof("TravelScheduleCommander.HandleCommand: unknown command name: %s", commandPath.CommandName)
	}
}

func init() {
	logrus.SetReportCaller(true)
}
