package notification

import (
	"fmt"
	"github.com/ozonmp/omp-bot/internal/model/communication"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CommunicationNotificationCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	notificationParts := strings.SplitN(args, " ", 3)
	if len(notificationParts) != 3 {
		c.SendErrorMessage(inputMessage.Chat.ID, "Wrong command format. Use /edit__communications__notifications <ID> <NewRecipient> <NewText>")
		return
	}
	notificationID, err := strconv.Atoi(notificationParts[0])
	if err != nil {
		c.SendErrorMessage(inputMessage.Chat.ID, "Wrong command format. Use /edit__communications__notifications <ID> <NewRecipient> <NewText>")
		return
	}

	newNotification := communication.Notification{
		ID:        uint64(notificationID),
		Title:     notificationParts[2],
		Recipient: notificationParts[1],
		Sender:    inputMessage.Chat.UserName,
	}
	err = c.notificationService.Update(uint64(notificationID), newNotification)
	if err != nil {
		log.Printf("Fail to edit notification with idx %d: %v", notificationID, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Edited notification with ID: %d", notificationID),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("CommunicationNotificationCommander.Edit: error sending reply message to chat - %v", err)
	}
}
