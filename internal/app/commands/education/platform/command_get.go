package platform

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c *PlatformBaseCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	platformID, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "platformID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	platform, err := c.service.Describe(platformID)
	if err != nil {
		log.Printf(err.Error())

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, platform.String())
	c.sendMessage(msg)
}
