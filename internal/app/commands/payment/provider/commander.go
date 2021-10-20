package provider

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/payment/provider"
)

type PaymentProviderCommander struct {
	bot             *tgbotapi.BotAPI
	providerService provider.ProviderService
}

func NewPaymentProviderCommander(bot *tgbotapi.BotAPI) *PaymentProviderCommander {
	providerService := provider.NewService()

	return &PaymentProviderCommander{
		bot:             bot,
		providerService: providerService,
	}
}

func (c *PaymentProviderCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PaymentProviderCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PaymentProviderCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.Create(msg)
	case "edit":
		break
	default:
		c.Default(msg)
	}
}
