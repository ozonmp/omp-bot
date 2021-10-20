package singlesubscription

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySingleSubscriptionCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	input := InputSingleSubscription{}

	err := json.Unmarshal([]byte(args), &input)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.New invalid body", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageNew))
		return
	}
	tmp, err := input.ToSingleSubscription()
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.New invalid time", args, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageNew))
		return
	}

	id, _ := c.service.Create(*tmp)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("%s\nid:%d", SuccessNew, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.New: error sending reply message to chat - %v", err)
	}
}
