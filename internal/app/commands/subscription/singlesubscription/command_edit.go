package singlesubscription

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/subscription"
)

func (c *DummySingleSubscriptionCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	tmp := subscription.SingleSubscription{}

	err := json.Unmarshal([]byte(args), &tmp)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.Edit invalid body", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageEdit))
		return
	}

	err = c.service.Update(tmp.ID, tmp)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Edit failed to update %+v: %v\n", tmp, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, ErrOnEdit))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		SuccessEdit,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Edit: error sending reply message to chat - %v", err)
	}
}
