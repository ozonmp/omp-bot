package singleSubscription

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

func (c *DummySingleSubscriptionCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	tmp := subscription.SingleSubscription{}

	err := json.Unmarshal([]byte(args), &tmp)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.New invalid body", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageNew))
		return
	}

	id, _ := c.service.Create(tmp)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("%s\nid:%d", SuccessNew, id),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.New: error sending reply message to chat - %v", err)
	}
}
