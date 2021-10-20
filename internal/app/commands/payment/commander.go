package payment

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/payment/provider"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type PaymentCommander struct {
	bot               *tgbotapi.BotAPI
	providerCommander Commander
}

func NewPaymentCommander(bot *tgbotapi.BotAPI) *PaymentCommander {
	return &PaymentCommander{
		bot:               bot,
		providerCommander: provider.NewPaymentProviderCommander(bot),
	}
}

func (c *PaymentCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "provider":
		c.providerCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("PaymentCommander.HandleCallback: unknown provider - %s", callbackPath.Subdomain)
	}
}

func (c *PaymentCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "provider":
		c.providerCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("PaymentCommander.HandleCommand: unknown provider - %s", commandPath.Subdomain)
	}
}
