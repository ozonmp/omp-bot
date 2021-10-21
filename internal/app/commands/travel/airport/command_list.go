package airport

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *TravelAirportCommander) List(inputMessage *tgbotapi.Message) {
	outputMessageText := "All the airports: \n\n"
	airports := c.airportService.List()
	for _, a := range airports {
		outputMessageText += fmt.Sprintf("%s\n", a)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessageText)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("TravelAirportCommander.List: error sending reply message to chat - %v", err)
	}
}
