package notification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/communication/notification"
)

type CommunicationNotificationCommander struct {
	bot                 *tgbotapi.BotAPI
	notificationService *notification.DummyNotificationService
}

func NewCommunicationNotificationCommander(
	bot *tgbotapi.BotAPI,
) *CommunicationNotificationCommander {
	notificationService := notification.NewDummyNotificationService()

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
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *CommunicationNotificationCommander) SendErrorMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander: error sending reply message to chat - %v", err)
	}
}