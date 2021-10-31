package intern

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *WorkInternCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		""+
			"/help__work__intern — print list of commands\n"+
			"/get__work__intern {id} — get an intern with id={id}\n"+
			"/list__work__intern — list all the interns \n"+
			"/delete__work__intern {id} — delete an intern with id={id}\n"+
			"/new__work__intern {name} — add new intern "+
			"/edit__work__intern {id} {name} — edit an intern with id={id}",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("InternCommander.Help: error sending reply message to chat - %v", err)
	}
}
