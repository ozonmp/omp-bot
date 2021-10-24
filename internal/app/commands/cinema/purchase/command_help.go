package purchase

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PurchaseCommanderImpl) Help(inputMsg *tgbotapi.Message) {
	replyToUser(
		`
		\help 					- list all commands
		\list 					- list all elements
		\get *id* 				- get element by id
		\delete *id* 			- delete element by id
		\new *json* 			- create new element with name
		\edit *json* 			- edit element by id
		`,
		inputMsg,
		c.bot,
	)
}
