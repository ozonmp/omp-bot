package operation

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c BankOperationCommander) Edit(inputMsg *tgbotapi.Message){
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"Not implemented",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("OperationCommander.Edit: error sending reply message to chat - %v", err)
	}
}
