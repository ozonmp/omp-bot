package subdomain

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/demo/subdomain"
)

var (
	ErrUnknownCallback = fmt.Errorf("unknown callback")
)

type DemoSubdomainCommander struct {
	subdomainService *subdomain.Service
}

func NewDemoSubdomainCommander() *DemoSubdomainCommander {
	subdomainService := subdomain.NewService()

	return &DemoSubdomainCommander{
		subdomainService: subdomainService,
	}
}

func (c *DemoSubdomainCommander) HandleCallback(
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

func (c *DemoSubdomainCommander) HandleCommand(
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
	default:
		resp, err = c.Default(command)
	}
	return
}
