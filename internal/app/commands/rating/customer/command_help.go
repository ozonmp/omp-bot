package customer

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *CustomerCommander) Help(inputMessage *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`
/help__rating__customer — print list of commands
/get__rating__customer — get a entity
/list__rating__customer — get a list of your entity
/delete__rating__customer — delete an existing entity

/new__rating__customer — create a new entity
/edit__rating__customer — edit a entity      // not implemented`,
	)

	_, err := c.bot.Send(msg)
	return err
}
