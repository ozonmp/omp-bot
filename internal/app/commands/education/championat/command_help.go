package championat

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ChampionatCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__education__championat — print list of commands\n"+
			"/get__education__championat - get an entity\n"+
			"/list__education__championat - get a list of entity\n"+
			"/delete__education__championat - delete an existing entity\n"+
			"/new__education__championat - create an entity\n"+
			"/edit__education__championat - edit an entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ChampionatCommander.Help: error sending reply message to chat - %v", err)
	}
}
