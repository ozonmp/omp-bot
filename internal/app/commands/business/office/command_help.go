package office

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *OfficeCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`/help__business__office - help
		/list__business__office - list products
		/get__business__office {id} - get entity by id
		/delete__business__office {id} - remove entity by id
		/create__business__office json:{"name":"name", "description"":"description"} - create new entity by json string
	`,
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OfficeCommander.Help: error sending reply message to chat - %v", err)
	}
}
