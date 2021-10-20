package provider

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PaymentProviderCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.ParseUint(args, 0, 64)

	if err != nil {
		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Provider ID is required!")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PaymentProviderCommander.Delete: error sending reply message to chat - %v", err)
		}
		return
	}
	var outputmsg string
	res, err := c.providerService.Remove(idx)

	if res == false {
		outputmsg = fmt.Sprintf("Provider with id:%d doesn't exists", idx)
	} else {
		outputmsg = fmt.Sprintf("Provider with id:%d successfully deleted", idx)
	}

	if err != nil {
		log.Printf("failed to delete payment provider by id %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputmsg)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentProviderCommander.Get: error sending reply message to chat - %v", err)
	}
}
