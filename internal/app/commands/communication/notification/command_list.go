package notification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommunicationNotificationCommander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the notifications: \n\n"

	notifications, _ := c.notificationService.List()
	for _, n := range notifications {
		outputMsgText += n.String() + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.List: error sending reply message to chat - %v", err)
	}
}
