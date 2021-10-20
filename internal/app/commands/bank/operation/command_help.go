package operation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *BankOperationCommander) Help(inputMsg *tgbotapi.Message){
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help__bank__operation - available commands \n"+
			"/get__bank__operation <id> - get operation \n"+
			"/list__bank__operation <first id> <limit> - list operations \n"+
			"/delete__bank__operation <id> - delete operation \n" +
			"/new__bank__operation <operationType> <transactionID> - create a new operation\n"+
			"/edit__bank__operation - edit an operation",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OperationCommander.Help: error sending reply message to chat - %v", err)
	}
}