package internship

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WorkInternshipCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__work__internship - this help\n"+
			"/list__work__internship - list of all internships\n"+
			"/new__work__internship - add new empty internship\n"+
			"/get__work__internship id - get internship by id\n"+
			"/edit__work__internship id - edit  info about internship, not implemented\n"+
			"/delete__work__internship id - delete info about internship by id\n\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("WorkInternshipCommander.Help: error sending reply message to chat - %v", err)
	}
}
