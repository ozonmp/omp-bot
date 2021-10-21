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
		/new__buy__order {"user_id": <number>, "address_id": <number>, "state_id":<number>, "paid": <true|false>} — place a new order
		/edit__buy__order {"id": <number>, "user_id": <number>, "address_id": <number>, "state_id":<number>, "paid": <true|false>} — edit an existing order`,
	)
}
