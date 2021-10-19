package communication

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/communication/notification"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CommunicationCommander struct {
	bot                   *tgbotapi.BotAPI
	notificationCommander Commander
}

func NewCommunicationCommander(
	bot *tgbotapi.BotAPI,
) *CommunicationCommander {
	return &CommunicationCommander{
		bot: bot,
		notificationCommander: notification.NewCommunicationNotificationCommander(bot),
	}
}

func (c *CommunicationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "notification":
		c.notificationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("CommunicationCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *CommunicationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "notification":
		c.notificationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("CommunicationCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
