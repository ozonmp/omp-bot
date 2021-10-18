package ground

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *GroundCommander) Help(inputMessage *tgbotapi.Message) {
	c.Send(inputMessage.Chat.ID,
		`
/help__autotransport__ground — print list of commands
/get__autotransport__ground — get a entity
/list__autotransport__ground — get a list of your entity (💎: with pagination via telegram keyboard)
/delete__autotransport__ground — delete an existing entity

/new__autotransport__ground — create a new entity // not implemented (💎: implement list fields via arguments)
/edit__autotransport__ground — edit a entity      // not implemented`,
	)
}
