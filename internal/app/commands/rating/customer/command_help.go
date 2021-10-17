package customer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Help(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`
/help__rating__customer — print list of commands
/get__rating__customer — get a entity
/list__rating__customer — get a list of your entity (💎: with pagination via telegram keyboard)
/delete__rating__customer — delete an existing entity

/new__rating__customer — create a new entity // not implemented (💎: implement list fields via arguments)
/edit__rating__customer — edit a entity      // not implemented`,
	)

	_, err := c.bot.Send(msg)
	return err
}
