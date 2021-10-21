package airport

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c *TravelAirportCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	nameAndLocation := strings.Split(args, ",")
	var outputMsg string
	if len(nameAndLocation) != 2 {
		outputMsg = "Wrong format. Please provide coma separated " +
			"name and location of the airport. " +
			"Example: 'LAX,Los-Angeles'"
	} else {
		airport, err := c.airportService.New(nameAndLocation[0], nameAndLocation[1])
		if err != nil {
			log.Printf("Couldn't create new airport")
			return
		}
		outputMsg = fmt.Sprintf("Created new airport: %s", airport)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("TravelAirportCommand.New: error sending reply - %v", err)
	}
}
