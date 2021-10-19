package notification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *CommunicationNotificationCommander) Help(inputMessage *tgbotapi.Message) {
	const helpMessage string = "/help__communication__notification - help\n" +
		"/get__communication__notification - get a notification\n" +
		"/list__communication__notification - notification list\n" +
		"/delete__communication__notification - delete an existing notification\n" +
		"/new__communication__notification - delete a new notification\n" +
		"/edit__communication__notification - edit notification"
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, helpMessage)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Help: error sending reply message to chat - %v", err)
	}
}
