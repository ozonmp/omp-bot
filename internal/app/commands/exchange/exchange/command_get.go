package exchange

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubdomainCommander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Wrong argument \"%v\".\n"+
				"Please input an existing ID number of your exchange request", args),
		)
		_, err = c.bot.Send(errorResponse)
		if err != nil {
			log.Printf("SubdomainCommander.Get: error sending reply message to chat - %v", err)
			return
		}
		log.Println("SubdomainCommander.Get: wrong args", args)
		return
	}

	exchangeRequest, err := c.exchangeService.Describe(uint64(idx))
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Exchange request with id \"%v\" doesn't exist.\n"+
				"Please input an existing ID number of your exchange request", idx),
		)
		_, err2 := c.bot.Send(errorResponse)
		if err2 != nil {
			log.Printf("SubdomainCommander.Get: error sending reply message to chat - %v", err2)
			return
		}
		log.Printf("fail to get exchangeRequest with idx %d: %v", idx, err)
		return
	}

	exchangeRequestText := fmt.Sprintf("Exchange request ID: %v\n"+
		"From: %v\n"+
		"To: %v\n"+
	    "Package: %v\n"+
		"Status: %v\n",
		exchangeRequest.Id,
		exchangeRequest.From,
		exchangeRequest.To,
		exchangeRequest.Package,
		exchangeRequest.Status,
	)

	response := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		exchangeRequestText,
	)

	_, err = c.bot.Send(response)
	if err != nil {
		log.Printf("SubdomainCommander.Get: error sending reply message to chat - %v", err)
	}
}
