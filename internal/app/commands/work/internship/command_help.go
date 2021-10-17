package internship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WorkInternshipCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__internship - this help\n"+
			"/get__work__internship - get internship by id\n"+
			"/list__work__internship - list of all internships\n"+
			"/delete__work__internship - delete info about internship by id\n"+
			"/new__work__internship - add info about internship\n"+
			"/edit__work__internship - edit  info about internship\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Help: error sending reply message to chat - %v", err)
	}
}
