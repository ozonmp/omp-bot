package return1

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Return1CommanderImpl) Help(inputMsg *tgbotapi.Message) {
	reply := func(text string, other ...interface{}) {
		for _, v := range other {
			log.Println("Return1CommanderImpl.Help:", v)
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("Return1CommanderImpl.Help: error sending reply message to chat [%v]", err)
		}
	}

	reply(
		`
		\help 					- list all commands
		\list 					- list all elements
		\get *id* 				- get element by id
		\delete *id* 			- delete element by id
		\new *name* 			- create new element with name
		\edit *id* *name* 		- edit element by id
		`,
	)
}
