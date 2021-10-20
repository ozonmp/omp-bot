package transaction

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BankTransactionCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var message string
	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("BankTransactionCommander.Get: wrong args", args)
		message = "wrong args"
	} else {
		transaction, err := c.transactionService.Describe(uint64(idx))
		if err != nil {
			log.Printf("BankTransactionCommander.Get: fail to get transaction with idx %d: %v", idx, err)
			message = "fail to get transaction"
		} else {
			message = transaction.String()
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		message,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankTransactionCommander.Get: error sending reply message to chat - %v", err)
	}
}
