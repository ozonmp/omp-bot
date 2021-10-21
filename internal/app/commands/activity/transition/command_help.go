package transition

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *ActivityTransitionCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__activity__transition - help\n"+
			"/list__activity__transition [offset limit] - list transitions\n"+
			"/get__activity__transition {transition id} - get transition with {transition id}"+
			"/delete__activity__transition {transition id} - delete transition with {transition id}",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityTransitionCommander.Help: error sending reply message to chat - %v", err)
	}
}
