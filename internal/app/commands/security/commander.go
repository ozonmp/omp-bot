package security

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/security/verification"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type SecurityCommander struct {
	bot                   *tgbotapi.BotAPI
	verificationCommander Commander
}

func NewSecurityCommander(
	bot *tgbotapi.BotAPI,
) *SecurityCommander {
	return &SecurityCommander{
		bot:                   bot,
		verificationCommander: verification.NewSecurityVerificationCommander(bot),
	}
}

func (c *SecurityCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "verification":
		c.verificationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("SecurityCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *SecurityCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "verification":
		c.verificationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("SecurityCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}