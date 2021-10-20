package transaction

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BankTransactionCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__bank__transaction - this help\n"+
			"/list__bank__transaction - list of all transactions\n"+
			"/new__bank__transaction - add transaction, , not implemented\n"+
			"/get__bank__transaction id - get transaction by id\n"+
			"/edit__bank__transaction id - edit info about transaction, not implemented\n"+
			"/delete__bank__transaction id - delete info about transaction by id\n\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("BankTransactionCommander.Help: error sending reply message to chat - %v", err)
	}
}
