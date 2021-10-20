package order

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *OrderCommander) Help(inputMessage *tgbotapi.Message) {
	c.Reply(
		inputMessage.Chat.ID,
		"Available commands:\n"+
			`		/help__buy__order — print list of commands
		/get__buy__order <id> (id >= 0) — get an order
		/list__buy__order   — get a list of orders
		/delete__buy__order <id> (id >= 0) — delete an existing order
		/new__buy__order {"title": <string>, "quantity": <number>} — place a new order
		/edit__buy__order {"id": <number>, "title": <string>, "quantity": <number>} — edit an existing order`,
	)
}
