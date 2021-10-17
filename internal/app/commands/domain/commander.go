package domain

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/domain/subdomain"
	"github.com/ozonmp/omp-bot/internal/app/path"
	svc "github.com/ozonmp/omp-bot/internal/service/domain/subdomain"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DomainCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewDomainCommander(
	bot *tgbotapi.BotAPI,
) *DomainCommander {
	subdomainService := svc.NewDummySubdomainService()
	return &DomainCommander{
		bot:                bot,
		subdomainCommander: subdomain.NewSubdomainCommander(bot, subdomainService),
	}
}

func (c *DomainCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DomainCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *DomainCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DomainCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
