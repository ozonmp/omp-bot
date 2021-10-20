package notification

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommunicationNotificationCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Get wrong args %v", args)
		c.SendErrorMessage(inputMessage.Chat.ID,"Wrong command format. Use /get__communication__notification <ID>")
		return
	}

	notification, err := c.notificationService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get notification with idx %d: %v", idx, err)
		c.SendErrorMessage(inputMessage.Chat.ID, "Notification does not exists")
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		notification.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Get: error sending reply message to chat - %v", err)
	}
}
