package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/service/verification"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ServiceCommander struct {
	bot                *tgbotapi.BotAPI
	verificationCommander Commander
}

func NewServiceCommander(
	bot *tgbotapi.BotAPI,
) *ServiceCommander {
	return &ServiceCommander{
		bot: bot,
		// subdomainCommander
		verificationCommander: verification.NewVerificationCommander(bot),
	}
}

func (c *ServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "verification":
		c.verificationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DomainCommander.HandleCallback: unknown verification - %s", callbackPath.Subdomain)
	}
}

func (c *ServiceCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "verification":
		c.verificationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DomainCommander.HandleCommand: unknown verification - %s", commandPath.Subdomain)
	}
}
