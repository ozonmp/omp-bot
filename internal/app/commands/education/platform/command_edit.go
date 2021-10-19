package platform

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *PlatformBaseCommander) Edit(inputMsg *tgbotapi.Message) {
	args := strings.SplitN(inputMsg.CommandArguments(), " ", 2)
	if len(args) != 2 {
		log.Println("wrong args", args)

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "must be 2 arguments separated by space")
		c.sendMessage(msg)

		return
	}

	platformID, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Println("wrong id", args[0])

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "platformID must be unsigned integer")
		c.sendMessage(msg)

		return
	}

	inputData := PlatformInput{}

	err = json.Unmarshal([]byte(args[1]), &inputData)
	if err != nil {
		log.Println("wrong data", args[1])

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

	err = c.service.Update(platformID, c.convertToEntity(inputData))
	if err != nil {
		log.Println("update error", err.Error())

		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "error update platform, try again")
		c.sendMessage(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("platform with ID %d successfully updated", platformID),
	)

	c.sendMessage(msg)
}
