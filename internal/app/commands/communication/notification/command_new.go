package notification

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/communication"
	"log"
	"strings"
)

func (c *CommunicationNotificationCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	inputParts := strings.SplitN(args, " ", 2)
	if len(inputParts) != 2 {
		c.SendErrorMessage(inputMessage.Chat.ID, "Wrong command format. Use /new__communications__notifications <Recipient> <Text>")
		return
	}

	newNotification := communication.Notification{
		Title:     inputParts[1],
		Recipient: inputParts[0],
		Sender:    inputMessage.Chat.UserName,
	}
	createdNotificationID, err := c.notificationService.Create(newNotification)
	if err != nil {
		log.Printf("Fail to create notification with idx %d: %v", createdNotificationID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Created notification with ID: %d", createdNotificationID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.New: error sending reply message to chat - %v", err)
	}
}
