package airport

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *TravelAirportCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list airports\n"+
			"/get - get airport py its index\n"+
			"/edit - set name and location of the airport\n"+
			"/delete - delete the airport by its index\n"+
			"/new - create new airport",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("TravelAirportCommander.Help: error sending reply message to chat - %v", err)
	}
}
