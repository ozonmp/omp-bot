package apartment

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/estate"
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

type ApartmentService interface {
	Describe(apartmentID uint64) (*estate.Apartment, error)
	List(cursor uint64, limit uint64) ([]estate.Apartment, error)
	Create(estate.Apartment) (uint64, error)
	Update(apartmentID uint64, apartment estate.Apartment) error
	Remove(apartmentID uint64) (bool, error)
}

type DummyApartmentCommander struct {
	service ApartmentService
}

func NewDummyApartmentCommander(service ApartmentService) *DummyApartmentCommander {
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
	case "new":
		resp, err = c.New(command)
	case "edit":
		resp, err = c.Edit(command)
	default:
		resp, err = c.Default(command)
	}
	return
}
