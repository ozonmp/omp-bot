package transaction

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *BankTransactionCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	isSuccess, err := c.transactionService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to get transaction with idx %d: %v", idx, err)
		return
	}

	var msgText string
	if isSuccess {
		msgText = fmt.Sprintf("TransactionID %v is deleted.", idx)
	} else {
		msgText = fmt.Sprintf("TransactionID %v is not found.", idx)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("BankTransactionCommander.Delete: error sending reply message to chat - %v", err)
	}
}
