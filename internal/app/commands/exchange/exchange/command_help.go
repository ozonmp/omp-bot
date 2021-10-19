package exchange

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *SubdomainCommander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__exchange__exchange — print list of commands\n"+
			"/get__exchange__exchange — get an entity\n"+
			"/list__exchange__exchange — get a list of your entities\n"+
			"/delete__exchange__exchange  — delete an existing entity\n"+
			"\n"+
			"/new__exchange__exchange  — create a new entity\n"+
			"/edit__exchange__exchange  — edit an entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SubdomainCommander.Help: error sending reply message to chat - %v", err)
	}
}
