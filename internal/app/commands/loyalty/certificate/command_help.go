package certificate

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *LoyaltyCertificateCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__loyalty__certificate  - help\n"+
			"/list__loyalty__certificate  - list of entities\n"+
			"/get__loyalty__certificate [id]  - get entity by index\n"+
			"/delete__loyalty__certificate [id]  - delete an entity\n"+
			"/new__loyalty__certificate [JSON]  - create a new entity\n"+
			"/edit__loyalty__certificate [JSON]  - edit a entity",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("LoyaltyCertificateCommander.Help: error sending reply message to chat - %v", err)
	}
}

