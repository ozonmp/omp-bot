package airport

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"strings"
)

func (c *TravelAirportCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idxNameLocation := strings.Split(args, ",")
	var outputMsg string
	if len(idxNameLocation) != 3 {
		outputMsg = "Wrong format. Please provide coma separated n" +
			"ame and location of the airport. " +
			"Example: '2,LAX,Los-Angeles'"
	} else {
		idx, err := strconv.Atoi(idxNameLocation[0])
		if err != nil {
			outputMsg = "Wrong format. Please provide coma separated " +
				"name and location of the airport. " +
				"Example: '2,LAX,Los-Angeles'\n\n" +
				"First symbol should me number of the airport to edit (starting from 0)"
		}
		airport, err := c.airportService.Edit(idx, idxNameLocation[1], idxNameLocation[2])
		if err != nil {
			log.Printf("Couldn't edit the airport")
			return
		}
		outputMsg = fmt.Sprintf("Edited #%d airport, not its: %s", idx, airport)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outputMsg,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("TravelAirportCommander.Edit: error sending reply - %v", err)
	}
}
