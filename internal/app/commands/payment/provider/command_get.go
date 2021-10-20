package provider

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PaymentProviderCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	provider, err := c.providerService.Get(idx)
	if err != nil {
		log.Printf("failed to get payment provider by id %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		c.providerService.LongDescription(provider),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Get: error sending reply message to chat - %v", err)
	}
}
