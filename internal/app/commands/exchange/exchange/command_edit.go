package exchange

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *SubdomainCommander) Edit(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	arguments := strings.Split(args, ",")
	numberOfArguments := len(arguments)
	if numberOfArguments != 2 {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Wrong number of arguments.\n"+
				"Please type the command in format \"ID, Status\""),
		)

		_, err := c.bot.Send(errorResponse)
		if err != nil {
			log.Printf("SubdomainCommander.Edit: error sending reply message to chat - %v", err)
		}
		log.Println("wrong args", args)
		return
	}

	editExchangeRequestId, err := strconv.Atoi(arguments[0])
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Wrong argument \"%v\".\n"+
				"Please input an existing ID number of your exchange request", arguments[0]),
		)
		_, err = c.bot.Send(errorResponse)
		if err != nil {
			log.Printf("SubdomainCommander.Edit: error sending reply message to chat - %v", err)
			return
		}
		log.Println("wrong args", arguments[0])
		return
	}

	exchangeRequest, err := c.exchangeService.Describe(uint64(editExchangeRequestId))
	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Exchange request with id \"%v\" doesn't exist.\n"+
				"Please input an existing ID number of your exchange request", editExchangeRequestId),
		)
		_, err2 := c.bot.Send(errorResponse)
		if err2 != nil {
			log.Printf("SubdomainCommander.Get: error sending reply message to chat - %v", err2)
			return
		}
		log.Printf("fail to get exchangeRequest with idx %d: %v", editExchangeRequestId, err)
		return
	}

	exchangeRequest.Status = strings.TrimSpace(arguments[1])
	_ = c.exchangeService.Update(uint64(editExchangeRequestId), *exchangeRequest)

	exchangeRequestText := fmt.Sprintf("Your exchange request with id %v successully updated",
		editExchangeRequestId)

	response := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		exchangeRequestText,
	)

	_, err = c.bot.Send(response)
	if err != nil {
		log.Printf("SubdomainCommander.Edit: error sending reply message to chat - %v", err)
	}
}
