package incident

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list incidents\n"+
			"/get - get incidents by index\n"+
			"/delete - delete incidents by index\n"+
			"/new - create new incident\n"+
			"/edit - edit exist incident\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("StreamingIncidentCommander.Help: error sending reply message to chat - %v", err)
	}
}
