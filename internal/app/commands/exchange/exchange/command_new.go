package exchange

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/model/exchange"
)

func (c *SubdomainCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	arguments := strings.Split(args, ",")
	numberOfArguments := len(arguments)
	if numberOfArguments != 3 {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Wrong number of arguments.\n"+
				"Please type the command in format \"Package, From, To\""),
		)

		_, err := c.bot.Send(errorResponse)
		if err != nil {
			log.Printf("SubdomainCommander.New: error sending reply message to chat - %v", err)
		}
		log.Println("wrong args", args)
		return
	}

	exchangeRequest := exchange.Exchange {
		Id:      0,
		Package: strings.TrimSpace(arguments[0]),
		From:    strings.TrimSpace(arguments[1]),
		To:      strings.TrimSpace(arguments[2]),
		Status:  "Registered",
	}

	exchangeRequestId, err := c.exchangeService.Create(exchangeRequest)

	if err != nil {
		errorResponse := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Unable to create exchange request with the same id - %v", exchangeRequestId),
		)
		_, err2 := c.bot.Send(errorResponse)
		if err2 != nil {
			log.Printf("SubdomainCommander.New: error sending reply message to chat - %v", err2)
			return
		}
		log.Printf("SubdomainCommander.New: error sending reply message to chat - %v", err)
	}

	exchangeRequestText := fmt.Sprintf("Your exchange request registred with id %v", exchangeRequestId)

	response := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		exchangeRequestText,
	)

	_, err = c.bot.Send(response)
	if err != nil {
		log.Printf("SubdomainCommander.New: error sending reply message to chat - %v", err)
	}
}