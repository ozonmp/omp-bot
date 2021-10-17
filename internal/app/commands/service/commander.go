package service

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/service/service"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ServiceCommander struct {
	bot              *tgbotapi.BotAPI
	serviceCommander Commander
}

func NewServiceCommander(
	bot *tgbotapi.BotAPI,
) *ServiceCommander {
	return &ServiceCommander{
		bot: bot,
		// serviceCommander
		serviceCommander: service.NewServiceServiceCommander(bot),
	}
}

func (c *ServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "service":
		c.serviceCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("ServiceCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "service":
		c.serviceCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("ServiceCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
