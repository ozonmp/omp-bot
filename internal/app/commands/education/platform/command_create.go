package platform

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PlatformBaseCommander) Create(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	inputData := PlatformInput{}

	err := json.Unmarshal([]byte(args), &inputData)
	if err != nil {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "invalid platform data")
		c.sendMessage(msg)

		return
	}

	err = c.validate(inputData)
	if err != nil {
		log.Println("validation error", err.Error())

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, err.Error())
		c.sendMessage(msg)

		return
	}

	platformID, err := c.service.Create(c.convertToEntity(inputData))
	if err != nil {
		log.Println("create error", err.Error())

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "error create platform, try again")
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("platform created with ID %d", platformID),
	)

	c.sendMessage(msg)
}
