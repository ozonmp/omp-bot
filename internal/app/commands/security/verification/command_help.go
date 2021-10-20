package verification

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *SecurityVerificationCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__security__verification - help\n"+
			"/list__security__verification - list products\n"+
			"/get__security__verification - get product\n"+
			"/delete__security__verification - delete product\n"+
			"/new__security__verification - create product\n"+
			"/edit__security__verification - update product\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("SecurityVerificationCommander.Help: error sending reply message to chat - %v", err)
	}
}
