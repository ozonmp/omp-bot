package verification

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ServiceVerificationCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	itemID, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Please enter a number as an argument.",
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("ServiceVerificationCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}

	status, err := c.verificationService.Remove(uint64(itemID))
	if err != nil {
		log.Printf("Deleting item ID: %d Error: %v", itemID, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("Can't delete item ID: %d, maybe it doesn't exist", itemID),
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("ServiceVerificationCommander.Remove: error sending reply message to chat - %v", err)
		}
		return
	}

	msgValue := ""
	if !status {
		msgValue = fmt.Sprintf("Can't delete item ID: %d, maybe it doesn't exist", itemID)
	} else {
		msgValue = fmt.Sprintf("Item ID %d is deleted", itemID)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgValue,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ServiceVerificationCommander.Remove: error sending reply message to chat - %v", err)
	}
}
