package singleSubscription

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySingleSubscriptionCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.Get invalid args", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageGet))
		return
	}

	elem, err := c.service.Describe(uint64(idx))
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Get fail to describe elem with idx %d: %v\n", idx, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, ErrNotFound))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		elem.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Get: error sending reply message to chat - %v", err)
	}
}
