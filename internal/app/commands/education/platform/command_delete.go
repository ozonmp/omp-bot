package platform

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *PlatformBaseCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	platformID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "platformID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	_, err = c.service.Remove(platformID)
	if err != nil {
		log.Printf(err.Error())

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("platform with ID %d successfully deleted", platformID),
	)

	c.sendMessage(msg)
}
