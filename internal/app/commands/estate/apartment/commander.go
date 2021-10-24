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
	switch CallbackName(callbackPath.CallbackName) {
	case CbNList:
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
	switch CommandName(commandPath.CommandName) {
	case CNHelp:
		resp, err = c.Help(command)
	case CNList:
		resp, err = c.List(command)
	case CNGet:
		resp, err = c.Get(command)
	case CNDelete:
		resp, err = c.Delete(command)
	case CNNew:
		resp, err = c.New(command)
	case CNEdit:
		resp, err = c.Edit(command)
	default:
		resp, err = c.Default(command)
	}
	return
}
