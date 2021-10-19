package estate

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/estate/apartment"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/estate/apartment"
)

var (
	ErrUnknownSubDomain = fmt.Errorf("unknown subdomain")
)

type Commander interface {
	HandleCallback(
		callback *tgbotapi.CallbackQuery,
		callbackPath path.CallbackPath,
	) (resp tgbotapi.MessageConfig, err error)

	HandleCommand(
		command *tgbotapi.Message,
		commandPath path.CommandPath,
	) (resp tgbotapi.MessageConfig, err error)
}

type EstateCommander struct {
	apartmentCommander Commander
}

func NewEstateCommander() *EstateCommander {
	return &EstateCommander{
		apartmentCommander: apartment.NewDummyApartmentCommander(service.NewDummyApartmentService()),
	}
}

func (c *EstateCommander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch callbackPath.Subdomain {
	case "apartment":
		resp, err = c.apartmentCommander.HandleCallback(callback, callbackPath)
	default:
		err = ErrUnknownSubDomain
	}
	return
}

func (c *EstateCommander) HandleCommand(
	command *tgbotapi.Message,
	commandPath path.CommandPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch commandPath.Subdomain {
	case "apartment":
		resp, err = c.apartmentCommander.HandleCommand(command, commandPath)
	default:
		err = ErrUnknownSubDomain
	}
	return
}
