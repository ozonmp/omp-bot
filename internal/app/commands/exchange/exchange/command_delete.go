package exchange

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SubdomainCommander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 10, 64)
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Wrong argument \"%v\".\n"+
				"Please input an existing ID number of your exchange request", args),
		)
		_, err2 := c.bot.Send(errorResponse)
		if err2 != nil {
			log.Printf("SubdomainCommander.Delete: error sending reply message to chat - %v", err2)
			return
		}
		log.Println("SubdomainCommander.Delete: wrong args", args)
		return
	}

	_, err = c.exchangeService.Remove(uint64(idx))
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Exchange request with id \"%v\" doesn't exist.\n"+
				"Please input an existing ID number of your exchange request", idx),
		)
		_, err2 := c.bot.Send(errorResponse)
		if err2 != nil {
			log.Printf("SubdomainCommander.Delete: error sending reply message to chat - %v", err2)
			return
		}
		log.Printf("fail to delete exchangeRequest with idx %d: %v", idx, err)
		return
	}

	responseText := fmt.Sprintf("Your exchange request with id %v successfully deleted", idx)
	response := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		responseText,
	)

	_, err = c.bot.Send(response)
	if err != nil {
		log.Printf("SubdomainCommander.Delete: error sending reply message to chat - %v", err)
	}
}
