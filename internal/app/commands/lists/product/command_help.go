package product

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ProductCommanderImpl) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__lists__product — print list of commands\n"+
			"/get__lists__product — get a entity\n"+
			"/list__lists__product — get a list of your entity (💎: with pagination via telegram keyboard)\n"+
			"/delete__lists__product — delete an existing entity\n"+
			"/new__lists__product — create a new entity // not implemented (💎: implement list fields via arguments)\n"+
			"/edit__lists__product — edit a entity      // not implemented\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ProductCommanderImpl.Help: error sending reply message to chat - %v", err)
	}
}
