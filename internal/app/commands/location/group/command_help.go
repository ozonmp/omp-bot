package group

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LocationGroupCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__location__group - print list of commands\n"+
			"/list__location__group - get a list of your entity (with pagination via telegram keyboard)\n"+
			"/get__location__group - get an entity (by index)\n"+
			"/delete__location__group - delete an existing entity\n"+
			"/new__location__group - create a new entity (implement list fields via arguments)\n"+
			"/edit__location__group - edit an entity (not implemented)",
	)
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("HelpLocationCommander.Help: error sending reply message to chat - %v", err)
	}
}
