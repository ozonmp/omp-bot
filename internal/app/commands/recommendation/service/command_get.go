package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RecommendationServiceCommander) Get(inputMessage *tgbotapi.Message) {
	args := strings.TrimSpace(inputMessage.CommandArguments())

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	var messageText string
	product, err := c.serviceService.Get(idx)
	if err != nil {
		log.Printf("fail to get product with idx %d: %v", idx, err)
		messageText = fmt.Sprintf("Service with Id = %d not found", idx)
	} else {
		messageText = product.String()
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		messageText,
	)

	c.bot.Send(msg)
}
