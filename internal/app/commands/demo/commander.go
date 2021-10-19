package demo

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/demo/subdomain"
	"github.com/ozonmp/omp-bot/internal/app/path"
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

type DemoCommander struct {
	subdomainCommander Commander
}

func NewDemoCommander() *DemoCommander {
	return &DemoCommander{
		// subdomainCommander
		subdomainCommander: subdomain.NewDemoSubdomainCommander(),
	}
}

func (c *DemoCommander) HandleCallback(
	callback *tgbotapi.CallbackQuery,
	callbackPath path.CallbackPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch callbackPath.Subdomain {
	case "subdomain":
		resp, err = c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		err = ErrUnknownSubDomain
	}
	return
}

func (c *DemoCommander) HandleCommand(
	command *tgbotapi.Message,
	commandPath path.CommandPath,
) (resp tgbotapi.MessageConfig, err error) {
	switch commandPath.Subdomain {
	case "subdomain":
		resp, err = c.subdomainCommander.HandleCommand(command, commandPath)
	default:
		err = ErrUnknownSubDomain
	}
	return
}
