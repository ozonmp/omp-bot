package airport

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *TravelAirportCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	airport, err := c.airportService.Delete(idx)
	if err != nil {
		log.Printf("fail to delete product with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		fmt.Sprintf("Deleted airport %s", airport),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("TravelAirportCommander.Delete: error sending reply message to chat - %v", err)
	}
}
