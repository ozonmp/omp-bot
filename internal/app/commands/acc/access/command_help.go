package access

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *AccAccessCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/get - list access by id\n"+
			"/list - list accesses\n"+
			"/delete - delete access\n"+
			"/new - new access\n"+
			"/edit - edit access\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("AccAccessCommander.Help: error sending reply message to chat - %v", err)
	}
}
