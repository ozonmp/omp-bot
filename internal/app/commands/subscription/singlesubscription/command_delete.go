package singlesubscription

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *DummySingleSubscriptionCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("DummySingleSubscriptionCommander.Delete invalid args", args)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, UsageDelete))
		return
	}

	if ok, err := c.service.Remove(uint64(idx)); !ok || err != nil {
		log.Printf("DummySingleSubscriptionCommander.Delete fail to remove elem with idx %d: %v\n", idx, err)
		c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID, ErrOnDelete))
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		SuccessDelete,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummySingleSubscriptionCommander.Delete: error sending reply message to chat - %v", err)
	}
}
