package verification

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"

)

func (c *ServiceVerificationCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Please enter a number as an argument.",
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("ServiceVerificationCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	item, err := c.verificationService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get item with idx %d: %v", idx, err)
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"Please enter a number as an argument.",
		)
		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("ServiceVerificationCommander.Get: error sending reply message to chat - %v", err)
		}
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		item.Name,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DemoSubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
