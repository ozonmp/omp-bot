package office

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *OfficeCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__business__office - help\n"+
			"/list__business__office - list products",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Help: error sending reply message to chat - %v", err)
	}
}
