package apartment

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/estate/apartment"
)

var (
	ErrUnknownCallback = fmt.Errorf("unknown callback")
)

type ApartmentCommander interface {
	Help(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig)
	Get(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error)
	List(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error)
	Delete(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error)

	New(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error)
	Edit(inputMsg *tgbotapi.Message) (resp tgbotapi.MessageConfig, err error)
}

type DummyApartmentCommander struct {
	service service.ApartmentService
}

func NewDummyApartmentCommander(service service.ApartmentService) *DummyApartmentCommander {
	return &DummyApartmentCommander{
		service: service,
	}
}

func (c *DummyApartmentCommander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch callbackPath.CallbackName {
	case "list":
		resp, err = c.CallbackList(callback, callbackPath)
	default:
		err = ErrUnknownCallback
	}
	return
}

func (c *DummyApartmentCommander) HandleCommand(
	command *tgbotapi.Message,
	commandPath path.CommandPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch commandPath.CommandName {
	case "help":
		resp, err = c.Help(command)
	case "list":
		resp, err = c.List(command)
	case "get":
		resp, err = c.Get(command)
	case "delete":
		resp, err = c.Delete(command)
	default:
		resp, err = c.Default(command)
	}
	return
}
