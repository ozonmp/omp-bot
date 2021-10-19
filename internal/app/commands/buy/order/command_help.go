package order

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OrderCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OrderCommander.Help: error sending reply message to chat - %v", err)
	}
}
