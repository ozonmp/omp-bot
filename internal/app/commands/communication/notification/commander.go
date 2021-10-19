package notification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/communication/notification"
)

type CommunicationNotificationCommander struct {
	bot                 *tgbotapi.BotAPI
	notificationService *notification.Service
}

func NewCommunicationNotificationCommander(
	bot *tgbotapi.BotAPI,
) *CommunicationNotificationCommander {
	notificationService := notification.NewService()

	return &CommunicationNotificationCommander{
		bot:                 bot,
		notificationService: notificationService,
	}
}

func (c *CommunicationNotificationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	//case "list":
	//	c.CallbackList(callback, callbackPath)
	default:
		log.Printf("CommunicationNotificationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *CommunicationNotificationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	//case "list":
	//	c.List(msg)
	//case "get":
	//	c.Get(msg)
	default:
		c.Default(msg)
	}
}
