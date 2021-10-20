package point

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *PointCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__loyalty__point - help\n"+
		"/list__loyalty__point - list products\n"+
		"/get__loyalty__point - get product by id\n"+
		"/delete__loyalty__point - remove entity by id\n" + 
		"/create__loyalty__point {name | description} - create new entity by args\n" + 
		"/edit__loyalty__point {id | name | description} - edit entity by args",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PointCommander.Help: error sending reply message to chat - %v", err)
	}
}