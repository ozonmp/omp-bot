package warehouse

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *EstateCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		`Estate / commands:
						/help__estate__warehouse - print list of commands
						/get__estate__warehouse $WAREHOUSE_ID - get a entity
						/list__estate__warehouse - get a list of your entity
						/delete__estate__warehouse $WAREHOUSE_ID - delete an existing entity
						/new__estate__warehouse  - create a new entity
						/edit__estate__warehouse -  $WAREHOUSE_ID — edit a entity
						`,
	)

	msg, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("EstateWarehouse.Help: error sending reply message to chat - %v", err)
	}
}
