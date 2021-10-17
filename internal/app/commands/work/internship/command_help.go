package internship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WorkInternshipCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__internship - help\n"+
			"/list__work__internship - internship list",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Help: error sending reply message to chat - %v", err)
	}
}
