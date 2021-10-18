package subscription

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/subscription/singleSubscription"
	"github.com/ozonmp/omp-bot/internal/app/path"
	svc "github.com/ozonmp/omp-bot/internal/service/subscription/singleSubscription"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type SubscriptionCommander struct {
	bot                         *tgbotapi.BotAPI
	singleSubscriptionCommander Commander
}

func NewSubscriptionCommander(
	bot *tgbotapi.BotAPI,
) *SubscriptionCommander {
	singleSubscriptionService := svc.NewDummySingleSubscriptionService()
	return &SubscriptionCommander{
		bot:                         bot,
		singleSubscriptionCommander: singleSubscription.NewSingleSubscriptionCommander(bot, singleSubscriptionService),
	}
}

func (c *SubscriptionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "singleSubscription":
		c.singleSubscriptionCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("SubscriptionCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *SubscriptionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "singleSubscription":
		c.singleSubscriptionCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("SubscriptionCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
