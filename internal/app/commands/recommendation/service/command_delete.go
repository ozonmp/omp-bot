package service

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationServiceCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var messageText string
	_, err = c.serviceService.Remove(idx)
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		messageText = fmt.Sprintf("Service with Id = %d can't be deleted", idx)
	} else {
		messageText = fmt.Sprintf("Service with Id = %d was deleted", idx)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		messageText,
	)

	c.bot.Send(msg)
}
