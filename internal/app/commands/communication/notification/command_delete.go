package notification

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommunicationNotificationCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Delete wrong args %v", args)
		c.SendErrorMessage(inputMessage.Chat.ID,"Wrong command format. Use /delete__communication__notification <ID>")
		return
	}

	_, err = c.notificationService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to remove notification with idx %d: %v", idx, err)
		c.SendErrorMessage(inputMessage.Chat.ID,"Notification does not exists")
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,"Notification deleted")
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Delete: error sending reply message to chat - %v", err)
	}
}
